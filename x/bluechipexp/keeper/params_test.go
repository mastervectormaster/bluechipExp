package keeper_test

import (
	"testing"

	testkeeper "github.com/mastervectormaster/bluechipExp/testutil/keeper"
	"github.com/mastervectormaster/bluechipExp/x/bluechipexp/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BluechipexpKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
