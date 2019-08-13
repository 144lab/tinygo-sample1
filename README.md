# tinygo-sample

# target

- papyr(nRF52840 + ePaper)
- with Adafruit-nRF52-Bootloader

# dependencies

- nRF52-SDK 15.3 or later
- docker(for build)
- uf2conv.py(convert binary for Adafruit-nRF52-Bootloader)

# build & deploy

1. make
2. target reset button double click.
3. app.uf2 drag and drop to NRF52BOOT volume.

# drivers folder

- softdevice/s140: stub for nRF52840 with softdevice
    - ble: BLE API
    - rtt: SEGGER_RTT logger
    - usbd: USB daemon API(WIP)
- waveshare-epd/epd2in13x: driver for papyr ePaper display
