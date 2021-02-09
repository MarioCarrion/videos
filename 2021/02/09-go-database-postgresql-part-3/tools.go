// +build tools

package main

// Install using:
//
// go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate
// go install github.com/MarioCarrion/complex-pipelines/part5
// go install github.com/kyleconroy/sqlc/cmd/sqlc

import (
	_ "github.com/MarioCarrion/complex-pipelines/part5"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/kyleconroy/sqlc/cmd/sqlc"
)
