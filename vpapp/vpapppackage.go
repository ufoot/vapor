// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015, 2016  Christian Mauduit <ufoot@ufoot.org>
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
	"github.com/ufoot/vapor/vpcommonapi"
	"strings"
)

// NewPackage creates a new package object.
func NewPackage(tarname, name, email, url, copyright, license string) *vpcommonapi.Package {
	return &vpcommonapi.Package{Tarname: tarname, Name: name, Email: email, URL: url, Copyright: copyright, License: license}
}

// DefaultPackage creates a new default object.
func DefaultPackage() *vpcommonapi.Package {
	return NewPackage(PackageTarname, PackageName, PackageEmail, PackageURL, PackageCopyright, PackageLicense)
}

// Compatible tells wether two packages refer to the same application,
// basically only the Tarname is checked, the rest is considered
// informative.
func Compatible(v1, v2 *vpcommonapi.Package) bool {
	return strings.EqualFold(v1.Tarname, v2.Tarname)
}

// PackageToString returns the package as a string
func PackageToString(p *vpcommonapi.Package) string {
	return fmt.Sprintf("%s (%s) email: %s, url:%s", p.Tarname, p.Name, p.Email, p.URL)
}
