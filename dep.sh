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

rm -f Makefile.dep && touch Makefile.dep

for t in "" "check-" "bench-" "lint-" "devel-" "doc-" ; do
    echo ".PHONY: ${t}vp" >> Makefile.dep
    echo -n "${t}vp:" >> Makefile.dep
    for i in src/ufoot.org/vapor/vp* ; do
	j=$(echo $i | sed "s/.*ufoot.org\/vapor\///")
        echo -n " ${t}$j" >> Makefile.dep
    done
    echo >> Makefile.dep
    echo >> Makefile.dep
done

echo "configure.ac: $(ls src/ufoot.org/vapor/vp*/*.go | sort | tr "\n" " ")" >> Makefile.dep
echo "\tcd \$(VP_TOPSRCDIR) && ./dep.sh && ./stamp.sh" >> Makefile.dep
echo >> Makefile.dep

for i in $(ls -d src/ufoot.org/vapor/vp* | sort | tr "\n" " ") ; do
    j=$(echo $i | sed "s/.*ufoot.org\/vapor\///")
    k=$(echo $i | sed "s/.*src\/ufoot\.org/ufoot.org/")
    l=$(ls $i/*.go | sort | tr "\n" " ")
    echo ".PHONY: $j" >> Makefile.dep
    echo "$j: configure.ac $l" >> Makefile.dep
    echo "\texport GOPATH=\$(VP_TOPSRCDIR):\$\$GOPATH && go install $k" >> Makefile.dep
    echo >> Makefile.dep
    echo ".PHONY: check-$j" >> Makefile.dep
    echo "check-$j: configure.ac $l" >> Makefile.dep
    echo "\texport GOPATH=\$(VP_TOPSRCDIR):\$\$GOPATH && go test $k" >> Makefile.dep
    echo >> Makefile.dep
    echo ".PHONY: bench-$j" >> Makefile.dep
    echo "bench-$j: configure.ac $l" >> Makefile.dep
    echo "\texport GOPATH=\$(VP_TOPSRCDIR):\$\$GOPATH && go test -bench . $k" >> Makefile.dep
    echo >> Makefile.dep
    echo ".PHONY: lint-$j" >> Makefile.dep
    echo "lint-$j: configure.ac $l" >> Makefile.dep
    echo "\texport GOPATH=\$(VP_TOPSRCDIR):\$\$GOPATH && go vet $k && \$(VP_TOPSRCDIR)/bin/golint $k" >> Makefile.dep
    echo >> Makefile.dep
    echo ".PHONY: devel-$j" >> Makefile.dep
    echo "devel-$j: configure.ac $l" >> Makefile.dep
    echo "\texport GOPATH=\$(VP_TOPSRCDIR):\$\$GOPATH && go vet $k && go test -v $k" >> Makefile.dep
    echo >> Makefile.dep
    echo ".PHONY: doc-$j" >> Makefile.dep
    echo "doc-$j: configure.ac $l" >> Makefile.dep
    echo "\texport GOPATH=\$(VP_TOPSRCDIR):\$\$GOPATH && godoc $k > doc/txt/$j.txt && godoc -html $k > doc/html/$j.html" >> Makefile.dep
    echo >> Makefile.dep
done

