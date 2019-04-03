package main

import (
	"github.com/catcatio/shio/app/database"
	"github.com/catcatio/shio/migrations"
)

func main() {
	db := database.MakeCloudSQLProxyDBConnection()
	migrations.Migrate(db)
}
