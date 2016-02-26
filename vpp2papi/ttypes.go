// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package vpp2papi

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/ufoot/vapor/vpcommonapi"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = vpcommonapi.GoUnusedProtection__
var GoUnusedProtection__ int

// HostInfo contains static informations about a host.
//
// Attributes:
//  - HostTitle
//  - HostURL
//  - HostPubKey
//  - HostSig
type HostInfo struct {
	HostTitle  string `thrift:"HostTitle,1" json:"HostTitle"`
	HostURL    string `thrift:"HostURL,2" json:"HostURL"`
	HostPubKey []byte `thrift:"HostPubKey,3" json:"HostPubKey"`
	HostSig    []byte `thrift:"HostSig,4" json:"HostSig"`
}

func NewHostInfo() *HostInfo {
	return &HostInfo{}
}

func (p *HostInfo) GetHostTitle() string {
	return p.HostTitle
}

func (p *HostInfo) GetHostURL() string {
	return p.HostURL
}

func (p *HostInfo) GetHostPubKey() []byte {
	return p.HostPubKey
}

func (p *HostInfo) GetHostSig() []byte {
	return p.HostSig
}
func (p *HostInfo) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *HostInfo) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.HostTitle = v
	}
	return nil
}

func (p *HostInfo) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.HostURL = v
	}
	return nil
}

func (p *HostInfo) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.HostPubKey = v
	}
	return nil
}

func (p *HostInfo) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.HostSig = v
	}
	return nil
}

func (p *HostInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("HostInfo"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *HostInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("HostTitle", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:HostTitle: ", p), err)
	}
	if err := oprot.WriteString(string(p.HostTitle)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.HostTitle (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:HostTitle: ", p), err)
	}
	return err
}

func (p *HostInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("HostURL", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:HostURL: ", p), err)
	}
	if err := oprot.WriteString(string(p.HostURL)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.HostURL (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:HostURL: ", p), err)
	}
	return err
}

func (p *HostInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("HostPubKey", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:HostPubKey: ", p), err)
	}
	if err := oprot.WriteBinary(p.HostPubKey); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.HostPubKey (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:HostPubKey: ", p), err)
	}
	return err
}

func (p *HostInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("HostSig", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:HostSig: ", p), err)
	}
	if err := oprot.WriteBinary(p.HostSig); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.HostSig (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:HostSig: ", p), err)
	}
	return err
}

func (p *HostInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HostInfo(%+v)", *p)
}

// NodeInfo contains static informations about a node.
//
// Attributes:
//  - NodeID
//  - HostPubKey
//  - RingID
type NodeInfo struct {
	NodeID     []byte `thrift:"NodeID,1" json:"NodeID"`
	HostPubKey []byte `thrift:"HostPubKey,2" json:"HostPubKey"`
	RingID     []byte `thrift:"RingID,3" json:"RingID"`
}

func NewNodeInfo() *NodeInfo {
	return &NodeInfo{}
}

func (p *NodeInfo) GetNodeID() []byte {
	return p.NodeID
}

func (p *NodeInfo) GetHostPubKey() []byte {
	return p.HostPubKey
}

func (p *NodeInfo) GetRingID() []byte {
	return p.RingID
}
func (p *NodeInfo) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *NodeInfo) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.NodeID = v
	}
	return nil
}

func (p *NodeInfo) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.HostPubKey = v
	}
	return nil
}

func (p *NodeInfo) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.RingID = v
	}
	return nil
}

