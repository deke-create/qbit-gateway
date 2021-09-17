package keeper

import (
	"github.com/deke-create/qbit-gateway/x/qbitgateway/types"
)

var _ types.QueryServer = Keeper{}
