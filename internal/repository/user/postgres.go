package user

import (
	"context"
	"fmt"
	"github.com/sprakhar77/faceit/internal/crypto"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	logger "github.com/sirupsen/logrus"

	"github.com/sprakhar77/faceit/internal/core/domain"
	"github.com/sprakhar77/faceit/internal/core/port"
)

// PostgresAdapter is the concrete implementation for ports.UserRepository
// using Postgres as the database
type PostgresAdapter struct {
	client       *sqlx.DB
	queryTimeout time.Duration
}

func NewPostgresAdapter(client *sqlx.DB, queryTimeout time.Duration) *PostgresAdapter {
	return &PostgresAdapter{client: client, queryTimeout: queryTimeout}

}

func (p PostgresAdapter) GetByID(ctx context.Context, userID int64) (*domain.User, error) {
	logger.Debugf("fetching user id: %d", userID)
	q := sq.Select("*").From("users").Where(sq.Eq{"id": userID})

	sql, args, err := q.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to convert to sql: %w", err)
	}

	ctx, cancelFunc := context.WithTimeout(ctx, p.queryTimeout)
	defer cancelFunc()

	user := domain.User{}
	logger.Debug("executing query: ", sql)
	err = p.client.GetContext(ctx, &user, sql, args...)
	if err != nil {
		logger.Error(err.Error())
		return nil, fmt.Errorf("failed to execute select query: %w", err)
	}

	user.Password = crypto.Decrypt(user.Password)

	return &user, nil
}

func (p PostgresAdapter) GetAll(ctx context.Context, filter port.GetUsersFilter) ([]domain.User, error) {
	logger.Debug("fetching all users with given filter")
	q := sq.Select("*").From("users")

	if len(filter.Country) != 0 {
		q = q.Where(sq.Eq{"country": filter.Country})
	}

	if filter.CreatedBefore != nil {
		q = q.Where(sq.LtOrEq{"created_before": filter.CreatedBefore})
	}

	if filter.CreatedAfter != nil {
		q = q.Where(sq.GtOrEq{"created_after": filter.CreatedAfter})
	}

	if filter.UpdatedBefore != nil {
		q = q.Where(sq.LtOrEq{"updated_before": filter.UpdatedBefore})
	}

	if filter.UpdatedAfter != nil {
		q = q.Where(sq.GtOrEq{"updated_after": filter.UpdatedAfter})
	}

	if filter.Limit != nil {
		q = q.Limit(*filter.Limit)
	}

	if filter.Offset != nil {
		q = q.Offset(*filter.Offset)
	}

	sql, args, err := q.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to convert to sql: %w", err)
	}

	ctx, cancelFunc := context.WithTimeout(ctx, p.queryTimeout)
	defer cancelFunc()

	var users []*domain.User
	logger.Debug("executing query: ", sql)
	err = p.client.SelectContext(ctx, &users, sql, args...)
	if err != nil {
		logger.Error(err.Error())
		return nil, fmt.Errorf("failed to execute select query: %w", err)
	}

	var result []domain.User
	for _, user := range users {
		if user == nil {
			continue
		}
		user.Password = crypto.Decrypt(user.Password)
		result = append(result, *user)
	}

	return result, nil
}

func (p PostgresAdapter) Create(ctx context.Context, user domain.User) (int64, error) {
	logger.Debug("creating a new user")

	user.Password = crypto.Encrypt(user.Password)
	var id int64
	q := sq.Insert("users").Columns("first_name, last_name, nickname, email, password, country").Values(
		user.FirstName, user.LastName, user.NickName, user.Email, user.Password, user.Country).Suffix("RETURNING id")

	sql, args, err := q.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to convert to sql: %w", err)
	}

	ctx, cancelFunc := context.WithTimeout(ctx, p.queryTimeout)
	defer cancelFunc()

	logger.Debug("executing query: ", sql)
	err = p.client.QueryRowxContext(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

func (p PostgresAdapter) Update(ctx context.Context, userID int64, user domain.User) error {
	logger.Debugf("updating user with id: %d", userID)

	user.Password = crypto.Encrypt(user.Password)

	q := sq.Update("users").Where(sq.Eq{"id": userID}).Set("first_name", user.FirstName).Set("last_name", user.LastName).Set(
		"nickname", user.NickName).Set("email", user.Email).Set("password", user.Password).Set(
		"country", user.Country)

	ctx, cancelFunc := context.WithTimeout(ctx, p.queryTimeout)
	defer cancelFunc()

	sql, args, err := q.ToSql()
	if err != nil {
		return fmt.Errorf("failed to convert to sql: %w", err)
	}

	logger.Debug("executing query: ", sql)
	_, err = p.client.ExecContext(ctx, sql, args...)
	return err
}

func (p PostgresAdapter) Delete(ctx context.Context, userID int64) error {
	logger.Debugf("deleting user with id: %d", userID)
	q := sq.Delete("users").From("users").Where(sq.Eq{"id": userID})

	ctx, cancelFunc := context.WithTimeout(ctx, p.queryTimeout)
	defer cancelFunc()

	sql, args, err := q.ToSql()
	if err != nil {
		return fmt.Errorf("failed to convert to sql: %w", err)
	}

	logger.Debug("executing query: ", sql)
	_, err = p.client.ExecContext(ctx, sql, args...)
	return err
}
