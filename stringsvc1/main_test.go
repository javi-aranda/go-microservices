package main

import "testing"

func TestUppercase(t *testing.T) {
	svc := stringService{}
	v := "Hello, World"
	got, error := svc.Uppercase(v)
	want := "HELLO, WORLD"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	if error != nil {
		t.Errorf(error.Error())
	}
}

func TestCount(t *testing.T) {
	svc := stringService{}
	v := "Hello, World"
	got := svc.Count(v)
	want := 12

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestReverse(t *testing.T) {
	svc := stringService{}
	v := "Hello, World"
	got, error := svc.Reverse(v)
	want := "dlroW ,olleH"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	if error != nil {
		t.Errorf(error.Error())
	}
}
