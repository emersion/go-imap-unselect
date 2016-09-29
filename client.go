package unselect

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type Client struct {
	client *client.Client
}

// Create a new client.
func NewClient(c *client.Client) *Client {
	return &Client{client: c}
}

// Check if the server supports the UNSELECT extension.
func (c *Client) SupportsUnselect() bool {
	return c.client.Caps[Capability]
}

// Unselect frees server's resources associated with the selected mailbox and
// returns the server to the authenticated state. This command performs the same
// actions as Close, except that no messages are permanently removed from the
// currently selected mailbox.
func (c *Client) Unselect() error {
	if c.client.State != imap.SelectedState {
		return client.ErrNoMailboxSelected
	}

	cmd := &Command{}

	if status, err := c.client.Execute(cmd, nil); err != nil {
		return err
	} else if err := status.Err(); err != nil {
		return err
	}

	c.client.Mailbox = nil
	c.client.State = imap.AuthenticatedState
	return nil
}
