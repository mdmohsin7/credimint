package keeper

import (
	"context"

	"credimint/x/credimint/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserCredibilityAll(c context.Context, req *types.QueryAllUserCredibilityRequest) (*types.QueryAllUserCredibilityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userCredibilitys []types.UserCredibility
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	userCredibilityStore := prefix.NewStore(store, types.KeyPrefix(types.UserCredibilityKey))

	pageRes, err := query.Paginate(userCredibilityStore, req.Pagination, func(key []byte, value []byte) error {
		var userCredibility types.UserCredibility
		if err := k.cdc.Unmarshal(value, &userCredibility); err != nil {
			return err
		}

		userCredibilitys = append(userCredibilitys, userCredibility)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserCredibilityResponse{UserCredibility: userCredibilitys, Pagination: pageRes}, nil
}

func (k Keeper) UserCredibility(c context.Context, req *types.QueryGetUserCredibilityRequest) (*types.QueryGetUserCredibilityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	userCredibility, found := k.GetUserCredibility(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetUserCredibilityResponse{UserCredibility: userCredibility}, nil
}
