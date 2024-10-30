package user

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetByEmail(email string) (*User, error)
	GetByID(id uuid.UUID) (*User, error)
	Add(User) error
}

type SQLUserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{
		DB: db,
	}
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt string    `json:"created_at"`
}

func (r *SQLUserRepository) Add(user User) error {
	query := "INSERT INTO users (username, email, password) VALUES ($1,$2,$3)"

	result, err := r.DB.Exec(query, user.Username, user.Email, user.Password)

	if err != nil {
		return err
	}

	num, err := result.RowsAffected()
	if num == 0 || err != nil {
		return err
	}

	return nil
}

func (r *SQLUserRepository) GetByEmail(email string) (*User, error) {
	query := "SELECT user_id, username, email, password, created_at FROM users WHERE email = $1"

	rows, err := r.DB.Query(query, email)

	if err != nil || !rows.Next() {
		return nil, err
	}

	defer rows.Close()

	u := new(User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	return u, nil
}

func (r *SQLUserRepository) GetByID(id uuid.UUID) (*User, error) {
	query := "SELECT user_id, username, email, password, created_at FROM users WHERE user_id= $1"

	rows, err := r.DB.Query(query, id)

	if err != nil || !rows.Next() {
		return nil, err
	}

	defer rows.Close()

	u := new(User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	return u, nil
}
func scanRowIntoUser(rows *sql.Rows) (*User, error) {
	user := new(User)

	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
