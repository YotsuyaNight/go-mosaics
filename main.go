package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

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
		},
		Action: func(ctx *cli.Context) error {
			Mosaicate(ctx.String("i"), ctx.String("d"), ctx.String("o"))
			return nil
		},
	}).Run(os.Args)
	if err != nil {
		log.Fatal("Could not proceed: ", err)
	}
}
