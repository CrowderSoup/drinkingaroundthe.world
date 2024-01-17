package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/CrowderSoup/drinkingaroundthe.world/web"
)

type ServerCmd struct {
	server *web.Server
	port   string
}

func NewServerCmd(server *web.Server) *ServerCmd {
	return &ServerCmd{
		server: server,
	}
}

func (c *ServerCmd) Start() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := c.server.Start(c.port)
		if err != nil {
			fmt.Println(err)
		}
	}
}
