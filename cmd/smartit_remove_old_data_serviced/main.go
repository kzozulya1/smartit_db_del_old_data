package main

import (
	"os"
	"os/signal"

	"github.com/kzozulya1/smartit_db_del_old_data/api/server"
	"github.com/kzozulya1/smartit_db_del_old_data/internal/cmdparams"

	"github.com/sirupsen/logrus"
)

func main() {
	//Process gen flag
	if cmdparams.ProcessGenSQL() {
		return
	}

	//Init ECHO HTTP  service and run it in non-blocked mode
	var httpServer = server.New()
	go func() {
		if err := httpServer.Run(); err != nil {
			logrus.Infof("http server run: %s", err.Error())
		}
	}()

	// Setting up signal capturing
	var stop = make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop

	logrus.Info("caught terminate signal, gracefully stopping server...")
	if err := httpServer.Shutdown(); err != nil {
		logrus.Errorf("http server shutdown err: %s", err.Error())
	}

}
