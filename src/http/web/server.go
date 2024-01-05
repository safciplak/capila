package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	writeTimeout = 60
	readTimeout  = 15
	idleTimeout  = 60
)

/*
Server starts a local webserver.
*/
func Server(port string, router http.Handler) {
	server := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Second * writeTimeout,
		ReadTimeout:  time.Second * readTimeout,
		IdleTimeout:  time.Second * idleTimeout,
		Handler:      router,
	}

	println("webserver started at port", port)

	// for a friendly shutdown
	done := make(chan bool)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("\nServer is shutting down...")
		StopServer(server)

		close(done)
	}()

	startWebServer(server)

	<-done
	log.Println("server stopped gracefully")
	os.Exit(0)
}

/*
startWebServer starts the webserver as go-routine
*/
func startWebServer(srv *http.Server) {
	// returns ErrServerClosed on graceful close
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Println("startWebServer error:")
		log.Fatalln(err)
	}
}

/*
StopServer stops the local web server.
*/
func StopServer(srv *http.Server) {
	// wait max. a minute before timing out
	locContext, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)

	if err := srv.Shutdown(locContext); err != nil {
		log.Println("StopServer: Could not gracefully shutdown the server:")
		log.Fatalln(err)
	}

	if err := srv.Shutdown(context.TODO()); err != nil {
		log.Println("StopServer:")
		log.Fatalln(err)
	}

	println("webserver stopped")
}
