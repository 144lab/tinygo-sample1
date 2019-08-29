package ble

/*
#define SVCALL_AS_NORMAL_FUNCTION
#include "nrf_error.h"
#include "ble.h"
#include "ble_for_go.h"
*/
import "C"

import (
	sd "sample1/drivers/softdevice"
)

var (
	secModeOpen C.ble_gap_conn_sec_mode_t
	deviceName  []byte
)

func init() {
	secModeOpen.set_bitfield_sm(1)
	secModeOpen.set_bitfield_lv(1)
}

// Enable ...
func Enable() error {
	appRAMBase := uint32(C.ApplicationRAMBaseAddress)
	return sd.NrfError(C.sd_ble_enable(&appRAMBase))
}

// GapDeviceNameSet ...
func GapDeviceNameSet(name string) error {
	deviceName = []byte(name)
	return sd.NrfError(C.sd_ble_gap_device_name_set(
		&secModeOpen, &deviceName[0],
		uint16(len(deviceName)),
	))
}

// UUIDVsAdd ...
func UUIDVsAdd(b [16]byte) (uint8, error) {
	var uuid128 C.ble_uuid128_t
	var typeUUID C.uint8_t
	uuid128.uuid128 = b
	return uint8(typeUUID), sd.NrfError(C.sd_ble_uuid_vs_add(
		&uuid128,
		&typeUUID,
	))
}

// UUIDVsRemove ...
func UUIDVsRemove(typeUUID uint8) error {
	return sd.NrfError(C.sd_ble_uuid_vs_remove(
		&typeUUID,
	))
}

// Version ...
type Version struct {
	VersionNumber    int
	CompanyID        int
	SubversionNumber int
}

// VersionGet ...
func VersionGet() (*Version, error) {
	var version C.ble_version_t
	err := sd.NrfError(C.sd_ble_version_get(&version))
	if err != nil {
		return nil, err
	}
	return &Version{
		VersionNumber:    int(version.version_number),
		CompanyID:        int(version.company_id),
		SubversionNumber: int(version.subversion_number),
	}, nil
}

// DefaultCfgSet ...
func DefaultCfgSet(tag uint8, params *C.cfg_params_t) error {
	return sd.NrfError(C.bleDefaultCfgSet(tag, params))
}

// EvtGet ...
func EvtGet(b []byte) (int, error) {
	n := C.uint16_t(len(b))
	err := C.sd_ble_evt_get(&b[0], &n)
	if err == C.ErrorNotFound {
		return 0, nil
	}
	return int(n), sd.NrfError(err)
}
