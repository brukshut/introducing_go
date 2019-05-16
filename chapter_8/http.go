package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {

	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
		  <head>
		    <title>Hello, World</title>
                  </head>
                </html>
                  <body>
                    Hello World!
                  </body>
		</html>`,
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.Handle(
		"/src/",
		http.StripPrefix(
			"/src/",
			http.FileServer(http.Dir("/usr/local/src")),
		),
	)
	http.ListenAndServe(":9000", nil)
}
