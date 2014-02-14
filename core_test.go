package core

import (
	"log"
	"testing"
)

func TestFunct(t *testing.T) {
	res := TestFunc()
	log.Println(res)
	//t.Errorf("bang")
}