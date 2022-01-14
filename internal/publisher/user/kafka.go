package user

import (
	"context"
	logger "github.com/sirupsen/logrus"
	"github.com/sprakhar77/faceit/internal/core/domain"
)

// KafkaAdapter is the concrete implementation of ports.UserPublisher using
// kafka as the message queue
type KafkaAdapter struct {
	// TODO: Add kafka client
}

func NewKafkaAdapter() KafkaAdapter {
	return KafkaAdapter{}
}

func (k KafkaAdapter) Created(ctx context.Context, user domain.User) error {
	logger.Info("Notify: User created:\n ", user)
	return nil
}

func (k KafkaAdapter) Updated(ctx context.Context, userId int64, user domain.User) error {
	logger.Infof("Notify: userID %d updated:\n %v ", userId, user)
	return nil
}

func (k KafkaAdapter) Deleted(ctx context.Context, userId int64) error {
	logger.Info("Notify: User deleted: ", userId)
	return nil
}
