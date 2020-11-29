package main

import "github.com/urfave/cli"

func main() {

	app := cli.NewApp()
	app.Name = "apikit"

	app.Commands = []cli.Command{
		{
			Name:   "generate",
			Action: generateAction,
		},
	}
}

func generateAction(ctx *cli.Context) error {

	panic("not implemented")
}
