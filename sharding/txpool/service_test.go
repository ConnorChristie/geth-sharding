package txpool

import "github.com/prysmaticlabs/geth-sharding/sharding/types"

// Verifies that TXPool implements the Service interface.
var _ = types.Service(&TXPool{})
