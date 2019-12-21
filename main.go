package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/goroute/compress"

	"github.com/goroute/route"
	"github.com/sirupsen/logrus"
)

var (
	port = flag.String("port", "1234", "Listen port. Default 1234")
	dir  = flag.String("dir", ".", "Static files dir directory. Default is current directory.")
	help = flag.Bool("help", false, "Show help")
)

const usage = `
Usage: static-server <options>
Options:
	--port <value>				Listen port. Default "1234".
	--dir <value>				Static files dir directory. Default ".".
`

func main() {
	flag.Parse()
	if *help {
		fmt.Print(usage)
		os.Exit(0)
	}

	if _, err := os.Stat(*dir); err != nil {
		logrus.Fatalf("could not open directory: %v", err)
	}

	mux := route.NewServeMux()
	mux.Use(compress.New())
	mux.Static("/", *dir)

	logrus.Infof(`Static files server running on "http://localhost:%s"`, *port)
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), mux))
}
