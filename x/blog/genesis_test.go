package blog_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zireael26/planet/testutil/keeper"
	"github.com/zireael26/planet/testutil/nullify"
	"github.com/zireael26/planet/x/blog"
	"github.com/zireael26/planet/x/blog/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		PostList: []types.Post{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PostCount: 2,
		SentPostList: []types.SentPost{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SentPostCount: 2,
		TimedoutPostList: []types.TimedoutPost{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TimedoutPostCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlogKeeper(t)
	blog.InitGenesis(ctx, *k, genesisState)
	got := blog.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.PostList, got.PostList)
	require.Equal(t, genesisState.PostCount, got.PostCount)
	require.ElementsMatch(t, genesisState.SentPostList, got.SentPostList)
	require.Equal(t, genesisState.SentPostCount, got.SentPostCount)
	require.ElementsMatch(t, genesisState.TimedoutPostList, got.TimedoutPostList)
	require.Equal(t, genesisState.TimedoutPostCount, got.TimedoutPostCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
