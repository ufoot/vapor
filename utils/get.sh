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

get() {
    echo "go get $1"
    go get $1
}

get github.com/golang/lint/golint
get github.com/tools/godep
get golang.org/x/crypto/ripemd160
get golang.org/x/crypto/openpgp
get golang.org/x/crypto/openpgp/packet
get git.apache.org/thrift.git/lib/go/thrift
get github.com/llgcode/draw2d

rm -rf src/github.com/ufoot/vapor
install -d src/github.com/ufoot/vapor
for i in vp*; do ln -s $(pwd)/$i src/github.com/ufoot/vapor/$i ; done

