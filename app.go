package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"tiny-server-go/Application"
)

func main() {

	fmt.Println("Redirect service started...")

	Application.InitEnv()

	Schema := Application.NewSchema()
	ServerHostAddr := Application.GetEnv("HOST_ADDR", ":80")

	rLen := len(Schema.Redirects)

	fmt.Println("Found redirects: " + strconv.Itoa(rLen))

	for i := 0; i < rLen; i++ {
		fmt.Println("Redirect from: " + Schema.Redirects[i].From)
		http.HandleFunc(Schema.Redirects[i].From, RedirectResponse(Schema.Redirects[i].To))
	}
	http.HandleFunc("/", IndexStub)

	fmt.Println("Listen at " + ServerHostAddr)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func RedirectResponse(to string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Redirect to: " + to)
		if strings.HasPrefix(to, "file://") {
			filename := to[7:]
			info, err := os.Stat(filename)
			if !os.IsNotExist(err) && !info.IsDir() {
				http.ServeFile(w, r, filename)
			} else {
				path, _ := os.Getwd()
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("File not found in path: " + path + "/" + filename))
			}
		} else {
			http.Redirect(w, r, to, 301)
		}
	}
}

func IndexStub(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Redirect service. Author: Mikhail Levi. (c) 2020. ")
	if err != nil {
		fmt.Println(err)
	}
}
