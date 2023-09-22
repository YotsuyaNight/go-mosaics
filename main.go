package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	fmt.Println("Press any button")
	var y [1200000000]uint8
	y = y
	for i := 0; i < 1200000000; i++ {
		y[i] = 1
	}
	x, _ := fmt.Scanln()
	fmt.Println("You pressed ", x)

	return

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
			Mosaicate(
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
