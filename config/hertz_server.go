package config

import (
	"fmt"
	"strings"
)

type Hertz struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"` // Hertz服务器地址
	Port string `mapstructure:"port" json:"port" yaml:"port"` // Hertz服务器端口
}

func (h *Hertz) WithHostPorts() string {
	if len(h.Host) == 0 || len(h.Port) == 0 {
		return "127.0.0.1:8080"
	}
	var b strings.Builder
	b.Grow(len(h.Host) + len(h.Port) + 1)
	if _, err := fmt.Fprintf(&b, "%s:%s", h.Host, h.Port); err != nil {
		return ""
	} else {
		return b.String()
	}
}
