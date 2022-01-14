package port

import (
	"context"
	"time"

	"github.com/sprakhar77/faceit/internal/core/domain"
)

// GetUsersFilter provides filter options to fetch users from the repository
type GetUsersFilter struct {
	Country       string
	CreatedBefore *time.Time
	CreatedAfter  *time.Time
	UpdatedBefore *time.Time
	UpdatedAfter  *time.Time
	Limit         *uint64
	Offset        *uint64
}

// UserRepository defines the database layer for the User resource
type UserRepository interface {
	GetByID(ctx context.Context, userID int64) (*domain.User, error)
	GetAll(ctx context.Context, filter GetUsersFilter) ([]domain.User, error)
	Create(ctx context.Context, user domain.User) (int64, error)
	Update(ctx context.Context, userID int64, user domain.User) error
	Delete(ctx context.Context, userID int64) error
}
