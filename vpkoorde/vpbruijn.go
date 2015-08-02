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

package vpkoorde

// Read https://en.wikipedia.org/wiki/De_Bruijn_graph for theory

import (
	"math/big"
	"fmt"
)

func checkMNX(m, n int, x *big.Int) (*big.Int, *big.Int, error) {
	if m<=1 {
		return nil, nil, fmt.Errorf("Invalid m=%d (base used) for Bruijn graph, should be >=2", m)
	}
	if n<=1 {
		return nil, nil, fmt.Errorf("Invalid n=%d (number of digits) for Bruijn graph, should be >=2", n)
	}
	bm:=big.NewInt(int64(m))
	bn:=big.NewInt(int64(n))
	max:=bm.Exp(bm,bn,nil)
	if x.Cmp(max)>=1 {
		return nil, nil, fmt.Errorf("Invalid x=%s (current node) for Bruijn graph with m=%d n=%d, should be <=%s", x.String(),m,n,max.String())
	}
	return bm, max, nil
}

func next0(x,bm,max *big.Int) *big.Int {
	p:=x.Mul(x,bm)
	return p.Mod(p,max)
}

/// BruijnNext0 returns on Bruijn node pointed by this nodes. Other
/// nodes might be deduced by just incrementing this one.
func BruijnNext0(m, n int, x *big.Int) (*big.Int, error) {
	bm,max,err:=checkMNX(m,n,x)
	if err!=nil {
		return nil,err
	}
	
	return next0(x,bm,max), nil
}

/// BruijnNextList returns the list of all Bruijn nodes pointed by
/// this node, the nodes following this one (we descend the graph).
func BruijnNextList(m, n int, x *big.Int) ([]*big.Int, error) {
	bm,max,err:=checkMNX(m,n,x)
	if err!=nil {
		return nil,err
	}
	next0:=next0(x,bm,max)

	ret:=make([]*big.Int, m)
	for i:=range ret {
		ret[i]=next0.Add(next0,big.NewInt(int64(i)))
		ret[i]=ret[i].Mod(ret[i],max)
	}
	
	return ret, nil
}
