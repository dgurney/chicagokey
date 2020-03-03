// Package generation handles the generation of beta site IDs and keys.
package generation

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
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/md4"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Used for md4 hash generation
func getText(build string) (string, error) {
	switch {
	case build == "73f":
		return "Microsoft Chicago PDK Release, November 1993", nil
	case build == "73g":
		return "Microsoft Chicago PDK2 Release, December 1993", nil
	case build == "81":
		return "Chicago Preliminary PDK Release, January 1994", nil
	case build == "99":
		// up to 116
		return "Chicago Preliminary Beta 1 Release, May 1994", nil
	case build == "122":
		// up to 189
		return "Chicago Beta 1 Release, May 1994", nil
	case build == "216":
		// Up to 275
		return "Windows 95 Beta 2 Release, October 1994", nil
	}
	return "", fmt.Errorf("incorrect build")
}

func genPass(site, build string, p uint) (string, error) {
	if p > 65535 {
		return "", errors.New("provided password value cannot be higher than 65535")
	}
	pass := ""
	switch {
	default:
		pass = fmt.Sprintf("%04x", r.Intn(65535))
	case p != 0:
		pass = fmt.Sprintf("%04x", p)
	}
	_, err := strconv.Atoi(site)
	if err != nil {
		return "", err
	}
	// Generate the MD4 hash.
	hash := md4.New()
	text, err := getText(build)
	if err != nil {
		// The hardcoded string is an integral component of the hash. We cannot proceed if it's incorrect.
		return "", err
	}
	io.WriteString(hash, site+pass+text)
	lasth := hash.Sum(nil)

	// Reorder the last segment which just consists of the first 4 characters of the hash.
	last := fmt.Sprintf("%x%x", lasth[1:2], lasth[0:1])

	// Sum all ascii character codes together.
	middle := 0
	for i := range site {
		middle += int(site[i])
	}
	for i := range pass {
		middle += int(pass[i])
	}
	for i := range last {
		middle += int(last[i])
	}

	// Middle digit must be mod 9'd.
	return fmt.Sprintf("%s%d%s", pass, middle%9, last), nil
}

// GenerateCredentials generates a beta site id and password.
func GenerateCredentials(build string, site, password uint) (string, string, error) {
	if site > 999999 {
		return "", "", errors.New("site cannot be higher than 999999")
	}
	s := strconv.Itoa(int(site))
	switch {
	default:
		s = fmt.Sprintf("%06d", r.Intn(999999))
	case site != 0:
		s = fmt.Sprintf("%06d", site)
	}
	pass, err := genPass(s, build, password)
	if err != nil {
		if site < 1000000 {
			return "", "", err
		}
	}
	return s, pass, nil
}
