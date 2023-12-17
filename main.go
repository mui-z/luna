package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "luna",
		Usage: "xcframework cache manager with s3",
		Commands: []*cli.Command{
			{
				Name:  "download",
				Usage: "download xcframeworks from s3 backets",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("downloading from s3!")
					return nil
				},
			},
			{
				Name:  "upload",
				Usage: "upload xcframeworks to s3 backets",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("uploading to s3!")
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
