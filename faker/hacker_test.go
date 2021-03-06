package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_HackerPhrase() {
	Global.Seed(20)
	fmt.Println(Global.HackerPhrase())
	// Output: Connecting the array won't do anything, we need to generate the haptic COM driver!
}

func BenchmarkHackerPhrase(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.HackerPhrase()
	}
}

func ExampleFaker_HackerAbbreviation() {
	Global.Seed(20)
	fmt.Println(Global.HackerAbbreviation())
	// Output: AGP
}

func BenchmarkHackerAbbreviation(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.HackerAbbreviation()
	}
}

func ExampleFaker_HackerAdjective() {
	Global.Seed(20)
	fmt.Println(Global.HackerAdjective())
	// Output: online
}

func BenchmarkHackerAdjective(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.HackerAdjective()
	}
}

func ExampleFaker_HackerNoun() {
	Global.Seed(20)
	fmt.Println(Global.HackerNoun())
	// Output: pixel
}

func BenchmarkHackerNoun(b *testing.B) {
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.HackerNoun()
	}
}

func ExampleFaker_HackerVerb() {
	Global.Seed(20)
	fmt.Println(Global.HackerVerb())
	// Output: connect
}

func BenchmarkHackerVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.HackerVerb()
	}
}

func ExampleFaker_HackerIngverb() {
	Global.Seed(20)
	fmt.Println(Global.HackerIngverb())
	// Output: navigating
}

func BenchmarkHackerIngverb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.HackerIngverb()
	}
}
