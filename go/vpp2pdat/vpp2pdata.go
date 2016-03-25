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

package vpp2pdat

import (
	"fmt"
	"github.com/ufoot/vapor/go/vpbruijn"
	"github.com/ufoot/vapor/go/vpcrypto"
	"github.com/ufoot/vapor/go/vpid"
	"github.com/ufoot/vapor/go/vpp2papi"
	"github.com/ufoot/vapor/go/vpsum"
)

const (
	// Host0Title is a default title for the root directory host.
	// It's a purely virtual host, never really instanciated, its
	// private key is forever lost once the public key is generated,
	// therefore it signs some artefacts once, and it's over.
	Host0Title = "Vapor Host 0"
	// Host0URL is the base url where to find the default directory
	// host. It should be a well known, usually up and running server.
	// Technically, once could use any address, but let's say this
	// is just a default seed.
	Host0URL = "http://ufoot.org:8777"
	// Ring0Title is the title of the default directory ring.
	Ring0Title = "Vapor Ring 0"
	// Ring0Description is the description of the default directory ring.
	Ring0Description = "The main Vapor directory, all rings can refer to this one to get the rings list. Think of it as a seed, a bootstrap, you're free to use another one."

	// Ring0Base64RingID contains the ring ID of the default ring0, encoded in base64
	Ring0Base64RingID = "QF6Xm658lNZhDt3Sev3QJtxYL19Odu0M_R-Gus6X3bLZA27of54Nbo2P-DS85phxWKIweZsoHQZ9w9hJKumEnA=="
	// Ring0Base64AppID contains the application ID of the default ring0, encoded in base64
	Ring0Base64AppID = "Td5Li6DajQ0VUT9BKyriaA=="
	// Ring0Base64HostPubKey contains the public key of the default ring0, encoded in base64
	Ring0Base64HostPubKey = "xsBNBFbW8k8BCACyz6zwlgqZK75FKB4Rq1e0lZy3a8T-x4Cm0dusvcHTmFvTtGOd5X7pNw831YbLXNvH2wUnOqieSh3GkabETtFxUAO1dCLIeMO_eZWcby5siHLUsubPUN19DPeTyf1wfFKSRJtVX0Nb7QFwn7q8a_SxQnoDnPrQLdCO63ZPAvEgjRs3yKvVs2IzwJdTMKAHOjl3bpfMuGDt9RW89KN9rrItgKjgDTawa95vvezSoszIajTCbReZEoAK-yVnQS4FqjMe9XqHPKHIjGsYOm3W9eN9KQbSWRJ7RAZJphbCrvwp4233iHDRu_JlivbMv7y-ODi1JF9nDALwlsAUF0kmpuwhABEBAAHNQFZhcG9yIFRvb2xraXQgKGh0dHBzOi8vZ2l0aHViLmNvbS91Zm9vdC92YXBvcikgPHVmb290QHVmb290Lm9yZz7CwGIEEwEIABYFAlbW8k8JEJDu2fRqTYXCAhsDAhkBAABmGggAD-2gLubjL8jCd1oArcpTJJ7Xmv-d8X5wDPR-BfHhs6n2cc--ADg5UdFGXSe_ztMBgYjrlaSFrrzv5bVx2WXt6WqtGWNe7Nx5FD8lCQ3TZt3qkN9a1F7Gv2v5SAfDB6tYjqT0fkIP6Df2ikAbs19f3_At-7NCTQV3Nc46c49F7a1DgupKhtbZxTjr6uD3egU5Su_cDGO2EnD08nYJi5vDPHrwTLDfBVuTU3piSXWrfxeciqK2E8aOhjqEmH71gTmcXD5f0J2VYR4GZKpZrqUr1CT0EMIXC8BgTPTgJEOhhuS1psyNrjVab6r9caS0SfP9fG8eqlLdNVxsXlRjEicZ0s7ATQRW1vJPAQgAzGV7ftrZYtI0iYDJmG_0z_Y9KpOs-qX0QUotlJ8z_G5lD2LiPXksiKMf6LqH1yqopy6pIwyDJDR9dniSID2SrkbqaVof4yDejG7DPyDRDR3qWizMAgpomaW_3r9ho35MnoxPtALu_y2WNpjGt4hjgKM294YxK_QorH3IfYjPWav6dikOVBSOAnyeRGY4F6AqYCa2bhrGiiPUUGsvKiKsVhBukpwyORq8_fKw6ADmjZCY5xDHPWODkoH7i7wEHyJvNoLK8Svtn0Pen2NXq5YwSXtvrdUC5_ilsVKFcw_YJ9IUYcHyCOpOPbIDsxs56hs2GFrfwx5K2XUQ9kWl8wJ4bwARAQABwsBfBBgBCAATBQJW1vJPCRCQ7tn0ak2FwgIbDAAAboEIAESsoVrVimdZF1PwBciGyq7snt1aGz-3asXaA22X_SXM4NIeO8bM8ZP6_cbvX0eZrnduJNJ-KapNCqnjwlker-wfRY6x7U4S2Mcg7wHI3pZtDXj4-4LFEKhf9kvmLD20k7mAgv6CXss8Uyvj5ZiJfbRPnVjqMTiuPo3XRuDHA_qWsye6jd6JMH9CwWsjZ356699YkrvFPE33vVxOt1Yh5dFKPcUPyVJ0dgwf3vJ3y6dluVahrFHYW5BUrtLXfX-dvJ-21srfQWUKii51MeUCOH5nVQVuE-8ViOi44fRJM8YQD60PL9yxFAH4UEq13woF4Myl4kc2YrM8-0okGyWE7j8="
	// Ring0Base64RingSig contains the siignature of the default ring0, encoded in base64
	Ring0Base64RingSig = "wsBcBAABCAAQBQJW1vJkCRCQ7tn0ak2FwgAAtYkIAK6D0IfKGLbPhk02tPeeO_NoWSSiWSeya8GaeCaV770sFZirmSO71ty8SNM839A7vUj7sulQCWePgXbYnljYVKF6ux1bMQpBhzCrAeT59jbNL1GyqRinK7jAbTITBRH1bKS3CSVNVMurWLbChv6C4rGrxjiZAPm4jIGgV4q-N-d02PvzoyLZbC7K0PhbZDLUTIgKLA7Z4lsptVY8sOcAUV2b5bRtmX38WyrI1ZD8taOcds7tdW1uYz6xQMULA-jFCdJbizduZ27v7dfnJIqXXk6L0IoocqDcqhQXyET4-2iuOiBoaZdbj2IcH-IcjxL1TuAD_bodOx4Go5JDI295vLM="

	// DefaultBruijnM default for the m parameter (base) used for Koorde/Bruijn ops.
	DefaultBruijnM = 16
	// DefaultBruijnN default for the n parameter (number of elements) used for Koorde/Bruijn ops.
	DefaultBruijnN = 64
	// DefaultNbCopy default for the number of copies of a key we store within the network.
	DefaultNbCopy = 3
	// DefaultNbStep default to optimizes Bruijn walk by considering only this number
	// of steps in the worst case.
	DefaultNbStep = 8
)

