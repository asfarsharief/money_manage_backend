package dbmigration

import (
	"database/sql"
	"fmt"
	"net/url"

	log "github.com/asfarsharief/money_management_backend/common/logingservice"
	"github.com/asfarsharief/money_management_backend/lib/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Dbmigrations(c *config.Configuration) {

	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s",
		c.Database.DBUser, url.QueryEscape(c.Database.DBPassword), c.Database.DBHost, c.Database.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Error(err)
		panic("Error while connecting to DB")
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Error(err)
		panic("Error while creating instance to DB")
	}

	migrationsFolder := fmt.Sprintf("file://%s", "./dbmigration/migrations")
	m, err := migrate.NewWithDatabaseInstance(
		migrationsFolder,
		"postgres", driver)

	//run the sql scripts for database migration
	err = m.Up()
	if err != nil {

		// suppress the error when no new sql files are addded to migrations
		if err == migrate.ErrNoChange {
			log.Info("No new migrations available for postgres database !!")
			return
		}

		log.Errorf("error in migration: %v", err.Error())
		return
	}

	log.Info("Migrations applied successfully for postgres database")

}
