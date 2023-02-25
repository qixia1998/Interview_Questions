package main

import (
	"fmt"

	"github.com/marmotedu/errors"
	"github.com/marmotedu/log"

	code "github.com/marmotedu/sample-code"
)

func main() {
	if err := getUser2(); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func getUser2() error {
	if err := queryDatabase2(); err != nil {
		return err
	}

	return nil
}

func queryDatabase2() error {
	opts := &log.Options{
		Level:            "info",
		Format:           "console",
		EnableColor:      true,
		EnableCaller:     true,
		OutputPaths:      []string{"test.log", "stdout"},
		ErrorOutputPaths: []string{},
	}

	log.Init(opts)
	defer log.Flush()

	err := errors.WithCode(code.ErrDatabase, "user 'qixia' not found.")
	if err != nil {
		log.Errorf("%v", err)
	}
	return err
}
