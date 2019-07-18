package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type SASSURVEYSUBSAMPLE struct {
	SERIAL              float64         `gorm:"column:SERIAL;primary_key"`
	AGE                 sql.NullFloat64 `gorm:"column:AGE"`
	AMPMNIGHT           sql.NullFloat64 `gorm:"column:AM_PM_NIGHT"`
	ANYUNDER16          sql.NullString  `gorm:"column:ANYUNDER16"`
	APORTLATDEG         sql.NullFloat64 `gorm:"column:APORTLATDEG"`
	APORTLATMIN         sql.NullFloat64 `gorm:"column:APORTLATMIN"`
	APORTLATSEC         sql.NullFloat64 `gorm:"column:APORTLATSEC"`
	APORTLATNS          sql.NullString  `gorm:"column:APORTLATNS"`
	APORTLONDEG         sql.NullFloat64 `gorm:"column:APORTLONDEG"`
	APORTLONMIN         sql.NullFloat64 `gorm:"column:APORTLONMIN"`
	APORTLONSEC         sql.NullFloat64 `gorm:"column:APORTLONSEC"`
	APORTLONEW          sql.NullString  `gorm:"column:APORTLONEW"`
	ARRIVEDEPART        sql.NullFloat64 `gorm:"column:ARRIVEDEPART"`
	BABYFARE            sql.NullFloat64 `gorm:"column:BABYFARE"`
	BEFAF               sql.NullFloat64 `gorm:"column:BEFAF"`
	CHANGECODE          sql.NullFloat64 `gorm:"column:CHANGECODE"`
	CHILDFARE           sql.NullFloat64 `gorm:"column:CHILDFARE"`
	COUNTRYVISIT        sql.NullFloat64 `gorm:"column:COUNTRYVISIT"`
	CPORTLATDEG         sql.NullFloat64 `gorm:"column:CPORTLATDEG"`
	CPORTLATMIN         sql.NullFloat64 `gorm:"column:CPORTLATMIN"`
	CPORTLATSEC         sql.NullFloat64 `gorm:"column:CPORTLATSEC"`
	CPORTLATNS          sql.NullString  `gorm:"column:CPORTLATNS"`
	CPORTLONDEG         sql.NullFloat64 `gorm:"column:CPORTLONDEG"`
	CPORTLONMIN         sql.NullFloat64 `gorm:"column:CPORTLONMIN"`
	CPORTLONSEC         sql.NullFloat64 `gorm:"column:CPORTLONSEC"`
	CPORTLONEW          sql.NullString  `gorm:"column:CPORTLONEW"`
	INTDATE             sql.NullString  `gorm:"column:INTDATE"`
	DAYTYPE             sql.NullFloat64 `gorm:"column:DAYTYPE"`
	DIRECTLEG           sql.NullFloat64 `gorm:"column:DIRECTLEG"`
	DVEXPEND            sql.NullFloat64 `gorm:"column:DVEXPEND"`
	DVFARE              sql.NullFloat64 `gorm:"column:DVFARE"`
	DVLINECODE          sql.NullFloat64 `gorm:"column:DVLINECODE"`
	DVPACKAGE           sql.NullFloat64 `gorm:"column:DVPACKAGE"`
	DVPACKCOST          sql.NullFloat64 `gorm:"column:DVPACKCOST"`
	DVPERSONS           sql.NullFloat64 `gorm:"column:DVPERSONS"`
	DVPORTCODE          sql.NullFloat64 `gorm:"column:DVPORTCODE"`
	EXPENDCODE          sql.NullString  `gorm:"column:EXPENDCODE"`
	EXPENDITURE         sql.NullFloat64 `gorm:"column:EXPENDITURE"`
	FARE                sql.NullFloat64 `gorm:"column:FARE"`
	FAREK               sql.NullFloat64 `gorm:"column:FAREK"`
	FLOW                sql.NullFloat64 `gorm:"column:FLOW"`
	HAULKEY             sql.NullFloat64 `gorm:"column:HAULKEY"`
	INTENDLOS           sql.NullFloat64 `gorm:"column:INTENDLOS"`
	KIDAGE              sql.NullFloat64 `gorm:"column:KIDAGE"`
	LOSKEY              sql.NullFloat64 `gorm:"column:LOSKEY"`
	MAINCONTRA          sql.NullFloat64 `gorm:"column:MAINCONTRA"`
	MIGSI               sql.NullInt64   `gorm:"column:MIGSI"`
	INTMONTH            sql.NullFloat64 `gorm:"column:INTMONTH"`
	NATIONALITY         sql.NullFloat64 `gorm:"column:NATIONALITY"`
	NATIONNAME          sql.NullString  `gorm:"column:NATIONNAME"`
	NIGHTS1             sql.NullFloat64 `gorm:"column:NIGHTS1"`
	NIGHTS2             sql.NullFloat64 `gorm:"column:NIGHTS2"`
	NIGHTS3             sql.NullFloat64 `gorm:"column:NIGHTS3"`
	NIGHTS4             sql.NullFloat64 `gorm:"column:NIGHTS4"`
	NIGHTS5             sql.NullFloat64 `gorm:"column:NIGHTS5"`
	NIGHTS6             sql.NullFloat64 `gorm:"column:NIGHTS6"`
	NIGHTS7             sql.NullFloat64 `gorm:"column:NIGHTS7"`
	NIGHTS8             sql.NullFloat64 `gorm:"column:NIGHTS8"`
	NUMADULTS           sql.NullFloat64 `gorm:"column:NUMADULTS"`
	NUMDAYS             sql.NullFloat64 `gorm:"column:NUMDAYS"`
	NUMNIGHTS           sql.NullFloat64 `gorm:"column:NUMNIGHTS"`
	NUMPEOPLE           sql.NullFloat64 `gorm:"column:NUMPEOPLE"`
	PACKAGEHOL          sql.NullFloat64 `gorm:"column:PACKAGEHOL"`
	PACKAGEHOLUK        sql.NullFloat64 `gorm:"column:PACKAGEHOLUK"`
	PERSONS             sql.NullFloat64 `gorm:"column:PERSONS"`
	PORTROUTE           sql.NullFloat64 `gorm:"column:PORTROUTE"`
	PACKAGE             sql.NullFloat64 `gorm:"column:PACKAGE"`
	PROUTELATDEG        sql.NullFloat64 `gorm:"column:PROUTELATDEG"`
	PROUTELATMIN        sql.NullFloat64 `gorm:"column:PROUTELATMIN"`
	PROUTELATSEC        sql.NullFloat64 `gorm:"column:PROUTELATSEC"`
	PROUTELATNS         sql.NullString  `gorm:"column:PROUTELATNS"`
	PROUTELONDEG        sql.NullFloat64 `gorm:"column:PROUTELONDEG"`
	PROUTELONMIN        sql.NullFloat64 `gorm:"column:PROUTELONMIN"`
	PROUTELONSEC        sql.NullFloat64 `gorm:"column:PROUTELONSEC"`
	PROUTELONEW         sql.NullString  `gorm:"column:PROUTELONEW"`
	PURPOSE             sql.NullFloat64 `gorm:"column:PURPOSE"`
	QUARTER             sql.NullFloat64 `gorm:"column:QUARTER"`
	RESIDENCE           sql.NullFloat64 `gorm:"column:RESIDENCE"`
	RESPNSE             sql.NullFloat64 `gorm:"column:RESPNSE"`
	SEX                 sql.NullFloat64 `gorm:"column:SEX"`
	SHIFTNO             sql.NullFloat64 `gorm:"column:SHIFTNO"`
	SHUTTLE             sql.NullFloat64 `gorm:"column:SHUTTLE"`
	SINGLERETURN        sql.NullFloat64 `gorm:"column:SINGLERETURN"`
	TANDTSI             sql.NullFloat64 `gorm:"column:TANDTSI"`
	TICKETCOST          sql.NullFloat64 `gorm:"column:TICKETCOST"`
	TOWNCODE1           sql.NullFloat64 `gorm:"column:TOWNCODE1"`
	TOWNCODE2           sql.NullFloat64 `gorm:"column:TOWNCODE2"`
	TOWNCODE3           sql.NullFloat64 `gorm:"column:TOWNCODE3"`
	TOWNCODE4           sql.NullFloat64 `gorm:"column:TOWNCODE4"`
	TOWNCODE5           sql.NullFloat64 `gorm:"column:TOWNCODE5"`
	TOWNCODE6           sql.NullFloat64 `gorm:"column:TOWNCODE6"`
	TOWNCODE7           sql.NullFloat64 `gorm:"column:TOWNCODE7"`
	TOWNCODE8           sql.NullFloat64 `gorm:"column:TOWNCODE8"`
	TRANSFER            sql.NullFloat64 `gorm:"column:TRANSFER"`
	UKFOREIGN           sql.NullFloat64 `gorm:"column:UKFOREIGN"`
	VEHICLE             sql.NullFloat64 `gorm:"column:VEHICLE"`
	VISITBEGAN          sql.NullString  `gorm:"column:VISITBEGAN"`
	WELSHNIGHTS         sql.NullFloat64 `gorm:"column:WELSHNIGHTS"`
	WELSHTOWN           sql.NullFloat64 `gorm:"column:WELSHTOWN"`
	AMPMNIGHTPV         sql.NullFloat64 `gorm:"column:AM_PM_NIGHT_PV"`
	APDPV               sql.NullFloat64 `gorm:"column:APD_PV"`
	ARRIVEDEPARTPV      sql.NullFloat64 `gorm:"column:ARRIVEDEPART_PV"`
	CROSSINGSFLAGPV     sql.NullFloat64 `gorm:"column:CROSSINGS_FLAG_PV"`
	STAYIMPCTRYLEVEL1PV sql.NullFloat64 `gorm:"column:STAYIMPCTRYLEVEL1_PV"`
	STAYIMPCTRYLEVEL2PV sql.NullFloat64 `gorm:"column:STAYIMPCTRYLEVEL2_PV"`
	STAYIMPCTRYLEVEL3PV sql.NullFloat64 `gorm:"column:STAYIMPCTRYLEVEL3_PV"`
	STAYIMPCTRYLEVEL4PV sql.NullFloat64 `gorm:"column:STAYIMPCTRYLEVEL4_PV"`
	DAYPV               sql.NullFloat64 `gorm:"column:DAY_PV"`
	DISCNTF1PV          sql.NullFloat64 `gorm:"column:DISCNT_F1_PV"`
	DISCNTF2PV          sql.NullFloat64 `gorm:"column:DISCNT_F2_PV"`
	DISCNTPACKAGECOSTPV sql.NullFloat64 `gorm:"column:DISCNT_PACKAGE_COST_PV"`
	DUR1PV              sql.NullFloat64 `gorm:"column:DUR1_PV"`
	DUR2PV              sql.NullFloat64 `gorm:"column:DUR2_PV"`
	DUTYFREEPV          sql.NullFloat64 `gorm:"column:DUTY_FREE_PV"`
	FAGEPV              sql.NullFloat64 `gorm:"column:FAGE_PV"`
	FARESIMPELIGIBLEPV  sql.NullFloat64 `gorm:"column:FARES_IMP_ELIGIBLE_PV"`
	FARESIMPFLAGPV      sql.NullFloat64 `gorm:"column:FARES_IMP_FLAG_PV"`
	FLOWPV              sql.NullFloat64 `gorm:"column:FLOW_PV"`
	FOOTORVEHICLEPV     sql.NullFloat64 `gorm:"column:FOOT_OR_VEHICLE_PV"`
	HAULPV              sql.NullString  `gorm:"column:HAUL_PV"`
	IMBALCTRYFACTPV     sql.NullFloat64 `gorm:"column:IMBAL_CTRY_FACT_PV"`
	IMBALCTRYGRPPV      sql.NullFloat64 `gorm:"column:IMBAL_CTRY_GRP_PV"`
	IMBALELIGIBLEPV     sql.NullFloat64 `gorm:"column:IMBAL_ELIGIBLE_PV"`
	IMBALPORTFACTPV     sql.NullFloat64 `gorm:"column:IMBAL_PORT_FACT_PV"`
	IMBALPORTGRPPV      sql.NullFloat64 `gorm:"column:IMBAL_PORT_GRP_PV"`
	IMBALPORTSUBGRPPV   sql.NullFloat64 `gorm:"column:IMBAL_PORT_SUBGRP_PV"`
	LOSPV               sql.NullFloat64 `gorm:"column:LOS_PV"`
	LOSDAYSPV           sql.NullFloat64 `gorm:"column:LOSDAYS_PV"`
	MIGFLAGPV           sql.NullFloat64 `gorm:"column:MIG_FLAG_PV"`
	MINSCTRYGRPPV       sql.NullFloat64 `gorm:"column:MINS_CTRY_GRP_PV"`
	MINSCTRYPORTGRPPV   sql.NullString  `gorm:"column:MINS_CTRY_PORT_GRP_PV"`
	MINSFLAGPV          sql.NullFloat64 `gorm:"column:MINS_FLAG_PV"`
	MINSNATGRPPV        sql.NullFloat64 `gorm:"column:MINS_NAT_GRP_PV"`
	MINSPORTGRPPV       sql.NullFloat64 `gorm:"column:MINS_PORT_GRP_PV"`
	MINSQUALITYPV       sql.NullFloat64 `gorm:"column:MINS_QUALITY_PV"`
	NRFLAGPV            sql.NullFloat64 `gorm:"column:NR_FLAG_PV"`
	NRPORTGRPPV         sql.NullFloat64 `gorm:"column:NR_PORT_GRP_PV"`
	OPERAPV             sql.NullFloat64 `gorm:"column:OPERA_PV"`
	OSPORT1PV           sql.NullFloat64 `gorm:"column:OSPORT1_PV"`
	OSPORT2PV           sql.NullFloat64 `gorm:"column:OSPORT2_PV"`
	OSPORT3PV           sql.NullFloat64 `gorm:"column:OSPORT3_PV"`
	OSPORT4PV           sql.NullFloat64 `gorm:"column:OSPORT4_PV"`
	PUR1PV              sql.NullFloat64 `gorm:"column:PUR1_PV"`
	PUR2PV              sql.NullFloat64 `gorm:"column:PUR2_PV"`
	PUR3PV              sql.NullFloat64 `gorm:"column:PUR3_PV"`
	PURPOSEPV           sql.NullFloat64 `gorm:"column:PURPOSE_PV"`
	QMFAREPV            sql.NullFloat64 `gorm:"column:QMFARE_PV"`
	RAILCNTRYGRPPV      sql.NullFloat64 `gorm:"column:RAIL_CNTRY_GRP_PV"`
	RAILEXERCISEPV      sql.NullFloat64 `gorm:"column:RAIL_EXERCISE_PV"`
	RAILIMPELIGIBLEPV   sql.NullFloat64 `gorm:"column:RAIL_IMP_ELIGIBLE_PV"`
	REGIMPELIGIBLEPV    sql.NullFloat64 `gorm:"column:REG_IMP_ELIGIBLE_PV"`
	SAMPPORTGRPPV       sql.NullString  `gorm:"column:SAMP_PORT_GRP_PV"`
	SHIFTFLAGPV         sql.NullFloat64 `gorm:"column:SHIFT_FLAG_PV"`
	SHIFTPORTGRPPV      sql.NullString  `gorm:"column:SHIFT_PORT_GRP_PV"`
	SPENDIMPFLAGPV      sql.NullFloat64 `gorm:"column:SPEND_IMP_FLAG_PV"`
	SPENDIMPELIGIBLEPV  sql.NullFloat64 `gorm:"column:SPEND_IMP_ELIGIBLE_PV"`
	STAYIMPELIGIBLEPV   sql.NullFloat64 `gorm:"column:STAY_IMP_ELIGIBLE_PV"`
	STAYIMPFLAGPV       sql.NullFloat64 `gorm:"column:STAY_IMP_FLAG_PV"`
	STAYPURPOSEGRPPV    sql.NullFloat64 `gorm:"column:STAY_PURPOSE_GRP_PV"`
	TOWNCODEPV          sql.NullString  `gorm:"column:TOWNCODE_PV"`
	TOWNIMPELIGIBLEPV   sql.NullFloat64 `gorm:"column:TOWN_IMP_ELIGIBLE_PV"`
	TYPEPV              sql.NullFloat64 `gorm:"column:TYPE_PV"`
	UKOSPV              sql.NullFloat64 `gorm:"column:UK_OS_PV"`
	UKPORT1PV           sql.NullFloat64 `gorm:"column:UKPORT1_PV"`
	UKPORT2PV           sql.NullFloat64 `gorm:"column:UKPORT2_PV"`
	UKPORT3PV           sql.NullFloat64 `gorm:"column:UKPORT3_PV"`
	UKPORT4PV           sql.NullFloat64 `gorm:"column:UKPORT4_PV"`
	UNSAMPPORTGRPPV     sql.NullString  `gorm:"column:UNSAMP_PORT_GRP_PV"`
	UNSAMPREGIONGRPPV   sql.NullFloat64 `gorm:"column:UNSAMP_REGION_GRP_PV"`
	WEEKDAYENDPV        sql.NullFloat64 `gorm:"column:WEEKDAY_END_PV"`
	DIRECT              sql.NullFloat64 `gorm:"column:DIRECT"`
	EXPENDITUREWT       sql.NullFloat64 `gorm:"column:EXPENDITURE_WT"`
	EXPENDITUREWTK      sql.NullString  `gorm:"column:EXPENDITURE_WTK"`
	FAREKEY             sql.NullString  `gorm:"column:FAREKEY"`
	OVLEG               sql.NullFloat64 `gorm:"column:OVLEG"`
	SPEND               sql.NullFloat64 `gorm:"column:SPEND"`
	SPEND1              sql.NullFloat64 `gorm:"column:SPEND1"`
	SPEND2              sql.NullFloat64 `gorm:"column:SPEND2"`
	SPEND3              sql.NullFloat64 `gorm:"column:SPEND3"`
	SPEND4              sql.NullFloat64 `gorm:"column:SPEND4"`
	SPEND5              sql.NullFloat64 `gorm:"column:SPEND5"`
	SPEND6              sql.NullFloat64 `gorm:"column:SPEND6"`
	SPEND7              sql.NullFloat64 `gorm:"column:SPEND7"`
	SPEND8              sql.NullFloat64 `gorm:"column:SPEND8"`
	SPEND9              sql.NullFloat64 `gorm:"column:SPEND9"`
	SPENDIMPREASON      sql.NullFloat64 `gorm:"column:SPENDIMPREASON"`
	SPENDK              sql.NullFloat64 `gorm:"column:SPENDK"`
	STAY                sql.NullFloat64 `gorm:"column:STAY"`
	STAYK               sql.NullFloat64 `gorm:"column:STAYK"`
	STAY1K              sql.NullString  `gorm:"column:STAY1K"`
	STAY2K              sql.NullString  `gorm:"column:STAY2K"`
	STAY3K              sql.NullString  `gorm:"column:STAY3K"`
	STAY4K              sql.NullString  `gorm:"column:STAY4K"`
	STAY5K              sql.NullString  `gorm:"column:STAY5K"`
	STAY6K              sql.NullString  `gorm:"column:STAY6K"`
	STAY7K              sql.NullString  `gorm:"column:STAY7K"`
	STAY8K              sql.NullString  `gorm:"column:STAY8K"`
	STAY9K              sql.NullString  `gorm:"column:STAY9K"`
	STAYTLY             sql.NullFloat64 `gorm:"column:STAYTLY"`
	STAYWT              sql.NullFloat64 `gorm:"column:STAY_WT"`
	STAYWTK             sql.NullString  `gorm:"column:STAY_WTK"`
	TYPEINTERVIEW       sql.NullFloat64 `gorm:"column:TYPEINTERVIEW"`
	UKLEG               sql.NullFloat64 `gorm:"column:UKLEG"`
	VISITWT             sql.NullFloat64 `gorm:"column:VISIT_WT"`
	VISITWTK            sql.NullString  `gorm:"column:VISIT_WTK"`
	SHIFTWT             sql.NullFloat64 `gorm:"column:SHIFT_WT"`
	NONRESPONSEWT       sql.NullFloat64 `gorm:"column:NON_RESPONSE_WT"`
	MINSWT              sql.NullFloat64 `gorm:"column:MINS_WT"`
	TRAFFICWT           sql.NullFloat64 `gorm:"column:TRAFFIC_WT"`
	UNSAMPTRAFFICWT     sql.NullFloat64 `gorm:"column:UNSAMP_TRAFFIC_WT"`
	IMBALWT             sql.NullFloat64 `gorm:"column:IMBAL_WT"`
	FINALWT             sql.NullFloat64 `gorm:"column:FINAL_WT"`
}

// TableName sets the insert table name for this struct type
func (s *SASSURVEYSUBSAMPLE) TableName() string {
	return "SAS_SURVEY_SUBSAMPLE"
}