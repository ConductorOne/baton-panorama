package main

import (
	cfg "github.com/conductorone/baton-panorama/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/config"
)

func main() {
	config.Generate("panorama", cfg.Config)
}
