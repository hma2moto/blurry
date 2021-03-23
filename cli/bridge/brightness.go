package bridge

import (
	"log"

	"gopkg.in/urfave/cli.v1"

	"github.com/octu0/blurry"
)

func brightnessAction(c *cli.Context) error {
	in, err := loadImage(c.String("input"))
	if err != nil {
		return err
	}

	out, err := blurry.Brightness(in, c.Float64("factor"))
	if err != nil {
		return err
	}

	path, err := saveImage(out)
	if err != nil {
		return err
	}
	log.Printf("info: open %s", path)
	return nil
}

func init() {
	addCommand(cli.Command{
		Name:   "brightness",
		Action: brightnessAction,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "i,input",
				Usage: "/path/to/input image",
				Value: "./testdata/src.png",
			},
			cli.Float64Flag{
				Name:  "f,factor",
				Usage: "brightness factor",
				Value: 1.5,
			},
		},
	})
}
