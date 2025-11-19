package controller

import "net/http"

type HTTPServer struct {
	Handlers HTTPHandler
}

func (s *HTTPServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("GET /question", s.Handlers.QuestionGetAll)
	router.HandleFunc("GET /question/{id}", s.Handlers.QuestionGetById)
	router.HandleFunc("POST /question", s.Handlers.QuestionCreate)
	router.HandleFunc("DELETE /question/{id}", s.Handlers.QuestionDelete)

	router.HandleFunc("GET /answer/{id}", s.Handlers.AnswerGetById)
	router.HandleFunc("POST /question/{id}/answer", s.Handlers.AnswerCreate)
	router.HandleFunc("DELETE /answer/{id}", s.Handlers.AnswerDelete)

	return http.ListenAndServe(":8080", router)
}
