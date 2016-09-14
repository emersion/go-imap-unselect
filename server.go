package unselect

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/server"
)

type handler struct {
	Command
}

func (h *handler) Handle(conn server.Conn) error {
	ctx := conn.Context()
	if ctx.Mailbox == nil {
		return server.ErrNoMailboxSelected
	}

	ctx.Mailbox = nil
	ctx.MailboxReadOnly = false
	return nil
}

type extension struct{}

func (ext *extension) Capabilities(c server.Conn) []string {
	if c.Context().State&imap.AuthenticatedState != 0 {
		return []string{Capability}
	}
	return nil
}

func (ext *extension) Command(name string) server.HandlerFactory {
	if name != commandName {
		return nil
	}

	return func() server.Handler {
		return &handler{}
	}
}

func NewExtension() server.Extension {
	return &extension{}
}
