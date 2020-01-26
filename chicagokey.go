package main

import (
	"flag"
	"fmt"

	"github.com/dgurney/chicagokey/generation"
)

// Used if chicagokey is not built using the makefile.
const version = "0.0.1"

// git describe --tags --dirty
var gitVersion string

func getVersion() string {
	if len(gitVersion) == 0 {
		return version
	}
	return gitVersion
}

func main() {
	b := flag.String("b", "", "Build to generate the site id and password for. Valid options: 73f, 73g, 81, 99 (for builds up to 116), 122 (for builds up to 189), 216 (for builds up to 275)")
	r := flag.Int("r", 1, "Repeat n times.")
	v := flag.Bool("v", false, "Print version information and exit.")
	flag.Parse()

	if *v {
		fmt.Printf("Chicagokey v%s by Daniel Gurney.\n", getVersion())
		return
	}

	build := ""
	switch {
	case *b == "73f" || *b == "73g" || *b == "81" || *b == "99" || *b == "122" || *b == "216":
		build = *b
	default:
		fmt.Println("Invalid build or no build specified! Usage:")
		flag.PrintDefaults()
		return
	}
	if *r <= 0 {
		*r = 1
	}

	for i := 0; i < *r; i++ {
		site, pass := generation.GenerateCredentials(build)
		fmt.Printf("Site ID: %s\nPassword: %s\n", site, pass)
	}
}
