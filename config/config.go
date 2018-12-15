package config

import (
	"fmt"
	"log"

	"time"

	gcfg "gopkg.in/gcfg.v1"
)

type MainConfig struct {
	Server struct {
		Port string
	}

	DatabasePostgre struct {
		SlaveDSN string
	}

	NSQ NSQCfg

	Redis struct {
		Connection string
		Timeout    time.Duration
		MaxIdle    int
	}
}

type NSQCfg struct {
	ListenAddress         []string
	PublishAddress        string
	PublishZendeskAddress string
	Prefix                string
}

type DBConfig struct {
	SlaveDSN string
}

var Debug *log.Logger

func ReadConfig(cfg interface{}, module string) interface{} {
	ok := ReadModuleConfig(cfg, "/etc/tokopedia", module) || ReadModuleConfig(cfg, "files/etc/tokopedia", module) || ReadModuleConfig(cfg, "../files/etc/loan", module) || ReadModuleConfig(cfg, "../../files/etc/loan", module) || ReadModuleConfig(cfg, "svc/loan/files/etc/loan", module)
	if !ok {
		log.Fatalln("failed to read config for ", module)
	}
	return cfg
}

func ReadModuleConfig(cfg interface{}, path string, module string) bool {
	// environ := os.Getenv("TKPENV")

	environ := ""
	if environ == "" {
		environ = "development"
	}

	debug := Debug.Println

	fname := path + "/" + module + "." + environ + ".ini"
	err := gcfg.ReadFileInto(cfg, fname)
	fmt.Println(err)
	if err == nil {
		debug("read config from ", fname)
		return true
	}
	debug(err)
	return false
}

// func MustReadModuleConfig(cfg interface{}, paths []string, module string) {
// 	res := false
// 	for _, path := range paths {
// 		res = ReadModuleConfig(cfg, path, module)
// 		if res == true {
// 			break
// 		}
// 	}

// 	if res == false {
// 		log.Fatalln("couldn't read config for ", os.Getenv("TKPENV"))
// 	}
// }
