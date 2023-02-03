package server

import (
	"SimpleCMServer/cmshandlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type CmsServer struct {
	logger   *log.Logger
	handlers []*cmshandlers.Handler
	server   *http.Server
}

func NewServer(p string, l *log.Logger, h []*cmshandlers.Handler) *CmsServer {
	server := &http.Server{
		Addr:         p,
		Handler:      nil,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	return &CmsServer{l, h, server}
}

func (r *CmsServer) GoServer() {
	mux := http.NewServeMux()
	for _, handler := range r.handlers {
		mux.Handle(handler.Pattern, handler.PatternHandler)
	}
	r.server.Handler = mux

	go func() {
		err := r.server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	r.logger.Println("Received signal, graceful shutdown. Signal:", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := r.server.Shutdown(tc)
	if err != nil {
		return
	}
}
