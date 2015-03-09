package fizzbuzz

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestCall(t *testing.T) {
	result := Call(1)
	if result != "1" {
		t.Errorf("Expect 1 but call %s", result)
	}

	result = Call(2)
	if result != "2" {
		t.Errorf("Expect 2 but call %s", result)
	}

	result = Call(3)
	if result != "Fizz" {
		t.Errorf("Expect Fizz but call %s", result)
	}

	result = Call(5)
	if result != "Buzz" {
		t.Errorf("Expect Buzz but call %s", result)
	}

	result = Call(15)
	if result != "Fizz Buzz" {
		t.Errorf("Expect Fizz Buzz but call %s", result)
	}
}

func TestRange(t *testing.T) {
	r := Range(1, 15)
	expect := []string{
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"Fizz Buzz",
	}
	if !reflect.DeepEqual(r, expect) {
		t.Errorf("Expect %+v but %+v", expect, r)
	}
}

func ExamplePrintRange() {
	fmt.Println(strings.Join(Range(1, 15), ","))
	// Output:
	// 1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,Fizz Buzz
}
