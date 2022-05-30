package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Port             int    `mapstructure:"port"`
	Mode             string `mapstructure:"mode"`
	*BaseInfo        `mapstructure:"base" json:"base" yaml:"base"`
	*LogConfig       `mapstructure:"log"`
	*LocalServerInfo `mapstructure:"localserver" json:"localserver" yaml:"localserver"`
	*RiceIPList      `mapstructure:"riceiplist" json:"riceiplist" yaml:"riceiplist"`
	*NcPortList      `mapstructure:"ncportlist" json:"ncportlist" yaml:"ncportlist"`
}

type BaseInfo struct {
	Port      int    `mapstructure:"port" json:"port" yaml:"port"`
	Auth      string `mapstructure:"auth" yaml:"auth"`
	PublicKey string `mapstructure:"publicKey" json:"publicKey" yaml:"publicKey"`
	User      string `mapstructure:"user" json:"user" yaml:"user"`
	Passwd    string `mapstructure:"passwd" json:"passwd" yaml:"passwd"`
	DesPath   string `mapstructure:"des_path" json:"des_path" yaml:"des_path"`
	SrcPath   string `mapstructure:"src_path" json:"src_path" yaml:"src_path"`
}

type LogConfig struct {
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
}

type LocalServerInfo struct {
	Servers []string `mapstructure:"servers" yaml:"servers"`
}

type RiceIPList struct {
	RiceList []string `mapstructure:"rice_list" yaml:"rice_list"`
}

type NcPortList struct {
	PortList []string `mapstructure:"nc_port_list" yaml:"nc_port_list"`
}

func Init() error {
	viper.SetConfigFile("./config/config.yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置更新。。。")
		viper.Unmarshal(&Conf)
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	return err
}
