package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Todo struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func InitRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", GetATodo)
	router.Delete("/{todoID}", DeleteTodo)
	router.Post("/", CreateTodo)
	router.Get("/", GetAllTodos)
	return router
}

func GetATodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		Id:    todoID,
		Title: "Hello world",
	}
	render.JSON(w, r, todos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted TODO successfully"
	render.JSON(w, r, response)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created TODO successfully"
	render.JSON(w, r, response)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{
			Id:    "1",
			Title: "Hello world",
		},
		{
			Id:    "2",
			Title: "Hello world",
		},
	}
	render.JSON(w, r, todos)
}
