package main

import "embed"

//go:embed templates/*
var templates embed.FS

type MessageTemplate struct {
	StartedAt  string
	Commit     string
	IsBuildURL bool
	BuildURL   string
}
