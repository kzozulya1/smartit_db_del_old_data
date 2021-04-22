package sqlutils

import (
	"fmt"
	"strings"
)

const (
	SQLNewTableSchemaTemplate = `
	CREATE TABLE IF NOT EXISTS %s (
		id             BIGSERIAL NOT NULL PRIMARY KEY,
		data           VARCHAR(128),
		created_at     TIMESTAMP
	);
	`
	SQLNewRowTemplate = ` INSERT INTO %s (data, created_at) VALUES ('day -%d', NOW() - INTERVAL '%d DAY'); 
	`
)

// GenerateTestData generates SQL for create big data tables
func GenerateTestData(tablename string, rowsCount int) string {
	var (
		strBuilder strings.Builder
		i          int = 1
	)

	//Write table definition
	strBuilder.WriteString(fmt.Sprintf(SQLNewTableSchemaTemplate, tablename))

	for i <= rowsCount {
		strBuilder.WriteString(fmt.Sprintf(SQLNewRowTemplate, tablename, i, i))
		i++
	}

	return strBuilder.String()
}
