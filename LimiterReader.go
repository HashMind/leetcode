package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func mainA() {
	r := strings.NewReader("some i\no.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)
	lr2 := io.LimitReader(r, 4)
	_, err := io.Copy(os.Stdout, lr)
	_, err2 := io.Copy(os.Stdout, lr2)
	if err != nil {
		log.Fatal(err)
	}
	if err2 != nil {
		log.Fatal(err)
	}
}
