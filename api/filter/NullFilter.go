package filter

import (
	"services/dataset"
	"services/types"
)

type NullFilter struct{}

func (n NullFilter) AddVariables() (int, error)                 { return 0, nil }
func (n NullFilter) SkipRow(row map[string]interface{}) bool    { return false }
func (n NullFilter) RenameColumns(column string) (string, bool) { return column, false }
func (n NullFilter) DropColumn(name string) bool                { return false }
func NewNullFilter(dataset *dataset.Dataset) types.Filter       { return NullFilter{} }
