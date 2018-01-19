package configuration

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	REMOTE struct {
		HOST      string `json:"HOST"`
		PORT      string    `json:"PORT"`
		ENDPOINTS struct {
			BASE string `json:"BASE"`
			ITEM   string `json:"ITEM"`
			REVIEW string `json:"REVIEW"`
			REVIEWTMP string `json:"REVIEW_TMP"`
			PRICE string `json:"PRICE"`
		} `json:"ENDPOINTS"`
	} `json:"REMOTE"`
}

var configuration Configuration

func InitConfiguration() Configuration {
	conf := Configuration{}
	err := gonfig.GetConf(getFileName(), &conf)
	if err != nil {
		fmt.Println("error " + err.Error())
		os.Exit(1)
	}
	configuration = conf
	return configuration
}

func GetConfiguration()  Configuration {
	return configuration
}

func getFileName() string {
	filename := "configuration.json"
	_, dirName, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirName), filename)

	return filePath
}
