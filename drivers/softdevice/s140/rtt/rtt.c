/*
#include <stdint.h>
#include <stddef.h>

#define SVCALL_AS_NORMAL_FUNCTION
#include "nrf_nvic.h"

nrf_nvic_state_t nrf_nvic_state;

void app_util_critical_region_enter(uint8_t *p_nested)
{
    (void) sd_nvic_critical_region_enter(p_nested);
}

void app_util_critical_region_exit(uint8_t nested)
{
    (void) sd_nvic_critical_region_exit(nested);
}
*/
