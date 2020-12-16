package main

import (
	"fmt"
	"net/url"

	"github.com/ory/dockertest"
)

func startDoc() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		fmt.Printf("Could not connect to docker: %s", err)
	}
	u, err := url.Parse(pool.Client.Endpoint())
	if err != nil {
		fmt.Printf("Could not parse the endpoint: %s", err)
	}
	resource, err := pool.Run("ebica-database", "latest", []string{
		"POSTGRES_PASSWORD=mysecretpassword",
		"POSTGRES_USER=postgres",
		"POSTGRES_DB=postgres",
	})
	host := u.Hostname()
	fmt.Println(u, resource, host)
}
