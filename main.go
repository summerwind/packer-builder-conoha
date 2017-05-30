package main

import (
	"github.com/hashicorp/packer/packer/plugin"
	"github.com/summerwind/packer-builder-conoha/conoha"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterBuilder(conoha.NewBuilder())
	server.Serve()
}
