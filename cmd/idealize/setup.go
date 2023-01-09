//go:build !prod

package main

import (
	"log"

	"github.com/haleyrc/vorpal"
)

func init() {
	log.Println("== RUNNING IN DEVELOPMENT MODE")
}

func newRouter() *vorpal.Router {
	r := vorpal.NewRouter(nil)
	r.EnableDevMode()
	return r
}
