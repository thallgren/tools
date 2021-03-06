// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmdtest

import (
	"fmt"
	"testing"

	"golang.org/x/tools/internal/lsp/tests"
	"golang.org/x/tools/internal/span"
)

func (r *runner) SuggestedFix(t *testing.T, spn span.Span, actionKinds []string) {
	uri := spn.URI()
	filename := uri.Filename()
	args := []string{"fix", "-a", fmt.Sprintf("%s", spn)}
	args = append(args, actionKinds...)
	got, _ := r.NormalizeGoplsCmd(t, args...)
	want := string(r.data.Golden("suggestedfix_"+tests.SpanName(spn), filename, func() ([]byte, error) {
		return []byte(got), nil
	}))
	if want != got {
		t.Errorf("suggested fixes failed for %s, expected:\n%v\ngot:\n%v", filename, want, got)
	}
}
