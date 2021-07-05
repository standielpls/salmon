package submit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"salmon/src/form"
	gh "salmon/src/github"

	"github.com/google/go-github/github"
)

func BuildSpareForm(ctx context.Context) string {
	ghClient := gh.GithubClient{
		Client: github.NewClient(http.DefaultClient),
	}
	prs, err := ghClient.ListPullRequests(ctx)
	if err != nil {
		fmt.Printf("unable to list pull requests: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(prs)

	return createDescription(prs)
}

func SubmitSpareForm(ctx context.Context, f form.SpareForm) error {
	_, err := f.SubmitForm()
	if err != nil {
		return errors.New(fmt.Sprintf("unable to submit spare form: %s", err.Error()))
	}
	return nil
}

func createDescription(prs []gh.PullRequest) string {
	var s strings.Builder

	productMap := make(map[string]struct{})
	var products []string
	var titles []string
	var bodies []string
	var statuses []string
	for _, pr := range prs {
		if _, ok := productMap[pr.Label]; !ok {
			productMap[pr.Label] = struct{}{}
			products = append(products, pr.Label)
		}

		titles = append(titles, pr.Title)
		bodies = append(bodies, pr.Description)
		var status string
		if pr.Status == "open" {
			status = "ongoing investigation"
		} else {
			status = "successfully implemented"
		}
		statuses = append(statuses, status)
	}
	s.WriteString("1.\t" + strings.Join(products, ", ") + "\n")
	s.WriteString("2.\t" + strings.Join(titles, ", ") + "\n")
	s.WriteString("3, 4.\t" + strings.Join(bodies, "\n") + "\n")
	s.WriteString("5.\t" + strings.Join(statuses, ", ") + "\n")
	return s.String()
}
