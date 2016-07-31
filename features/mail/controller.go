package mail

import (
	"bytes"
	"fmt"
	"github.com/magleff/gobro/features/expense"
	"github.com/mxk/go-imap/imap"
	"log"
	"net/mail"
	"os"
	"strings"
	"time"
)

const PREFIX = "Gobro"

type MailController struct{}

func NewMailController() *MailController {
	instance := new(MailController)
	return instance
}

func populateExpensesFromGobroMail(expenses *[]expense.Expense, subject string, body string) {
	chunks := strings.Split(subject, " ")
	bodyChunks := strings.Split(body, "=")
	if len(chunks) == 2 && chunks[0] == PREFIX {
		fmt.Println("Amount:", chunks[1], "Body:", bodyChunks[0])
		newExpense := *expense.NewExpense(chunks[1], bodyChunks[0])
		*expenses = append(*expenses, newExpense)
	}
}

func fetchExpensesFromMailData(c *imap.Client, cmd *imap.Command) []expense.Expense {
	var expenses []expense.Expense

	for cmd.InProgress() {
		// Wait for the next response (no timeout)
		c.Recv(-1)

		// Process command data
		for _, rsp := range cmd.Data {
			header := imap.AsBytes(rsp.MessageInfo().Attrs["RFC822.HEADER"])
			body := imap.AsBytes(rsp.MessageInfo().Attrs["BODY[1]"])

			if msg, _ := mail.ReadMessage(bytes.NewReader(header)); msg != nil {
				populateExpensesFromGobroMail(&expenses, msg.Header.Get("Subject"), string(body))
			}
		}

		cmd.Data = nil

		// Process unilateral server data
		for _, rsp := range c.Data {
			fmt.Println("Server data:", rsp)
		}
		c.Data = nil
	}

	// Check command completion status
	if rsp, err := cmd.Result(imap.OK); err != nil {
		if err == imap.ErrAborted {
			fmt.Println("Fetch command aborted")
		} else {
			fmt.Println("Fetch error:", rsp.Info)
		}
	}

	return expenses
}

// Imports the last Gobro expenses from a mailbox
// The codes relies on https://godoc.org/github.com/mxk/go-imap/imap#Client
// Many thanks to mxk for this package \o/
func (self MailController) ImportFromMail() []expense.Expense {
	var (
		c   *imap.Client
		cmd *imap.Command
	)

	if os.Getenv("GOBRO_MAIL_SERVER") == "" || os.Getenv("GOBRO_MAIL_USER") == "" || os.Getenv("GOBRO_MAIL_PASSWORD") == "" {
		log.Fatal("Please set the GOBRO_MAIL_{SERVER, USER, PASSWORD} environment variables.")
	}

	// Connect to the server
	c, _ = imap.Dial(os.Getenv("GOBRO_MAIL_SERVER"))

	// Remember to log out and close the connection when finished
	defer c.Logout(30 * time.Second)

	// Print server greeting (first response in the unilateral server data queue)
	fmt.Println("Server says hello:", c.Data[0].Info)
	c.Data = nil

	// Enable encryption, if supported by the server
	if c.Caps["STARTTLS"] {
		c.StartTLS(nil)
	}

	// Authenticate
	if c.State() == imap.Login {
		c.Login(os.Getenv("GOBRO_MAIL_USER"), os.Getenv("GOBRO_MAIL_PASSWORD"))
	}

	// Open a mailbox (synchronous command - no need for imap.Wait)
	c.Select("INBOX", false)

	cmd, _ = imap.Wait(c.Search("UNSEEN", "SUBJECT", c.Quote("Gobro")))
	set, _ := imap.NewSeqSet("")
	set.AddNum(cmd.Data[0].SearchResults()...)

	cmd, _ = c.Fetch(set, "RFC822.HEADER", "BODY[1]")

	// Process responses while the command is running
	fmt.Println("\nExtracting data from the mailbox:")

	expenses := fetchExpensesFromMailData(c, cmd)

	c.Store(set, "+FLAGS", imap.NewFlagSet(`\Seen`))

	return expenses
}
