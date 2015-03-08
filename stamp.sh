#!/bin/sh

# Vapor is a toolkit designed to support Liquid War 7.
# Copyright (C)  2015  Christian Mauduit <ufoot@ufoot.org>
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

PACKAGE_TARNAME="vapor"
PACKAGE_NAME="Vapor Toolkit"
PACKAGE_EMAIL="ufoot@ufoot.org"
PACKAGE_URL="http:\\/\\/www.ufoot.org\\/liquidwar\\/v7\\/vapor"
VERSION_MAJOR=0
VERSION_MINOR=1

usage () {
    echo "usage:"
    echo "        ./stamp.sh"
}

find_configure_ac () {
    if [ -f configure.ac ] ; then
	    CONFIGURE_AC="$(readlink -f configure.ac || readlink configure.ac)"
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

find_vpversion_go () {
    if [ -f src/ufoot.org/vapor/vpbuild/vpversion.go ] ; then
	    VPVERSION_GO="$(readlink -f src/ufoot.org/vapor/vpbuild/vpversion.go || readlink src/ufoot.org/vpbuild/vpversion.go)"
	    if [ -f ${CONFIGURE_AC} ] ; then
            true
	    else
            echo "unable to open ${VPVERSION_GO}"
            exit 2
	    fi
    else
	    echo "unable to find src/ufoot.org/vapor/vpbuild/vpversion.go"
	    exit 1
    fi
}

find_vppackage_go () {
    if [ -f src/ufoot.org/vapor/vpbuild/vppackage.go ] ; then
	    VPPACKAGE_GO="$(readlink -f src/ufoot.org/vapor/vpbuild/vppackage.go || readlink src/ufoot.org/vapor/vpbuild/vppackage.go)"
	    if [ -f ${CONFIGURE_AC} ] ; then
            true
	    else
            echo "unable to open ${VPPACKAGE_GO}"
            exit 2
	    fi
    else
	    echo "unable to find src/ufoot.org/vapor/vpbuild/vppackage.go"
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

calc_commits () {
    SRC_GO=$(ls -d src/ufoot.org/vapor/vp* | grep -v "vapor/vpbuild" | sort | tr "\n" " ")
    VERSION_COMMITS=$(git log --oneline --color=never -- ${SRC_GO} | wc -l)
}

calc_branch () {
    VERSION_BRANCH=$(git branch --color=never | grep "* " | cut -c 3- | sed "s/ //g")
    if [ "x${VERSION_BRANCH}" = "x" ] ; then
        VERSION_BRANCH=unknown
    fi
    if [ "x${VERSION_BRANCH}" = "xmaster" ] ; then
        VERSION_BRANCH=
    fi
}

do_patch () {
    VERSION_STAMP=${VERSION_COMMITS}${VERSION_BRANCH}
    VERSION_DOT=${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_STAMP}
    if grep -q ${VERSION_DOT} ${CONFIGURE_AC} ; then
        echo "current version is ${VERSION_DOT}"
        touch ${CONFIGURE_AC}
    else
        echo "patching ${VPPACKAGE_GO} with package tarname=${PACKAGE_TARNAME} name=${PACKAGE_NAME} email=${PACKAGE_EMAIL} url=${PACKAGE_URL}"
        sed -i "s/const.*\/\/.*PACKAGE_TARNAME.*stamp.sh/const PACKAGE_TARNAME = \"${PACKAGE_TARNAME}\" \/\/ PACKAGE_TARNAME set by stamp.sh/g" ${VPPACKAGE_GO}
        sed -i "s/const.*\/\/.*PACKAGE_NAME.*stamp.sh/const PACKAGE_NAME = \"${PACKAGE_NAME}\" \/\/ PACKAGE_NAME set by stamp.sh/g" ${VPPACKAGE_GO}
        sed -i "s/const.*\/\/.*PACKAGE_EMAIL.*stamp.sh/const PACKAGE_EMAIL = \"${PACKAGE_EMAIL}\" \/\/ PACKAGE_EMAIL set by stamp.sh/g" ${VPPACKAGE_GO}
        sed -i "s/const.*\/\/.*PACKAGE_URL.*stamp.sh/const PACKAGE_URL = \"${PACKAGE_URL}\" \/\/ PACKAGE_URL set by stamp.sh/g" ${VPPACKAGE_GO}
        go vet ${VPPACKAGE_GO}
        go fmt ${VPPACKAGE_GO}
        echo "patching ${VPVERSION_GO} with version major=${VERSION_MAJOR} minor=${VERSION_MINOR} stamp=${VERSION_STAMP}"
        sed -i "s/const.*\/\/.*VERSION_MAJOR.*stamp.sh/const VERSION_MAJOR = ${VERSION_MAJOR} \/\/ VERSION_MAJOR set by stamp.sh/g" ${VPVERSION_GO}
        sed -i "s/const.*\/\/.*VERSION_MINOR.*stamp.sh/const VERSION_MINOR = ${VERSION_MINOR} \/\/ VERSION_MINOR set by stamp.sh/g" ${VPVERSION_GO}
        sed -i "s/const.*\/\/.*VERSION_STAMP.*stamp.sh/const VERSION_STAMP = \"${VERSION_STAMP}\" \/\/ VERSION_STAMP set by stamp.sh/g" ${VPVERSION_GO}
        go vet ${VPVERSION_GO}
        go fmt ${VPVERSION_GO}
        echo "patching ${CONFIGURE_AC} with version ${VERSION_DOT}"
        sed -i "s/^AC_INIT.*/AC_INIT([${PACKAGE_NAME}],[${VERSION_DOT}],[${PACKAGE_EMAIL}],[${PACKAGE_TARNAME}],[${PACKAGE_URL}])/g" ${CONFIGURE_AC}
    fi
}

if [ x"$1" = "x" ] ; then
    find_configure_ac
    find_vpversion_go
    find_vppackage_go
    git_check
    git_changelog
    calc_commits
    calc_branch
    do_patch
else
    usage
fi

