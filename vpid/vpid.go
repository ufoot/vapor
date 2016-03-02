// Vapor is a toolkit designed to support Liquid War 7.
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
// Vapor homepage: https://github.com/ufoot/vapor
// Contact author: ufoot@ufoot.org

package vpid

import (
	"github.com/ufoot/vapor/vpcrypto"
	"github.com/ufoot/vapor/vperror"
	"github.com/ufoot/vapor/vprand"
	"github.com/ufoot/vapor/vpsum"
	"math/big"
	"time"
)

// NanoScale is the number we should use to transform seconds to nanoseconds.
const NanoScale = 1000000000

// FilterChecker is used to filter and check wether a number,
// typically an id, verifies a property or not.
type FilterChecker interface {
	// Filter processes the value and returns the
	// filtered value.
	Filter(*big.Int) *big.Int
	// Check should return true if number matches property,
	// false if not.
	Check(*big.Int) bool
}

// BytesTransformer is used to process a bytes buffer before
// signing, it
type BytesTransformer interface {
	// Transform takes a bytes buffer and returns a transformed bytes buffer
	Transform([]byte) []byte
}

// GenerateID512 generates a 512 bits id, and signs it.
// If key is nil, no signature is generated.
// If fc is not nil, it is garanteed that the property
// is verified by the id.
// If maxSeconds is greater than 0, will wait for this amount of
// time and try and find an id that has a signature with a
// checksum containing a maximum of zeroes. This allows deep
// per-key personnalisation.
// If minZeroes is greater than 0, will wait until at least that
// amount of zeroes is achieved, regardless of the maxSeconds setting.
// This can take a lot of time...
func GenerateID512(key *vpcrypto.Key, fc FilterChecker, bt BytesTransformer, maxSeconds, minZeroes int) (*big.Int, []byte, int, error) {
	r := vprand.NewRand()
	var ret, tmpInt *big.Int
	var tmpZ, z int
	var tmpData, tmpSig, sig []byte
	var err error

	for start := time.Now().UnixNano(); ret == nil || (key != nil && (time.Now().UnixNano() < start+(int64(maxSeconds)*NanoScale) || z < minZeroes)); {
		tmpInt = vprand.Rand512(r, nil)
		if fc != nil {
			tmpInt = fc.Filter(tmpInt)
		}
		if fc == nil || fc.Check(tmpInt) {
			if key != nil {
				tmpData = vpsum.IntToBuf512(tmpInt)
				if bt != nil {
					tmpData = bt.Transform(tmpData)
				}
				tmpSig, err = key.Sign(tmpData)
				if err != nil {
					return nil, nil, 0, vperror.Chain(err, "can't sign id")
				}
			}
			tmpZ = ZeroesInBuf(vpsum.Checksum512(tmpSig))
			if tmpZ >= z {
				ret = tmpInt
				z = tmpZ
				sig = tmpSig
			}
		}
	}

	return ret, sig, z, nil
}

// GenerateID256 generates a 256 bits id, and signs it.
// If key is nil, no signature is generated.
// If fc is not nil, it is garanteed that the property
// is verified by the id.
// If maxSeconds is greater than 0, will wait for this amount of
// time and try and find an id that has a signature with a
// checksum containing a maximum of zeroes. This allows deep
// per-key personnalisation.
// If minZeroes is greater than 0, will wait until at least that
// amount of zeroes is achieved, regardless of the maxSeconds setting.
// This can take a lot of time...
func GenerateID256(key *vpcrypto.Key, fc FilterChecker, bt BytesTransformer, maxSeconds, minZeroes int) (*big.Int, []byte, int, error) {
	r := vprand.NewRand()
	var ret, tmpInt *big.Int
	var tmpZ, z int
	var tmpData, tmpSig, sig []byte
	var err error

	for start := time.Now().UnixNano(); ret == nil || (key != nil && (time.Now().UnixNano() < start+(int64(maxSeconds)*NanoScale) || z < minZeroes)); {
		tmpInt = vprand.Rand256(r, nil)
		if fc != nil {
			tmpInt = fc.Filter(tmpInt)
		}
		if fc == nil || fc.Check(tmpInt) {
			if key != nil {
				tmpData = vpsum.IntToBuf256(tmpInt)
				if bt != nil {
					tmpData = bt.Transform(tmpData)
				}
				tmpSig, err = key.Sign(tmpData)
				if err != nil {
					return nil, nil, 0, vperror.Chain(err, "can't sign id")
				}
			}
			tmpZ = ZeroesInBuf(vpsum.Checksum256(tmpSig))
			if tmpZ >= z {
				ret = tmpInt
				z = tmpZ
				sig = tmpSig
			}
		}
	}

	return ret, sig, z, nil
}

