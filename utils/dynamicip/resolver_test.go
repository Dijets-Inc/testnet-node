// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package dynamicip

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewResolver(t *testing.T) {
	type test struct {
		service      string
		validService bool
	}
	tests := []test{
		{
			service:      OpenDNSName,
			validService: true,
		},
		{
			service:      IFConfigName,
			validService: true,
		},
		{
			service:      IFConfigCoName,
			validService: true,
		},
		{
			service:      IFConfigMeName,
			validService: true,
		},
		{
			service:      strings.ToUpper(IFConfigMeName),
			validService: true,
		},
		{
			service:      "not a valid resolution service name",
			validService: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.service, func(t *testing.T) {
			require := require.New(t)
			_, err := NewResolver(tt.service)
			if tt.validService {
				require.NoError(err)
			} else {
				require.Error(err)
			}
		})
	}
}
