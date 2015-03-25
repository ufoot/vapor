Vapor is a toolkit designed to support Liquid War 7.

While truely dedicated to power the multiplayer game Liquid War 7, it
may be used to develop other games. The idea is to provide all the
components that are not directly linked to the Liquid War gameplay,
including, but not limited to, network tools, massive parallel
computation, game loops, fixed-point math arithmetics, etc.

Overview
========

Technology
----------

Vapor is a set of [golang](http://golang.org/) packages.

Install
-------

Packages can be installed from [github](http://github.com/) with
a standard `go get` command. For instance:

`go get github.com/ufoot/vapor/vpsys`

Alternatively, on UNIXish platforms, a `./configure` script sets
up some Makefiles which allow the usual `./configure && make` sequence
to work. Use `make help` to discover standard targets.

Project status
--------------

Experimental. Unstable. Under heavy work. Use at your own risks. Period.

News
====

* September 2014 : package creation, nothing yet, only ideas.
  Just an almost empty golang package.

* March 2015 : some basic functionnalities such as matrix computation
  are well advanced, code published on github.

License
=======

Copyright (C)  2015  Christian Mauduit <ufoot@ufoot.org>

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it wil/l be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
Contact author: ufoot@ufoot.org


