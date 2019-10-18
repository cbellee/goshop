package app

func initRoutes(s *server) {
	s.CRUD("order", s.orderController)
}
