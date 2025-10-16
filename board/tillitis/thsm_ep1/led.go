// THSM EP1 support for tamago/arm
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package thsm_ep1

import (
	"errors"
	"strings"

	"github.com/usbarmory/tamago/soc/nxp/gpio"
	"github.com/usbarmory/tamago/soc/nxp/imx6ul"
	"github.com/usbarmory/tamago/soc/nxp/iomuxc"
)

// LED configuration constants
//
// On the USB armory Mk II the following LEDs are connected:
//   - pad CSI_DATA00, GPIO4_IO21: white
//   - pad CSI_DATA01, GPIO4_IO22: blue
//
// On the USB armory Mk II LAN the RJ45 connector LEDs can be controlled
// through the Ethernet PHY.
const (
	GPIO_MODE = 5

	// GPIO number
	RED = 11
	// mux control
	IOMUXC_SW_MUX_CTL_PAD_ENET2_TX_DATA00 = 0x020e00f0
	// pad control
	IOMUXC_SW_PAD_CTL_PAD_ENET2_TX_DATA00 = 0x020e037c

	// GPIO number
	BLUE = 12
	// mux control
	IOMUXC_SW_MUX_CTL_PAD_ENET2_TX_DATA01 = 0x020e00f4
	// pad control
	IOMUXC_SW_PAD_CTL_PAD_ENET2_TX_DATA01 = 0x020e0380

	// GPIO number
	GREEN = 13
	// mux control
	IOMUXC_SW_MUX_CTL_PAD_ENET2_TX_EN = 0x020e00f8
	// pad control
	IOMUXC_SW_PAD_CTL_PAD_ENET2_TX_EN = 0x020e0384
)

var (
	red *gpio.Pin
	blue *gpio.Pin
	green *gpio.Pin
)

func init() {
	var err error

	ctl := uint32((1 << iomuxc.SW_PAD_CTL_PKE) |
		(iomuxc.SW_PAD_CTL_SPEED_100MHZ << iomuxc.SW_PAD_CTL_SPEED) |
		(iomuxc.SW_PAD_CTL_DSE_2_R0_6 << iomuxc.SW_PAD_CTL_DSE))

	if red, err = imx6ul.GPIO2.Init(RED); err != nil {
		panic(err)
	}

	red.Out()

	p := iomuxc.Init(
		IOMUXC_SW_MUX_CTL_PAD_ENET2_TX_DATA00,
		IOMUXC_SW_PAD_CTL_PAD_ENET2_TX_DATA00,
		GPIO_MODE)
	p.Ctl(ctl)

	if blue, err = imx6ul.GPIO2.Init(BLUE); err != nil {
		panic(err)
	}

	blue.Out()

	p = iomuxc.Init(
		IOMUXC_SW_MUX_CTL_PAD_ENET2_TX_DATA01,
		IOMUXC_SW_PAD_CTL_PAD_ENET2_TX_DATA01,
		GPIO_MODE)
	p.Ctl(ctl)

	if green, err = imx6ul.GPIO2.Init(GREEN); err != nil {
		panic(err)
	}

	green.Out()

	p = iomuxc.Init(
		IOMUXC_SW_MUX_CTL_PAD_ENET2_TX_EN,
		IOMUXC_SW_PAD_CTL_PAD_ENET2_TX_EN,
		GPIO_MODE)
	p.Ctl(ctl)
}

// LED turns on/off an LED by name.
func LED(name string, on bool) (err error) {
	var led *gpio.Pin

	switch {
	case strings.EqualFold(name, "red"):
		led = red
	case strings.EqualFold(name, "blue"):
		led = blue
	case strings.EqualFold(name, "green"):
		led = green
	default:
		return errors.New("invalid LED")
	}

	if led != nil {
		if on {
			led.Low()
		} else {
			led.High()
		}
	}

	return
}
