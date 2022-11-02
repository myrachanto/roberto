package main

import (
	"log"

	"github.com/myrachanto/roberto/routes"
)

func init() {
	log.SetPrefix("tag microservice ")
}
func main() {
	// defer os.Exit(0)
	// cli := cmd.CommandLine{}
	// cli.Run()
	routes.ApiLoader()
}
