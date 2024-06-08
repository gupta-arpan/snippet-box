package main

import(
	"net/http";
	"log";
	"flag";
	"os";
)

type application struct {
	infoLog *log.Logger;
	errorLog *log.Logger;
}

func main() {
	addr := flag.String("addr", ":4000", "Http network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		infoLog: infoLog,
		errorLog: errorLog,
	}	
	serv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}
	infoLog.Printf("Starting server on :%s", *addr)
	err := serv.ListenAndServe()
	errorLog.Fatal(err)
}