// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package wgengine

import (
	"testing"
	"time"
)

func TestWatchdog(t *testing.T) {
	t.Parallel()

	var maxWaitMultiple time.Duration = 1
	// if runtime.GOOS == "darwin" {
	// 	// Work around slow close syscalls on Big Sur with content filter Network Extensions installed.
	// 	// See https://github.com/tailscale/tailscale/issues/1598.
	// 	maxWaitMultiple = 15
	// }

	t.Run("default watchdog does not fire", func(t *testing.T) {
		t.Parallel()
		e, err := NewFakeUserspaceEngine(t.Logf, 0)
		if err != nil {
			t.Fatal(err)
		}

		e = NewWatchdog(e)
		e.(*watchdogEngine).maxWait = maxWaitMultiple * 150 * time.Millisecond
		e.(*watchdogEngine).logf = t.Logf
		e.(*watchdogEngine).fatalf = t.Fatalf

		e.RequestStatus()
		e.RequestStatus()
		e.RequestStatus()
		e.Close()
	})
}
