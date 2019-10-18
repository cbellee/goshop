package app

func initRoutes(s *server) {
	s.CRUD("customer", s.customerController)
}
