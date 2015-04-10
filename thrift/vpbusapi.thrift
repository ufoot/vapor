// Vapor is a toolkit designed to support Liquid War 7.
// Copyright (C)  2015  Christian Mauduit <ufoot@ufoot.org>
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
// Vapor homepage: https://github.com/ufoot/vapor
// Contact author: ufoot@ufoot.org


# VpBusApi
# Christian Mauduit (ufoot@ufoot.org)
#
# Thrift protocol between Golang backend server and C++ GUI client.

namespace cpp vpbusapi
namespace go vpbusapi

struct Version {
  1: i32 Major, 
  2: i32 Minor,
  3: string Stamp,
}

struct Package {
  1: string Tarname,
  2: string Name,
  3: string Email,
  4: string URL,
}

service VpBusApi
{
  void ping (),
  Version getVersion (),
  Package getPackage (),
  i64 uptime (),
  oneway void halt ()
}
