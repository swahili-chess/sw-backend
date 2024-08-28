// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateToken(ctx context.Context, arg CreateTokenParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) error
	DeleteToken(ctx context.Context, arg DeleteTokenParams) error
	DeleteUserById(ctx context.Context, id uuid.UUID) error
	GetLichessTeamMembers(ctx context.Context) ([]Lichess, error)
	GetUserByPhone(ctx context.Context, phoneNumber string) (GetUserByPhoneRow, error)
	GetUserByToken(ctx context.Context, arg GetUserByTokenParams) (GetUserByTokenRow, error)
	GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error)
	UpdateUserById(ctx context.Context, arg UpdateUserByIdParams) error
}

var _ Querier = (*Queries)(nil)