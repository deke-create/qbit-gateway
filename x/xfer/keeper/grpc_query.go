package keeper

import (
	"github.com/deke-create/qbit-gateway/x/xfer/types"
)

var _ types.QueryServer = Keeper{}
