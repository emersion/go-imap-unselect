package unselect

import (
	"github.com/emersion/go-imap"
)

// An UNSELECT command.
// See RFC 3691 section 2.
type Command struct{}

func (cmd *Command) Command() *imap.Command {
	return &imap.Command{Name: commandName}
}

func (cmd *Command) Parse(fields []interface{}) error {
	return nil
}
