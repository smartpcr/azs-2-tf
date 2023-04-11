package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/smartpcr/azs-2-tf/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}
}
