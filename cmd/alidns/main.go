// Copyright (C) 2026 Joey Kot <joey.kot.x@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed WITHOUT ANY WARRANTY; without even the
// implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See <https://www.gnu.org/licenses/> for more details.

package main

import (
	"fmt"
	"os"

	"alidns/internal/cli"
)

func main() {
	if err := cli.Run(os.Args[1:], cli.NewDefaultDeps(os.Stdout, os.Stderr)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
