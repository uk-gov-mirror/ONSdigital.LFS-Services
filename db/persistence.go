package db

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"services/config"
	"services/db/postgres"
	"services/types"
	"sync"
)

var cachedConnection Persistence
var connectionMux = &sync.Mutex{}

func GetDefaultPersistenceImpl() (Persistence, error) {
	connectionMux.Lock()
	defer connectionMux.Unlock()

	if cachedConnection != nil {
		log.Info().
			Str("databaseName", config.Config.Database.Database).
			Msg("Returning cached database connection")
		return cachedConnection, nil
	}

	cachedConnection = &postgres.Postgres{nil}

	if err := cachedConnection.Connect(); err != nil {
		log.Info().
			Err(err).
			Str("databaseName", config.Config.Database.Database).
			Msg("Cannot connect to database")
		cachedConnection = nil
		return nil, fmt.Errorf("cannot connect to database")
	}

	return cachedConnection, nil
}

type Persistence interface {
	Connect() error
	Close()

	// Dashboard
	GetAnnualBatches() ([]types.Dashboard, error)
	GetQuarterlyBatches() ([]types.Dashboard, error)
	GetMonthlyBatches() ([]types.Dashboard, error)

	// Import
	PersistSurvey(vo types.SurveyVO) error
	PersistVariableDefinitions([]types.Header, types.FileSource) error
	PersistDVChanges(definitions []types.VariableDefinitions) error
	PersistAddresses(headers []string, rows [][]string, status *types.WSMessage) error

	// User
	GetUserID(user string) (types.UserCredentials, error)

	// New Batch
	MonthlyBatchExists(month, year int) bool
	AnnualBatchExists(year int) bool
	QuarterBatchExists(quarter, year int) bool

	ValidateMonthsForQuarterlyBatch(period, year int) ([]types.MonthlyBatch, error)
	ValidateMonthsForAnnualBatch(year int) ([]types.MonthlyBatch, error)
	ValidateQuartersForAnnualBatch(year int) ([]types.QuarterlyBatch, error)

	CreateMonthlyBatch(batch types.MonthlyBatch) error
	CreateQuarterlyBatch(batch types.QuarterlyBatch) error
	CreateAnnualBatch(batch types.AnnualBatch) error

	FindGBBatchInfo(week, year int) (types.GBBatchItem, error)
	FindNIBatchInfo(month, year int) (types.NIBatchItem, error)

	// Batch IDs
	GetIdsByYear(year types.Year) ([]types.YearID, error)
	GetIdsByQuarter(year types.Year, quarter types.Quarter) ([]types.QuarterID, error)
	GetIdsByMonth(year types.Year, quarter types.Month) ([]types.MonthID, error)

	// Audits
	GetAllAudits() ([]types.Audit, error)
	GetAuditsByYear(year types.Year) ([]types.Audit, error)
	GetAuditsByYearMonth(month types.Month, year types.Year) ([]types.Audit, error)
	GetAuditsByYearWeek(week types.Week, year types.Year) ([]types.Audit, error)

	// Variable Definitions
	GetAllDefinitions() ([]types.VariableDefinitionsQuery, error)
	PersistDefinitions(d types.VariableDefinitions) error
	GetDefinitionsForVariable(variable string) ([]types.VariableDefinitionsQuery, error)
	GetAllGBDefinitions() ([]types.VariableDefinitions, error)
	GetAllNIDefinitions() ([]types.VariableDefinitions, error)

	// Common SQL statements
	DeleteFrom(table string) error
	CountRows(table string) (int, error)

	// Value Labels
	GetAllValueLabels() ([]types.ValueLabelsView, error)
	GetLabelsForValue(variable string) ([]types.ValueLabelsView, error)
	PersistValues(types.ValueLabelsRow) error
	PersistValueLabels([]types.ValueLabelsRow) error
	PersistSavValueLabels(map[string][]types.Labels, types.FileSource) error
}
