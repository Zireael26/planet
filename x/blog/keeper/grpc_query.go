package keeper

import (
	"github.com/zireael26/planet/x/blog/types"
)

var _ types.QueryServer = Keeper{}
