package unselect

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

// Client is an UNSELECT client.
type Client struct {
	c *client.Client
}

// NewClient creates a new client.
func NewClient(c *client.Client) *Client {
	return &Client{c: c}
}

// SupportUnselect checks if the server supports the UNSELECT extension.
func (c *Client) SupportUnselect() (bool, error) {
	return c.c.Support(Capability)
}

// Unselect frees server's resources associated with the selected mailbox and
// returns the server to the authenticated state. This command performs the same
// actions as Close, except that no messages are permanently removed from the
// currently selected mailbox.
func (c *Client) Unselect() error {
	if c.c.State != imap.SelectedState {
		return client.ErrNoMailboxSelected
	}

	cmd := &Command{}

	if status, err := c.c.Execute(cmd, nil); err != nil {
		return err
	} else if err := status.Err(); err != nil {
		return err
	}

	c.c.Mailbox = nil
	c.c.State = imap.AuthenticatedState
	return nil
}
