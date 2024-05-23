package main

import (
	"database/sql"
	"kissandeat/cmd/api"
	"kissandeat/config"
	"kissandeat/internal/repository"
	"log"
)

func main() {
	// Load configuration
	dbConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// configFile, err := ioutil.ReadFile("config.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var config config.Config
	// err = json.Unmarshal(configFile, &config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// host := "localhost"
	// port := 3306
	// user := "my_user"
	// password := "my_password"
	// dbName := "my_database"

	// Create a database connection variable
	var db *sql.DB

	// Create a repository instance (replace with your actual implementation)
	repo := repository.NewSQLRepository(db) // Assuming SQL

	// Start the API server
	log.Printf("Starting API server on port %d", dbConfig.Port)
	err = api.StartServer(dbConfig, &repo)
	if err != nil {
		log.Fatal(err)
	}
}
