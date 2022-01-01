package base

import (
	"github.com/ZMuSiShui/ipsearch-go/conf"
	"github.com/ZMuSiShui/ipsearch-go/util"
	log "github.com/sirupsen/logrus"
)

func InitIPDB() {
	log.Infof("Check IPIP IPDB File: %s", conf.IPIPFile)
	if !CheckIPDBFlie(conf.IPIPFile, conf.IPIPURL) {
		log.Fatalf("Check IPIP File Failed")
	}
	log.Infof("Check Maxmind IPDB File: %s", conf.MaxmindFile)
	if !CheckIPDBFlie(conf.MaxmindFile, conf.MaxmindURL) {
		log.Fatalf("Check IPIP File Failed")
	}

	log.Infof("Check CZ88 IPDB File: %s", conf.CZ88File)
	if !CheckIPDBFlie(conf.CZ88File, conf.CZ88URL) {
		log.Fatalf("Check IPIP File Failed")
	}
}

func CheckIPDBFlie(path string, ipdburl string) bool {
	if !util.FileExists(path) {
		log.Infof("IPDB file not exists, Download default IPDB file")
		err := util.DownloadIPDB(path, ipdburl)
		if err != nil {
			log.Fatalf("Failed to download ipdb file. Error: %s", err)
		}
		return true
	}
	return true
}