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

for m in Makefile.dep Makefile.help ; do
    rm -f $m && touch $m
done

echo >> Makefile.help
echo "Below are listed global Makefile targets:" >> Makefile.help
echo "vp: builds all binaries (alias for all)." >> Makefile.help
echo "check: runs test suite." >> Makefile.help
echo "bench: runs test suite in bench mode." >> Makefile.help
echo "lint: lints (cleanup & check) source code." >> Makefile.help
echo "devel: developers target, indents, lints and tests code in verbose mode." >> Makefile.help
echo "doc: generate documentation." >> Makefile.help
echo "clean: cleans up build." >> Makefile.help
echo "distclean: cleans up build, removing even more files." >> Makefile.help
echo "get: get go libraries from github.com or golang.org." >> Makefile.help
echo "dep: generate dependencies."  >> Makefile.help
echo "help: display help about Makefile targets." >> Makefile.help
echo "stamp: update build stamp." >> Makefile.help
echo "indent: automatically indent code." >> Makefile.help
echo >> Makefile.help
echo "Below are listed per-package Makefile targets:" >> Makefile.help

for t in "" "check-" "bench-" "lint-" "devel-" "doc-" ; do
    echo ".PHONY: ${t}vp" >> Makefile.dep
    echo -n "${t}vp:" >> Makefile.dep
    for i in vp* ; do
        echo -n " ${t}${i}" >> Makefile.dep
	echo -n "${t}${i} " >> Makefile.help 
    done
    echo >> Makefile.dep
    echo >> Makefile.dep
    echo >> Makefile.help
done

echo "configure.ac: $(ls vp*/*.go | sort -u | tr "\n" " ")" >> Makefile.dep
echo "\tcd \$(VP_TOPSRCDIR) && ./dep.sh && ./stamp.sh" >> Makefile.dep
echo >> Makefile.dep

for i in $(ls -d vp* | sort -u | tr "\n" " ") ; do
    j="github.com/ufoot/vapor/$i"
    k=$(ls $i/*.go | sort -u | tr "\n" " ")
    echo ".PHONY: $i" >> Makefile.dep
    echo "$i: configure.ac $k" >> Makefile.dep
    echo "\texport GOPATH=\$(VP_TOPSRCDIR):\$\$GOPATH && go install $j" >> Makefile.dep
    echo >> Makefile.dep
    echo ".PHONY: check-$i" >> Makefile.dep
    echo "check-$i: configure.ac $k" >> Makefile.dep
    echo "\texport GOPATH=\$(VP_TOPSRCDIR):\$\$GOPATH && go test $j" >> Makefile.dep
    echo >> Makefile.dep
    echo ".PHONY: bench-$i" >> Makefile.dep
    echo "bench-$i: configure.ac $k" >> Makefile.dep
    echo "\texport GOPATH=\$(VP_TOPSRCDIR):\$\$GOPATH && go test -bench . $j" >> Makefile.dep
    echo >> Makefile.dep
    echo ".PHONY: lint-$i" >> Makefile.dep
    echo "lint-$i: configure.ac $k" >> Makefile.dep
    echo "\texport GOPATH=\$(VP_TOPSRCDIR):\$\$GOPATH && go vet $j && \$(VP_TOPSRCDIR)/bin/golint $j" >> Makefile.dep
    echo >> Makefile.dep
    echo ".PHONY: devel-$i" >> Makefile.dep
    echo "devel-$i: configure.ac $k" >> Makefile.dep
    echo "\texport GOPATH=\$(VP_TOPSRCDIR):\$\$GOPATH && go vet $j && go test -v $j" >> Makefile.dep
    echo >> Makefile.dep
    echo ".PHONY: doc-$i" >> Makefile.dep
    echo "doc-$i: configure.ac $k" >> Makefile.dep
    echo "\texport GOPATH=\$(VP_TOPSRCDIR):\$\$GOPATH && godoc $j > doc/txt/$i.txt && godoc -html $j > doc/html/$i.html" >> Makefile.dep
    echo >> Makefile.dep
done

