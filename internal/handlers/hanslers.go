package handlers

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {
}

// Реализация интерфейса Serve
func (h *HelloHandler) Serve(w http.ResponseWriter, r *http.Request) error {
	// Запишите сообщение "Привет, мир!" в ответ
	fmt.Fprint(w, "Привет, мир!")

	// Возвратите nil, чтобы указать на отсутствие ошибок
	return nil
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}
