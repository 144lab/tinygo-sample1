package ble

/*
#define SVCALL_AS_NORMAL_FUNCTION
#include "ble_gap.h"
*/
import "C"
import sd "sample1/drivers/softdevice"

var (
	gapConnParams C.ble_gap_conn_params_t
)

// GapPpcpSet ...
func GapPpcpSet() error {
	gapConnParams.min_conn_interval = C.BLE_GAP_CP_MIN_CONN_INTVL_MIN
	gapConnParams.max_conn_interval = C.BLE_GAP_CP_MIN_CONN_INTVL_MAX
	gapConnParams.slave_latency = 0
	gapConnParams.conn_sup_timeout = C.BLE_GAP_CP_CONN_SUP_TIMEOUT_NONE
	return sd.NrfError(C.sd_ble_gap_ppcp_set(&gapConnParams))
}
