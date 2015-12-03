// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015  Christian Mauduit <ufoot@ufoot.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Vapor homepage: https://github.com/ufoot/vapor
// Contact author: ufoot@ufoot.org

package vpapp

import (
	"fmt"
	"github.com/ufoot/vapor/vpbuild"
	"strings"
)

// Package is used to identify the program package, think
// of it as old good old autotools PACKAGE (name on steroids).
type Package struct {
	// Tarname, for instance myapp
	Tarname string
	// Readable name, for instance My App
	Name string
	// Email of maintainer
	Email string
	// URL of program, to get info about it
	URL string
}

// NewPackage creates a new package object.
func NewPackage(tarname, name, email, url string) *Package {
	return &Package{tarname, name, email, url}
}

// BuildPackage creates a new default object.
func BuildPackage() *Package {
	return NewPackage(vpbuild.PackageTarname, vpbuild.PackageName, vpbuild.PackageEmail, vpbuild.PackageURL)
}

// Compatible tells wether two packages refer to the same application,
// basically only the Tarname is checked, the rest is considered
// informative.
func Compatible(v1, v2 *Package) bool {
	return strings.EqualFold(v1.Tarname, v2.Tarname)
}

// String returns the package as a string
func (p *Package) String() string {
	return fmt.Sprintf("%s (%s) email: %s, url:%s", p.Tarname, p.Name, p.Email, p.URL)
}
