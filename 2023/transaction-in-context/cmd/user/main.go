package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/jackc/pgx/v5"

	cmdinternal "github.com/MarioCarrion/videos/2023/transaction-in-context/cmd/internal"
	"github.com/MarioCarrion/videos/2023/transaction-in-context/internal/postgresql"
)

func main() {
	var name string

	flag.StringVar(&name, "name", "", "name of the user to insert")
	flag.Parse()

	if name == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	//-

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cmdinternal.NewConnString())
	if err != nil {
		log.Fatalln("Connection error:", err)
	}

	userRepo := postgresql.NewUser(conn)

	user, err := userRepo.Insert(ctx, name)
	if err != nil {
		log.Fatalln("Insert error:", err)
	}

	cmdinternal.Print(&user)
}