// GenerateID128 generates a 128 bits id, and signs it.
// If key is nil, no signature is generated.
// If fc is not nil, it is garanteed that the property
// is verified by the id.
// If maxSeconds is greater than 0, will wait for this amount of
// time and try and find an id that has a signature with a
// checksum containing a maximum of zeroes. This allows deep
// per-key personnalisation.
// If minZeroes is greater than 0, will wait until at least that
// amount of zeroes is achieved, regardless of the maxSeconds setting.
// This can take a lot of time...
func GenerateID128(key *vpcrypto.Key, fc FilterChecker, bt BytesTransformer, maxSeconds, minZeroes int) (*big.Int, []byte, int, error) {
	r := vprand.NewRand()
	var ret, tmpInt *big.Int
	var tmpZ, z int
	var tmpData, tmpSig, sig []byte
	var err error

	for start := time.Now().UnixNano(); ret == nil || (key != nil && (time.Now().UnixNano() < start+(int64(maxSeconds)*NanoScale) || z < minZeroes)); {
		tmpInt = vprand.Rand128(r, nil)
		if fc != nil {
			tmpInt = fc.Filter(tmpInt)
		}
		if fc == nil || fc.Check(tmpInt) {
			if key != nil {
				tmpData = vpsum.IntToBuf128(tmpInt)
				if bt != nil {
					tmpData = bt.Transform(tmpData)
				}
				tmpSig, err = key.Sign(tmpData)
				if err != nil {
					return nil, nil, 0, vperror.Chain(err, "can't sign id")
				}
			}
			tmpZ = ZeroesInBuf(vpsum.Checksum128(tmpSig))
			if tmpZ >= z {
				ret = tmpInt
				z = tmpZ
				sig = tmpSig
			}
		}
	}

	return ret, sig, z, nil
}

// GenerateID64 generates a 64 bits id, and signs it.
// If key is nil, no signature is generated.
// If fc is not nil, it is garanteed that the property
// is verified by the id.
// If maxSeconds is greater than 0, will wait for this amount of
// time and try and find an id that has a signature with a
// checksum containing a maximum of zeroes. This allows deep
// per-key personnalisation.
// If minZeroes is greater than 0, will wait until at least that
// amount of zeroes is achieved, regardless of the maxSeconds setting.
// This can take a lot of time...
func GenerateID64(key *vpcrypto.Key, fc FilterChecker, bt BytesTransformer, maxSeconds, minZeroes int) (uint64, []byte, int, error) {
	r := vprand.NewRand()
	var ret, tmpInt uint64
	var tmpBig big.Int
	var tmpZ, z int
	var tmpData, tmpSig, sig []byte
	var err error

	for start := time.Now().UnixNano(); ret == 0 || (key != nil && (time.Now().UnixNano() < start+(int64(maxSeconds)*NanoScale) || z < minZeroes)); {
		tmpInt = vprand.Rand64(r, 0)
		tmpBig.SetUint64(tmpInt)
		if fc != nil {
			tmpBig = *fc.Filter(&tmpBig)
		}
		tmpInt = tmpBig.Uint64()
		if fc == nil || fc.Check(&tmpBig) {
			if key != nil {
				tmpData = vpsum.IntToBuf64(tmpInt)
				if bt != nil {
					tmpData = bt.Transform(tmpData)
				}
				tmpSig, err = key.Sign(tmpData)
				if err != nil {
					return 0, nil, 0, vperror.Chain(err, "can't sign id")
				}
			}
			tmpZ = ZeroesInBuf(vpsum.Checksum64(tmpSig))
			if tmpZ >= z {
				ret = tmpInt
				z = tmpZ
				sig = tmpSig
			}
		}
	}

	return ret, sig, z, nil
}

// GenerateID32 generates a 32 bits id, and signs it.
// If key is nil, no signature is generated.
// If fc is not nil, it is garanteed that the property
// is verified by the id.
// If maxSeconds is greater than 0, will wait for this amount of
// time and try and find an id that has a signature with a
// checksum containing a maximum of zeroes. This allows deep
// per-key personnalisation.
// If minZeroes is greater than 0, will wait until at least that
// amount of zeroes is achieved, regardless of the maxSeconds setting.
// This can take a lot of time...
func GenerateID32(key *vpcrypto.Key, fc FilterChecker, bt BytesTransformer, maxSeconds, minZeroes int) (uint32, []byte, int, error) {
	r := vprand.NewRand()
	var ret, tmpInt uint32
	var tmpBig big.Int
	var tmpZ, z int
	var tmpData, tmpSig, sig []byte
	var err error

	for start := time.Now().UnixNano(); ret == 0 || (key != nil && (time.Now().UnixNano() < start+(int64(maxSeconds)*NanoScale) || z < minZeroes)); {
		tmpInt = vprand.Rand32(r, 0)
		tmpBig.SetUint64(uint64(tmpInt))
		if fc != nil {
			tmpBig = *fc.Filter(&tmpBig)
		}
		tmpInt = uint32(tmpBig.Uint64())
		if fc == nil || fc.Check(&tmpBig) {
			if key != nil {
				tmpData = vpsum.IntToBuf32(tmpInt)
				if bt != nil {
					tmpData = bt.Transform(tmpData)
				}
				tmpSig, err = key.Sign(tmpData)
				if err != nil {
					return 0, nil, 0, vperror.Chain(err, "can't sign id")
				}
			}
			tmpZ = ZeroesInBuf(vpsum.Checksum32(tmpSig))
			if tmpZ >= z {
				ret = tmpInt
				z = tmpZ
				sig = tmpSig
			}
		}
	}

	return ret, sig, z, nil
}
