package main

import (
	"fmt"
	"log"
	"net/http"
)

const content = `|  software engineer & designer
|
|  kamil#6027 886739001812332564
|  hi@kamil.codes
|  https://kamil.codes
|  https://github.com/k9mil
|  https://www.linkedin.com/in/kamilzak00
`

type App struct{}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, content)
}

func main() {
	app := &App{}
	http.Handle("/", app)

	port := 8080
	log.Printf("Listening on port %d", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
