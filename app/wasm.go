package app

import (
	wasmkeeper "github.com/Generative-Labs/wasmd/x/wasm/keeper"
)

// Deprecated: Use BuiltInCapabilities from github.com/Generative-Labs/wasmd/x/wasm/keeper
func AllCapabilities() []string {
	return wasmkeeper.BuiltInCapabilities()
}
