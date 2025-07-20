package main

import "embed"

//go:embed templates/*
var templates embed.FS

type VdfTemplate struct {
	AppId       string
	Description string
	ContentRoot string
	HasSetLive  bool
	SetLive     string
	DepotId     string
	LocalPath   string
}
