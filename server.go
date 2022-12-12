package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/issy20/go-task/graph/generated"
	"github.com/issy20/go-task/graph/interface/resolver"
	"github.com/issy20/go-task/graph/middleware"
	"github.com/issy20/go-task/graph/pkg/db"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// dsn := os.Getenv("MYSQL_DSN")
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	// if db == nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	if db != nil {
	// 		sqlDb, err := db.DB()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		sqlDb.Close()
	// 	}
	// }()

	db := db.InitDB()

	defer func() {
		if db != nil {
			sqlDb, err := db.DB()
			if err != nil {
				panic(err)
			}
			sqlDb.Close()
		}
	}()

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver.NewResolver(db),
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.Middleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
