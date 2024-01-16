package main

import (
	"fmt"
	"os"

	"github.com/asfarsharief/money_management_backend/common/configservice"
	log "github.com/asfarsharief/money_management_backend/common/logingservice"
	"github.com/asfarsharief/money_management_backend/dbmigration"
	"github.com/asfarsharief/money_management_backend/lib/config"
	"github.com/asfarsharief/money_management_backend/lib/server"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "<appName>"
	app.Usage = "<Usage>"
	app.UsageText = "<Usage Text>"
	app.Version = "latest"

	app.Commands = []cli.Command{
		cli.Command{
			Name:      "serve",
			Usage:     "Run an API Server",
			UsageText: "serve -e ENV -p PORT --host HOST --api-path API_PATH",
			Action:    RunServer,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "ENV, e",
					Usage:  "environment (dev | test | stage | load | prod)",
					EnvVar: "ENV",
				},
				cli.StringFlag{
					Name:   "HOST",
					Usage:  "The host to listen on.",
					EnvVar: "HOST",
				},
				cli.StringFlag{
					Name:   "API_PATH",
					Usage:  "url path prefix for mounting api router",
					EnvVar: "API_PATH",
				},
				cli.StringFlag{
					Name:   "PORT, p",
					Usage:  "port to listen on.",
					EnvVar: "PORT",
				},
				cli.StringFlag{
					Name:   "WEB_DIR",
					Usage:  "Specify path to local web assests (e.g. Swagger UI)",
					EnvVar: "WEB_DIR",
				},
				cli.StringFlag{
					Name:   "DBHOST",
					Usage:  "Database Host",
					EnvVar: "DBHOST",
				},
				cli.StringFlag{
					Name:   "DBNAME",
					Usage:  "Database Name",
					EnvVar: "DBNAME",
				},
				cli.StringFlag{
					Name:   "DBUSER",
					Usage:  "Database username",
					EnvVar: "DBUSER",
				},
				cli.StringFlag{
					Name:   "DBPORT",
					Usage:  "Database port number",
					EnvVar: "DBPORT",
				},
				cli.StringFlag{
					Name:   "DBPASSWORD",
					Usage:  "Database password",
					EnvVar: "DBPASSWORD",
				},
			},
		},
		cli.Command{
			Name:      "migrate_db",
			Usage:     "migrate db",
			UsageText: "serve -e ENV -p PORT --host HOST --api-path API_PATH",
			Action:    RunDBMigrations,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "ENV, e",
					Usage:  "environment (dev | test | stage | load | prod)",
					EnvVar: "ENV",
				},
				cli.StringFlag{
					Name:   "HOST",
					Usage:  "The host to listen on.",
					EnvVar: "HOST",
				},
				cli.StringFlag{
					Name:   "API_PATH",
					Usage:  "url path prefix for mounting api router",
					EnvVar: "API_PATH",
				},
				cli.StringFlag{
					Name:   "PORT, p",
					Usage:  "port to listen on.",
					EnvVar: "PORT",
				},
				cli.StringFlag{
					Name:   "WEB_DIR",
					Usage:  "Specify path to local web assests (e.g. Swagger UI)",
					EnvVar: "WEB_DIR",
				},
				cli.StringFlag{
					Name:   "DBHOST",
					Usage:  "Database Host",
					EnvVar: "DBHOST",
				},
				cli.StringFlag{
					Name:   "DBNAME",
					Usage:  "Database Name",
					EnvVar: "DBNAME",
				},
				cli.StringFlag{
					Name:   "DBUSER",
					Usage:  "Database username",
					EnvVar: "DBUSER",
				},
				cli.StringFlag{
					Name:   "DBPORT",
					Usage:  "Database port number",
					EnvVar: "DBPORT",
				},
				cli.StringFlag{
					Name:   "DBPASSWORD",
					Usage:  "Database password",
					EnvVar: "DBPASSWORD",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Errorf("%s", err)
		os.Exit(1)
	}
}

// RunServer - Function that will run the server
func RunServer(c *cli.Context) error {

	runEnvMap := getContextDefaultMap(c)
	cfg, err := configservice.LoadConfiguration(runEnvMap, c.String("env"))

	if err != nil {
		log.Error("Error while loading configuration !!")
		return err
	}

	log.Infof("VERSION:  %s", "latest")
	log.Infof("ENV:      %s", cfg.Server.Env)
	log.Infof("API PATH: %s", cfg.Server.APIPath)
	log.Infof("DB HOST: %s", cfg.Database.DBHost)
	log.Infof("DB Name: %s", cfg.Database.DBName)
	log.Infof("DB User: %s", cfg.Database.DBUser)

	// db, err := dbservice.GetDbInstance(cfg.Database)
	// if err != nil {
	// 	return err
	// }

	if err := RunHTTPServer(cfg, nil); err != nil {
		return err
	}

	return nil
}

func getContextDefaultMap(c *cli.Context) map[string]string {

	var m map[string]string
	m = make(map[string]string)
	m["server.env"] = c.String("ENV")
	m["server.host"] = c.String("HOST")
	m["server.port"] = c.String("PORT")
	m["server.apipath"] = c.String("API_PATH")
	m["server.webdir"] = c.String("WEB_DIR")
	m["database.dbhost"] = c.String("DBHOST")
	m["database.dbname"] = c.String("DBNAME")
	m["database.dbuser"] = c.String("DBUSER")
	m["database.dbport"] = c.String("DBPORT")
	m["database.dbpassword"] = c.String("DBPASSWORD")
	return m
}

// RunHTTPServer - Function that will run HTTP Server
func RunHTTPServer(cfg *config.Configuration, db *gorm.DB) error {
	server := server.NewServer(cfg, db)
	err := server.Start()
	if err != nil {
		log.Errorf("Error while starting HTTP Server. Err:%+v", err)
		return fmt.Errorf("server.Start failed - %s", err.Error())
	}
	return nil
}

func RunDBMigrations(c *cli.Context) error {
	runEnvMap := getContextDefaultMap(c)
	cfg, err := configservice.LoadConfiguration(runEnvMap, c.String("env"))
	if err != nil {
		return err
	}
	dbmigration.Dbmigrations(cfg)
	return nil
}
