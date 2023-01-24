package pkg2

import (
	"os"

	"github.com/willyrgf/errors-pattern-demo/pkg/errorx"
)

func doSomething() error {
	// something goes wrong
	_, err := os.OpenFile("foo.bar", os.O_WRONLY, 0600)
	if err != nil {
		return newPkg2Error(err)
	}
	return nil
}

func DoABunchOfThings() error {
	// run a lot of things here

	err := doSomething()
	if err != nil {
		// here a know that it's a standard error
		return errorx.NewSpecificError(err)
	}

	return nil
}
