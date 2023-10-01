package main

import (
	"log"
	_ "net/http/pprof"
	"os"

	"github.com/urfave/cli/v2"
	"gomosaics"
)

func main() {
	err := (&cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "iconsDir",
				Aliases:  []string{"d"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Required: true,
			},
			&cli.IntFlag{
				Name:  "source-block",
				Value: 16,
			},
			&cli.IntFlag{
				Name:  "icon-block",
				Value: 16,
			},
		},
		Action: func(ctx *cli.Context) error {
			gomosaics.Mosaicate(
				ctx.String("i"),
				ctx.String("d"),
				ctx.String("o"),
				ctx.Int("source-block"),
				ctx.Int("icon-block"),
			)
			return nil
		},
	}).Run(os.Args)
	if err != nil {
		log.Fatal("Could not proceed: ", err)
	}
}
