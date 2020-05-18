package main

import (
	"io"
	"net/http"
	"io/ioutil"
)

func startWeb() {
	webLogger.Info("Starting Web Server!")
	http.HandleFunc("/", servePage)
	http.ListenAndServe(":8080", nil)
}

func servePage(writer http.ResponseWriter, request *http.Request) {
	content, err := ioutil.ReadFile("web/index.html")
    if err != nil {
        webLogger.Fatal(err)
    }

    // Convert []byte to string and print to screen
    html := string(content)
	io.WriteString(writer, html)
}
