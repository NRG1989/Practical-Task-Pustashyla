package main

import (
	"fmt"
	"testing"
)

var (
	args = []string{"exec arg0 arg1 arg2 arg3"}
)

func TestEcho1(t *testing.T) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		echo1(args)
	} else {
		fmt.Println("no args")
	}
}

func BenchmarkEcho1(b *testing.B) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			echo1(args)
		}
	} else {
		fmt.Println("no args")
	}
}

func BenchmarkEcho2(b *testing.B) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			echo2(args)
		}
	} else {
		fmt.Println("no args")
	}
}

func BenchmarkEcho3(b *testing.B) {
	fmt.Println("args:", args)
	if len(args) > 0 {
		for i := 0; i < b.N; i++ {
			echo3(args)
		}
	} else {
		fmt.Println("no args")
	}
}
