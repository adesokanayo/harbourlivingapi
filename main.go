package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	db "github.com/BigListRyRy/harbourlivingapi/db/sqlc"
	"github.com/BigListRyRy/harbourlivingapi/graphql"
	"github.com/BigListRyRy/harbourlivingapi/middleware"
	"github.com/BigListRyRy/harbourlivingapi/rest"
	"github.com/BigListRyRy/harbourlivingapi/token"
	"github.com/BigListRyRy/harbourlivingapi/util"
)

const defaultPort = "3007"

func main() {
	start()
}

type dependencies struct {
	Config       util.Config
	Repo         *db.Store
	EmailService *util.EmailService
	TokenService token.TokenService
}

func start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	deps := getDependencies()

	httpSvr, err := rest.NewHTTPServer(rest.HttpServerOpts{
		Store:        deps.Repo,
		Config:       deps.Config,
		TokenService: deps.TokenService,
	})
	if err != nil {
		log.Fatal("unable to start http server", err)
	}

	graphQLsrv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{
		Config:       deps.Config,
		Repo:         deps.Repo,
		EmailService: deps.EmailService,
		TokenMaker:   deps.TokenService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.AuthMiddleware(CorsMiddleware(graphQLsrv)))

	http.HandleFunc("/api/v1/verifyemail", httpSvr.VerifyEmail)
	http.HandleFunc("/api/v1/users", httpSvr.ListUsers)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getDependencies() dependencies {
	// Load Config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot find config ", err)
	}

	//Open DB Conn
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to database, ", err)
	}

	store := db.NewStore(conn)

	// Create Token Svc
	tokenService, err := token.NewJWTService(config.TokenSymmetricKey)
	if err != nil {
		log.Fatalln("cannot create a token maker ", err)
	}

	// Create Email Svr
	sendActualMail := true

	if config.ENVIRONMENT == "DEV" {
		sendActualMail = false
	}

	emailServiceOpts := util.EmailServiceOpts{
		APIKey:     config.SibAPIKey,
		PartnerKey: "partner_kay",
		LiveEMail:  sendActualMail,
		URL:        config.EmailURL,
	}

	emailSvc, err := util.NewEmailService(emailServiceOpts)
	if err != nil {
		log.Fatalln("cannot create emailSvc ", err)
	}

	return dependencies{
		Config:       *config,
		Repo:         store,
		EmailService: emailSvc,
		TokenService: tokenService,
	}
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}
