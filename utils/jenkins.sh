#!/bin/bash

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

# setting up Jenkins env if needed, should be set by caller
if test x$WORKSPACE = x ; then
    if test x$TMP = x ; then
	if test x$TMPDIR = x ; then
	    WORKSPACE=/tmp
	else
	    WORKSPACE=$TMPDIR
	fi
    else
	WORKSPACE=$TMP
    fi
fi

# run various typicall build commands

echo "******** $0 $(date) ********"
#rm -rf src/git.apache.org src/github.com src/golang.org
rm -rf test doc/txt doc/html doc/cover
git clean -d -f -x
rm -rf $HOME/.vapor

echo "******** $0 $(date) ********"
if ./bootstrap ; then
    echo "./bootstrap OK"
else
    echo "./bootstrap failed"
    exit 1
fi

echo "******** $0 $(date) ********"
if ./configure --prefix=$WORKSPACE/local ; then
    echo "./configure OK"
else
    echo "./configure failed"
    exit 2
fi

echo "******** $0 $(date) ********"
if make ; then
    echo "make OK"
else
    echo "make failed"
    exit 3
fi

echo "******** $0 $(date) ********"
if make devel ; then
    echo "make devel OK"
else
    echo "make devel failed"
    exit 4
fi

echo "******** $0 $(date) ********"
if make check ; then
    echo "make check OK"
else
    echo "make check failed"
    exit 5
fi

echo "******** $0 $(date) ********"
if make bench ; then
    echo "make bench OK"
else
    echo "make bench failed"
    exit 6
fi

echo "******** $0 $(date) ********"
if make doc ; then
    echo "make doc OK"
else
    echo "make doc failed"
    exit 7
fi

