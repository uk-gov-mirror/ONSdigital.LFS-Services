package filter

import (
	"github.com/rs/zerolog/log"
	conf "services/config"
	"services/types"
)

type Filter interface {
	DropColumn(string) bool
	AddVariables(*types.SavImportData) error
	SkipRowsFilter(*types.SavImportData) error
	RenameColumns(string) (string, bool)
}

var dropColumns = conf.Config.DropColumns.Survey
var renameColumns map[string]string

/*
Load the columns to rename from the configuration
*/
func init() {
	cols := conf.Config.Rename.Survey
	renameColumns = make(map[string]string, len(cols))

	for _, v := range cols {
		renameColumns[v.From] = v.To
	}
}

type BaseFilter struct{}

/*
Generic drop columns functionality - based on the name of columns to drop in the configuration file
*/
func (bf BaseFilter) DropColumn(name string) bool {
	for _, j := range dropColumns.ColumnNames {
		if j == name {
			log.Debug().Str("columnName", name).Msg("Dropping column")
			return true
		}
	}
	return false
}

/*
Generic rename columns functionality - based on the name of columns to drop in the configuration file
*/
func (bf BaseFilter) RenameColumns(column string) (string, bool) {
	item, ok := renameColumns[column]
	if ok {
		log.Debug().Str("from", column).Str("to", item).Msg("Renaming column")
		return item, true
	}
	return "", false
}
