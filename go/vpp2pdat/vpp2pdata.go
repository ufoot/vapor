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
	Ring0Base64RingID = "rBYTLlG4OWNQi9Tedd_q0YnMjteYCsTIz0Y6qBOrH205FoVGCDZo4GGtRTbljbMxycElhsnOMTHnkVXIqlTK_w=="
	// Ring0Base64AppID contains the application ID of the default ring0, encoded in base64
	Ring0Base64AppID = "K3kU9SU3EAEiqkAgsYo72g=="
	// Ring0Base64HostPubKey contains the public key of the default ring0, encoded in base64
	Ring0Base64HostPubKey = "xsBNBFb67QwBCADYrYiUK_P1RDTCdewQFtAAAAQ6XQmsvrv8VCkErv-jqyH-6mvzAHJ4OCO6ibMR9QUU7S6D97jmh0YFStt-gdItjVEiLk0RN-yzMXJdR1XLTa9mNVy7ECqP4WHuDpfMhYhX7VVcNXEgDCmOOzaQMZsiL398MlVmyzJqnsD5s1xxm3OXTJa6n1OYXQpZI_w-URP-sCgpXhka5zpk9hQkTz9BjxV-rTEnW7UI67H7f4PGUnL-1qejFJ9Cw3TvMY-cFrhthSRUIZx4VeL-qRVpFQVhph_uKqTWU6O6P7XffO0D7wVIjF7jbmm6RpqpxAJOEkxjkCr2ozD_PoH2D5LHCME5ABEBAAHNQFZhcG9yIFRvb2xraXQgKGh0dHBzOi8vZ2l0aHViLmNvbS91Zm9vdC92YXBvcikgPHVmb290QHVmb290Lm9yZz7CwGIEEwEIABYFAlb67QwJEN79NI3QRWkVAhsDAhkBAABqhwgAzX8jBfuVhlxhqnCHuiD03UjPFKBXffi_YY_WL0tZelQibCVGGk4AAZMRiPUPHfM-3zlKQaQcmt74ZuMgTCwrA8Tj6jbC4VHiBkxBwb47WiBK2tHmgYmA4q_2FixYpMVKrfTpjguB0nJffmGS1QGE0V1rRZYtwaUMI-CaXXiAtDsD5Ax7y_LM0UTlHfMxU9hBYlVVlF6C6q4aE57wtx0OTANpVGOtzpCwpXUCrXYJnsNC2xrA5MxjhRcTBvNJj99eIDwIf2VJ_1eLF7dSrfux_fZJujsUrCottJbBrV7KpwMysYi34jkonwVVcoBgSMYI_pvnkwU2BdV3ufm06XqqXs7ATQRW-u0MAQgAp_isNP7DpuP-hXTuSOG8yLrX9n9wxqoFpBljDGoraw015X-C5zhD7-rsteagShEq3Bg2dtqr2FOe5mkKDUHDLFFDgxxePa0o-Ybr549YeIpETVmL862iPOy9cWlIprTOo0leGspTvmkPG6ot1py4ht21sJrKM6R45s_mzLJPx_6136nIGcFYDZPNLr-qGWQ6-gZ1srSJVS3T1dwPxRHJW9M1R4AG5FZQkwh_MCx_ZYQ_xncTbn1mC9eexqcFSsTSc-9uHqn3_TkX1lKVpwzVnCrw9tk6ose7bx-tY5-X5ZaJ6EHALaXXn-VypgjGIRFFUzEFZDz6K-qxl5OGDS36awARAQABwsBfBBgBCAATBQJW-u0MCRDe_TSN0EVpFQIbDAAAJ50IAEjL793i6asp7zBnYy2gcvlUKkAmH2C5ejln8HoOZIBRertaGjNYZg7DDs1xUrwmrrxemkaKZQFHrFr6956weO8o1fXzWQgHOvF9RP1KP5f3TcBR4wG_Gyj8zHcuqoSZVUVbLH7oGi6fdB20SK9-eKvUC0pcBa3tKsmk9QTMsu7p-rFgAKRdWJJFvKoIWlAe9nkjOLla3rQ_zOptTFJAnmWBtjNVsYwSdJ5uQQsxwjzfds_RAEdcFECAWEDNpZxeGVCIi8nWG6V0_BeiziEi4E4KI1GB7QwKq_Ub5RwLv2J8k-8WGrKb2D1HzvZk_z-YZCZgLgDaNT9K3bvRVGwf0HY="
	// Ring0Base64RingSig contains the siignature of the default ring0, encoded in base64
	Ring0Base64RingSig = "wsBcBAABCAAQBQJW-u0UCRDe_TSN0EVpFQAAMMkIAIbKsbQ4ovvA5wN0s7V3r1-YVyTT0wCDSjAJjy58GG_YN-9iILfCLlmBoutmA_TtQxB3f9AG14oXJTNA9-y0rDMhZzq8cXDteEaNWKQEZOvYe9fjxeoh1lpHwVkXK2CWI5ngn6sD3gX1lNv0LrnDWWy2aLDWxdLqbASJi8L0NkpEXkrlnsQymcHbq68FjxC3P6F9I-oVMJF8TM2F-_h86WeptJ0xZtGEsCe2aApZtptaEatbeFsWiwQAWdu5oDTD1gxUqnZ8glxN1AoY-M5C146wXKPSr1aHYrHahHe1QHhTTI1Da4E-Ld_Iss8QBRCmjn5CcL2FGPT9ZBycrZ_b4HQ="

	// DefaultBruijnM default for the m parameter (base) used for Koorde/Bruijn ops.
	DefaultBruijnM = 16
	// DefaultBruijnN default for the n parameter (number of elements) used for Koorde/Bruijn ops.
	DefaultBruijnN = 64
	// DefaultNbCopy default for the number of copies of a key we store within the network.
	DefaultNbCopy = 3
	// DefaultNbStep default to optimizes Bruijn walk by considering only this number
	// of steps in the worst case.
	DefaultNbStep = 8
	// DefaultCallTimeout is the default delay after which a network call is considered
	// stall and cancelled.
	DefaultCallTimeout = 5
	// DefaultSyncDelay is the default delay between two node synchronizations.
	DefaultSyncDelay = 15
	// DefaultDisconnectTimeout is the default delay after which, if a node has been
	// unreachable, it's considered totally unjoinable.
	DefaultDisconnectTimeout = 60
	// DefaultDataLifetime is the amount of time after which keys are automatically deleted
	// to purge the data store, no matter what. Any update on a key extends this duration.
	DefaultDataLifetime = 86400
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

