package sqlutils

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

const SQLTableRemoveDataTemplate = `
	DELETE FROM %s WHERE id IN (SELECT id FROM %s WHERE created_at < NOW() - INTERVAL '%d DAY' LIMIT %d);
	`

// TableDelRecordsBatch removes count records from table
func TableDelRecordsBatch(db *pg.DB, table string, lifetimeDays, count int) (int, error) {
	var sql = fmt.Sprintf(SQLTableRemoveDataTemplate, table, table, lifetimeDays, count)
	res, err := db.Exec(sql)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected(), nil
}
