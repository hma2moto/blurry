package genrun

import (
	"gopkg.in/urfave/cli.v1"
)

func sobelAction(c *cli.Context) error {
	runtimePath := c.String("runtime")
	generateOutFilePath, err := generate(runtimePath, c.String("file"))
	if err != nil {
		return err
	}

	if err := runLocal(runtimePath, generateOutFilePath, c.String("input"), "sobel", nil); err != nil {
		return err
	}

	return nil
}

func init() {
	addCommand(cli.Command{
		Name:   "sobel",
		Action: sobelAction,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "i,input",
				Usage: "/path/to/input image",
				Value: "./testdata/src.png",
			},
			cli.StringFlag{
				Name:  "r,runtime",
				Usage: "halide runtime path",
				Value: "./Halide-Runtime",
			},
			cli.StringFlag{
				Name:  "f,file",
				Usage: "/path/to/blurry.cpp path",
				Value: "./blurry.cpp",
			},
		},
	})
}
