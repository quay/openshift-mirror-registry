package cmd

import (
	_ "embed" // embed package is used to embed service files
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//go:embed assets/Dockerfile
var dockerfile string

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return true
	}
	return false
}

func check(err error) {
	if err != nil {
		log.Error("An error occurred: %s", err.Error())
		cleanup()
		os.Exit(1)
	}
}

func cleanup() {
	os.RemoveAll("/tmp/ansible")
	os.RemoveAll("/tmp/quay-ansible")
}

// verbose is the optional command that will display INFO logs
var verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Display verbose logs")
}

var (
	rootCmd = &cobra.Command{
		Use: "quay-installer",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if verbose {
				log.SetLevel(log.InfoLevel)
			} else {
				log.SetLevel(log.WarnLevel)
			}
		},
	}
)

// Execute executes the root command.
func Execute() error {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	fmt.Println(`
   __   __
  /  \ /  \     ______   _    _     __   __   __
 / /\ / /\ \   /  __  \ | |  | |   /  \  \ \ / /
/ /  / /  \ \  | |  | | | |  | |  / /\ \  \   /
\ \  \ \  / /  | |__| | | |__| | / ____ \  | |
 \ \/ \ \/ /   \_  ___/  \____/ /_/    \_\ |_|
  \__/ \__/      \ \__
                  \___\ by Red Hat
 Build, Store, and Distribute your Containers
	`)
	return rootCmd.Execute()
}
