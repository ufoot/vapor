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

if [ -d ../utils ] && [ ! -d utils ] ; then
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

find_version_go_in () {
    if [ -f $1/version.go.in ] ; then
	VERSION_GO_IN="$1/version.go.in"
	if [ -f ${VERSION_GO_IN} ] ; then
	    true
	else
            echo "unable to open ${VERSION_GO_IN}"
            exit 2
	fi
    else
	echo "unable to find version.go in $1 ($(pwd))"
	exit 1
    fi
}

git_check () {
    if [ -d .git ] && git status > /dev/null 2>&1 ; then
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
    export VERSION_MINOR=$(git log --oneline --color=never -- $1 | wc -l)
}

calc_stamp () {
    export VERSION_STAMP=$(git log --oneline --color=never -- $1 | head -n 1 | awk '{print $1}')
}

patch_autotools () {
    calc_minor go
    calc_stamp go
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
    find_version_go_in $1
    calc_minor $1
    calc_stamp $1
    sed -i "s/const.*\/\/.*PackageTarname.*sh/const PackageTarname = \"${PACKAGE_TARNAME}\" \/\/ PackageTarname set by version.sh/g" ${VERSION_GO_IN}
    sed -i "s/const.*\/\/.*PackageName.*sh/const PackageName = \"${PACKAGE_NAME}\" \/\/ PackageName set by version.sh/g" ${VERSION_GO_IN}
    sed -i "s/const.*\/\/.*PackageEmail.*sh/const PackageEmail = \"${PACKAGE_EMAIL}\" \/\/ PackageEmail set by version.sh/g" ${VERSION_GO_IN}
    sed -i "s/const.*\/\/.*PackageURL.*sh/const PackageURL = \"${PACKAGE_URL}\" \/\/ PackageURL set by version.sh/g" ${VERSION_GO_IN}
    sed -i "s/const.*\/\/.*PackageCopyright.*sh/const PackageCopyright = \"${PACKAGE_COPYRIGHT}\" \/\/ PackageCopyright set by version.sh/g" ${VERSION_GO_IN}
    sed -i "s/const.*\/\/.*PackageLicense.*sh/const PackageLicense = \"${PACKAGE_LICENSE}\" \/\/ PackageLicense set by version.sh/g" ${VERSION_GO_IN}
    sed -i "s/const.*\/\/.*VersionMajor.*sh/const VersionMajor = ${VERSION_MAJOR} \/\/ VersionMajor set by version.sh/g" ${VERSION_GO_IN}
    sed -i "s/const.*\/\/.*VersionMinor.*sh/const VersionMinor = 0 \/\/ VersionMinor set by version.sh/g" ${VERSION_GO_IN}
    sed -i "s/const.*\/\/.*VersionStamp.*sh/const VersionStamp = \"0000000\" \/\/ VersionStamp set by version.sh/g" ${VERSION_GO_IN}
    #go vet ${VERSION_GO_IN}
    #go fmt ${VERSION_GO_IN}
}

find_configure_ac
git_check
git_changelog
patch_autotools
for i in go/vp* ; do
    patch_go $i
done

