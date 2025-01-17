package api

import (
	"github.com/rs/zerolog/log"
	"services/db"
	"services/types"
)

func (a AuditHandler) GetAllAudits() ([]types.Audit, error) {

	dbase, err := db.GetDefaultPersistenceImpl()
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	res, err := dbase.GetAllAudits()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a AuditHandler) GetAuditsForYear(year types.Year) ([]types.Audit, error) {
	dbase, err := db.GetDefaultPersistenceImpl()
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	res, err := dbase.GetAuditsByYear(year)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a AuditHandler) GetAuditsForWeek(week types.Week, year types.Year) ([]types.Audit, error) {
	dbase, err := db.GetDefaultPersistenceImpl()
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	res, err := dbase.GetAuditsByYearWeek(week, year)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a AuditHandler) GetAuditsForMonth(month types.Month, year types.Year) ([]types.Audit, error) {
	dbase, err := db.GetDefaultPersistenceImpl()
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	res, err := dbase.GetAuditsByYearMonth(month, year)
	if err != nil {
		return nil, err
	}

	return res, nil
}
