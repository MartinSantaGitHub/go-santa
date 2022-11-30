package db

import (
	"log"
	"os"

	df "db/file"
	dns "db/nosql"
)

type DbAdapter interface {
	// Connection
	Connect() error
	IsConnection() bool

	//Greetings
	SaveName(name string) (bool, error)
	GetNames() ([]string, error)
}

/* DbConn is the connection to the database */
var DbConn DbAdapter

/* SetDataBaseConnector sets the connector to the database type */
func SetDataBaseConnector(dbType string) {
	switch dbType {
	case "NoSql":
		dbNoSql := new(dns.DbNoSql)

		dbNoSql.Connect()

		//DbConn = dbNoSql
	case "File":
		dbFile := new(df.DbFile)

		dbFile.FilePath = os.Getenv("DB_FILE_PATH")

		dbFile.Connect()

		DbConn = dbFile
	default:
		log.Fatal("No database connector selected")
	}
}
