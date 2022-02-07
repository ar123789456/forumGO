package migrations

import (
	"database/sql"
	"forum/config"
	"log"
)

//Run migrates
func Run() {
	migrate(config.DB, Users)
	migrate(config.DB, Posts)
	migrate(config.DB, Categories)
	migrate(config.DB, Tags)
	migrate(config.DB, Comments)
	migrate(config.DB, LikePost)
	migrate(config.DB, CategoryPost)
	migrate(config.DB, TagPosts)
	migrate(config.DB, UserSession)
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
