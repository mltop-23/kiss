package main

import (
	"fmt"
	"kissandeat/internal/handlers"
	"kissandeat/internal/structs"
	"net/http"
)

func main() {
	user := structs.User{
		ID:        1,
		Username:  "john_doe",
		Password:  "password123",
		Email:     "johndoe@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Gender:    "male",
		Role:      "husband",
		FamilyID:  1,
	}
	fmt.Println(user)

	helloHandler := handlers.NewHelloHandler() // Вызов функции NewHelloHandler

	// Сопоставьте обработчик с маршрутом "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helloHandler.Serve(w, r) // Передайте параметры в метод Serve
	})
	// Запустите сервер на порту 8080
	fmt.Println("Сервер запущен на localhost:8080")
	http.ListenAndServe(":8080", nil)

}
