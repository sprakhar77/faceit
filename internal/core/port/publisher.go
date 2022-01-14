package port

import (
	"context"

	"github.com/sprakhar77/faceit/internal/core/domain"
)

// UserPublisher is the publisher/writer for user events
type UserPublisher interface {
	Created(ctx context.Context, user domain.User) error
	Updated(ctx context.Context, userId int64, user domain.User) error
	Deleted(ctx context.Context, userId int64) error
}
