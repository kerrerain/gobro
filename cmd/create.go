package cmd

import (
	"errors"
	"github.com/magleff/gobro/features/account"
	"github.com/spf13/cobra"
)

type GobroCreateCommand struct {
	Command           *cobra.Command
	AccountController account.AccountController
}

func (self *GobroCreateCommand) Init() {
	self.Command = &cobra.Command{
		Use:   "create [account] [name] or create [budget]",
		Short: "Creates an account / a budget",
		Long:  `Creates an account / a budget. For creating an account, you must provide a name.`,
		RunE:  self.Run,
	}
	self.AccountController = account.NewAccountController()
}

func (self *GobroCreateCommand) Run(cmd *cobra.Command, args []string) error {
	var err error

	if len(args) == 0 {
		return errors.New("The type of the element to create is mandatory.")
	}

	if typeOfElement := args[0]; typeOfElement == "account" {
		err = self.CreateAccount(args)
	} else if typeOfElement == "budget" {
		err = self.CreateBudget(args)
	} else {
		err = errors.New("The type " + typeOfElement + " is not recognized.")
	}

	return err
}

func (self *GobroCreateCommand) CreateAccount(args []string) error {
	var err error

	if len(args) < 2 {
		err = errors.New("A name should be provided for the new account.")
	} else {
		err = self.AccountController.Create(args[1])
	}
	return err
}

func (self *GobroCreateCommand) CreateBudget(args []string) error {
	return nil
}

func init() {
	command := GobroCreateCommand{}
	command.Init()
	RootCmd.AddCommand(command.Command)
}
