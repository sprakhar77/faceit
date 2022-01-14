package server

import (
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/sprakhar77/faceit/internal/config"
	userservice "github.com/sprakhar77/faceit/internal/core/service/user"
	"github.com/sprakhar77/faceit/internal/handler"
	userepub "github.com/sprakhar77/faceit/internal/publisher/user"
	userrepo "github.com/sprakhar77/faceit/internal/repository/user"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Dependency struct {
	UserHandler *handler.UserHandler
}

func initDependencies(cfg config.Application) (*Dependency, error) {
	// Support of Postgres
	sq.StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	d := Dependency{}

	db, err := sqlx.Connect("pgx", connString(cfg.Database))
	if err != nil {
		return nil, fmt.Errorf("failed to create database connection: %w", err)
	}

	//TODO: Add concrete Kafka connection and pass client to KafkaAdapter, just like DB
	userPublisher := userepub.NewKafkaAdapter()

	userRepository := userrepo.NewPostgresAdapter(db, cfg.Database.QueryTimeout)
	userService := userservice.NewService(userRepository, userPublisher)
	d.UserHandler = handler.NewUserHandler(userService)

	return &d, nil
}

func connString(cfg config.Database) string {
	var opts []string
	for name, val := range map[string]string{
		"dbname":   cfg.Name,
		"user":     cfg.User,
		"password": cfg.Password,
		"host":     cfg.Host,
		"port":     cfg.Port,
		"sslmode":  cfg.SSLMode,
	} {
		if val != "" {
			opts = append(opts, fmt.Sprintf("%s=%s", name, val))
		}
	}
	return strings.Join(opts, " ")
}
