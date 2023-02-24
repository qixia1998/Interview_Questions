package main

import (
	"github.com/pkg/errors"
	"log"
)

func main() {
	if err := funcAA(); err != nil {
		log.Fatalf("call func got failed: %v", err)
		return
	}

	log.Println("call func success")
}

func funcAA() error {
	if err := funcBB(); err != nil {
		return errors.Wrap(err, "call funcB failed")
	}

	return errors.New("func called error")
}

func funcBB() error {
	return errors.New("func called error")
}
