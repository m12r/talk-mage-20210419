package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

//go:embed presentation
var rootFS embed.FS

func main() {
	if err := run(rootFS); err != nil {
		fmt.Fprintf(os.Stderr, "http server: %v\n", err)
		os.Exit(1)
	}
}

func run(rootFS fs.FS) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	webFS, err := fs.Sub(rootFS, "presentation")
	if err != nil {
		return err
	}

	http.Handle("/", http.FileServer(http.FS(webFS)))

	hs := http.Server{}

	listener, listenerErr := net.Listen("tcp", "127.0.0.1:0")
	if listenerErr != nil {
		return listenerErr
	}

	addr := fmt.Sprintf("http://%s", listener.Addr())

	errC := make(chan error, 1)

	go func() {
		errC <- hs.Serve(listener)
	}()

	if err := exec.Command("xdg-open", addr).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "could not open browser: %v\n", err)
	}
	fmt.Printf("Serving at %s\n", addr)
	fmt.Println("Press ctrl+c to stop...")

	select {
	case startErr := <-errC:
		if startErr != nil && startErr != http.ErrServerClosed {
			return startErr
		}
	case <-ctx.Done():
		hs.SetKeepAlivesEnabled(false)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if shutdownErr := hs.Shutdown(shutdownCtx); shutdownErr != nil {
			fmt.Fprintf(os.Stderr, "http server: cannot shutdown gracefully: %v\n", shutdownErr)
		}
	}

	return nil
}
