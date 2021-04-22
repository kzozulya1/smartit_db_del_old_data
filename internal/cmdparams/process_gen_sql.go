package cmdparams

import (
	"smartit_db_del_old_data/internal/sqlutils"

	"github.com/sirupsen/logrus"

	"flag"
	"fmt"
)

const (
	genSQLDumpErrTemplate   = "failed to create dump for %s table"
	sqlDumpFilenameTemplate = "../../scripts/sql/%s.sql"

	numRowsToGenerate = 11000000 //11M
	tablename1        = "tmp1"
	tablename2        = "tmp2"
)

// ProcessGenSQL generates SQL dump data
func ProcessGenSQL() (retExit bool) {
	var gen bool

	flag.BoolVar(&gen, "gen", false, "Generate SQL dump for 2 test tables")
	flag.Parse()
	if gen {

		if err := genSQLDump(tablename1); err != nil {
			logrus.Errorf(genSQLDumpErrTemplate, tablename1)
		}

		if err := genSQLDump(tablename2); err != nil {
			logrus.Errorf(genSQLDumpErrTemplate, tablename1)
		}

		retExit = true
	}
	return
}

// genSQLDump generates table dump
func genSQLDump(t string) error {
	logrus.Infof("generating SQL dump for %s DB table ...", t)
	return sqlutils.DumpTestData(getTableDumpFilename(t),
		sqlutils.GenerateTestData(t, numRowsToGenerate))
}

// getTableDumpFilename returns full filename of table
func getTableDumpFilename(t string) string {
	return fmt.Sprintf(sqlDumpFilenameTemplate, t)
}
