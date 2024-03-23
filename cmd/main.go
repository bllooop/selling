package main

import (
	"flag"
	"log"
	"os"
	selling "selling"
	"selling/pkg/handler"
	"selling/pkg/repository"
	"selling/pkg/service"

	_ "github.com/lib/pq"
)

func main() {
	port := "8000"
	addr := flag.String("addr", port, "web-server address")
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	dbpool, err := repository.NewPostgresDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	repos := repository.NewRepository(dbpool)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	infoLog.Printf("app is starting on %s port", *addr)
	srv := new(selling.Server)
	if err := srv.RunServer(port, handlers.InitRoutes()); err != nil {
		errorLog.Fatal(err)
	}

}
