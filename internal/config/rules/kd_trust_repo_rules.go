// SPDX-License-Identifier: Apache-2.0
// Copyright 2026 alibaba/open-code-review Contributors

package rules

import (
	"fmt"
	"os"
	"path/filepath"
)

// TrustRepoRulesEnv opts back in to loading the reviewed repository's own
// .opencodereview/rule.json.
//
// kd-hardening: that file is written by whoever controls the repository under
// review, and its text is handed verbatim to the reviewing model as
// instructions. Reviewing a third-party checkout would therefore let the
// repository author inject prompts. Operator-chosen layers (--rule and
// ~/.opencodereview/rule.json) are unaffected.
const TrustRepoRulesEnv = "OCR_TRUST_REPO_RULES"

func repoRulesTrusted() bool {
	return os.Getenv(TrustRepoRulesEnv) == "1"
}

func warnIgnoredProjectRule(repoDir string) {
	path := filepath.Join(repoDir, ".opencodereview", "rule.json")
	if _, err := os.Lstat(path); err == nil {
		fmt.Fprintf(os.Stderr, "[ocr] WARNING: ignoring repository rule file %s (set %s=1 to trust it)\n", path, TrustRepoRulesEnv)
	}
}
