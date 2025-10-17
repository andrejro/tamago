// THSM EP1 support for tamago/arm
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

// Package thsm_ep1 provides hardware initialization, automatically on import,
// for the Tillitis HSM EP1 board.
//
// This package is only meant to be used with `GOOS=tamago GOARCH=arm` as
// supported by the TamaGo framework for bare metal Go, see
// https://github.com/usbarmory/tamago.
package thsm_ep1

import (
	_ "unsafe"

	"github.com/usbarmory/tamago/soc/nxp/imx6ul"
)

// Peripheral instances
var (
	UART1 = imx6ul.UART1
	UART3 = imx6ul.UART3
	USB1 = imx6ul.USB1
	USB2 = imx6ul.USB2
)

// Init takes care of the lower level initialization triggered early in runtime
// setup (post World start).
//
//go:linkname Init runtime.hwinit1
func Init() {
	imx6ul.Init()

	// ECP5 expects a 500k baudrate
	imx6ul.UART3.Baudrate = 500000
	// ECP5 uses CTS
	imx6ul.UART3.Flow = true
	imx6ul.UART3.Init()

	// initialize console
	imx6ul.UART1.Init()
}
