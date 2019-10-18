package app

func initRoutes(s *server) {
	s.CRUD("product", s.productController)
}
