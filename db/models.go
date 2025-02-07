// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserRole string

const (
	UserRoleCustomer UserRole = "customer"
	UserRoleAdmin    UserRole = "admin"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole `json:"user_role"`
	Valid    bool     `json:"valid"` // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

func (e UserRole) Valid() bool {
	switch e {
	case UserRoleCustomer,
		UserRoleAdmin:
		return true
	}
	return false
}

func AllUserRoleValues() []UserRole {
	return []UserRole{
		UserRoleCustomer,
		UserRoleAdmin,
	}
}

type Expense struct {
	ID          int                `json:"id"`
	UserID      int                `json:"user_id"`
	Description string             `json:"description"`
	Amount      pgtype.Numeric     `json:"amount"`
	Date        pgtype.Timestamptz `json:"date"`
}

type User struct {
	ID        int                `json:"id"`
	Email     *string            `json:"email"`
	Username  string             `json:"username"`
	Role      NullUserRole       `json:"role"`
	Password  string             `json:"password"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}
