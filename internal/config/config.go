package config

import (
	"github.com/Ladence/golang_base_rest/internal/server"
	"github.com/Ladence/golang_base_rest/internal/storage"
)

type Config struct {
	ServerCfg server.Config `json:"server" yaml:"server"`
	StorageCfg storage.Config `json:"storage" yaml:"storage"`
}
