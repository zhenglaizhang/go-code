package main


// Notes
//  * Primitive types (numeric, boolean, string) are immutable, they should be passed/returned by value

import (
	"fmt"
)

type user struct {
	name  string
	email string
}


// a method with a value receiver
// u is a copy of receiver
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}


// a method with a pointer receiver
// u is a copy of pointer receiver
func (u *user) updateEmail(email string) {
	u.email = email
}


// ## Reference types
//  * the set of slice, map, channel, interface, and function types
//  * header value
//  * string is also a reference type value

// a slice of bytes
// Very useful when we want to declare behavior around a built-in or reference type.
// The compiler only let's declare methods for user-defined types that are named
type IP []byte


func main() {
    // Value of type user
	bill := user{"Billy", "bill@email.com"}
	bill.notify()

    // Pointer of type user
	lisa := &user{"Lisa", "lisa@email.com"}
    // goc is doing: (*lisa).notify()
	lisa.notify()

	bill.updateEmail("bill@new.email.com")
    // goc is doing: (&bill).notify()
	bill.notify()

	lisa.updateEmail("lisa@new.email.com")
	lisa.notify()
}

