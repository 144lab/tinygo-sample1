#define static
#include "ble.h"
#include <stddef.h>
#include <stdint.h>
#include "ble_gap.h"
#include "ble_for_go.h"

extern uint32_t ApplicationRAMBaseAddress;
uint32_t ErrorNotFound = NRF_ERROR_NOT_FOUND;

void* memset(void* s, int c, size_t n);

void setProperties(ble_gap_adv_properties_t* prop) {
  prop->type = BLE_GAP_ADV_TYPE_CONNECTABLE_SCANNABLE_UNDIRECTED;
}

static cfg_params_t defaultParams = {
  .totalConCount = 1,
  .eventLength = 6,
  .prepheralLinkCount = 1,
  .centralLinkCount = 1,
  .vsUUIDCount = 10,
  .gattAttrTabSize = 1408,
};

uint32_t bleDefaultCfgSet(uint8_t tag, cfg_params_t* params) {
  uint32_t ret_code;
  uint32_t ram_start = ApplicationRAMBaseAddress;
  ble_cfg_t ble_cfg;
  if(params==NULL) {
    params = &defaultParams;
  }
  memset(&ble_cfg, 0, sizeof(ble_cfg));
  ble_cfg.conn_cfg.conn_cfg_tag = tag;
  ble_cfg.conn_cfg.params.gap_conn_cfg.conn_count = params->totalConCount;
  ble_cfg.conn_cfg.params.gap_conn_cfg.event_length = params->eventLength;
  ret_code = sd_ble_cfg_set(BLE_CONN_CFG_GAP, &ble_cfg, ram_start);
  if (ret_code != NRF_SUCCESS) {
    return ret_code;
  }
  memset(&ble_cfg, 0, sizeof(ble_cfg));
  ble_cfg.gap_cfg.role_count_cfg.periph_role_count = params->prepheralLinkCount;
  ret_code = sd_ble_cfg_set(BLE_GAP_CFG_ROLE_COUNT, &ble_cfg, ram_start);
  if (ret_code != NRF_SUCCESS) {
    return ret_code;
  }
  memset(&ble_cfg, 0, sizeof(ble_cfg));
  ble_cfg.common_cfg.vs_uuid_cfg.vs_uuid_count = params->vsUUIDCount;
  ret_code = sd_ble_cfg_set(BLE_COMMON_CFG_VS_UUID, &ble_cfg, ram_start);
  if (ret_code != NRF_SUCCESS) {
    return ret_code;
  }
  memset(&ble_cfg, 0x00, sizeof(ble_cfg));
  ble_cfg.gatts_cfg.attr_tab_size.attr_tab_size = params->gattAttrTabSize;
  ret_code = sd_ble_cfg_set(BLE_GATTS_CFG_ATTR_TAB_SIZE, &ble_cfg, ram_start);
  if (ret_code != NRF_SUCCESS) {
    return ret_code;
  }
    memset(&ble_cfg, 0x00, sizeof(ble_cfg));
    ble_cfg.gatts_cfg.service_changed.service_changed = 0;
    ret_code = sd_ble_cfg_set(BLE_GATTS_CFG_SERVICE_CHANGED, &ble_cfg, ram_start);
    if (ret_code != NRF_SUCCESS)
    {
    return ret_code;
    }
      return NRF_SUCCESS;
}
