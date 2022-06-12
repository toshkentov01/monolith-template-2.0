package migration

import (
	"fmt"
	"log"
	"strings"
	
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"       //file is needed for migration url
	"github.com/toshkentov01/template/pkg/utils"
)

// Up ...
func Up() {
	url, err := utils.ConnectionURLBuilder("migration")
	if err != nil {
		log.Println("Error generating migration url: ", err.Error())
	}
	m, err := migrate.New("file://migrations", url)
	if err != nil {
		log.Fatal("error in creating migrations: ", err.Error())
	}
	fmt.Printf("")
	if err := m.Up(); err != nil {
		if !strings.Contains(err.Error(), "no change") {
			log.Println("Error in migrating ", err.Error())
		}
	}
}