// HostInfoSigBytes returns the byte buffer that needs to be signed.
func HostInfoSigBytes(hostInfo *vpp2papi.HostInfo) []byte {
	return []byte(fmt.Sprintf("%s;%s", hostInfo.HostTitle, hostInfo.HostURL))
}

// HostInfoIsSigned returns true if the has been self-signed.
// It does not checks if the signature is valid.
func HostInfoIsSigned(hostInfo *vpp2papi.HostInfo) bool {
	if hostInfo.HostSig == nil || len(hostInfo.HostSig) <= 0 {
		return false
	}
	return true
}

// HostInfoCheckSig checks if the host signature is OK, if it's not, returns false and an error.
func HostInfoCheckSig(hostInfo *vpp2papi.HostInfo) (bool, error) {
	var ok bool

	if hostInfo.HostPubKey == nil || len(hostInfo.HostPubKey) <= 0 {
		return false, fmt.Errorf("no public key")
	}
	if !HostInfoIsSigned(hostInfo) {
		if IsPubKeyExpectedToSign(hostInfo.HostPubKey) {
			return false, fmt.Errorf("no signature")
		}
		// no signature but we don't expect such a key to sign anything
		return true, nil
	}

	key, err := vpcrypto.ImportPubKey(hostInfo.HostPubKey)
	if err != nil {
		return false, err
	}
	ok, err = key.CheckSig(HostInfoSigBytes(hostInfo), hostInfo.HostSig)
	if err != nil {
		return false, err
	}

	return ok, nil
}

