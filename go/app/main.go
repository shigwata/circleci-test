package main

import (
	"net/http"

	"github.com/shigwata/circleci-test/go/hello"
)

func init() {
	http.HandleFunc("/", hello.SayHello)
}
