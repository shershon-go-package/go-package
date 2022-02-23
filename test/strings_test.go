package test

import (
	"fmt"
	"testing"
)

func TestStrCompare(t *testing.T) {
	a := "2"
	b := "1"
	if a > b {
		fmt.Println("ok")
	} else {
		fmt.Println("false")
	}
}
