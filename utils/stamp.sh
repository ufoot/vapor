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
VERSION_MAJOR=0
VERSION_MINOR=1

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

find_vpversion_go () {
    if [ -f vpbuild/vpversion.go ] ; then
	    VPVERSION_GO="vpbuild/vpversion.go"
	    if [ -f ${CONFIGURE_AC} ] ; then
            true
	    else
            echo "unable to open ${VPVERSION_GO}"
            exit 2
	    fi
    else
	    echo "unable to find vpbuild/vpversion.go"
	    exit 1
    fi
}

find_vppackage_go () {
    if [ -f vpbuild/vppackage.go ] ; then
	    VPPACKAGE_GO="vpbuild/vppackage.go"
	    if [ -f ${CONFIGURE_AC} ] ; then
            true
	    else
            echo "unable to open ${VPPACKAGE_GO}"
            exit 2
	    fi
    else
	    echo "unable to find vpbuild/vppackage.go"
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

calc_branch () {
    VERSION_BRANCH=$(git rev-parse --abbrev-ref HEAD)
    if [ "x${VERSION_BRANCH}" = "x" ] ; then
        VERSION_BRANCH=unknown
    fi
    if [ "x${VERSION_BRANCH}" = "xmaster" ] ; then
        VERSION_BRANCH=
    fi
}

calc_commits () {
    if [ "x${VERSION_BRANCH}" = "x" ] ; then
	SRC_GO=$(ls -d vp* | grep -v "vpbuild" | sort | tr "\n" " ")
	VERSION_COMMITS=$(git log --oneline --color=never -- ${SRC_GO} | wc -l)
    else
	VERSION_COMMITS=
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
        sed -i "s/const.*\/\/.*PackageTarname.*stamp.sh/const PackageTarname = \"${PACKAGE_TARNAME}\" \/\/ PackageTarname set by stamp.sh/g" ${VPPACKAGE_GO}
        sed -i "s/const.*\/\/.*PackageName.*stamp.sh/const PackageName = \"${PACKAGE_NAME}\" \/\/ PackageName set by stamp.sh/g" ${VPPACKAGE_GO}
        sed -i "s/const.*\/\/.*PackageEmail.*stamp.sh/const PackageEmail = \"${PACKAGE_EMAIL}\" \/\/ PackageEmail set by stamp.sh/g" ${VPPACKAGE_GO}
        sed -i "s/const.*\/\/.*PackageURL.*stamp.sh/const PackageURL = \"${PACKAGE_URL}\" \/\/ PackageURL set by stamp.sh/g" ${VPPACKAGE_GO}
        go vet ${VPPACKAGE_GO}
        go fmt ${VPPACKAGE_GO}
        echo "patching ${VPVERSION_GO} with version major=${VERSION_MAJOR} minor=${VERSION_MINOR} stamp=${VERSION_STAMP}"
        sed -i "s/const.*\/\/.*VersionMajor.*stamp.sh/const VersionMajor = ${VERSION_MAJOR} \/\/ VersionMajor set by stamp.sh/g" ${VPVERSION_GO}
        sed -i "s/const.*\/\/.*VersionMinor.*stamp.sh/const VersionMinor = ${VERSION_MINOR} \/\/ VersionMinor set by stamp.sh/g" ${VPVERSION_GO}
        sed -i "s/const.*\/\/.*VersionStamp.*stamp.sh/const VersionStamp = \"${VERSION_STAMP}\" \/\/ VersionStamp set by stamp.sh/g" ${VPVERSION_GO}
        go vet ${VPVERSION_GO}
        go fmt ${VPVERSION_GO}
        echo "patching ${CONFIGURE_AC} with version ${VERSION_DOT}"
        sed -i "s/^AC_INIT.*/AC_INIT([${PACKAGE_NAME}],[${VERSION_DOT}],[${PACKAGE_EMAIL}],[${PACKAGE_TARNAME}],[${PACKAGE_URL}])/g" ${CONFIGURE_AC}
    fi
}

find_configure_ac
find_vpversion_go
find_vppackage_go
git_check
git_changelog
calc_branch
calc_commits
do_patch

