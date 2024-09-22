package routes

import (
	"api/handlers"
	"libs/utils"
)

func RegisterFiberRoutes(s *utils.FiberServer) {
	s.Post("/user", s.WithClient(handlers.CreateUserHandler))
	s.Get("/user/:id", s.WithClient(handlers.GetUserHandler))
	s.Post("/user/:id", s.WithClient(handlers.UpdateUserHandler))
	s.Post("/product", s.WithClient(handlers.CreateProductHandler))
	s.Get("/product/:id", s.WithClient(handlers.GetProductHandler))
	s.Post("/product/:id", s.WithClient(handlers.UpdateProductHandler))
}
