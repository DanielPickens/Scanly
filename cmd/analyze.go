package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"github.com/spf13/cobra"

)

// AnalyzeCmd takes in a docker image tag, digest, or id and passes scan to display the
// image analysis to the screen
func AnalyzeCmd(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		printVersionFlag, err := cmd.PersistentFlags().GetBool("version")
		if err == nil && printVersionFlag {
			printVersion(cmd, args)
			return
		}

		fmt.Println("No image argument given")
		os.Exit(1)
	}

	userImage := args[0]
	if userImage == "" {
		fmt.Println("No image argument given")
		os.Exit(1)
	}

	initLogging()

	isCi, ciConfig, err := configureCi()

	if err != nil {
		fmt.Printf("ci configuration error: %v\n", err)
		os.Exit(1)
	}

	var sourceType scanly.ImageSource
	var imageStr string

	sourceType, imageStr = scanly.DeriveImageSource(userImage)

	if sourceType == scanly.SourceUnknown {
		sourceStr := viper.GetString("source")
		sourceType = scanly.ParseImageSource(sourceStr)
		if sourceType == scanly.SourceUnknown {
			fmt.Printf("unable to determine image source: %v\n", sourceStr)
			os.Exit(1)
		}

		imageStr = userImage
	}

	ignoreErrors, err := cmd.PersistentFlags().GetBool("ignore-errors")
	if err != nil {
		logrus.Error("unable to get 'ignore-errors' option:", err)
	}

	runtime.Run(runtime.Options{
		Ci:           isCi,
		Source:       sourceType,
		Image:        imageStr,
		ExportFile:   exportFile,
		CiConfig:     ciConfig,
		IgnoreErrors: viper.GetBool("ignore-errors") || ignoreErrors,
	})
}