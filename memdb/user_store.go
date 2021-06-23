package memdb

import (
	"context"

	pb "github.com/ts-new/model"
	. "github.com/ts-new/user_service"
	. "github.com/ts-new/utils/log"

	db "github.com/boltdb/bolt"
)

// In memory user store.
type UserMemDBStore struct {
	db *db.DB
}

// Creates a new user record.
func (umds *UserMemDBStore) Create(ctx context.Context, in *pb.User) (
	*pb.User, error) {
}

// Updates an existing user.
func (umds *UserMemDBStore) Update(ctx context.Context, org *pb.User,
	in *pb.User) (*pb.User, error) {
}

// Delete an existing user.
func (umds *UserMemDBStore) Delete(ctx context.Context, in *pb.User) error {
}

// Reads an existing user.
func (umds *UserMemDBStore) Read(ctx context.Context, in *pb.User) (
    *pb.User, error) {
}

func InitUserMemDB(_ interface{}) (UserStoreService, error) {
	var err error
	umds := &UserMemDBStore{}
	umds.db, err = db.Open("/tmp/user_in_mem.db", 0600, nil)
	if err != nil {
		Log.FATAL("Failed to open user bolt db file, err=%s", err.Error())
	}

	return umds, nil
}

var _ UserStoreService = &UserMemDBStore{}
