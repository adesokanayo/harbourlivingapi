package rest

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	db "github.com/BigListRyRy/harbourlivingapi/db/sqlc"
	_ "github.com/BigListRyRy/harbourlivingapi/docs"

	"github.com/BigListRyRy/harbourlivingapi/graphql"
	"github.com/BigListRyRy/harbourlivingapi/token"
	"github.com/BigListRyRy/harbourlivingapi/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//HTTPServer defines our store and router field
type HTTPServer struct {
	store        *db.Store
	tokenService token.TokenService
	router       *gin.Engine
	config       util.Config
}

type HttpServerOpts struct {
	Store        *db.Store
	Config       util.Config
	TokenService token.TokenService
}

//NewHTTPServer creates a server instance
func NewHTTPServer(opts HttpServerOpts) (*HTTPServer, error) {

	server := &HTTPServer{
		config:       opts.Config,
		store:        opts.Store,
		tokenService: opts.TokenService,
	}
	router := gin.Default()
	router.POST("/api/v1/users", server.CreateUser)
	router.GET("/api/v1/users", server.ListUsers)
	router.GET("/api/v1/users/:id", server.GetUser)

	router.POST("/api/v1/login", server.Login)

	router.POST("/api/v1/events", server.CreateEvent)
	router.GET("/api/v1/events", server.ListEvents)
	router.GET("/api/v1/api/v1/events/:id", server.GetEvent)

	router.POST("/api/v1/venues", server.CreateVenue)
	router.GET("/api/v1/venues", server.ListVenues)
	router.GET("/api/v1/venues/:id", server.GetVenue)

	//server.router = router
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())

	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//============ Start ================//
func (s *HTTPServer) Start(address string) error {
	return s.router.Run(address)
}

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GetDatabase() {

}
