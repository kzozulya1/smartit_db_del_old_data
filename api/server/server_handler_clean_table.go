package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kzozulya1/smartit_db_del_old_data/internal/sqlutils"
	"github.com/kzozulya1/smartit_db_del_old_data/internal/types"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// CleanTableHandler handles cleanup of tables
func (s *Server) CleanTableHandler(c echo.Context) error {
	var tableName = c.Param("tablename")

	go func() {
		for {
			n, err := sqlutils.TableDelRecordsBatch(s.DB, tableName, s.Cfg.TableRecordsLifeTimeDays, s.Cfg.TableRecordsDelBatchSize)
			if err != nil {
				logrus.Errorf("table %s delete records batch err: %s", tableName, err.Error())
				break
			}
			//if no records were removed
			if n == 0 {
				logrus.Infof("table %s was successfully cleaned up", tableName)
				break
			}
			logrus.Infof("successfull remove %d record in table %s", n, tableName)

			//Non blocked removal - make pauses between delete transactions
			time.Sleep(time.Duration(s.Cfg.TableRecordsRemovePause) * time.Millisecond)
		}
	}()

	resp := types.NewJSONResponse(false, fmt.Sprintf("table %s cleanup was scheduled. For more information please refer to app console logs.", tableName))
	return c.JSON(http.StatusOK, resp)
}
