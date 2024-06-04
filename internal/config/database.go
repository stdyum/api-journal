package config

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/cassandra"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	URL      string `env:"URL"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Keyspace string `env:"KEYSPACE"`
}

func ConnectToDatabase(config DatabaseConfig) (*gocql.Session, error) {
	retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
		Min:        time.Second,
		Max:        10 * time.Second,
		NumRetries: 50,
	}

	cluster := gocql.NewCluster(config.URL)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: config.Username,
		Password: config.Password,
	}
	cluster.Keyspace = config.Keyspace
	cluster.Timeout = 60 * time.Second
	cluster.RetryPolicy = retryPolicy
	cluster.ReconnectionPolicy = &gocql.ConstantReconnectionPolicy{
		MaxRetries: 100,
		Interval:   1 * time.Second,
	}
	cluster.ConnectTimeout = 10 * time.Second
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	if err = runMigrations(session, config.Keyspace); err != nil {
		return nil, err
	}

	return session, nil
}

func runMigrations(db *gocql.Session, keyspace string) error {
	driver, err := cassandra.WithInstance(db, &cassandra.Config{
		MigrationsTable: "a_migrations", // always on top
		KeyspaceName:    keyspace,
	})
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		"file://scheme/migrations",
		"cassandra", driver)
	if err != nil {
		return err
	}

	// ignore error because of 'no change'
	err = migrations.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}
