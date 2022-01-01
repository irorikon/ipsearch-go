package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/ZMuSiShui/ipsearch-go/cmd/base"
	"github.com/ZMuSiShui/ipsearch-go/conf"
	"github.com/ZMuSiShui/ipsearch-go/util"
)

func init() {
	flag.BoolVar(&conf.Debug, "debug", false, "start with debug mode")
	flag.BoolVar(&conf.Version, "version", false, "print version info")
	flag.BoolVar(&conf.SkipUpdate, "skip-update", false, "skip update")
	flag.StringVar(&conf.ConfigFile, "config-file", "data/config.json", "config file")
	flag.Parse()
}

func Init() bool {
	base.InitLog()
	base.InitConfig()
	return true
}

func main() {
	if conf.Version {
		fmt.Printf("Built At: %s\nGo Version: %s\nVersion: %s\n", conf.BuiltAt, conf.GoVersion, conf.VERSION)
		return
	}
	if !Init() {
		return
	}
	if conf.Debug {
		log.Info("Set Debug Mode")
	}
	util.PrintLogo()
	log.Info("Starting Client")
}
