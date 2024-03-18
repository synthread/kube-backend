package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const indexHtml = `<html>
	<head>
		<title>Synthread Labs - Not Found</title>
		<link rel="preconnect" href="https://fonts.googleapis.com">
		<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
		<link href="https://fonts.googleapis.com/css2?family=Josefin+Sans:ital,wght@0,400;1,700&display=swap" rel="stylesheet">
		<style type="text/css">
			html, body {
				margin: 0;
				padding: 0;
				height: 100%;
				background: #222;
				color: #eee;
				text-align: center;
			}

			div.col {
				display: flex;
				flex-direction: column;
				align-items: center;
				justify-content: center;
				height: 100%;
			}

			img {
				max-height: 8em;
				opacity: 0.85;
				filter: invert(1);
			}

			h1 {
				margin: 1em 0 0 0;
				font-size: 300%;
				font-family: "Josefin Sans", sans-serif;
				font-weight: 400;
				font-style: normal;
				font-optical-sizing: auto;
			}
		</style>
	</head>
	<body>

		<div class="col">
			<div>
				<img src="/static/logo.svg" alt="Synthread Labs">
				<h1>not found</h1>
			</div>
		</div>

	</body>
</html>`

func main() {
	addr := os.Getenv("LISTEN")
	if addr == "" {
		addr = "0.0.0.0:8080"
	}

	r := mux.NewRouter()

	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte(indexHtml))
	})

	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
