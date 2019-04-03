package database

import (
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panicf("%s environment variable not set.", k)
	}
	return v
}
func MakeCloudSQLProxyDBConnection() *gorm.DB {
	var (
		connectionName = mustGetenv("CLOUDSQL_CONNECTION_NAME")
		user           = mustGetenv("CLOUDSQL_USER")
		password       = os.Getenv("CLOUDSQL_PASSWORD")
	)
	dbURI := fmt.Sprintf("%s:%s@cloudsql(%s)/default", user, password, connectionName)
	db, dbErr := gorm.Open("mysql", dbURI)
	if dbErr != nil {
		log.Panicf("Initializer database connection error %v", dbErr)
	}
	return db
}
