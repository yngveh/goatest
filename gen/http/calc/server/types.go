// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen github.com/yngveh/goatest/design

package server

import (
	calcsvc "github.com/yngveh/goatest/gen/calc"
)

// NewAddAddPayload builds a calc service add endpoint payload.
func NewAddAddPayload(a int, b int) *calcsvc.AddPayload {
	return &calcsvc.AddPayload{
		A: a,
		B: b,
	}
}
