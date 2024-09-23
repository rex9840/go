package test

import (
	"go_modules/helloworld"
	"testing"
)

func Test_for_helloworld(T *testing.T) {
	want := "hellow world"
	if got := helloworld.Hello(); got != want {
		T.Errorf("Hello() = %q, want %q", got, want)
	}

}
