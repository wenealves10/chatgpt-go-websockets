package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/wenealves10/chatgpt-go-websockets/internal/handlers"
	"golang.org/x/net/websocket"
)

func main() {

	log.Println("Starting server on :4000")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tasksHandler := handlers.NewTalks(os.Getenv("OPENAI_API_KEY"))

	mux := http.NewServeMux()

	mux.Handle("/", websocket.Handler(tasksHandler.Handle))

	fileServer := http.FileServer(http.Dir("./public"))

	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal(err)
	}

}
