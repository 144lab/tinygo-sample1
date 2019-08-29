package softdevice

/*
#define SVCALL_AS_NORMAL_FUNCTION
#include "nrf_sdm.h"
#include "nrf_soc.h"

extern uint32_t ApplicationRAMBaseAddress;

void SoftdeviceAssertHandler(uint32_t a, uint32_t b, uint32_t c);
void SetupTimer1(uint32_t us);
*/
import "C"
import (
	"time"

	"device/arm"
	"device/nrf"
)

var (
	clockConfig C.nrf_clock_lf_cfg_t = C.nrf_clock_lf_cfg_t{
		source:       C.NRF_CLOCK_LF_SRC_SYNTH,
		rc_ctiv:      0,
		rc_temp_ctiv: 0,
		accuracy:     0,
	}
)

func init() {
	C.ApplicationRAMBaseAddress = 0x2003ffbc //0x200039c0
}

// SetupTimer1 ...
func SetupTimer1(t time.Duration) {
	C.SetupTimer1(uint32(t / 1000))
	//nrf.TIMER1.INTENSET.Set(nrf.TIMER_INTENSET_COMPARE0)
	nrf.TIMER1.INTENSET.Set(
		nrf.TIMER_INTENSET_COMPARE0_Enabled << nrf.TIMER_INTENSET_COMPARE0_Pos,
	)
	nrf.TIMER1.SHORTS.Set(
		nrf.TIMER_SHORTS_COMPARE0_CLEAR_Enabled << nrf.TIMER_SHORTS_COMPARE0_CLEAR_Pos,
	)
	nrf.TIMER1.TASKS_START.Set(1)
	arm.SetPriority(nrf.IRQ_TIMER1, 15)
}

// Enable ...
func Enable() error {
	return NrfError(C.sd_softdevice_enable(&clockConfig,
		C.nrf_fault_handler_t(C.SoftdeviceAssertHandler)))
}

// TempGet unit:Celsius（℃）
func TempGet() (float32, error) {
	var t C.int32_t
	err := NrfError(C.sd_temp_get(&t))
	return float32(t) * 0.25, err
}
