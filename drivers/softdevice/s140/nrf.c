#define static
#include "nrf_sdm.h"
#include "nrf_soc.h"

extern uint32_t ApplicationRAMBaseAddress;

uint32_t ApplicationRAMBaseAddress;

void SetupTimer1(uint32_t us) {
    NRF_TIMER1->TASKS_STOP = 1;
    // Create an Event-Task shortcut to clear TIMER0 on COMPARE[0] event
    NRF_TIMER1->MODE        = TIMER_MODE_MODE_Timer;
    NRF_TIMER1->BITMODE     = (TIMER_BITMODE_BITMODE_24Bit << TIMER_BITMODE_BITMODE_Pos);
    NRF_TIMER1->PRESCALER   = 4;  // 1us resolution
    NRF_TIMER1->TASKS_CLEAR = 1;         // clear the task first to be usable for later
    NRF_TIMER1->CC[0] = us; //timeout
    NRF_TIMER1->INTENSET    = TIMER_INTENSET_COMPARE0_Enabled << TIMER_INTENSET_COMPARE0_Pos;
    /* Create an Event-Task shortcut to clear TIMER0 on COMPARE[0] event. */
    NRF_TIMER1->SHORTS      = TIMER_SHORTS_COMPARE0_CLEAR_Enabled << TIMER_SHORTS_COMPARE0_CLEAR_Pos;
}