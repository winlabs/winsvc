// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package eventlog_test

import (
	"github.com/winlabs/winsvc/eventlog"
	"testing"
)

func TestLog(t *testing.T) {
	const name = "mylog"
	const supports = eventlog.Error | eventlog.Warning | eventlog.Info
	err := eventlog.InstallAsEventCreate(name, supports)
	if err != nil {
		t.Fatalf("Install failed: %s", err)
	}
	l, err := eventlog.Open(name)
	if err != nil {
		t.Fatalf("Open failed: %s", err)
	}
	defer l.Close()
	err = l.Info(1, "info")
	if err != nil {
		t.Fatalf("Info failed: %s", err)
	}
	err = l.Warning(2, "warning")
	if err != nil {
		t.Fatalf("Warning failed: %s", err)
	}
	err = l.Error(3, "error")
	if err != nil {
		t.Fatalf("Error failed: %s", err)
	}
	err = eventlog.Remove(name)
	if err != nil {
		t.Fatalf("Remove failed: %s", err)
	}
}
