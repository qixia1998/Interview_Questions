package main

import (
	"fmt"
	"github.com/marmotedu/errors"
	code "github.com/marmotedu/sample-code"
)

func main() {
	if err := getUser1(); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func getUser1() error {
	if err := queryDatabase1(); err != nil {
		return errors.Wrap(err, "get user failed.")
	}

	return nil
}

func queryDatabase1() error {
	return errors.WithCode(code.ErrDatabase, "user 'Lingfei Kong' not found.")
}
