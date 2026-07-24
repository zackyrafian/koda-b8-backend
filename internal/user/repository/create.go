package repository

import (
	"belimudah/internal/user/dto"
	"context"
)

func (r *UserRepository) Create(ctx context.Context, req dto.CreateUserRequest) (int64, error) { 
  tx, err := r.db.Begin(ctx) 
  if err != nil { 
    return 0, err 
  }
  defer tx.Rollback(ctx) 
  var userID int64 
  
  err = tx.QueryRow(
    ctx, `
      INSERT INTO users(
      email, 
      password
      ) 
     VALUES ($1, $2) 
     RETURNING id 
    `, 
    req.Email, 
    req.Password, 
  ).Scan(&userID)

  if err != nil { 
    return 0, err
  }

  _, err = tx.Exec(
    ctx, `
      INSERT INTO user_profiles(
        user_id, 
        fullname
      )
      VALUES ($1, $2)
    `, userID, req.FullName,
  ) 
  if err != nil { 
    return 0, err
  }

  if err := tx.Commit(ctx); err != nil { 
    return 0, err
  }

  return userID, err
}