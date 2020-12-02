package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ganselmo/go-book-store-api/internal/config"
	"github.com/ganselmo/go-book-store-api/internal/database"
	book "github.com/ganselmo/go-book-store-api/internal/service/books"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {

	// Lectura de configuracion
	cfg := readConfig()

	// Instanciacion de db
	db, err := database.NewDatabase(cfg)
	defer db.Close()
	createSchema(db)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Instanciacion un servicio y le inyecta la configuracion y la base de datos
	service, _ := book.NewBookService(db, cfg)
	httpService := book.NewHTTPtransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "config service")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}

func createSchema(db *sqlx.DB) error {

	schema := `CREATE TABLE IF NOT EXISTS books(
		id integer primary key autoincrement,
		name text,
		description text,
		rating integer,
		pages integer,
		author text);`
	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	return nil

}
