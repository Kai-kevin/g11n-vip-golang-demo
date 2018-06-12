package conf

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

var vipConfigInstance = new(vipConfig)
var once sync.Once

type vipConfig struct {
	ProductName     string
	ProductId string
	Locales string
	Components string
	Version         string
	VipServer       string
	InitializeCache string
	presudo         string
	collectSource   string
	EnableCache     string
	NumOfThread     int
}

func GetVipConfigInstance() *vipConfig {
	return vipConfigInstance
}

func init() {
	vipConfigMap := ReadProperties("../../resource/conf/vipconfig.properties")
	vipConfigInstance.ProductName = vipConfigMap["productName"]
	vipConfigInstance.InitializeCache = vipConfigMap["initializeCache"]
	vipConfigInstance.ProductId = vipConfigMap["productId"]
	vipConfigInstance.Locales = vipConfigMap["locales"]
	vipConfigInstance.Components = vipConfigMap["components"]
	vipConfigInstance.EnableCache = vipConfigMap["enableCache"]
	vipConfigInstance.Version = vipConfigMap["version"]
	vipConfigInstance.VipServer = vipConfigMap["vipServer"]
	vipConfigInstance.NumOfThread, _ = strconv.Atoi(vipConfigMap["numOfThread"])
}

func ReadProperties(path string) map[string]string {
	properties := make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		property := strings.TrimSpace(string(b))

		if strings.Index(property, "#") == 0 || len(property) == 0 {
			continue
		}

		index := strings.Index(property, "=")

		key := strings.TrimSpace(property[:index])
		value := strings.TrimSpace(property[index+1:])

		properties[key] = value
	}

	return properties
}