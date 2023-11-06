package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	cmdinternal "github.com/MarioCarrion/videos/2023/transaction-in-context/cmd/internal"
	"github.com/MarioCarrion/videos/2023/transaction-in-context/internal/postgresql"
)

func main() {
	var (
		userID  string
		roleIDs []uuid.UUID
	)

	flag.StringVar(&userID, "id", "", "id of user to select")
	flag.Func("roles", "comma separated list of Role IDs", func(s string) error {
		for _, str := range strings.Split(s, ",") {
			id, err := uuid.Parse(str)
			if err != nil {
				return fmt.Errorf("UUID Parsing error %w", err)
			}

			roleIDs = append(roleIDs, id)
		}

		return nil
	})
	flag.Parse()

	if userID == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		log.Fatalln("UUID Parsing error:", err)
	}

	//-

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cmdinternal.NewConnString())
	if err != nil {
		log.Fatalln("Connection error:", err)
	}

	userRoleRepo := postgresql.NewUserRole(conn)

	if err := userRoleRepo.Insert(ctx, id, roleIDs...); err != nil {
		log.Fatalln("userRoleRepo.Insert", err)
	}

	fmt.Println("Roles inserted")
}
