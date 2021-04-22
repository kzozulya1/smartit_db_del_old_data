package server

import (
	"context"
	"fmt"
	"time"

	"smartit_db_del_old_data/config"
	"smartit_db_del_old_data/internal/storage"

	"github.com/go-pg/pg/v10"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

const shutdownWaitTime time.Duration = 5

// Server contains instance details for the server
type Server struct {
	Cfg    *config.Configuration
	Server *echo.Echo
	DB     *pg.DB
}

// New returns a new instance of the server based on the specified configuration.
func New() *Server {
	var server = &Server{
		Cfg: &config.Configuration{},
	}

	server.initConfig()
	server.initDBConn()
	server.initServer()
	server.useMiddleware()
	server.initRoutes()
	return server
}

// initRouter inits routers
func (s *Server) initRoutes() {
	logrus.Infof("http server: init routes...")

	//Handle clean tables
	s.Server.GET("/cleantable/:tablename", s.CleanTableHandler)
}

// initConfig inits config
func (s *Server) initConfig() {
	if err := envconfig.Process("", s.Cfg); err != nil {
		panic(fmt.Errorf("load configuration err: %s", err.Error()))
	}
}

// initServer initialized HTTP server
func (s *Server) initDBConn() {
	var err error
	logrus.Infof("http server: init SDB conn...")
	s.DB, err = storage.InitDB(s.Cfg.DBConn, s.Cfg.DBSQLQueryLog)
	if err != nil {
		panic(fmt.Errorf("db init err: %s", err.Error()))
	}
}

// initServer initialized HTTP server
func (s *Server) initServer() {
	logrus.Infof("http server: init server...")
	s.Server = echo.New()
	s.Server.HideBanner = true
}

// initRouter inits routers
func (s *Server) useMiddleware() {
	logrus.Infof("http server: use CORS middleware...")

	s.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "HEAD", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
}

// Run starts service
func (s *Server) Run() error {
	if s.Server == nil {
		return fmt.Errorf("echo server is not initialized")
	}
	logrus.Infof("starting HTTP service listen with %s", s.Cfg.ListenAddr)
	return s.Server.Start(s.Cfg.ListenAddr)
}

// Shutdown() gracefully stops serving
func (s *Server) Shutdown() error {
	var ctx, cancel = context.WithTimeout(context.Background(), shutdownWaitTime*time.Second)
	defer cancel()
	logrus.Infof("http server: shutdown...")
	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}

	if err := s.closeDB(); err != nil {
		return fmt.Errorf("close db err: %s", err.Error())
	}

	return nil
}

// closeDB closes DB conns
func (s *Server) closeDB() error {
	if s.DB != nil {
		if err := s.DB.Close(); err != nil {
			return err
		}
	}
	return nil
}
