package main

import (
	"SimpleCMServer/config"
	"SimpleCMServer/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "[SimpleCMServer]", log.LstdFlags)

	serverArguments := config.ParseArguments(os.Args)

	root := handlers.NewRoot(l)

	mux := http.NewServeMux()
	mux.Handle("/", root)

	cmServer := &http.Server{
		Addr:         serverArguments.Port,
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := cmServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received signal, graceful shutdown. Signal:", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cmServer.Shutdown(tc)
}
