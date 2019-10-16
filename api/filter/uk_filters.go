package filter

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"reflect"
	"strconv"
	"time"
)

type UKFilter struct {
	BaseFilter
}

func (sf UKFilter) findLocation(headers []string, column string) (int, error) {
	for i, j := range headers {
		if j == column {
			return i, nil
		}
	}
	return 0, fmt.Errorf("column %s not found in findLoaction()", column)
}

func (sf UKFilter) addHSerial() error {
	column, err := sf.dataset.AddColumn("HSERIAL", reflect.Int64)
	if err != nil {
		return err
	}

	startAllrows := time.Now()
	header, items := sf.dataset.GetAllRows()
	log.Debug().
		TimeDiff("elapsedTime", time.Now(), startAllrows).
		Msg("Get all rows")

	// get indexes of items we are interested in for the calculation
	quotaInx, err := sf.findLocation(header, "QUOTA")
	if err != nil {
		return err
	}
	weekInx, err := sf.findLocation(header, "WEEK")
	if err != nil {
		return err
	}
	w1yrInx, err := sf.findLocation(header, "W1YR")
	if err != nil {
		return err
	}
	qrtrInx, err := sf.findLocation(header, "QRTR")
	if err != nil {
		return err
	}
	addrInx, err := sf.findLocation(header, "ADD")
	if err != nil {
		return err
	}
	wavfndInx, err := sf.findLocation(header, "WAVFND")
	if err != nil {
		return err
	}
	hhldInx, err := sf.findLocation(header, "HHLD")
	if err != nil {
		return err
	}

	for i := range column.Rows {
		row := items[i]

		quota, err := strconv.ParseFloat(row[quotaInx], 64)
		if err != nil {
			return err
		}

		week, err := strconv.ParseFloat(row[weekInx], 64)
		if err != nil {
			return err
		}

		w1yr, err := strconv.ParseFloat(row[w1yrInx], 64)
		if err != nil {
			return err
		}

		qrtr, err := strconv.ParseFloat(row[qrtrInx], 64)
		if err != nil {
			return err
		}

		addr, err := strconv.ParseFloat(row[addrInx], 64)
		if err != nil {
			return err
		}

		wavfnd, err := strconv.ParseFloat(row[wavfndInx], 64)
		if err != nil {
			return err
		}

		hhld, err := strconv.ParseFloat(row[hhldInx], 64)
		if err != nil {
			return err
		}

		n := (quota * 1000000000) + (week * 10000000) + (w1yr * 1000000) +
			(qrtr * 100000) + (addr * 1000) + (wavfnd * 100) + (hhld + 1)
		column.Rows[i] = int64(n)
	}

	return nil
}

func (sf UKFilter) addCASENO() error {

	column, err := sf.dataset.AddColumn("CASENO", reflect.Int64)
	if err != nil {
		return err
	}

	startAllrows := time.Now()
	header, items := sf.dataset.GetAllRows()
	log.Debug().
		TimeDiff("elapsedTime", time.Now(), startAllrows).
		Msg("Get all rows")

	// get indexes of items we are interested in for the calculation
	quotaInx, err := sf.findLocation(header, "QUOTA")
	if err != nil {
		return err
	}
	weekInx, err := sf.findLocation(header, "WEEK")
	if err != nil {
		return err
	}
	w1yrInx, err := sf.findLocation(header, "W1YR")
	if err != nil {
		return err
	}
	qrtrInx, err := sf.findLocation(header, "QRTR")
	if err != nil {
		return err
	}
	addrInx, err := sf.findLocation(header, "ADD")
	if err != nil {
		return err
	}
	wavfndInx, err := sf.findLocation(header, "WAVFND")
	if err != nil {
		return err
	}
	hhldInx, err := sf.findLocation(header, "HHLD")
	if err != nil {
		return err
	}
	persnoInx, err := sf.findLocation(header, "PERSON")
	if err != nil {
		return err
	}

	for i := range column.Rows {
		row := items[i]

		quota, err := strconv.ParseFloat(row[quotaInx], 64)
		if err != nil {
			return err
		}

		week, err := strconv.ParseFloat(row[weekInx], 64)
		if err != nil {
			return err
		}

		w1yr, err := strconv.ParseFloat(row[w1yrInx], 64)
		if err != nil {
			return err
		}

		qrtr, err := strconv.ParseFloat(row[qrtrInx], 64)
		if err != nil {
			return err
		}

		addr, err := strconv.ParseFloat(row[addrInx], 64)
		if err != nil {
			return err
		}

		wavfnd, err := strconv.ParseFloat(row[wavfndInx], 64)
		if err != nil {
			return err
		}

		hhld, err := strconv.ParseFloat(row[hhldInx], 64)
		if err != nil {
			return err
		}

		persno, err := strconv.ParseFloat(row[persnoInx], 64)
		if err != nil {
			return err
		}

		n := (quota * 100000000000) + (week * 1000000000) + (w1yr * 100000000) +
			(qrtr * 10000000) + (addr * 100000) + (wavfnd * 10000) + (hhld * 100) + persno
		column.Rows[i] = int64(n)
	}

	return nil
}