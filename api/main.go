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
		os.Exit(1)
	}
  fmt.Println(appConfig.Port)
  fmt.Println(appConfig.DBurl)

	// db, err := config.InitDB(appConfig.DBurl)
  db := repo.NewDAO()
	if err != nil {
		log.Printf("error initialzing db: %v", err)
		os.Exit(1)
	}

  srv := server.NewServer(db)

  // server := &http.Server{
  //   Addr: ":"+appConfig.Port,
  //   Handler: srv,
  // }

  if err := http.ListenAndServe(":"+appConfig.Port, srv); err != nil {
    log.Fatal(err)
  }

	return nil
}
