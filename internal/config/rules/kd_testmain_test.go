// SPDX-License-Identifier: Apache-2.0
// Copyright 2026 alibaba/open-code-review Contributors

package rules

import (
	"os"
	"testing"
)

// Upstream tests in this package assume the repository rule.json is always
// loaded. Keep that assumption for them; the kd tests override it per test.
func TestMain(m *testing.M) {
	os.Setenv(TrustRepoRulesEnv, "1")
	os.Exit(m.Run())
}
