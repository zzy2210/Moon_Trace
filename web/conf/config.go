package config

import "github.com/go-ini/ini"

type Conf struct {
	Web  *Web  `ini:"web"`
	Grpc *Grpc `ini:"grpc"`
}

type Grpc struct {
	GrpcAddr []string `ini:"grpcAddr"delim:","`
}
type Web struct {
	Port int `ini:"port"`
}

func Load(path string) (*Conf, error) {
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
