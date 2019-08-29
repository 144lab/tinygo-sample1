package ble

/*
#define SVCALL_AS_NORMAL_FUNCTION
#include "ble_gap.h"

extern void setProperties(ble_gap_adv_properties_t *prop);
*/
import "C"
import sd "sample1/drivers/softdevice"

// Advertisement encapsulates a single advertisement instance.
type Advertisement struct {
	handle uint8
}

// AdvertiseOptions configures everything related to BLE advertisements.
type AdvertiseOptions struct {
	Interval AdvertiseInterval
}

// AdvertiseInterval is the advertisement interval in 0.625µs units.
type AdvertiseInterval uint32

// NewAdvertiseInterval returns a new advertisement interval, based on an
// interval in milliseconds.
func NewAdvertiseInterval(intervalMillis uint32) AdvertiseInterval {
	// Convert an interval to units of
	return AdvertiseInterval(intervalMillis * 8 / 5)
}

// NewAdvertisement creates a new advertisement instance but does not configure
// it. It can be called before the SoftDevice has been initialized.
func NewAdvertisement() *Advertisement {
	return &Advertisement{
		handle: C.BLE_GAP_ADV_SET_HANDLE_NOT_SET,
	}
}

// Configure this advertisement. Must be called after SoftDevice initialization.
func (a *Advertisement) Configure(broadcastData, scanResponseData []byte, options *AdvertiseOptions) error {
	data := C.ble_gap_adv_data_t{}
	if broadcastData != nil {
		data.adv_data = C.ble_data_t{
			p_data: &broadcastData[0],
			len:    uint16(len(broadcastData)),
		}
	}
	if scanResponseData != nil {
		data.scan_rsp_data = C.ble_data_t{
			p_data: &scanResponseData[0],
			len:    uint16(len(scanResponseData)),
		}
	}
	params := C.ble_gap_adv_params_t{
		properties: C.ble_gap_adv_properties_t{
			//_type: C.BLE_GAP_ADV_TYPE_CONNECTABLE_SCANNABLE_UNDIRECTED,
		},
		interval: uint32(options.Interval),
	}
	C.setProperties(&params.properties)
	return sd.NrfError(C.sd_ble_gap_adv_set_configure(&a.handle, &data, &params))
}

// Start advertisement. May only be called after it has been configured.
func (a *Advertisement) Start(tag uint8) error {
	return sd.NrfError(C.sd_ble_gap_adv_start(a.handle, tag))
}

// Stop advertisement.
func (a *Advertisement) Stop() error {
	return sd.NrfError(C.sd_ble_gap_adv_stop(a.handle))
}
