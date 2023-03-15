package pkg

import (
	"go-skeleton/pkg/config"
	"go-skeleton/pkg/logger"
)

type Bootable interface {
	Boot()
}

var Config = config.NewConfig()
var Logger = logger.NewLogger(
	Config.Environment,
	Config.App,
	Config.Version,
)

var KernelDependencies = map[string]Bootable{
	"config": Config,
	"logger": Logger,
}

var ServerDependencies = map[string]Bootable{
	"config": Config,
	"logger": Logger,
}

var CliDependencies = map[string]Bootable{
	"config": Config,
	"logger": Logger,
}