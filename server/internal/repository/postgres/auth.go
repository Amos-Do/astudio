package postgres

import (
	"context"
	"database/sql"

	"github.com/Amos-Do/astudio/server/domain"
	"go.uber.org/zap"
)

type AuthRepo struct {
	DB *sql.DB
}

// NewAuthRepo will create an implementation of auth.Repository
func NewAuthRepo(db *sql.DB) domain.IAuthRepo {
	return &AuthRepo{DB: db}
}

func (m AuthRepo) fetch(c context.Context, query string, args ...interface{}) ([]domain.Auth, error) {
	rows, err := m.DB.QueryContext(c, query, args...)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			zap.S().Error(err)
		}
	}()

	result := make([]domain.Auth, 0)
	for rows.Next() {
		t := domain.Auth{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Birthday,
			&t.Account,
			&t.Password,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			zap.S().Error(err)
			return nil, err
		}

		result = append(result, t)
	}
	return result, nil
}

func (m *AuthRepo) GetByEmail(c context.Context, email string) (domain.Auth, error) {
	var res = domain.Auth{}

	query := `SELECT id, name, birthday, account, password, created_at, updated_at
				FROM users WHERE account = $1`
	list, err := m.fetch(c, query, email)
	if err != nil {
		return res, err
	}

	if len(list) > 0 {
		res = list[0]
		return res, nil
	} else {
		return res, domain.ErrNotFound
	}
}

func (m *AuthRepo) Create(c context.Context, auth *domain.Auth) error {
	query := `INSERT INTO users (name, account, password)
				VALUES ($1, $2, $3) RETURNING id`
	stmt, err := m.DB.PrepareContext(c, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var userID int64
	err = stmt.QueryRow(auth.Name, auth.Account, auth.Password).Scan(&userID)
	if err != nil {
		return err
	}

	auth.ID = userID
	return nil

	// postgres not support LastInsertId()
	// res, err := stmt.ExecContext(c, auth.Account, auth.Password)
	// if err != nil {
	// 	return err
	// }

	// lastID, err := res.LastInsertId()
	// if err != nil {
	// 	return err
	// }
	// auth.ID = lastID
	// return nil
}
