package usbd

/*
#define NRF52840_XXAA
#include "nrfx_usbd.h"
*/
import "C"

// Enable ...
func Enable() {
	C.nrfx_usbd_enable()
}
