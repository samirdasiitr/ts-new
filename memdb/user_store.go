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

// Data type for User as stored in mem db.
type UserInDB struct {
	Name string
	LastName string
	Address1 string
	Address2 string
	CityName string
	State string
	Email string
	MobileNumber string
	CityIndex int
	Uuid UUID
	Password string
}

// Converts proto based user record to DB style.
func ToDB(in *pb.User) *UserInDB {
	return &UserInDB {
		Name: in.GetName(),
		LastName: in.GetLastName(),
		Address1: in.GetAddress1(),
		Address2: in.GetAddress2(),
		CityName: in.GetCityName(),
		State: in.GetState(),
		Email: in.GetEmail(),
		MobileNumber: in.GetMobileNumber(),
		CityIndex: in.GetCityIndex(),
		Uuid: in.GetUuid(),
		Password: in.GetPassword(),
	}
}

// Converts DB style to protobuf record.
func FromDB(in *UserInDB) *pb.User {
	return &pb.User{
		Name: in.Name,
		LastName: in.LastName,
		Address1: in.Address1,
		Address2: in.Address2,
		CityName: in.CityName,
		State: in.State,
		Email: in.Email,
		MobileNumber: in.MobileNumber,
		CityIndex: in.CityIndex,
		Uuid: in.Uuid,
		Password: in.Password,
	}
}

// Creates a new user record.
func (umds *UserMemDBStore) Create(ctx context.Context, in *pb.User) (
	*pb.User, error) {

	dbUser := ToDB(in)
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
