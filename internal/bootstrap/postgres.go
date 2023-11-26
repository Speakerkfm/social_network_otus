package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/bxcodec/dbresolver/v2"
	_ "github.com/lib/pq"
)

func CreatePostgres() (dbresolver.DB, error) {
	// TODO конфиг когда-нибудь...
	hostMaster := "localhost"
	portMaster := 5432
	hostSyncReplica := "localhost"
	portSyncReplica := 5432
	hostAsyncReplica := "localhost"
	portAsyncReplica := 5432
	user := "social-network-user"
	password := "social-network-password"
	dbname := "social_network_otus"
	rwPrimary := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostMaster, portMaster, user, password, dbname)
	readOnlySyncReplica := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostSyncReplica, portSyncReplica, user, password, dbname)
	readOnlyAsyncReplica := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostAsyncReplica, portAsyncReplica, user, password, dbname)

	// open database for primary
	dbPrimary, err := sql.Open("postgres", rwPrimary)
	if err != nil {
		return nil, err
	}

	// open database for sync replica
	dbReadOnlySyncReplica, err := sql.Open("postgres", readOnlySyncReplica)
	if err != nil {
		return nil, err
	}

	// open database for async replica
	dbReadOnlyAsyncReplica, err := sql.Open("postgres", readOnlyAsyncReplica)
	if err != nil {
		return nil, err
	}

	connectionDB := dbresolver.New(
		dbresolver.WithPrimaryDBs(dbPrimary),
		dbresolver.WithReplicaDBs(dbReadOnlySyncReplica, dbReadOnlyAsyncReplica),
		dbresolver.WithLoadBalancer(dbresolver.RoundRobinLB))

	return connectionDB, nil
}
