package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/fatih/color"
)

const address = ":8382"

type ServeOpts struct {
	outpath   string
	keepAlive bool
}

func serve(opts *ServeOpts) {
	server := &http.Server{Addr: address}
	idleConnections := make(chan struct{})

	hasLog := LogExists(opts.outpath)
	hasResult := ResultExists(opts.outpath)

	if hasResult {
		http.HandleFunc("/result", func(resp http.ResponseWriter, req *http.Request) {
			handleResult(resp, req, opts.outpath)
			if !opts.keepAlive {
				shutdown()
			}
		})
	}

	if hasLog {
		http.HandleFunc("/log", func(resp http.ResponseWriter, req *http.Request) {
			handleLog(resp, req, opts.outpath)
			if !opts.keepAlive {
				shutdown()
			}
		})
	}

	title("--- serve results ---")
	fmt.Printf("starting server at %v\n", address)
	fmt.Printf("keep-alive: %v\n", opts.keepAlive)

	go interuptShutdown(server, idleConnections)
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndSever Error: %v", err)
		}

		fmt.Println("successfully served results, shutting down...")
	}()

	fmt.Println()

	arrow := color.New(color.FgGreen, color.Bold).SprintFunc()
	url := color.New(color.FgBlue, color.Bold).SprintFunc()
	if hasResult {
		fmt.Println(arrow("➜  Results:"), url("https://srsim.app/local"))
	}

	if hasLog {
		fmt.Println(arrow("➜  Logs:   "), url("https://srsim.app/debug"))
	}

	fmt.Println()

	<-idleConnections
}

func interuptShutdown(server *http.Server, connections chan struct{}) {
	defer close(connections)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Printf("HTTP server shutdown error: %v\n", err)
	}
}

func shutdown() {
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Signal(os.Interrupt); err != nil {
		log.Fatal(err)
	}
}

func handleResult(resp http.ResponseWriter, req *http.Request, path string) {
	if !prepare(resp, req) {
		return
	}

	compressed, err := os.ReadFile(ResultFile(path))
	if err != nil {
		log.Printf("error reading gz data: %v\n", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	flush(resp, compressed)
}

func handleLog(resp http.ResponseWriter, req *http.Request, path string) {
	if !prepare(resp, req) {
		return
	}

	compressed, err := os.ReadFile(LogFile(path))
	if err != nil {
		log.Printf("error reading gz data: %v\n", err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	flush(resp, compressed)
}

func prepare(resp http.ResponseWriter, req *http.Request) bool {
	if req.Method == http.MethodOptions {
		optionsResponse(resp)
		return false
	}

	if req.Method != http.MethodGet {
		resp.WriteHeader(http.StatusForbidden)
		return false
	}

	return true
}

func flush(resp http.ResponseWriter, compressed []byte) {
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Encoding", "gzip")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.WriteHeader(http.StatusOK)
	resp.Write(compressed)

	if f, ok := resp.(http.Flusher); ok {
		f.Flush()
	}
}

func optionsResponse(resp http.ResponseWriter) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	resp.Header().Set(
		"Access-Control-Allow-Headers",
		"Accept, Access-Control-Allow-Origin, Content-Type, "+
			"Content-Length, Accept-Encoding, Authorization")
	resp.WriteHeader(http.StatusNoContent)
}
