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

package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/ufoot/vapor/go/vplog"
	"github.com/ufoot/vapor/go/vploop"
)

var showHelp=false
var showPackage=false
var showVersion=false

func usage() {
	if len(os.Args)>0 {
		fmt.Printf("usage: %s <options>\n\n",os.Args[0])
	}
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&showHelp, "help", false, "show usage information")
	flag.BoolVar(&showVersion, "version", false, "show version information")
	flag.BoolVar(&showPackage, "package", false, "show package information")
}

func main() {
	flag.Parse()

	if showHelp {
		usage()
		return
	}
	if showVersion {
		fmt.Printf("%d.%d.%s\n",VersionMajor,VersionMinor,VersionStamp)
		return
	}
	if showPackage {
		fmt.Printf("%s-%d.%d.%s (%s %d.%d.%s)\n",PackageTarname,VersionMajor,VersionMinor,VersionStamp,PackageName,VersionMajor,VersionMinor,VersionStamp)
		fmt.Printf("%s - %s\n",PackageURL,PackageEmail)
		fmt.Printf("%s (%s)\n",PackageCopyright,PackageLicense)
		return
	}

	var state NibblesState
	var server NibblesServer

	vplog.LogInit("vpdemo")
	vploop.MainLoop(state, server)

	return
}
