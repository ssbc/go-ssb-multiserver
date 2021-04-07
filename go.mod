module go.mindeco.de/ssb-multiserver

go 1.14

require (
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	go.cryptoscope.co/netwrap v0.1.1
	go.cryptoscope.co/secretstream v1.2.2
	go.mindeco.de/ssb-refs v0.0.0-20200615115708-7b8437f55cef
)

// We need our internal/extra25519 since agl pulled his repo recently.
// Issue: https://github.com/cryptoscope/ssb/issues/44
// Ours uses a fork of x/crypto where edwards25519 is not an internal package,
// This seemed like the easiest change to port agl's extra25519 to use x/crypto
// Background: https://github.com/agl/ed25519/issues/27#issuecomment-591073699
// The branch in use: https://github.com/cryptix/golang_x_crypto/tree/non-internal-edwards
replace golang.org/x/crypto => github.com/cryptix/golang_x_crypto v0.0.0-20200303113948-2939d6771b24
