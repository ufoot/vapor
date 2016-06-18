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

for m in Makefile.dep Makefile.help ; do
    rm -f $m && touch $m
done

echo >> Makefile.help
echo "Below are listed global Makefile targets:" >> Makefile.help
echo "vp: builds all binaries (alias for all)." >> Makefile.help
echo "version: patch autoconf and go source files with current version." >> Makefile.help
echo "stamp: copy stamp.sh to all go package directories." >> Makefile.help
echo "generate: process .go.in files with go generate." >> Makefile.help
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
echo "indent: automatically indent code." >> Makefile.help
echo >> Makefile.help
echo "Below are listed per-package Makefile targets:" >> Makefile.help

cd go || exit 1

for t in "" "generate-" "check-" "bench-" "lint-" "devel-" "doc-" ; do
    echo ".PHONY: ${t}vp" >> ../Makefile.dep
    echo -n "${t}vp:" >> ../Makefile.dep
    for i in vp* ; do
	if ([ "x$t" = "xcheck-" ] || [ "x$t" = "xbench-" ] || [ "x$t" = "xlint-" ] || [ "x$t" = "xdevel-" ]) && grep -q "Autogenerated by Thrift Compiler" $i/*.go ; then
	    true
	else 
            echo -n " ${t}${i}" >> ../Makefile.dep
	    echo -n "${t}${i} " >> ../Makefile.help
	fi
    done
    echo >> ../Makefile.dep
    echo >> ../Makefile.dep
    echo >> ../Makefile.help
done

cd ..
echo "configure.ac: $(ls go/vp*/*.go | sort -u | tr "\n" " ")" >> Makefile.dep
printf "\tcd \$(VP_TOPSRCDIR) && ./utils/dep.sh && ./utils/version.sh\n" >> Makefile.dep
echo >> Makefile.dep
cd go

for i in $(ls -d vp* | sort -u | tr "\n" " ") ; do
    j="github.com/ufoot/vapor/go/$i"
    cd ..
    k=$(ls go/$i/*.go | sort -u | tr "\n" " ")
    cd go
    m=""
    for l in $(ls -d vp* | sort -u | tr "\n" " ") ; do 
        if test $i !=  $l ; then
            if $(grep "\"github.com/ufoot/vapor/$l\"" $i/vp*.go > /dev/null) ; then
                m="$m $l"
            fi
        fi
    done
    echo ".PHONY: $i" >> ../Makefile.dep
    echo "$i: $k generate-$i # $m" >> ../Makefile.dep
    printf "\texport GOPATH=\$(VP_TOPSRCDIR)/go && go install $j\n" >> ../Makefile.dep
    echo >> ../Makefile.dep
    echo ".PHONY: stamp-$i" >> ../Makefile.dep
    echo "stamp-$i: go/$i/stamp.sh" >> ../Makefile.dep
    echo >> ../Makefile.dep
    echo "go/$i/stamp.sh: utils/stamp.sh" >> ../Makefile.dep
    printf "\tcp $< $@\n" >> ../Makefile.dep
    echo >> ../Makefile.dep
    echo ".PHONY: generate-$i" >> ../Makefile.dep
    echo "generate-$i: $k go/$i/stamp.sh" >> ../Makefile.dep
    printf "\texport GOPATH=\$(VP_TOPSRCDIR)/go && go generate $j\n" >> ../Makefile.dep
    echo >> ../Makefile.dep
    echo ".PHONY: check-$i" >> ../Makefile.dep
    echo "check-$i: $k" >> ../Makefile.dep
    printf "\texport GOPATH=\$(VP_TOPSRCDIR)/go && go test $j\n" >> ../Makefile.dep
    echo >> ../Makefile.dep
    echo ".PHONY: bench-$i" >> ../Makefile.dep
    echo "bench-$i: $k" >> ../Makefile.dep
    printf "\texport GOPATH=\$(VP_TOPSRCDIR)/go && go test -bench=. $j\n" >> ../Makefile.dep
    echo >> ../Makefile.dep
    echo ".PHONY: lint-$i" >> ../Makefile.dep
    echo "lint-$i: $k" >> ../Makefile.dep
    printf "\texport GOPATH=\$(VP_TOPSRCDIR)/go && (export PATH=go/bin:\$\$PATH && gometalinter --dupl-threshold=1000 --cyclo-over=20 go/src/$j || true)\n" >> ../Makefile.dep
    echo >> ../Makefile.dep
    echo ".PHONY: devel-$i" >> ../Makefile.dep
    echo "devel-$i: $k" >> ../Makefile.dep
    printf "\techo check $i\n" >> ../Makefile.dep
    printf "\texport GOPATH=\$(VP_TOPSRCDIR)/go && ((go test -v -o test/$i-check.bin -coverprofile=test/$i-cover.cov $j | tee test/$i-check.log) || true)\n" >> ../Makefile.dep
    printf "\texport GOPATH=\$(VP_TOPSRCDIR)/go && ((\$(VP_TOPSRCDIR)/go/bin/go-junit-report < test/$i-check.log > test/$i-junit.xml) || true)\n" >> ../Makefile.dep
    printf "\texport GOPATH=\$(VP_TOPSRCDIR)/go && ((go tool cover -html=test/$i-cover.cov -o doc/cover/$i-cover.html) || true)\n" >> ../Makefile.dep
    printf "\texport GOPATH=\$(VP_TOPSRCDIR)/go && ((\$(VP_TOPSRCDIR)/go/bin/gocov convert test/$i-cover.cov | \$(VP_TOPSRCDIR)/go/bin/gocov-xml > test/$i-cobertura.xml) || true)\n" >> ../Makefile.dep
    printf "\texport GOPATH=\$(VP_TOPSRCDIR)/go && ((go test -v -bench=. $j | tee test/$i-bench.log) || true)\n" >> ../Makefile.dep
    printf "\texport GOPATH=\$(VP_TOPSRCDIR)/go && ((\$(VP_TOPSRCDIR)/go/bin/gobench2plot < test/$i-bench.log > test/$i-plot.xml) || true)\n" >> ../Makefile.dep
    echo >> ../Makefile.dep
    echo ".PHONY: doc-$i" >> ../Makefile.dep
    echo "doc-$i: $k" >> ../Makefile.dep
    printf "\texport GOPATH=\$(VP_TOPSRCDIR)/go && godoc $j > doc/txt/$i.txt && godoc -html $j > doc/html/$i.html\n" >> ../Makefile.dep
    echo >> ../Makefile.dep
done

cd ..

