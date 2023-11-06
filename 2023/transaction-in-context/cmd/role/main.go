package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/jackc/pgx/v5"

	cmdinternal "github.com/MarioCarrion/videos/2023/transaction-in-context/cmd/internal"
	"github.com/MarioCarrion/videos/2023/transaction-in-context/internal"
	"github.com/MarioCarrion/videos/2023/transaction-in-context/internal/postgresql"
)

func main() {
	var (
		name                             string
		createP, readP, deleteP, updateP bool
	)

	flag.StringVar(&name, "name", "", "name of the role to insert")
	flag.BoolVar(&createP, "c", false, "indicates whether 'create' permission is enabled or not")
	flag.BoolVar(&readP, "r", false, "indicates whether 'read' permission is enabled or not")
	flag.BoolVar(&deleteP, "d", false, "indicates whether 'delete' permission is enabled or not")
	flag.BoolVar(&updateP, "u", false, "indicates whether 'update' permission is enabled or not")

	flag.Parse()

	if name == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	//-

	permissions := make([]internal.Permission, 0, 4)

	updatePermissions := func(enabled bool, ptype string) {
		if !enabled {
			return
		}

		permissions = append(permissions,
			internal.Permission{
				Type: internal.PermissionType(ptype),
			})
	}

	updatePermissions(createP, "create")
	updatePermissions(readP, "read")
	updatePermissions(updateP, "update")
	updatePermissions(deleteP, "delete")

	//-

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cmdinternal.NewConnString())
	if err != nil {
		log.Fatalln("Connection error", err)
	}

	roleRepo := postgresql.NewRole(conn)

	r, err := roleRepo.Insert(ctx, name, permissions)
	if err != nil {
		log.Fatalln("Inserting error:", err)
	}

	cmdinternal.Print(&r)
}
