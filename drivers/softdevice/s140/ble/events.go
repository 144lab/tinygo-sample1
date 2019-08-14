package ble

import (
	"encoding/binary"
)

/*
event:   10(44): 044B6EC78F7B5001180018000000480000000000E08900200B0000000000000000000000
event:   23(16): FB00FB0048084808
event:   11(9): 13
event:   10(44): 04CC549F24D75401180018000000480000000000E08900200B0000000000000000000000
event:   23(16): FB00FB0048084808
event:   11(9): 13

adv: 0x20006340(11)
version: 9/89/182
temp.: 53.333332
event:   10(44): 04464E97FDE34401120012000000480000000000406300200B0000000000000000000000
event:   23(16): FB00FB0048084808

*/

// EvtFrame ...
type EvtFrame struct {
	ID         uint16
	Length     uint16
	ConnHandle uint16
	_          uint16
	Payload    []byte
}

// UnmarshalBinary ...
func (e *EvtFrame) UnmarshalBinary(b []byte) error {
	e.ID = binary.LittleEndian.Uint16(b[:2])
	e.Length = binary.LittleEndian.Uint16(b[2:4])
	e.ConnHandle = binary.LittleEndian.Uint16(b[4:6])
	e.Payload = b[8:e.Length]
	return nil
}

// MarshalBinary ...
func (e *EvtFrame) MarshalBinary() ([]byte, error) {
	b := make([]byte, len(e.Payload)+6)
	binary.LittleEndian.PutUint16(b[:2], e.ID)
	binary.LittleEndian.PutUint16(b[2:4], uint16(len(b)))
	binary.LittleEndian.PutUint16(b[4:6], e.ConnHandle)
	copy(b[6:], e.Payload)
	return b, nil
}

// Addr ...
type Addr struct {
	Type uint8
	Addr [6]uint8
}

// UnmarshalBinary ...
func (e *Addr) UnmarshalBinary(b []byte) error {
	e.Type = b[0]
	copy(e.Addr[:], b[1:])
	return nil
}

// GapConnParams ...
type GapConnParams struct {
	MinConnInterval uint16 /* Minimum Connection Interval in 1.25 ms units, see @ref BLE_GAP_CP_LIMITS.*/
	MaxConnInterval uint16 /* Maximum Connection Interval in 1.25 ms units, see @ref BLE_GAP_CP_LIMITS.*/
	SlaveLatency    uint16 /* Slave Latency in number of connection events, see @ref BLE_GAP_CP_LIMITS.*/
	ConnSupTimeout  uint16 /* Connection Supervision Timeout in 10 ms units, see @ref BLE_GAP_CP_LIMITS.*/
}

// UnmarshalBinary ...
func (e *GapConnParams) UnmarshalBinary(b []byte) error {
	e.MinConnInterval = binary.LittleEndian.Uint16(b[:2])
	e.MaxConnInterval = binary.LittleEndian.Uint16(b[2:4])
	e.SlaveLatency = binary.LittleEndian.Uint16(b[4:6])
	e.ConnSupTimeout = binary.LittleEndian.Uint16(b[6:8])
	return nil
}

/* for Common Events */

// UserMemRequest ...
type UserMemRequest struct {
	Type uint8
}

// UnmarshalBinary ...
func (e *UserMemRequest) UnmarshalBinary(b []byte) error {
	e.Type = b[0]
	return nil
}

// UserMemRelease ...
type UserMemRelease struct {
	Type     uint8
	MemBlock []byte
}

// UnmarshalBinary ...
func (e *UserMemRelease) UnmarshalBinary(b []byte) error {
	e.Type = b[0]
	e.MemBlock = b[1:]
	return nil
}

/* for Gap Events */

// GapConnected ...
type GapConnected struct {
	PeerAddr   Addr
	Role       uint8
	ConnParams GapConnParams
	AdvHandle  uint8
	_          [3]uint8
	AdvData    uint32
	AdvLength  uint16
	_          [2]uint8
	RspData    uint32
	RspLength  uint16
	_          [2]uint8
}

// UnmarshalBinary ...
func (e *GapConnected) UnmarshalBinary(b []byte) error {
	e.PeerAddr.UnmarshalBinary(b[:7])
	e.Role = b[7]
	e.ConnParams.UnmarshalBinary(b[8:16])
	e.AdvHandle = b[16]
	e.AdvData = binary.LittleEndian.Uint32(b[20:24])
	e.AdvLength = binary.LittleEndian.Uint16(b[24:26])
	e.RspData = binary.LittleEndian.Uint32(b[28:32])
	e.RspLength = binary.LittleEndian.Uint16(b[32:34])
	return nil
}

// GapDisconnected ...
type GapDisconnected struct {
	Reason uint8
}

// UnmarshalBinary ...
func (e *GapDisconnected) UnmarshalBinary(b []byte) error {
	e.Reason = b[0]
	return nil
}
