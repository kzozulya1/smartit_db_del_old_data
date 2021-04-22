package server

import (
	"net/http"
	"smartit_db_del_old_data/internal/types"

	"github.com/labstack/echo"
)

// RunAddressStockHandler ...
func (s *Server) CleanTableHandler(c echo.Context) error {
	var tableName = c.QueryParam("tablename")

	resp := types.NewJSONResponse(false, "cleaning table "+tableName)
	return c.JSON(http.StatusOK, resp)
}
