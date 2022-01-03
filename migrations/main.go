package migrations

import (
	"database/sql"
	"forum/config"
	"log"
)

//Run migrates
func Run() {
	migrate(config.DB, Notes)
}

func migrate(dbDriver *sql.DB, query string) {
	statement, err := dbDriver.Prepare(query)
	if err == nil {
		_, creationError := statement.Exec()
		if creationError == nil {
			log.Println("Table created successfully")
		} else {
			log.Println(creationError.Error())
		}
	} else {
		log.Println(err.Error())
	}

}
