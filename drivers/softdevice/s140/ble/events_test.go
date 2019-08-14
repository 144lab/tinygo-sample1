package ble

import (
	"fmt"
	"encoding/hex"
)

func ExampleGapConnected() {
	b, err := hex.DecodeString(
	"10002C0000000000044B6EC78F7B5001180018000000480000000000E08900200B0000000000000000000000",
	)
	if err != nil {
		panic(err)
	}
	var frame EvtFrame
	frame.UnmarshalBinary(b)
	var gapCon GapConnected
	gapCon.UnmarshalBinary(frame.Payload)
	fmt.Printf("%v", gapCon)
	// output:
	// {{4 [75 110 199 143 123 80]} 1 {24 24 0 72} 0 [0 0 0] 536906208 11 [0 0] 0 0 [0 0]}
}