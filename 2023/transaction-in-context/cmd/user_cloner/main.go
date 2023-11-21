package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	cmdinternal "github.com/MarioCarrion/videos/2023/transaction-in-context/cmd/internal"
	"github.com/MarioCarrion/videos/2023/transaction-in-context/internal/postgresql"
)

func main() {
	var userID, name string

	flag.StringVar(&userID, "id", "", "id of user to clone")
	flag.StringVar(&name, "name", "", "name of the new user")
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

	userClonerRepo := postgresql.NewUserCloner(conn)

	user, err := userClonerRepo.Clone(ctx, id, name)
	if err != nil {
		log.Fatalln("userClonerRepo.Clone", err)
	}

	cmdinternal.Print(&user)
}
