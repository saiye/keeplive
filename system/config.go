package system

import (
	"fmt"
	"github.com/spf13/viper"
)

var _cfg *viper.Viper

func GetCfg(dir string) (*viper.Viper, error) {
	_cfg, err := NewConfig(dir, "keeplive", "ini")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return _cfg, err
}

func NewConfig(dir string, fileName string, fileType string) (*viper.Viper, error) {
	if dir == "" {
		dir = "./"
	}
	//1.读取配置文件
	config := viper.New()
	config.AddConfigPath(dir)      // 文件所在目录
	config.SetConfigName(fileName) // 文件名
	config.SetConfigType(fileType) // 文件类型
	if err := config.ReadInConfig(); err != nil {
		fmt.Errorf("读取配置文件异常.. %v", err)
		return nil, err
	}
	return config, nil
}
