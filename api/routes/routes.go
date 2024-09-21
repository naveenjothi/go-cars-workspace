package routes

import (
	"api/handlers"
	"libs/utils"
)

func RegisterFiberRoutes(s *utils.FiberServer) {
	s.Post("/user", s.WithClient(handlers.CreateUserHandler))
	s.Get("/user/:id", s.WithClient(handlers.GetUserHandler))
	s.Post("/user/:id", s.WithClient(handlers.UpdateUserHandler))
}
