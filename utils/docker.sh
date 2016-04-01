#!/bin/bash

# Vapor is a toolkit designed to support Liquid War 7.
# Copyright (C)  2015, 2016  Christian Mauduit <ufoot@ufoot.org>
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

if [ -d ../utils ] && [ ! -d utils ] ; then
    cd ..
fi
if [ ! -d utils ] ; then
    echo "$0 should be run in srcdir"
    exit 1
fi
if [ ! -x utils/env.sh ] || [ ! -f docker/Dockerfile-build ] || [ ! -f docker/Dockerfile-vpdemo ] ; then
    echo "run ./configure to generate ./utils/env.sh"
    exit 1
fi

. utils/env.sh

make vapor-${PACKAGE_VERSION}.tar.gz && cp -f vapor-${PACKAGE_VERSION}.tar.gz docker/
for i in build vpdemo ; do
    cd docker && docker build -f Dockerfile-$i -t ufoot/vapor-$i:${PACKAGE_VERSION} . && cd ..
    cd docker && docker build -f Dockerfile-$i -t ufoot/vapor-$i:latest . && cd ..
done
rm -f docker/vapor-${PACKAGE_VERSION}.tar.gz

