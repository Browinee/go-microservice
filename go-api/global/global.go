package global

import (
	"go-api/config"

	ut "github.com/go-playground/universal-translator"
)

// NOTE:global variables
var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	Trans ut.Translator
)