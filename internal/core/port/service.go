package port

import (
	"context"

	"github.com/sprakhar77/faceit/internal/core/domain"
)

// UserService is responsible for all business logic related to user
type UserService interface {
	GetByID(ctx context.Context, userID string) (domain.User, error)
	GetAll(ctx context.Context, filter GetUsersFilter) ([]domain.User, error)
	Create(ctx context.Context, user domain.User) (int64, error)
	Update(ctx context.Context, userID string, user domain.User) error
	Delete(ctx context.Context, userID string) error
}
