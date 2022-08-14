package bluechipexp_test

import (
	"testing"

	keepertest "github.com/mastervectormaster/bluechipExp/testutil/keeper"
	"github.com/mastervectormaster/bluechipExp/testutil/nullify"
	"github.com/mastervectormaster/bluechipExp/x/bluechipexp"
	"github.com/mastervectormaster/bluechipExp/x/bluechipexp/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BluechipexpKeeper(t)
	bluechipexp.InitGenesis(ctx, *k, genesisState)
	got := bluechipexp.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
