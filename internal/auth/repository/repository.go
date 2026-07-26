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


func (r *AuthRepository) Register (ctx context.Context, req dto.RegisterRequest) (int64, error) { 
  tx, err := r.db.Begin(ctx) 
  if err != nil { 
    return 0, err
  }
  var UserID int64
  defer tx.Rollback(ctx)
  err = tx.QueryRow(
    ctx, `
      INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id
    `, req.Email, req.Password,
  ).Scan(&UserID)

  _, err = tx.Exec(
    ctx, `
      INSERT INTO user_profiles (user_id, fullname) VALUES ($1, $2)
    `, UserID, req.FullName,
  ) 

  if err := tx.Commit(ctx); err != nil { 
    return 0, err 
  }
  return UserID, err
}

func (r *AuthRepository) Login (ctx context.Context ,req dto.LoginRequest) (userDomain.User, error) { 
  var user = userDomain.User{}
  err := r.db.QueryRow(ctx ,`SELECT id, email, password, fullname FROM users JOIN user_profiles ON user.id = user_profiles.user_id WHERE email = $1`, req.Email).Scan(&user.ID, &user.Email, &user.Password, &user.Fullname)

  if err != nil { 
    return user, errors.New("Failed")
  }
  return user, err
}

