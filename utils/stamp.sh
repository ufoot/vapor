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
# Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
# Contact author: ufoot@ufoot.org

if [ -d ../utils ] ; then
    cd ..
fi
if [ ! -d utils ] ; then
    echo "$0 should be run in srcdir"
    exit 1
fi

PACKAGE_TARNAME="vapor"
PACKAGE_NAME="Vapor Toolkit"
PACKAGE_EMAIL="ufoot@ufoot.org"
PACKAGE_URL="https:\\/\\/github.com\\/ufoot\\/vapor"
PACKAGE_COPYRIGHT="Copyright (C)  2015, 2016  Christian Mauduit <ufoot@ufoot.org>"
PACKAGE_LICENSE="GNU GPL v3"

VERSION_MAJOR=0

find_configure_ac () {
    if [ -f configure.ac ] ; then
	    CONFIGURE_AC="configure.ac"
	    if [ -f ${CONFIGURE_AC} ] ; then
            true
	    else
            echo "unable to open ${CONFIGURE_AC}"
            exit 2
	    fi
    else
	    echo "unable to find configure.ac"
	    exit 1
    fi
}

find_version_go () {
    if [ -f $1/version.go ] ; then
	    VERSION_GO="$1/version.go"
	    if [ -f ${VERSION_GO} ] ; then
            true
	    else
            echo "unable to open ${VERSION_GO}"
            exit 2
	    fi
    else
	    echo "unable to find version.go"
	    exit 1
    fi
}

git_check () {
    if git status > /dev/null 2>&1 ; then
        true
    else
        echo "this is not a git repo, $0 won't do anything"
        exit 0
    fi
}

git_changelog () {
    if which git2cl > /dev/null ; then
        GIT_FILES=$(ls -d * | grep -v "ChangeLog" | sort | tr "\n" " ")
        git log --pretty --numstat --summary -- $GIT_FILES | git2cl > ChangeLog
    fi
}

calc_minor () {
    export VERSION_MINOR=$(git log --oneline --color=never -- $(find $1 -name "vp*.go" | grep -v version | sort) | wc -l)
    export VERSION_GO="$1/version.go"
}

calc_stamp () {
    export VERSION_STAMP=$(git log --oneline --color=never -- $(find $1 -name "vp*.go" | grep -v version | sort) | head -n 1 | awk '{print $1}')
    export VERSION_GO="$1/version.go"
}

patch_autotools () {
calc_minor .
calc_stamp .
    VERSION_DOT=${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_STAMP}
    if grep -q ${VERSION_DOT} ${CONFIGURE_AC} ; then
        echo "current version is ${VERSION_DOT}"
        touch ${CONFIGURE_AC}
    else
        echo "patching ${CONFIGURE_AC} with version ${VERSION_DOT}"
        sed -i "s/^AC_INIT.*/AC_INIT([${PACKAGE_NAME}],[${VERSION_DOT}],[${PACKAGE_EMAIL}],[${PACKAGE_TARNAME}],[${PACKAGE_URL}])/g" ${CONFIGURE_AC}
    fi
}

patch_go () {
	find_version_go $1
	calc_minor $1
	calc_stamp $1
        sed -i "s/const.*\/\/.*PackageTarname.*stamp.sh/const PackageTarname = \"${PACKAGE_TARNAME}\" \/\/ PackageTarname set by stamp.sh/g" ${VERSION_GO}
        sed -i "s/const.*\/\/.*PackageName.*stamp.sh/const PackageName = \"${PACKAGE_NAME}\" \/\/ PackageName set by stamp.sh/g" ${VERSION_GO}
        sed -i "s/const.*\/\/.*PackageEmail.*stamp.sh/const PackageEmail = \"${PACKAGE_EMAIL}\" \/\/ PackageEmail set by stamp.sh/g" ${VERSION_GO}
        sed -i "s/const.*\/\/.*PackageURL.*stamp.sh/const PackageURL = \"${PACKAGE_URL}\" \/\/ PackageURL set by stamp.sh/g" ${VERSION_GO}
        sed -i "s/const.*\/\/.*PackageCopyright.*stamp.sh/const PackageCopyright = \"${PACKAGE_COPYRIGHT}\" \/\/ PackageCopyright set by stamp.sh/g" ${VERSION_GO}
        sed -i "s/const.*\/\/.*VersionMajor.*stamp.sh/const VersionMajor = ${VERSION_MAJOR} \/\/ VersionMajor set by stamp.sh/g" ${VERSION_GO}
        sed -i "s/const.*\/\/.*VersionMinor.*stamp.sh/const VersionMinor = ${VERSION_MINOR} \/\/ VersionMinor set by stamp.sh/g" ${VERSION_GO}
        sed -i "s/const.*\/\/.*VersionStamp.*stamp.sh/const VersionStamp = \"${VERSION_STAMP}\" \/\/ VersionStamp set by stamp.sh/g" ${VERSION_GO}
        go vet ${VERSION_GO}
        go fmt ${VERSION_GO}
}

find_configure_ac
git_check
git_changelog
patch_autotools
for i in vp* ; do
	patch_go $i
done

