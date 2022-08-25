package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sabino-ramirez/jems/api/config"
	"github.com/sabino-ramirez/jems/api/repo"
	"github.com/sabino-ramirez/jems/api/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() error {
	appConfig, err := config.NewAppConfig()
	if err != nil {
		log.Printf("%v", err)
		// os.Exit(1)
	}

	dao := repo.NewDAO()
  db, err := config.InitDB("./jems_db")
	if err != nil {
		log.Printf("error initialzing db: %v", err)
		// os.Exit(1)
	}

	srv := server.NewServer(db, dao)

	if err := http.ListenAndServe(":"+appConfig.Port, srv); err != nil {
		log.Fatal(err)
	}

	return nil
}
