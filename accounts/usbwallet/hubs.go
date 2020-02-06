// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// +build usbhubs

package usbwallet

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
)

// NewLedgerHub creates a new hardware wallet manager for Ledger devices.
func NewLedgerHub() (*Hub, error) {
	return newHub(LedgerScheme, 0x2c97, []uint16{
		// Original product IDs
		0x0000, /* Ledger Blue */
		0x0001, /* Ledger Nano S */
		0x0004, /* Ledger Nano X */

		// Upcoming product IDs: https://www.ledger.com/2019/05/17/windows-10-update-sunsetting-u2f-tunnel-transport-for-ledger-devices/
		0x0015, /* HID + U2F + WebUSB Ledger Blue */
		0x1015, /* HID + U2F + WebUSB Ledger Nano S */
		0x4015, /* HID + U2F + WebUSB Ledger Nano X */
		0x0011, /* HID + WebUSB Ledger Blue */
		0x1011, /* HID + WebUSB Ledger Nano S */
		0x4011, /* HID + WebUSB Ledger Nano X */
	}, 0xffa0, 0, newLedgerDriver)
}

// NewTrezorHubWithHID creates a new hardware wallet manager for Trezor devices.
func NewTrezorHubWithHID() (*Hub, error) {
	return newHub(TrezorScheme, 0x534c, []uint16{0x0001 /* Trezor HID */}, 0xff00, 0, newTrezorDriver)
}

// NewTrezorHubWithWebUSB creates a new hardware wallet manager for Trezor devices with
// firmware version > 1.8.0
func NewTrezorHubWithWebUSB() (*Hub, error) {
	return newHub(TrezorScheme, 0x1209, []uint16{0x53c1 /* Trezor WebUSB */}, 0xffff /* No usage id on webusb, don't match unset (0) */, 0, newTrezorDriver)
}
