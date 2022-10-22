package queries

const (
	PREPSTMT_GET_ENTRY_COUNT_BY_UIC      = `SELECT COUNT(id) AS c,ui_c FROM engine_workdays WHERE date=?`
	PREPSTMT_GET_ENTRIES_FOR_UIC_BY_DATE = `select date,ui_c,job_type,train_number from engine_workdays where ui_c=? and date=?`
)
