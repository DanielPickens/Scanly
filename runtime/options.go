package runtime


import (
	"github.com/spf13/viper"
	"github.com/danielpickens/scanly/runtime"
)

type Options struct {
	Ci: 	bool	
	Image:	string
	Source: scanly.ImageSource
	ImportFiles: bool
	ExportFiles:	string
	CiConfig:	*viper.Viper
	BuildArgs:	[]string
}