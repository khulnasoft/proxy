// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Khulnasoft

package test

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Tmpdir string

func init() {
	var err error
	Tmpdir, err = os.MkdirTemp("", "khulnasoft_envoy_go_test")
	if err != nil {
		logrus.Fatal("Failed to create a temporary directory for testing")
	}
}
