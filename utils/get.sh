#!/bin/sh

# Vapor is a toolkit designed to support Liquid War 7.
# Copyright (C)  2015  Christian Mauduit <ufoot@ufoot.org>
#
# This program is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it wil/l be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.
#
# Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
# Contact author: ufoot@ufoot.org

if [ -d ../utils ] ; then
    cd ..
fi
if [ ! -d utils ] ; then
    echo "$0 should be run in srcdir"
    exit 1
fi

export GOPATH=$(pwd)

if [ -d /usr/share/gocode/src/golang.org/x/crypto/openpgp ] ; then
    true # Debian package golang-go.crypto-dev installed
else
    go get golang.org/x/crypto/ripemd160
    go get golang.org/x/crypto/openpgp
    go get golang.org/x/crypto/openpgp/packet
fi
go get github.com/golang/lint/golint
go get github.com/tools/godep
go get git.apache.org/thrift.git/lib/go/thrift
go get github.com/llgcode/draw2d

rm -rf src/github.com/ufoot/vapor
install -d src/github.com/ufoot/vapor
for i in vp*; do ln -s $(pwd)/$i src/github.com/ufoot/vapor/$i ; done

