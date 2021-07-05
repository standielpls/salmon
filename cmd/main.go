package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"salmon/src/config"
	"salmon/src/ctx_values"
	"salmon/src/form"
	"salmon/src/io"
	"salmon/src/submit"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "Form submission",
		Commands: []*cli.Command{
			{
				Name:    "prs",
				Aliases: []string{"p"},
				Usage:   "submit pull requests",
				Action: func(c *cli.Context) error {
					ctx := config.ParseConfig()

					draft := submit.BuildSpareForm(ctx)
					f := form.SpareForm{
						Name: ctx_values.Get(ctx, "name"),
						Date: time.Now(),
						Body: draft,
					}
					fmt.Printf("Submission for %s at %s with body: \n%s\n", f.Name, f.Date.Format("2006-01-02"), f.Body)
					fmt.Printf("Is this ok? ")
					ok := io.ReadFrom(os.Stdin)
					if ok == "y" {
						yes(ctx, f)
					}
					if ok == "n" {
						draft := no(ctx)
						f.Body = draft
						yes(ctx, f)
					}
					fmt.Println("Submitted!")
					return nil
				},
			},
		},
		UsageText: "salmon submit",
	}

	app.Run(os.Args)
}

func yes(ctx context.Context, draft form.SpareForm) {
	err := submit.SubmitSpareForm(ctx, draft)
	if err != nil {
		panic(err)
	}
}
func no(ctx context.Context) string {
	fmt.Printf("Provide an updated draft:\n")
	return io.ReadFrom(os.Stdin)
}
