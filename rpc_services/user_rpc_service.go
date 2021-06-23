package userrpcservice

import (
	"context"

	pb "github.com/ts-new/model"
	. "github.com/ts-new/utils/log"
	. "github.com/ts-new/user_service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// User RPC Service.
type UserRpcService struct {
	pb.UnimplementedUserServiceServer
}

var internalError = status.Error(codes.Internal, "Contact Administrator")

// Register a new user.
func (us *UserRpcService) Register(ctx context.Context,
	in *pb.User) (*pb.User, error) {

	Log.INFO("Got a new user registeration request: %v", in)

	// Get store from context, it is added by middleware.
	store := ctx.Value("user_store").(UserStoreService)
	if store == nil {
		Log.ERROR("Failed to read user store from context")
		return nil, internalError
	}

	// Create the user in store.
	user, err := store.Create(ctx, in)
	if err != nil {
		Log.ERROR("Failed to create user in store, err=%s", err.Error())
		return nil, internalError
	}

	Log.DEBUG("Created user with %v", user)
	return user, nil
}

// Change password for an user.
func (us *UserRpcService) ChangePassword(ctx context.Context,
	in *pb.User) (*pb.User, error) {

	Log.INFO("Updating password for user:%v", in)

	store := ctx.Value("user_store").(UserStoreService)
	if store == nil {
		Log.ERROR("Failed to read user store from context")
		return nil, internalError
	}

	// Read the user details first.
	user, err := store.Read(ctx, in)
	if err != nil {
		Log.ERROR("Failed to read user from store, err=%s", err.Error())
		return nil, internalError
	}

	outUser, err := store.Update(ctx, user, in)
	if err != nil {
		Log.ERROR("Failed to update user in store, err=%s", err.Error())
		return nil, internalError
	}

	return outUser, nil
}

// Update user info for an user.
func (us *UserRpcService) UserInfoUpdate(ctx context.Context,
	in *pb.User) (*pb.User, error) {

	Log.INFO("Updating info for user: %v", in)

	store := ctx.Value("user_store").(UserStoreService)
	if store == nil {
		Log.ERROR("Failed to read user store from context")
		return nil, internalError
	}

	// Read the user details first.
	user, err := store.Read(ctx, in)
	if err != nil {
		Log.ERROR("Failed to read user from store, err=%s", err.Error())
		return nil, internalError
	}

	outUser, err := store.Update(ctx, user, in)
	if err != nil {
		Log.ERROR("Failed to update user in store, err=%s", err.Error())
		return nil, internalError
	}

	return outUser, nil
}

// Delete an user.
func (us *UserRpcService) UserDelete(ctx context.Context,
	in *pb.User) (*pb.User, error) {

	Log.INFO("Deleting user:%v", in)

	store := ctx.Value("user_store").(UserStoreService)
	if store == nil {
		Log.ERROR("Failed to read user store from context")
		return nil, internalError
	}

	// Read the user details first.
	user, err := store.Read(ctx, in)
	if err != nil {
		Log.ERROR("Failed to read user from store, err=%s", err.Error())
		return nil, internalError
	}

	err = store.Delete(ctx, user)
	if err != nil {
		Log.ERROR("Failed to delete  user in store, err=%s", err.Error())
		return nil, internalError
	}

	return user, nil
}

// Get user details.
func (us *UserRpcService) GetUser(ctx context.Context,
	in *pb.User) (*pb.User, error) {

	Log.INFO("Fetching user: %v", in)

	store := ctx.Value("user_store").(UserStoreService)
	if store == nil {
		Log.ERROR("Failed to read user store from context")
		return nil, internalError
	}

	// Read the user details first.
	user, err := store.Read(ctx, in)
	if err != nil {
		Log.ERROR("Failed to read user from store, err=%s", err.Error())
		return nil, internalError
	}

	return user, nil
}
