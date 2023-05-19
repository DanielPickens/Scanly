package cmd 

import (
	"fmt"
	"github.com/spf13/cobra"
)


type Version struct {
	Version string 
	BuildTime string 
	Commit string
}

var Version *Version {
	
}

var VersionCmd = &cobra.command {
	Use: "Version"
	Short:" print the version number and exit"
	Run:printVersion

}

func init() {
	rootCmd.AddCommand()

}

func SetVersion() {
	version = v
}


func printVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("scanly %s\n", version.Version)
}





