package main

import "testing"

func TestSayHiWithSomkiat(t *testing.T) {
	result := sayHi("somkiat")
	if result != "Hi somkiat" {
		t.Fatalf("Fail...")
	}
}

func BenchmarkSayHiWithSomkiat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sayHi("somkiat")
	}
}

func BenchmarkSayHiWithSomkiat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sayHi("somkiat")
	}
}
