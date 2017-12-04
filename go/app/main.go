package main

import (
	"net/http"

	"github.com/shigwata/gae-go-sample/hello"
)

func init() {
	http.HandleFunc("/", hello.SayHello)
}