// NodeInfoSigBytes returns the byte buffer that needs to be signed.
func NodeInfoSigBytes(nodeInfo *vpp2papi.NodeInfo) []byte {
	ret := make([]byte, len(nodeInfo.NodeID)+len(nodeInfo.RingID))
	copy(ret[0:len(nodeInfo.NodeID)], nodeInfo.NodeID)
	copy(ret[len(nodeInfo.NodeID):len(nodeInfo.NodeID)+len(nodeInfo.RingID)], nodeInfo.RingID)
	return ret
}

// NodeInfoIsSigned returns true if the node has been signed by corresponding host.
// It does not checks if the signature is valid.
func NodeInfoIsSigned(nodeInfo *vpp2papi.NodeInfo) bool {
	if nodeInfo.NodeSig == nil || len(nodeInfo.NodeSig) <= 0 {
		return false
	}
	return true
}

// NodeInfoCheckSig checks if the node signature is OK, if it's not, returns 0 and an error.
// If it's OK, returns the number of zeroes in the signature hash.
func NodeInfoCheckSig(nodeInfo *vpp2papi.NodeInfo) (int, error) {
	var z int

	if nodeInfo.HostPubKey == nil || len(nodeInfo.HostPubKey) <= 0 {
		return 0, fmt.Errorf("no public key")
	}
	if !NodeInfoIsSigned(nodeInfo) {
		if IsPubKeyExpectedToSign(nodeInfo.HostPubKey) {
			return 0, fmt.Errorf("no signature")
		}
		// no signature but we don't expect such a key to sign anything
		return 0, nil
	}

	key, err := vpcrypto.ImportPubKey(nodeInfo.HostPubKey)
	if err != nil {
		return 0, err
	}
	_, err = key.CheckSig(NodeInfoSigBytes(nodeInfo), nodeInfo.NodeSig)
	if err != nil {
		return 0, err
	}
	z = vpid.ZeroesInBuf(vpsum.Checksum256(nodeInfo.NodeSig))

	return z, nil
}

// DefaultRingConfig returns a default ring configuration, with 256-bit keys
func DefaultRingConfig() *vpp2papi.RingConfig {
	return &vpp2papi.RingConfig{BruijnM: DefaultBruijnM, BruijnN: DefaultBruijnN, NbCopy: DefaultNbCopy, NbStep: DefaultNbStep}
}

// RingInfoSigBytes returns the byte buffer that needs to be signed.
func RingInfoSigBytes(ringInfo *vpp2papi.RingInfo) []byte {
	bufTitle := []byte(ringInfo.RingTitle)
	bufDescription := []byte(ringInfo.RingDescription)
	bufScalar := fmt.Sprintf("(%d,%d,%d,%d,%t)", ringInfo.Config.BruijnM, ringInfo.Config.BruijnN, ringInfo.Config.NbCopy, ringInfo.Config.NbStep, ringInfo.HasPassword)
	ret := make([]byte, len(ringInfo.RingID)+len(bufTitle)+len(bufDescription)+len(ringInfo.AppID)+len(bufScalar))
	begin := 0
	end := begin + len(ringInfo.RingID)
	copy(ret[begin:end], ringInfo.RingID)
	begin += len(ringInfo.RingID)
	end = begin + len(bufTitle)
	copy(ret[begin:end], bufTitle)
	begin += len(bufTitle)
	end = begin + len(bufDescription)
	copy(ret[begin:end], bufDescription)
	begin += len(bufDescription)
	end = begin + len(ringInfo.AppID)
	copy(ret[begin:end], ringInfo.AppID)
	begin += len(ringInfo.AppID)
	end = begin + len(bufScalar)
	copy(ret[begin:end], bufScalar)

	return ret
}

