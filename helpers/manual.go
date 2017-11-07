/*Package helpers containts helper functions used within project */
package helpers

import (
	"errors"
	"fmt"
	"os"

	logic "github.com/xDarkicex/Logic"
)

// vars for use of logic ops
var (
	Equal = logic.Eq
	And   = logic.And

	argOne bool
	argTwo bool
)

/*
Manual - Accepts two boolean values and
Prints out avalible flags for new project generation
Manual is an exclusive or function
Manual(true, false) prints output
Manual(true, true) doesn't print output
public function
*/
func Manual(a, b bool) error {
	if Equal(a, b) {
		argOne = (a == false)
		argTwo = (b == false)
		if And(argOne, argTwo) {
			return errors.New("can not set both -m and -man")
		}
	}
	fmt.Printf("getGoing: options require an argument\nUsage: [-new application name] [-template defines templates options: basic, mvc]\n")
	os.Exit(0)
	return nil
}
