package model

import (
	"context"
	"time"
	"workspace_booking/migration"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  []byte    `json:"-"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

type Recipient struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

// Recipients struct
type Recipients struct {
	Recipients []Recipient `json:"recipients"`
}

func (u *User) InsertUser() error {
	dt := time.Now()
	// var role_id *int
	err := migration.DbPool.QueryRow(context.Background(), "SELECT id, name, created_at, updated_at FROM roles where roles.name=$1", "Employee").Scan(&u.Role.Id, &u.Role.Name, &u.Role.CreatedAt, &u.Role.UpdatedAt)
	if err != nil {
		return err
	}
	query := "INSERT INTO users (name, email, encrypted_password, role_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at"
	err = migration.DbPool.QueryRow(context.Background(), query, u.Name, u.Email, u.Password, u.Role.Id, dt, dt).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (users *Users) FetchUsers() (*Users, error) {
	rows, err := migration.DbPool.Query(context.Background(),
		"SELECT users.id, users.name, users.email, users.encrypted_password as password, roles.id as role_id, roles.name as role_name, roles.created_at as role_created_at, roles.updated_at as role_updated_at, users.created_at, users.updated_at FROM users LEFT JOIN roles ON users.role_id = roles.id")
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		u := User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role.Id, &u.Role.Name, &u.Role.CreatedAt, &u.Role.UpdatedAt, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return users, err
		}
		users.Users = append(users.Users, u)
	}

	return users, err
}

func (u *User) FetchUser() error {
	return migration.DbPool.QueryRow(context.Background(),
		"SELECT users.id, users.name, users.email, users.encrypted_password as password, roles.id as role_id, roles.name as role_name, roles.created_at as role_created_at, roles.updated_at as role_updated_at, users.created_at, users.updated_at FROM users LEFT JOIN roles ON users.role_id = roles.id WHERE users.id= $1",
		u.ID).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role.Id, &u.Role.Name, &u.Role.CreatedAt, &u.Role.UpdatedAt, &u.CreatedAt, &u.UpdatedAt)
}

func (u *User) UpdateUser() error {
	dt := time.Now()
	_, err := migration.DbPool.Exec(context.Background(), "UPDATE users SET name=$1, email=$2, updated_at=$3 WHERE id=$4",
		u.Name, u.Email, dt, u.ID)
	return err
}

func (u *User) DeleteUser() error {
	_, err := migration.DbPool.Exec(context.Background(), "DELETE FROM users WHERE id=$1", u.ID)
	return err
}

func (u *User) LoginUser() error {
	return migration.DbPool.QueryRow(context.Background(),
		"SELECT users.id, users.name, users.email, users.encrypted_password as password, roles.id as role_id, roles.name as role_name, roles.created_at as role_created_at, roles.updated_at as role_updated_at, users.created_at, users.updated_at FROM users LEFT JOIN roles ON users.role_id = roles.id WHERE users.email= $1 ",
		u.Email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role.Id, &u.Role.Name, &u.Role.CreatedAt, &u.Role.UpdatedAt, &u.CreatedAt, &u.UpdatedAt)
}
