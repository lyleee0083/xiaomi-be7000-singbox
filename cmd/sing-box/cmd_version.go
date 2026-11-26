package main

import (
	"os"
	"runtime"
	"runtime/debug"

	C "github.com/sagernet/sing-box/constant"

	"github.com/spf13/cobra"
)

var commandVersion = &cobra.Command{
	Use:   "version",
	Short: "Print current version of sing-box",
	Run:   printVersion,
	Args:  cobra.NoArgs,
}

var nameOnly bool
var distributionVersion string

func init() {
	commandVersion.Flags().BoolVarP(&nameOnly, "name", "n", false, "print version name only")
	mainCommand.AddCommand(commandVersion)
}

func printVersion(cmd *cobra.Command, args []string) {
	reportedVersion := C.Version
	if distributionVersion != "" {
		reportedVersion = distributionVersion
	}
	if nameOnly {
		os.Stdout.WriteString(reportedVersion + "\n")
		return
	}
	version := "sing-box version " + reportedVersion + "\n\n"
	if distributionVersion != "" {
		version += "Upstream version: " + C.Version + "\n"
	}
	version += "Environment: " + runtime.Version() + " " + runtime.GOOS + "/" + runtime.GOARCH + "\n"

	var tags string
	var revision string

	debugInfo, loaded := debug.ReadBuildInfo()
	if loaded {
		for _, setting := range debugInfo.Settings {
			switch setting.Key {
			case "-tags":
				tags = setting.Value
			case "vcs.revision":
				revision = setting.Value
			}
		}
	}

	if tags != "" {
		version += "Tags: " + tags + "\n"
	}
	if revision != "" {
		version += "Revision: " + revision + "\n"
	}

	if C.CGO_ENABLED {
		version += "CGO: enabled\n"
	} else {
		version += "CGO: disabled\n"
	}

	os.Stdout.WriteString(version)
}
