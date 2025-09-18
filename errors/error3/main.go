package main

import "fmt"

// ! An interface is considered nil only if it isn't associated with any underlying nil value.

type MyError struct {
	A       int
	B       int
	Message string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("values %d and %d produced error %s", me.A, me.B, me.Message)
}

func reallyNil() error {
	var e error
	fmt.Println("e is nil:", e == nil)
	return e
}

func notReallyNil() error {
	var me *MyError
	fmt.Println("me is nil:", me == nil)
	return me
}

func main() {
	e := reallyNil()
	// The value of me will be a nil MyError
	me := notReallyNil()

	fmt.Println("in main, e is nil:", e == nil)
	// This will be false because the error interface is associated with a MyError type
	fmt.Println("in main, me is nil:", me == nil)
	fmt.Println(e)
	fmt.Println(me)
}
