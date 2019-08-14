package main

// tinygo build -target custom.json -o app.bin .
// uf2conv.py app.bin -c -b 0x26000 -f 0xADA52840 -o app.uf2

import (
	"image/color"
	"time"

	"machine"
	"device/arm"
	"device/nrf"

	"sample1/drivers/softdevice/s140"
	"sample1/drivers/softdevice/s140/ble"
	"sample1/drivers/softdevice/s140/rtt"
	//_ "github.com/tinygo-org/softdevice/usbd"

	"github.com/conejoninja/tinyfont"
	"github.com/conejoninja/tinyfont/freemono"
	"sample1/drivers/waveshare-epd/epd2in13x"
	//"tinygo.org/x/drivers/waveshare-epd/epd2in13x"
)

// ConfigTag ...
const ConfigTag = 1

var (
	debug      machine.Pin
	colorLED   [3]machine.Pin
	logger     = rtt.New()
	display    epd2in13x.Device
	eInkEnable machine.Pin
	white      = color.RGBA{0, 0, 0, 255}
	red        = color.RGBA{255, 0, 0, 255}
	black      = color.RGBA{1, 1, 1, 255}
)

func init() {
	debug = machine.Pin(8)
	debug.Configure(machine.PinConfig{Mode: machine.PinOutput})
	colorLED[0] = machine.Pin(14)
	colorLED[0].Configure(machine.PinConfig{Mode: machine.PinOutput})
	colorLED[1] = machine.Pin(13)
	colorLED[1].Configure(machine.PinConfig{Mode: machine.PinOutput})
	colorLED[2] = machine.Pin(15)
	colorLED[2].Configure(machine.PinConfig{Mode: machine.PinOutput})

	eInkEnable = machine.Pin(11)
	eInkEnable.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
		SCK:       31,
		MOSI:      29,
		MISO:      33,
	})
	display = epd2in13x.New(machine.SPI0, 30, 28, 2, 3)
	eInkEnable.Low()
	display.Configure(epd2in13x.Config{200, 200})
}

func beginDisplay() {
	display.ClearBuffer()
}

func finishDisplay() {
	eInkEnable.Low()
	display.Display()
	display.WaitUntilIdle()
	display.DeepSleep()
	eInkEnable.High()
}

func drawText(x, y int, s string, col color.RGBA) {
	tinyfont.WriteLine(&display, &freemono.Bold9pt7b, int16(x), int16(y), []byte(s), col)
}

func setColorLED(n int) {
	for i := uint8(0); i < 3; i++ {
		if (n>>i)&1 != 0 {
			colorLED[i].Low()
		} else {
			colorLED[i].High()
		}
	}
}

// Morse ...
func Morse(bits uint32) {
	Blink(500*time.Millisecond, 2)
	for i := 24; i < 32; i++ {
		if (bits>>uint(31-i))&0x1 > 0 {
			Blink(200*time.Millisecond, 1, 0, 0, 0)
		} else {
			Blink(200*time.Millisecond, 1, 1, 1, 0)
		}
	}
	Blink(500*time.Millisecond, 4, 0)
}

// Blink ...
func Blink(t time.Duration, n ...int) {
	for _, v := range n {
		setColorLED(v)
		time.Sleep(t)
	}
}

