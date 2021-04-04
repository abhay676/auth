package controllers

func (s *Server) initializeRoutes()  {
//	 Heath check route
	s.Router.HandleFunc("/health", s.HeathCheck).Methods("GET")

//	SignUp route
	s.Router.HandleFunc("/create", s.CreateUser).Methods("POST")
}