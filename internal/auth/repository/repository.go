package repository

import (
	"belimudah/internal/auth/dto"
	userDomain "belimudah/internal/user/domain"
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct { 
  db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository{ 
  return &AuthRepository{db: db}
}


func (r *AuthRepository) Register (ctx context.Context, req dto.RegisterRequest) (string, error) { 
  tx, err := r.db.Begin(ctx) 
  if err != nil { 
    return "", err
  }
  var UserID int64
  defer tx.Rollback(ctx)
  err = tx.QueryRow(
    ctx, `
      INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id
    `, req.Email, req.Password,
  ).Scan(&UserID)
  if err != nil {
    return "", err
  }

  _, err = tx.Exec(
    ctx, `
      INSERT INTO user_profiles (user_id, fullname) VALUES ($1, $2)
    `, UserID, req.FullName,
  )
  if err != nil {
    return "", err
  }

  if err := tx.Commit(ctx); err != nil {
    return "", err
  }
  return req.FullName, nil
}

func (r *AuthRepository) Login(ctx context.Context, req dto.LoginRequest) (userDomain.User, error) {
	var user userDomain.User
	err := r.db.QueryRow(ctx, `
		SELECT u.id, u.email, u.password, u.role, p.fullname
		FROM users u
		JOIN user_profiles p ON u.id = p.user_id
		WHERE u.email = $1
	`, req.Email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.Fullname,
	)
	if err != nil {
		return user, errors.New("invalid email or password")
	}
	return user, nil
}

