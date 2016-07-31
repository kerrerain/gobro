package cmd

import (
	"github.com/spf13/cobra"
)

type GobroCommand interface {
	Init()
	Run(*cobra.Command, []string) error
}
