package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/danielpickens/scanly/scanly
	"github.com/danielpickens/scanly/runtime"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:                "build [any valid `docker build` arguments]",
	Short:              "Builds and analyzes a docker image from a Dockerfile (this is a thin wrapper for the `docker build` command).",
	DisableFlagParsing: true,
	Run:                doBuildCmd,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

// doBuildCmd represents the steps taken for the build command
func doBuildCmd(cmd *cobra.Command, args []string) {
	initLogging()

	// there are no cli options during build allowed, only config can be supplied
	engine := viper.GetString("container-engine")

	runtime.Run(runtime.Options{
		Ci:         isCi,
		Source:     scanly.ParseImageSource(engine),
		BuildArgs:  args,
		ExportFile: exportFile,
		CiConfig:   ciConfig,
	})
}