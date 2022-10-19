// SPDX-FileCopyrightText: 2021 Henry Bubert
//
// SPDX-License-Identifier: MIT

package multiserver

import (
	"testing"

	refs "github.com/ssbc/go-ssb-refs"
	"github.com/stretchr/testify/require"
)

func TestParseTunnelAddressSimple(t *testing.T) {
	r := require.New(t)

	intermediary, err := refs.ParseFeedRef("@7MG1hyfz8SsxlIgansud4LKM57IHIw2Okw/hvOdeJWw=.ed25519")
	r.NoError(err)

	target, err := refs.ParseFeedRef("@1b9KP8znF7A4i8wnSevBSK2ZabI/Re4bYF/Vh3hXasQ=.ed25519")
	r.NoError(err)

	want := &TunnelAddress{
		Intermediary: intermediary,
		Target:       target,
	}

	input := `tunnel:@7MG1hyfz8SsxlIgansud4LKM57IHIw2Okw/hvOdeJWw=.ed25519:@1b9KP8znF7A4i8wnSevBSK2ZabI/Re4bYF/Vh3hXasQ=.ed25519~shs:1b9KP8znF7A4i8wnSevBSK2ZabI/Re4bYF/Vh3hXasQ=`

	addr, err := ParseTunnelAddress(input)
	r.NoError(err)
	r.Equal(want.String(), addr.String())
}
