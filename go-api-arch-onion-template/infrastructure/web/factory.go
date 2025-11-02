package web

import (
	"errors"

	"gorm.io/gorm"
)

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceGin  int = iota
	InstanceEcho int = iota
)

func StartWebServer(instance int, db *gorm.DB) {
	switch instance {
	case InstanceGin:
		config := NewConfigWeb()
		StartGinServer(config.Host, config.Port, db)
	case InstanceEcho:
		config := NewConfigWeb()
		StartEchoServer(config.Host, config.Port, db)
	default:
		panic(errInvalidWebServerInstance)
	}
}
