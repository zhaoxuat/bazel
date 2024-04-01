package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/zhaoxuat/bazel/pkg"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)

	s := pkg.NewService()
	log.Info(s.Run())
}
