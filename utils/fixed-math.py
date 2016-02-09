#!/usr/bin/python

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

# This is a basic Python (http://www.python.org/) program which
# generates cos, sin and square root tables for use in fixed point
# algorithms.
#
# Indeed, we don't want to rely on floating point routines for
# core algorithms, since most of these must be 100.00000000000%
# predictable, which is not garanted with floating point routines.
#
# The script generates code on standard output, pipe it into
# whatever file you want. It is not run in the standard build, else
# it would depend on a local floating point implementation with its
# limitations. Instead it's run once and the result is stored in
# the source control manager, garanting consistency among all builds
# and platforms.

import math

X32_1 = 0x10000
X64_1 = 0x100000000

X32_2PI = int(round(float(X32_1) * (2.0 * math.pi)))
X32_PI = int(round(float(X32_1) * (math.pi)))
X32_PI2 = int(round(float(X32_1) * (math.pi / 2.0)))
X32_PI4 = int(round(float(X32_1) * (math.pi / 4.0)))
X64_2PI = int(round(float(X64_1) * (2.0 * math.pi)))
X64_PI = int(round(float(X64_1) * (math.pi)))
X64_PI2 = int(round(float(X64_1) * (math.pi / 2.0)))
X64_PI4 = int(round(float(X64_1) * (math.pi / 4.0)))

X32_TABLE_SIZE = 0x400
X32_TABLE_MASK = 0x3FF
X32_TABLE_SHIFT = 10
X64_TABLE_SIZE = 0x2000
X64_TABLE_MASK = 0x1FFF
X64_TABLE_SHIFT = 13

content_begin="""// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015, 2016  Christian Mauduit <ufoot@ufoot.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Vapor homepage: http://www.ufoot.org/liquidwar/v7/vapor
// Contact author: ufoot@ufoot.org

// Generated automatically by utils/fixed-math.py

package vpmath

import (
        "ufoot.org/vapor/vpnumber"
)

"""

content_consts="""
// X32Const2Pi contains 2*PI, that is, about 6.28, as a fixed point integer on 32 bits.
const X32Const2Pi vpnumber.X32 = %d

// X32ConstPi contains PI, that is, about 3.14, as a fixed point integer on 32 bits.
const X32ConstPi vpnumber.X32 = %d

// X32ConstPi2 contains PI/2, that is, about 1.57, as a fixed point integer on 32 bits.
const X32ConstPi2 vpnumber.X32 = %d

// X32ConstPi4 contains PI/4, that is, about 0.78, as a fixed point integer on 32 bits.
const X32ConstPi4 vpnumber.X32 = %d

// X64Const2Pi contains 2*PI, that is, about 6.28, as a fixed point integer on 64 bits.
const X64Const2Pi vpnumber.X64 = %d

// X64ConstPi contains PI, that is, about 3.14, as a fixed point integer on 64 bits.
const X64ConstPi vpnumber.X64 = %d

// X64ConstPi2 contains PI/2, that is, about 1.57, as a fixed point integer on 64 bits.
const X64ConstPi2 vpnumber.X64 = %d

// X64ConstPi4 contains PI/4, that is, about 0.78, as a fixed point integer on 64 bits.
const X64ConstPi4 vpnumber.X64 = %d

const x32TableSize int = 0x%08x
const x32TableMask uint = 0x%08x
const x32TableShift uint = %d
const x64TableSize int = 0x%016x
const x64TableMask uint = 0x%016x
const x64TableShift uint = %d

""" % (X32_2PI,X32_PI,X32_PI2,X32_PI4,X64_2PI,X64_PI,X64_PI2,X64_PI4,X32_TABLE_SIZE,X32_TABLE_MASK,X32_TABLE_SHIFT,X64_TABLE_SIZE,X64_TABLE_MASK,X64_TABLE_SHIFT)

X32_SIN=",\n".join(["%d" % int(round(float(X32_1)*math.sin(float(i)*math.pi/float(2*X32_TABLE_SIZE)))) for i in range(X32_TABLE_SIZE+1)])
X32_ATAN=",\n".join(["%d" % int(round(float(X32_1)*math.atan(float(i)/float(X32_TABLE_SIZE)))) for i in range(X32_TABLE_SIZE+1)])
X32_SQRT=",\n".join(["%d" % int(round(float(X32_1)*math.sqrt(1.0+3.0*float(i)/float(X32_TABLE_SIZE)))) for i in range(X32_TABLE_SIZE+1)])
X64_SIN=",\n".join(["%d" % int(round(float(X64_1)*math.sin(float(i)*math.pi/float(2*X64_TABLE_SIZE)))) for i in range(X64_TABLE_SIZE+1)])
X64_ATAN=",\n".join(["%d" % int(round(float(X64_1)*math.atan(float(i)/float(X64_TABLE_SIZE)))) for i in range(X64_TABLE_SIZE+1)])
X64_SQRT=",\n".join(["%d" % int(round(float(X64_1)*math.sqrt(1.0+3.0*float(i)/float(X64_TABLE_SIZE)))) for i in range(X64_TABLE_SIZE+1)])

content_tables="""
var x32SinTable = [%d]vpnumber.X32 { %s }

var x32AtanTable = [%d]vpnumber.X32 { %s }

var x32SqrtTable = [%d]vpnumber.X32 { %s }

var x64SinTable = [%d]vpnumber.X64 { %s }

var x64AtanTable = [%d]vpnumber.X64 { %s }

var x64SqrtTable = [%d]vpnumber.X64 { %s }

// EOF
""" % (X32_TABLE_SIZE+1, X32_SIN, X32_TABLE_SIZE+1, X32_ATAN, X32_TABLE_SIZE+1, X32_SQRT, X64_TABLE_SIZE+1, X64_SIN, X64_TABLE_SIZE+1, X64_ATAN, X64_TABLE_SIZE+1, X64_SQRT)

print "%s%s%s" % (content_begin, content_consts, content_tables)
