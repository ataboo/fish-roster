package graph

import "github.com/ataboo/fish-roster/fish-roster-api/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreatureRepo *db.CreatureRepo
}
