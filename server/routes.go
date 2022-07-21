package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"todo/handler"
	"todo/middleware"
)

type Server struct {
	chi.Router
}

func SetUpRoutes() *Server {
	router := chi.NewRouter()
	router.Route("/todo", func(todo chi.Router) {
		todo.Post("/sign-up", handler.SignUp)
		todo.Post("/sign-in", handler.SignIn)

		todo.Route("/user", func(user chi.Router) {
			user.Use(middleware.AuthMiddleware)

			user.Delete("/delete-user", handler.DeleteUser)
			user.Put("/update-user", handler.UpdateUser)

		})
		todo.Route("/task", func(task chi.Router) {
			task.Use(middleware.AuthMiddleware)
			task.Post("/create-task", handler.CreateTask)
			task.Put("/update-task", handler.UpdateTask)
			task.Get("/fetch-task", handler.FetchTask)
			task.Delete("/delete-task", handler.DeleteTask)
			task.Put("/log-out", handler.LogOut)

		})
	})
	return &Server{router}

}

func (svr *Server) Run(port string) error {
	return http.ListenAndServe(port, svr)

}
