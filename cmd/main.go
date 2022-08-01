package main

import (
	"cleanArch/config"
	"cleanArch/internal/registry"
)

func main() {
	config.ReadConfig()

	registry.NewRegistry()
}
