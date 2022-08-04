package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr      string
	StaticDir string
}

func main() {

	cfg := new(Config)

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir(cfg.StaticDir))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	//err := http.ListenAndServe(cfg.Addr, mux)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
