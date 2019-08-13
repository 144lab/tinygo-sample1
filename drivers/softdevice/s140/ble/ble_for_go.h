#ifndef BLE_FOR_GO_H__
#define BLE_FOR_GO_H__

typedef struct {
  uint8_t totalConCount;
  uint8_t eventLength;
  uint8_t prepheralLinkCount;
  uint8_t centralLinkCount;
  uint8_t vsUUIDCount;
  uint32_t gattAttrTabSize;
} cfg_params_t;

extern uint32_t ErrorNotFound;
extern uint32_t ApplicationRAMBaseAddress;

uint32_t bleDefaultCfgSet(uint8_t tag, cfg_params_t* params);

#endif // BLE_FOR_GO_H__
