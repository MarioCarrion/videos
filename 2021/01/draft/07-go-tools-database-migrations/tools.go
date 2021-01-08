// +build tools

package tools

// Run
// go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate

import (
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
)