// BlinkLoop ...
func BlinkLoop(n ...int) {
	for {
		for _, v := range n {
			setColorLED(v)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

//go:export SoftdeviceAssertHandler
func assertHandler(id, pc, info uint32) {
	setColorLED(1)
	for {
	}
}

var count = 0

//go:export TIMER1_IRQHandler
func timerHandler(ptr uint32) {
	logger.Println("timer tick!")
	debug.Set(!debug.Get())
	count++
	//setColorLED(count % 8)
	if nrf.TIMER1.EVENTS_COMPARE[0].Get() != 0 {
		nrf.TIMER1.EVENTS_COMPARE[0].Set(0)
	}
}

func greeting(x,y int, s string, col color.RGBA) {
	beginDisplay()
	defer finishDisplay()
	drawText(x, y, s, col)
}

func evtCheck(adv *ble.Advertisement) error {
	defer setColorLED(0)
	b := make([]byte, 1408)
	for {
		setColorLED(0)
		n, err := ble.EvtGet(b)
		if err != nil {
			return err
		}
		if n == 0 {
			return nil
		}
		setColorLED(2)
		var frame ble.EvtFrame
		frame.UnmarshalBinary(b)
		logger.Printf("event: %X", b[:n])
		switch frame.ID {
		case 0x10: // gap-connected
			var ev ble.GapConnected
			ev.UnmarshalBinary(frame.Payload)
			logger.Println("connected:", ev.PeerAddr.Type, ev.Role, ev.AdvHandle)
		case 0x11: // gap-disconnected
			var ev ble.GapDisconnected
			ev.UnmarshalBinary(frame.Payload)
			logger.Println("disconnected:", ev.Reason)
			logger.Println("advertize start")
			if err := adv.Start(ConfigTag); err != nil {
				logger.Println(err)
				BlinkLoop(1, 1, 6, 6)
			}
		}
	}
}

func main() {
	msg := "Welcome to papyr!"
	setColorLED(1)
	logger.Println(msg)
	//greeting(10,10, msg, black)
	debug.Low()
	setColorLED(2)
	s140.SetupTimer1(1 * time.Second)
	if err := s140.Enable(); err != nil {
		logger.Println("sd_softdevice_enable failed:", err)
		BlinkLoop(1, 1, 0)
	}
	if err := ble.DefaultCfgSet(ConfigTag, nil); err != nil {
		logger.Println("sd_ble_cfg_set failed:", err)
		BlinkLoop(1, 2, 0)
	}
	if err := ble.Enable(); err != nil {
		logger.Println("sd_ble_enable failed:", err)
		BlinkLoop(1, 0, 2, 0)
	}
	if err := ble.GapDeviceNameSet("TinyGo"); err != nil {
		logger.Println("sd_ble_gap_device_name_set failed:", err)
		BlinkLoop(1, 0, 4, 0)
	}
	if err := ble.GapPpcpSet(); err != nil {
		logger.Println("sd_ble_gap_ppcp_set failed:", err)
		BlinkLoop(1, 0, 1, 0)
	}
	adv := ble.NewAdvertisement()
	options := &ble.AdvertiseOptions{
		Interval: ble.NewAdvertiseInterval(100),
	}
	advPayload := []byte("\x02\x01\x06" + "\x07\x09TinyGo")
	var rspPayload []byte //:= []byte("\x02\x01\x06" + "\x07\x09tINYgO")
	logger.Printf("adv: %p(%d)", &advPayload[0], len(advPayload))
	if err := adv.Configure(advPayload, rspPayload, options); err != nil {
		logger.Println(err)
		BlinkLoop(1, 1, 1, 0)
	}
	logger.Println("advertize start")
	if err := adv.Start(ConfigTag); err != nil {
		logger.Println(err)
		BlinkLoop(1, 1, 6, 6)
	}
	ver, err := ble.VersionGet()
	if err != nil {
		logger.Println(err)
		BlinkLoop(1, 1, 4, 4)
	}
	logger.Printf("version: %d/%d/%d", ver.VersionNumber, ver.CompanyID, ver.SubversionNumber)
	t, err := s140.TempGet()
	if err != nil {
		logger.Println("temp_get failed:", err)
		BlinkLoop(1, 1, 2, 2)
	}
	logger.Printf("temp.: %4.1f", t)
	arm.EnableIRQ(nrf.IRQ_TIMER1)
	setColorLED(0)
	for {
		if err := evtCheck(adv); err != nil {
			logger.Println(err)
		}
		arm.Asm("wfi")
	}
}
