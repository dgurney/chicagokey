package main

/*
   Copyright (C) 2020 Daniel Gurney
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"flag"
	"fmt"

	"github.com/dgurney/chicagokey/pkg/generator"
)

// Used if chicagokey is not built using the makefile.
const version = "0.3.2"

// git describe --tags --dirty
var gitVersion string

func getVersion() string {
	if len(gitVersion) == 0 {
		return version
	}
	return gitVersion
}

func main() {
	b := flag.String("b", "", "Build to generate the site id and password for. Valid Chicago builds: 73f, 73g, 81 (up to 90c), 99 (for builds up to 116), 122 (for builds up to 189), 216 (for builds up to 302). Valid IE builds: ie4july (4.70.1169), ie4sept (4.71.0225)")
	s := flag.Uint("s", 0, "Use a custom site, can be up to 6 digits long.")
	p := flag.Uint("p", 0, "Use a constant integer for the first 4 characters of the password. Can be any positive number up to 65535.")
	r := flag.Int("r", 1, "Repeat n times.")
	v := flag.Bool("v", false, "Print version information and exit.")
	flag.Parse()

	if *v {
		fmt.Printf("Chicagokey v%s by Daniel Gurney. Licensed under GPLv3.\n", getVersion())
		return
	}

	build := ""
	switch {
	case *b == "73f" || *b == "73g" || *b == "81" || *b == "99" || *b == "122" || *b == "216" || *b == "ie4july" || *b == "ie4sept":
		build = *b
	default:
		fmt.Println("Invalid build or no build specified! Usage:")
		flag.PrintDefaults()
		return
	}
	if *r <= 0 {
		*r = 1
	}
	if *p > 65535 {
		*p = 0
		fmt.Println("password cannot be more than 65535, ignoring...")
	}
	if *s > 999999 {
		*s = 0
		fmt.Println("Site cannot be more than 999999, ignoring...")
	}
	for i := 0; i < *r; i++ {
		site, pass, err := generator.GenerateCredentials(build, *s, *p)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Site ID: %s\nPassword: %s\n", site, pass)
	}
}
