Vapor is a toolkit designed to support Liquid War 7.

While truely dedicated to power the multiplayer game Liquid War 7, it
may be used to develop other similar games. The idea is to provide all
the components that are not directly linked to the
[Liquid War](http://www.ufoot.org/liquidwar) gameplay, including, but not
limited to, network tools, massive parallel computation, game loops,
fixed-point math arithmetics, etc.  Design might evolve but as of
today, the most probable option is that this package would handle all
game logic while a GUI client, typically written in C++, would connect
to it.  A candidate name for this client is
[Fumes](https://github.com/ufoot/fumes).

Overview
========

Technology
----------

Vapor is a set of [golang](http://golang.org/) packages.

Install
-------

Packages can be installed from [github](https://github.com/ufoot/vapor)
with a standard `go get` command. For instance:

`go get github.com/ufoot/vapor/go/vpdemo`

Alternatively, on UNIXish platforms, a `./configure` script sets
up some Makefiles which allow a quite usual `./boostrap && ./configure && make`
sequence to work. 

Project is build automatically on 
[Travis](https://travis-ci.org/ufoot/vapor).

Documentation
-------------

Source code documentation is browsable online 
on [godoc.org](http://godoc.org/github.com/ufoot/vapor/go).

Project status
--------------

Experimental. Unstable. Under heavy work. Use at your own risks. Period.

[![Build Status](https://travis-ci.org/ufoot/vapor.svg?branch=master)](https://travis-ci.org/ufoot/vapor)

Why is it named Vapor?
----------------------

The main author of Liquid War struggles to find free time to hack arround,
so most of this is [vaporware](https://en.wikipedia.org/wiki/Vaporware).

News
====

* September 2014 : package creation, nothing yet, only ideas.
  Just an almost empty golang package.

* March 2015 : some basic functionnalities such as matrix computation
  are well advanced, code published on github.

Authors
=======

* Christian Mauduit <ufoot@ufoot.org> : main developper, project
  maintainer.

* All Liquid War 3, Liquid War 5 and Liquid War 6 contributors have
  indirectly helped for this project to exist. No real code snippets,
  but ideas where reused whenever possible.

License
=======

Copyright (C)  2015, 2016  Christian Mauduit <ufoot@ufoot.org>

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

Vapor homepage: https://github.com/ufoot/vapor
Contact author: ufoot@ufoot.org


