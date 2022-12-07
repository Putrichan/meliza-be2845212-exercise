package main

import (
	"fmt"
	"net/http"
	"quoteapp/controller"
	"quoteapp/db"
	"quoteapp/model"
	"quoteapp/routes"
)

func main() {
	db := db.NewDB(
		Host:         "localhost",
		Username:     "postgres",
		Password:     "asd123",
		DatabaseName: "pgadmin",
		Port:         5432,
		Schema:       "public",
	)

	db.Migrate()
	conn := db.DB()

	model := model.NewQuoteModel(conn)
	controller := controller.NewQuoteController(model)

	router := routes.NewRoute(controller)

	fmt.Println("starting api server at http://localhost:8080")
	http.ListenAndServe(":8080", router.Run())
}
