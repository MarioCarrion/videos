package postgres_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/url"
	"runtime"
	"testing"
	"time"

	migrate "github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	postgres "github.com/MarioCarrion/videos/2025/three-things-i-dont-like/02-test-dependencies/without-dependencies"
)

func TestUserRepository_Create(t *testing.T) {
	// XXX: Make all of this table-based!
	repo := postgres.NewUserRepository(newConn(t))
	user, err := repo.Create(context.Background(), "Mario", 1969)
	if err != nil {
		t.Fatalf("got err: %v", err)
	}

	found, err := repo.Get(context.Background(), user.ID)
	if err != nil {
		t.Fatalf("got err: %v", err)
	}

	if found.BirthYear != user.BirthYear {
		t.Fatalf("got %d, expected %d", found.BirthYear, user.BirthYear)
	}

	if found.Name != user.Name {
		t.Fatalf("got %s, expected %s", found.Name, user.Name)
	}
}

func newConn(tb testing.TB) *pgx.Conn {
	tb.Helper()

	dsn := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("username", "password"),
		Path:   "todo",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	//-

	pool, err := dockertest.NewPool("")
	if err != nil {
		tb.Fatalf("Couldn't connect to docker: %s", err)
	}

	pool.MaxWait = 10 * time.Second

	pw, _ := dsn.User.Password()

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "17.4-bookworm",
		Env: []string{
			fmt.Sprintf("POSTGRES_USER=%s", dsn.User.Username()),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", pw),
			fmt.Sprintf("POSTGRES_DB=%s", dsn.Path),
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		tb.Fatalf("Couldn't start resource: %s", err)
	}

	_ = resource.Expire(60)

	tb.Cleanup(func() {
		if err := pool.Purge(resource); err != nil {
			tb.Fatalf("Couldn't purge container: %v", err)
		}
	})

	dsn.Host = fmt.Sprintf("%s:5432", resource.Container.NetworkSettings.IPAddress)
	if runtime.GOOS == "darwin" { // MacOS-specific
		dsn.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
	}

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		tb.Fatalf("Couldn't open DB: %s", err)
	}

	defer db.Close()

	if err := pool.Retry(func() (err error) {
		return db.Ping()
	}); err != nil {
		tb.Fatalf("Couldn't ping DB: %s", err)
	}

	//-

	instance, err := migratepostgres.WithInstance(db, &migratepostgres.Config{})
	if err != nil {
		tb.Fatalf("Couldn't migrate (1): %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://../migrations/", "pgx5", instance)
	if err != nil {
		tb.Fatalf("Couldn't migrate (2): %s", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		tb.Fatalf("Couldn't migrate (3): %s", err)
	}

	//-

	conn, err := pgx.Connect(context.Background(), dsn.String())
	if err != nil {
		tb.Fatalf("Couldn't open DB Pool: %s", err)
	}

	tb.Cleanup(func() {
		conn.Close(context.Background())
	})

	return conn
}