// CheckRingConfig checks that the ring config values
// are within reasonnable ranges.
func CheckRingConfig(config *vpp2papi.RingConfig) (bool, error) {
	if config.BruijnM < vpbruijn.GenericMinM || config.BruijnM > vpbruijn.GenericMaxM {
		return false, fmt.Errorf("bad BruijnM param %d, should be between %d and %d", config.BruijnM, vpbruijn.GenericMinM, vpbruijn.GenericMaxM)
	}
	if config.BruijnN < vpbruijn.GenericMinN || config.BruijnN > vpbruijn.GenericMaxN {
		return false, fmt.Errorf("bad BruijnN param %d, should be between %d and %d", config.BruijnN, vpbruijn.GenericMinN, vpbruijn.GenericMaxN)
	}
	if config.NbCopy <= 0 || config.NbCopy > config.BruijnM*config.BruijnN {
		return false, fmt.Errorf("bad NbCopy param %d, should be between 0 and %d (the latter is BruijnN*BruijnM, while this number does not technically has a direct meaning, it should really be bigger than copies, just check your settings, usually, keeping more than a dozen copies is overkill)", config.NbCopy, config.BruijnN*config.BruijnM)
	}
	if config.NbStep < 1 || config.NbStep > config.BruijnN {
		return false, fmt.Errorf("bad NbStep param %d, should be between 1 and BruijnN which is %d", config.NbStep, config.BruijnN)
	}
	return true, nil
}

// RingInfoIsSigned returns true if the ring has been signed by corresponding host.
// It does not checks if the signature is valid.
func RingInfoIsSigned(ringInfo *vpp2papi.RingInfo) bool {
	if ringInfo.RingSig == nil || len(ringInfo.RingSig) <= 0 {
		return false
	}
	return true
}

// RingInfoCheckSig checks if the ring signature is OK, if it's not, returns 0 and an error.
// If it's OK, returns the number of zeroes in the signature hash.
func RingInfoCheckSig(ringInfo *vpp2papi.RingInfo) (int, error) {
	var z int

	if ringInfo.HostPubKey == nil || len(ringInfo.HostPubKey) <= 0 {
		return 0, fmt.Errorf("no public key")
	}
	if !RingInfoIsSigned(ringInfo) {
		if IsPubKeyExpectedToSign(ringInfo.HostPubKey) {
			return 0, fmt.Errorf("no signature")
		}
		// no signature but we don't expect such a key to sign anything
		return 0, nil
	}

	key, err := vpcrypto.ImportPubKey(ringInfo.HostPubKey)
	if err != nil {
		return 0, err
	}
	_, err = key.CheckSig(RingInfoSigBytes(ringInfo), ringInfo.RingSig)
	if err != nil {
		return 0, err
	}
	z = vpid.ZeroesInBuf(vpsum.Checksum512(ringInfo.RingSig))

	return z, nil
}

// CheckContextInfo checks wether a contextinfo struct is viable
func CheckContextInfo(context *vpp2papi.ContextInfo) (bool, error) {
	var err error

	_, err = HostInfoCheckSig(context.SourceHost)
	if err != nil {
		return false, err
	}
	_, err = RingInfoCheckSig(context.SourceRing)
	if err != nil {
		return false, err
	}
	_, err = NodeInfoCheckSig(context.SourceNode)
	if err != nil {
		return false, err
	}

	return true, nil
}
