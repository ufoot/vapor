#!/bin/bash

# Vapor is a toolkit designed to support Liquid War 7.
# Copyright (C)  2015, 2016  Christian Mauduit <ufoot@ufoot.org>
#
# This program is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.
#
# Vapor homepage: https://github.com/ufoot/vapor
# Contact author: ufoot@ufoot.org

d=$(dirname $0)
cd $d
cp version.go.in version.go
if which sed > /dev/null ; then
    if which git > /dev/null ; then
	nbcommits=$(git log --oneline -- . | wc -l)
	shortref=$(git log --oneline -- . | head -n 1 | awk '{print $1}')
	echo "$0: nbcommits=$nbcommits shortref=$shortref"
	if [ x$nbcommits != x ] ; then
	    sed -i "s/VersionMinor = .*/VersionMinor = \"$nbcommits\" \/\/ VersionMinor set by version.sh/g" version.go
	fi
	if [ x$shortref != x ] ; then
	    sed -i "s/VersionStamp = .*/VersionStamp = \"$shortref\" \/\/ VersionStamp set by version.sh/g" version.go
	fi
    fi
fi
