package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"woonbeaj/snippetbox/pkg/models/pgsql"
	_ "github.com/lib/pq"
)

type neuteredFileSystem struct {
    fs http.FileSystem
}

type application struct {
    errlog *log.Logger
    infolog *log.Logger
    orders *pgsql.OrderModel
    myCache map[string]*template.Template
}
    
func main() {
    addr := flag.String("addr", ":8080", "Сетевой адрес HTTP")
    dsn := flag.String("dsn", "user=postgres password=1 dbname=postgres sslmode=disable", "Данные для подключнеия к бд")
    flag.Parse()
    infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
    
    db, err := openDB(*dsn)
	if err != nil {
		errlog.Fatal(err)
	}
    defer db.Close()

    tmp, err := newCache("./ui/html/")
    if err != nil {
        errlog.Fatal(err)
    }

    app := application {
        errlog: errlog,
        infolog: infolog,
        orders: &pgsql.OrderModel{DB: db},
        myCache: tmp,
    }

    infolog.Println("Запуск веб-сервера на ", *addr)
    srv := &http.Server{
        Addr: *addr,
        ErrorLog: errlog,
        Handler: app.routes(),
    }
    err = srv.ListenAndServe()
    if err != nil {
        errlog.Fatal(err)
    }
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
    f, err := nfs.fs.Open(path)
    if err != nil {
        return nil, err
    }
    s, err := f.Stat()
    if s.IsDir() {
        index := filepath.Join(path, "index.html")
        if _, err := nfs.fs.Open(index); err != nil {
            if closeErr := f.Close(); closeErr != nil {
                return nil, closeErr
            }
            return nil, err
        }
    }
    return f, nil
}

func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}