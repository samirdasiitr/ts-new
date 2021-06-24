package scylladb

import (
	"fmt"
	"time"

	. "github.com/ts-new/utils/log"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
)

// ScyllaDB handle type.
type ScyllaDBStore struct {
	cluster *gocql.ClusterConfig
}

// User by mobile number table.
var userByMobileNumber = table.Metadata {
	Name:    "ts.users_by_mobile_number",
	Columns: []string {"name", "lastname", "address1", "address2", "city_name",
		"state", "email", "mobile_number", "city_index", "id", "password"},
	PartKey: []string {"mobile_number"},
	SortKey: []string{"city_name", "state"},
}

// User by uuid table.
var userByUuid = table.Metadata {
	Name:    "ts.users_by_uuid",
	Columns: []string {"name", "lastname", "address1", "address2", "city_name",
		"state", "email", "mobile_number", "city_index", "id", "password"},
	PartKey: []string {"uuid"},
	SortKey: []string{"city_name", "state"},
}

// User DB type
type UserInDB struct {
	Name         string
	LastName     string
	Address1     string
	Address2     string
	CityName     string
	State        string
	Email        string
	MobileNumber string
	CityIndex    int
	id           gocql.UUID
	Password     string
}

// Config for ScyllaDB client.
type ScyllaDBClientConfig struct {
	Endpoint string
	Timeout time.Duration
	NumRetries int
	Compressor string
	ReplicationFactor int
	CQLVersion string
	ProtoVersion int
}

// Returns default config for client.
func GetDefaultScyllaDBClientConfig() *ScyllaDBClientConfig {
	return &ScyllaDBClientConfig {
		Endpoint: "localhost",
		Timeout: time.Duration(5 * time.Second),
		NumRetries: 5,
		ReplicationFactor: 1,
		CQLVersion: "3.0.0",
		ProtoVersion: 0,
		Compressor: "",
	}
}

// Initialize user store for scylladb.
func (db *ScyllaDBStore) InitUserStore(config *ScyllaDBClientConfig) error {
	// Create cluster
	db.cluster = gocql.NewCluster(config.Endpoint)
	db.cluster.ProtoVersion = config.ProtoVersion
	db.cluster.CQLVersion = config.CQLVersion
	db.cluster.Timeout = config.Timeout
	db.cluster.RetryPolicy = &gocql.SimpleRetryPolicy{
		NumRetries: config.NumRetries,
	}

	session, err := gocqlx.WrapSession(db.cluster.CreateSession())
	if err != nil {
		Log.ERROR("Failed to create session, err=%s", err.Error())
		return err
	}
	defer session.Close()

	// Create KeySpace for User.
	stmt := fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS ts` +
		` WITH replication = {'class': 'SimpleStrategy', ` +
		`'replication_factor': %d}`, config.ReplicationFactor)
	err = session.ExecStmt(stmt)
	if err != nil {
		err = fmt.Errorf("Failed to create ts keyspace, err=%s", err.Error())
		Log.ERROR(err.Error())
		return err
	}

	// Create tables
	stmt = `CREATE TABLE IF NOT EXISTS ts.user_by_uuid (
		id uuid,
		name text, lastname text, address1 text, address2 text,
		city_name text, state text, email text, mobile_number text,
		city_index int, password string,
		PRIMARY KEY (uuid, city_name, state))`
	err = session.ExecStmt(stmt)
	if err != nil {
		err = fmt.Errorf("Failed to create user by uuid table, err=%s", err.Error())
		Log.ERROR(err.Error())
		return err
	}

	stmt = `CREATE TABLE IF NOT EXISTS ts.user_by_mobile_number(
		id uuid,
		name text, lastname text, address1 text, address2 text,
		city_name text, state text, email text, mobile_number text,
		city_index int, password string,
		PRIMARY KEY (mobile_number, city_name, state))`
	err = session.ExecStmt(stmt)
	if err != nil {
		err = fmt.Errorf("Failed to create user by uuid table, err=%s", err.Error())
		Log.ERROR(err.Error())
		return err
	}

	return nil
}
