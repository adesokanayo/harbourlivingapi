package main

import (
	"database/sql"
	"github.com/BigListRyRy/harbourlivingapi/api"
	db "github.com/BigListRyRy/harbourlivingapi/db/sqlc"
	"github.com/BigListRyRy/harbourlivingapi/util"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {

	// @title Swagger Example API
	// @version 1.0
	// @description This is a sample server celler server.
	// @termsOfService http://swagger.io/terms/

	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email support@swagger.io

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @host localhost:8080
	// @BasePath /api/v1

	// @securityDefinitions.basic BasicAuth

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	// @securitydefinitions.oauth2.application OAuth2Application
	// @tokenUrl https://example.com/oauth/token
	// @scope.write Grants write access
	// @scope.admin Grants read and write access to administrative information

	// @securitydefinitions.oauth2.implicit OAuth2Implicit
	// @authorizationUrl https://example.com/oauth/authorize
	// @scope.write Grants write access
	// @scope.admin Grants read and write access to administrative information

	// @securitydefinitions.oauth2.password OAuth2Password
	// @tokenUrl https://example.com/oauth/token
	// @scope.read Grants read access
	// @scope.write Grants write access
	// @scope.admin Grants read and write access to administrative information

	// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
	// @tokenUrl https://example.com/oauth/token
	// @authorizationUrl https://example.com/oauth/authorize
	// @scope.admin Grants read and write access to administrative information

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot local config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to database, ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store, *config)
	if err != nil {
		log.Fatal("unable to create a new server ", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = config.Port
	}
	err = server.Start(":" +port)
	if err != nil {
		log.Fatal("unable to start server", err)
	}
}
