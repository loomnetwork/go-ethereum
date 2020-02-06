// +build !usbhubs

package usbwallet

import "errors"

// NewLedgerHub creates a new hardware wallet manager for Ledger devices.
func NewLedgerHub() (*Hub, error) {
	return nil, errors.New("not supported")
}

// NewTrezorHubWithHID creates a new hardware wallet manager for Trezor devices.
func NewTrezorHubWithHID() (*Hub, error) {
	return nil, errors.New("not supported")
}

// NewTrezorHubWithWebUSB creates a new hardware wallet manager for Trezor devices with
// firmware version > 1.8.0
func NewTrezorHubWithWebUSB() (*Hub, error) {
	return nil, errors.New("not supported")
}
