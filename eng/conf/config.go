package conf

import "github.com/go-ini/ini"

type Conf struct {
	PGConf *PGConf `ini:pg`
}

type PGConf struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	PassWord string `ini:"password"`
	DBName   string `ini:"dbname"`
}

func load(path string) (*Conf, error) {
	f, err := ini.Load(path)
	if err != nil {
		return nil, err
	}
	conf := &Conf{}
	if err := f.MapTo(conf); err != nil {
		return nil, err
	}
	return conf, nil
}
