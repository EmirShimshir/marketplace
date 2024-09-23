package postgres

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate"
	migratepg "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	testpg "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"path/filepath"
	"runtime"
	"time"
)

type Config struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

var (
	postgresConfig = Config{
		Database: "marketplace",
		User:     "postgres",
		Password: "password",
	}
)

func newPostgresContainer(ctx context.Context) (*testpg.PostgresContainer, error) {
	container, err := testpg.RunContainer(
		ctx,
		testpg.WithDatabase(postgresConfig.Database),
		testpg.WithUsername(postgresConfig.User),
		testpg.WithPassword(postgresConfig.Password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	_, path, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get caller path")
	}

	sourceUrl := "file://" + filepath.Dir(path) + "/migrations"
	url, err := container.ConnectionString(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get postgres db url: %s", err)
	}

	db, err := newPostgresDB(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres db: %s", err)
	}
	defer db.Close()

	driver, err := migratepg.WithInstance(db.DB, &migratepg.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to get db driver from instance: %s", err)
	}

	mig, err := migrate.NewWithDatabaseInstance(sourceUrl, postgresConfig.Database, driver)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrator driver: %s", err)
	}

	err = mig.Up()
	if err != nil {
		return nil, fmt.Errorf("failed to up migrations: %s", err)
	}

	err = container.Snapshot(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to make a snapshot of postgres db: %s", err)
	}

	return container, nil
}

const (
	maxConn         = 100
	maxConnIdleTime = 1 * time.Minute
	maxConnLifetime = 3 * time.Minute
)

func newPostgresDB(url string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect postgres db: %s", err)
	}

	db.SetMaxOpenConns(maxConn)
	db.SetConnMaxLifetime(maxConnLifetime)
	db.SetConnMaxIdleTime(maxConnIdleTime)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping postgres db: %s", err)
	}

	return db, nil
}
