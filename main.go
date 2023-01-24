package main

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/willyrgf/errors-pattern-demo/pkg/errorx"
	"github.com/willyrgf/errors-pattern-demo/pkg/pkg1"
	"github.com/willyrgf/errors-pattern-demo/pkg/pkg2"
)

func main() {
	err := pkg1.DoABunchOfThings()
	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", errors.Unwrap(err))
	fmt.Printf("%+v\n", errors.Is(err, errorx.ErrStandard))
	fmt.Printf("%+v\n", errors.Is(err, pkg1.ErrPkg1))
	fmt.Printf("%+v\n", errors.Is(err, fs.ErrNotExist))

	var errStandard *errorx.StandardError
	if errors.As(err, &errStandard) {
		fmt.Printf("os.Exit(): %+v\n", errStandard.ExitCode())
	}

	fmt.Println()

	err2 := pkg2.DoABunchOfThings()
	fmt.Printf("%+v\n", err2)
	fmt.Printf("%+v\n", errors.Unwrap(err2))
	fmt.Printf("%+v\n", errors.Is(err2, errorx.ErrSpecific))
	fmt.Printf("%+v\n", errors.Is(err2, pkg2.ErrPkg2))
	fmt.Printf("%+v\n", errors.Is(err2, fs.ErrNotExist))

	var errSpecific *errorx.SpecificError
	if errors.As(err2, &errSpecific) {
		fmt.Printf("os.Exit(): %+v\n", errSpecific.ExitCode())
	}
}