// CheckHostInfo checks that a host info struct is filled with correct data.
func CheckHostInfo(host *vpp2papi.HostInfo) (bool, error) {
	var ok bool
	var err error

	ok, err = CheckTitle(host.HostTitle)
	if err != nil || !ok {
		return false, err
	}
	ok, err = CheckURL(host.HostURL)
	if err != nil || !ok {
		return false, err
	}
	ok, err = CheckPubKey(host.HostPubKey)
	if err != nil {
		return false, err
	}
	ok, err = CheckSig(host.HostSig)
	if err != nil {
		return false, err
	}
	_, err = HostInfoCheckSig(host)
	if err != nil {
		return false, err
	}

	return true, nil
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

// CheckNodeInfo checks that a node info struct is filled with correct data.
func CheckNodeInfo(node *vpp2papi.NodeInfo) (bool, error) {
	var ok bool
	var err error

	ok, err = CheckNodeID(node.NodeID)
	if err != nil || !ok {
		return false, err
	}
	ok, err = CheckRingID(node.RingID)
	if err != nil || !ok {
		return false, err
	}
	ok, err = CheckPubKey(node.HostPubKey)
	if err != nil {
		return false, err
	}
	ok, err = CheckSig(node.NodeSig)
	if err != nil {
		return false, err
	}
	_, err = NodeInfoCheckSig(node)
	if err != nil {
		return false, err
	}

	return true, nil
}

// DefaultRingConfig returns a default ring configuration, with 256-bit keys
func DefaultRingConfig() *vpp2papi.RingConfig {
	return &vpp2papi.RingConfig{BruijnM: DefaultBruijnM, BruijnN: DefaultBruijnN, NbCopy: DefaultNbCopy, NbStep: DefaultNbStep, CallTimeout: DefaultCallTimeout, SyncDelay: DefaultSyncDelay, DisconnectTimeout: DefaultDisconnectTimeout, DataLifetime: DefaultDataLifetime}
}

// RingInfoSigBytes returns the byte buffer that needs to be signed.
func RingInfoSigBytes(ringInfo *vpp2papi.RingInfo) []byte {
	bufTitle := []byte(ringInfo.RingTitle)
	bufDescription := []byte(ringInfo.RingDescription)
	// not using a JSON Marshall since technically, we do not want any
	// compatible JSON representations (it could differ depending
	// on spaces and implementation details...) but a unique projection
	// of the data.
	bufScalar := fmt.Sprintf("(%d,%d,%d,%d,%d,%d,%d,%d,%t)", ringInfo.Config.BruijnM, ringInfo.Config.BruijnN, ringInfo.Config.NbCopy, ringInfo.Config.NbStep, ringInfo.Config.CallTimeout, ringInfo.Config.SyncDelay, ringInfo.Config.DisconnectTimeout, ringInfo.Config.DataLifetime, ringInfo.HasPassword)
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

// CheckRingInfo checks that a ring info struct is filled with correct data.
func CheckRingInfo(ring *vpp2papi.RingInfo) (bool, error) {
	var ok bool
	var err error

	ok, err = CheckRingID(ring.RingID)
	if err != nil || !ok {
		return false, err
	}
	ok, err = CheckTitle(ring.RingTitle)
	if err != nil || !ok {
		return false, err
	}
	ok, err = CheckDescription(ring.RingDescription)
	if err != nil || !ok {
		return false, err
	}
	ok, err = CheckID(ring.AppID)
	if err != nil || !ok {
		return false, err
	}
	ok, err = CheckRingConfig(ring.Config)
	if err != nil {
		return false, err
	}
	ok, err = CheckPubKey(ring.HostPubKey)
	if err != nil {
		return false, err
	}
	ok, err = CheckSig(ring.RingSig)
	if err != nil {
		return false, err
	}
	_, err = RingInfoCheckSig(ring)
	if err != nil {
		return false, err
	}

	return true, nil
}

// CheckContextInfo checks wether a contextinfo struct is viable
func CheckContextInfo(context *vpp2papi.ContextInfo) (bool, error) {
	var err error

	_, err = CheckHostInfo(context.SourceHost)
	if err != nil {
		return false, err
	}
	_, err = CheckRingInfo(context.SourceRing)
	if err != nil {
		return false, err
	}
	_, err = CheckNodeInfo(context.SourceNode)
	if err != nil {
		return false, err
	}
	_, err = CheckNodeID(context.TargetNodeID)
	if err != nil {
		return false, err
	}

	return true, nil
}
