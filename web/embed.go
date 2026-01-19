package web

import "embed"

// DistFS provides the embedded Vite build output.
//
//go:embed dist/* dist/assets/*
var DistFS embed.FS
