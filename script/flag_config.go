package script

import (
	"path"
	"strings"

	"github.com/small-tk/pathlib"
	"github.com/spf13/pflag"
)

//InitFlagConfig 从命令行获取配置
func InitFlagConfig() (ConfigType, error) {
	loglevel := pflag.StringP("loglevel", "l", "", "log的等级")
	address := pflag.StringP("address", "a", "", "要启动的服务器地址")
	confPath := pflag.StringP("config", "c", "", "配置文件位置")
	pflag.Parse()
	var flagConfig = ConfigType{}

	if *confPath != "" {
		p, err := pathlib.New(*confPath).Absolute()
		if err != nil {
			return flagConfig, err
		}
		if p.Exists() && p.IsFile() {
			filenameWithSuffix := path.Base(*confPath)
			fileSuffix := path.Ext(filenameWithSuffix)
			fileName := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
			dir, err := p.Parent()
			if err != nil {
				return flagConfig, err
			}
			filePaths := []string{dir.Path}
			targetfileconf, err := SetFileConfig(fileName, filePaths)
			if err != nil {
				return flagConfig, err
			}
			for k, v := range targetfileconf {
				flagConfig[k] = v
			}
		}
	}
	if *loglevel != "" {
		flagConfig["log_level"] = *loglevel
	}

	if *address != "" {
		flagConfig["address"] = *address
	}
	return flagConfig, nil
}
