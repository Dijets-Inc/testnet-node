// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package hierarchycodec

import (
	"testing"

	"github.com/lasthyphen/dijetsnodego/codec"
)

func TestVectors(t *testing.T) {
	for _, test := range codec.Tests {
		c := NewDefault()
		test(c, t)
	}
}

func TestMultipleTags(t *testing.T) {
	for _, test := range codec.MultipleTagsTests {
		c := New([]string{"tag1", "tag2"}, defaultMaxSliceLength)
		test(c, t)
	}
}
