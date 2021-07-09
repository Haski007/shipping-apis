package main

import (
	"fmt"
	"os"

	"github.com/Haski007/shipping-apis/internal/shipping"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var Version string

func main() {
	app := cli.App{
		Name:    "shipping",
		Usage:   "Mock shipping APIs to get best deal",
		Version: Version,
		Action: func(c *cli.Context) error {
			if err := shipping.Run(); err != nil {
				return fmt.Errorf("run: %w", err)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
