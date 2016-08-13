// This package contains the cobra commands of the Gobro CLI application.
package cmd

import (
	"github.com/spf13/cobra"
)

type GobroCommand interface {
	Init()
	Run(*cobra.Command, []string) error
}
