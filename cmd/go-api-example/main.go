package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/Kitanotori/go-api-example/pkg/db"
	"github.com/Kitanotori/go-api-example/pkg/routes"
)

func InitRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set "Content-Type: application/json"
		middleware.Logger,                             // Log API request calls
		middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes,                    // Redirect slashes to no slash URL versions
		middleware.Recoverer,                          // Recover from panics without crashing server
	)

	router.Route("/api", func(r chi.Router) {
		r.Mount("/v1/todo", routes.InitRoutes())
	})

	return router
}

func main() {
	router := InitRouter()
	db.InitDb()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":8088", router))
}