func (p *NodeInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("NodeInfo"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *NodeInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("NodeID", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:NodeID: ", p), err)
	}
	if err := oprot.WriteBinary(p.NodeID); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.NodeID (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:NodeID: ", p), err)
	}
	return err
}

func (p *NodeInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("HostPubKey", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:HostPubKey: ", p), err)
	}
	if err := oprot.WriteBinary(p.HostPubKey); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.HostPubKey (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:HostPubKey: ", p), err)
	}
	return err
}

func (p *NodeInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("RingID", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:RingID: ", p), err)
	}
	if err := oprot.WriteBinary(p.RingID); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.RingID (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:RingID: ", p), err)
	}
	return err
}

func (p *NodeInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NodeInfo(%+v)", *p)
}

// RingInfo contains static informations about a ring.
//
// Attributes:
//  - RingID
//  - RingTitle
//  - AppID
type RingInfo struct {
	RingID    []byte `thrift:"RingID,1" json:"RingID"`
	RingTitle string `thrift:"RingTitle,2" json:"RingTitle"`
	AppID     []byte `thrift:"AppID,3" json:"AppID"`
}

func NewRingInfo() *RingInfo {
	return &RingInfo{}
}

func (p *RingInfo) GetRingID() []byte {
	return p.RingID
}

func (p *RingInfo) GetRingTitle() string {
	return p.RingTitle
}

func (p *RingInfo) GetAppID() []byte {
	return p.AppID
}
func (p *RingInfo) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RingInfo) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.RingID = v
	}
	return nil
}

func (p *RingInfo) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.RingTitle = v
	}
	return nil
}

func (p *RingInfo) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.AppID = v
	}
	return nil
}

func (p *RingInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("RingInfo"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RingInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("RingID", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:RingID: ", p), err)
	}
	if err := oprot.WriteBinary(p.RingID); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.RingID (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:RingID: ", p), err)
	}
	return err
}

func (p *RingInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("RingTitle", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:RingTitle: ", p), err)
	}
	if err := oprot.WriteString(string(p.RingTitle)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.RingTitle (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:RingTitle: ", p), err)
	}
	return err
}

func (p *RingInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("AppID", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:AppID: ", p), err)
	}
	if err := oprot.WriteBinary(p.AppID); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.AppID (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:AppID: ", p), err)
	}
	return err
}

func (p *RingInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RingInfo(%+v)", *p)
}

// Used to store results when doing Lookup-like requests.
//
// Attributes:
//  - Path
//  - Errcode
type LookupData struct {
	Path    []*NodeInfo `thrift:"path,1" json:"path"`
	Errcode int32       `thrift:"errcode,2" json:"errcode"`
}

func NewLookupData() *LookupData {
	return &LookupData{}
}

func (p *LookupData) GetPath() []*NodeInfo {
	return p.Path
}

func (p *LookupData) GetErrcode() int32 {
	return p.Errcode
}
func (p *LookupData) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *LookupData) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*NodeInfo, 0, size)
	p.Path = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &NodeInfo{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.Path = append(p.Path, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *LookupData) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Errcode = v
	}
	return nil
}

func (p *LookupData) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("LookupData"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *LookupData) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("path", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:path: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Path)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Path {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:path: ", p), err)
	}
	return err
}

func (p *LookupData) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("errcode", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:errcode: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Errcode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.errcode (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:errcode: ", p), err)
	}
	return err
}

func (p *LookupData) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LookupData(%+v)", *p)
}

// Used to store results when doing Get-like requests.
//
// Attributes:
//  - Value
//  - Path
//  - Errcode
type GetData struct {
	Value   []byte      `thrift:"value,1" json:"value"`
	Path    []*NodeInfo `thrift:"path,2" json:"path"`
	Errcode int32       `thrift:"errcode,3" json:"errcode"`
}

func NewGetData() *GetData {
	return &GetData{}
}

func (p *GetData) GetValue() []byte {
	return p.Value
}

func (p *GetData) GetPath() []*NodeInfo {
	return p.Path
}

func (p *GetData) GetErrcode() int32 {
	return p.Errcode
}
func (p *GetData) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetData) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *GetData) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*NodeInfo, 0, size)
	p.Path = tSlice
	for i := 0; i < size; i++ {
		_elem1 := &NodeInfo{}
		if err := _elem1.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem1), err)
		}
		p.Path = append(p.Path, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *GetData) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Errcode = v
	}
	return nil
}

func (p *GetData) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetData"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetData) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:value: ", p), err)
	}
	if err := oprot.WriteBinary(p.Value); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.value (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:value: ", p), err)
	}
	return err
}

func (p *GetData) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("path", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:path: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Path)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Path {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:path: ", p), err)
	}
	return err
}

func (p *GetData) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("errcode", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:errcode: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Errcode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.errcode (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:errcode: ", p), err)
	}
	return err
}

func (p *GetData) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetData(%+v)", *p)
}
