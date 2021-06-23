package userservice

import (
	"context"

	pb "github.com/ts-new/model"
)

// User Store Service
type UserStoreService interface {
    // Create a new user.
	Create(context.Context, *pb.User) (*pb.User, error)
    // Update an existing user.
    Update(context.Context, *pb.User, *pb.User) (*pb.User, error)
    // Delete an existing user.
    Delete(context.Context, *pb.User) error
    // Read detail of an existing user.
    Read(context.Context, *pb.User) (*pb.User, error)
}
