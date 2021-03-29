package database

import (
	"testing"

	"github.com/grafana/grafana/pkg/services/sqlstore"
	"github.com/grafana/grafana/pkg/services/sqlstore/migrator"
)

// accessControlStoreTestImpl is a test store implementation which additionally executes a database migrations
type accessControlStoreTestImpl struct {
	AccessControlStore
}

func (ac *accessControlStoreTestImpl) AddMigration(mg *migrator.Migrator) {
	AddAccessControlMigrations(mg)
}

func setupTestEnv(t testing.TB) *accessControlStoreTestImpl {
	t.Helper()

	sqlStore := sqlstore.InitTestDB(t)
	store := accessControlStoreTestImpl{
		AccessControlStore: AccessControlStore{
			SQLStore: sqlStore,
		},
	}

	return &store
}
