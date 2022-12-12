package main

import (
	"fmt"
	"server/utils"
	"testing"
)

func Test(t *testing.T) {
	iban := "  DE12 5001 0517 0648 4898 90  "

	err := utils.ParseIBAN(&iban)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(iban)
}
