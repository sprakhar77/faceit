package user

import (
	"context"
	"fmt"
	"strconv"

	logger "github.com/sirupsen/logrus"
	"github.com/sprakhar77/faceit/internal/core/domain"
	"github.com/sprakhar77/faceit/internal/core/port"
)

// service is the concrete implementation for ports.UserService
type service struct {
	repository port.UserRepository
	publisher  port.UserPublisher
}

func NewService(repository port.UserRepository, publisher port.UserPublisher) *service {
	return &service{repository: repository, publisher: publisher}
}

// GetByID gets the user with the given userID
func (srv *service) GetByID(ctx context.Context, userID string) (domain.User, error) {
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return domain.User{}, fmt.Errorf("invalid id: %w", err)
	}

	user, err := srv.repository.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	return *user, nil
}

// GetAll gets all the users satisfying the given filter
func (srv *service) GetAll(ctx context.Context, filter port.GetUsersFilter) ([]domain.User, error) {
	users, err := srv.repository.GetAll(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}

	return users, nil
}

// Create creates a new user from the data given
func (srv *service) Create(ctx context.Context, user domain.User) (int64, error) {
	user.TrimSpaces()
	id, err := srv.repository.Create(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("create user: %w", err)
	}

	return id, srv.publisher.Created(ctx, user)
}

// Update updates the user with the given userID to new user data
func (srv *service) Update(ctx context.Context, userID string, user domain.User) error {
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	user.TrimSpaces()
	if err := srv.repository.Update(ctx, id, user); err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	logger.Debug("publishing changes for userID: ", id)
	return srv.publisher.Updated(ctx, id, user)
}

// Delete deletes the user with the given userID
func (srv *service) Delete(ctx context.Context, userID string) error {
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	if err := srv.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return srv.publisher.Deleted(ctx, id)
}
