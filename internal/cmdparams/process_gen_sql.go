package cmdparams

import (
	"github.com/kzozulya1/smartit_db_del_old_data/internal/sqlutils"

	"github.com/sirupsen/logrus"

	"flag"
	"fmt"
)

const (
	genSQLDumpErrTemplate   = "failed to create dump for %s table"
	sqlDumpFilenameTemplate = "../../scripts/sql/%s.sql"
	numRowsToGenerate       = 333000 //0.33М
)

var tablenames = []string{"tmp1", "tmp2"}

// ProcessGenSQL generates SQL dump data
func ProcessGenSQL() (retExit bool) {
	var gen bool

	flag.BoolVar(&gen, "gen", false, "Generate SQL dump for 2 test tables")
	flag.Parse()
	if gen {

		for _, tablename := range tablenames {
			if err := genSQLDump(tablename); err != nil {
				logrus.Errorf(genSQLDumpErrTemplate, tablename)
			}
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
