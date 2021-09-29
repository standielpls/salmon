package gh

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"salmon/src/ctx_values"

	"github.com/google/go-github/github"
)

type GithubClient struct {
	Client *github.Client
}

type PullRequest struct {
	Title       string
	Description string
	Label       string
	Status      string
	User        string
	Date        time.Time
}

const GITHUB_URL = `https://api.github.com/repos/%s/%s/pulls`

func (g *GithubClient) ListPullRequests(ctx context.Context) ([]PullRequest, error) {

	var prs []PullRequest
	page := 1
	for page < 3 {
		pullRequests, shouldContinue, err := g.Do(ctx, page)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("unable to fetch pull requests: %s", err.Error()))
		}
		if !shouldContinue {
			break
		}
		prs = append(prs, pullRequests...)
		page++
	}

	return prs, nil
}

func (g *GithubClient) Do(ctx context.Context, page int) ([]PullRequest, bool, error) {
	ONE_WEEK_AGO := time.Now().AddDate(0, 0, -7)

	fullUrl := fmt.Sprintf(GITHUB_URL, ctx_values.Get(ctx, "owner"), ctx_values.Get(ctx, "repo"))
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return []PullRequest{}, false, errors.New(fmt.Sprintf("unable to create request: %s", err.Error()))
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	v := url.Values{}
	v.Set("page", strconv.Itoa(page))
	v.Set("state", "all")

	req.URL.RawQuery = v.Encode()

	req.SetBasicAuth(ctx_values.Get(ctx, "user"), ctx_values.Get(ctx, "token"))
	fmt.Println(req.URL.String())

	client := http.DefaultClient
	data, err := client.Do(req)
	if err != nil {
		return []PullRequest{}, false, errors.New(fmt.Sprintf("unable to list pull request: %s", err.Error()))
	}

	var githubPRs []*github.PullRequest
	err = json.NewDecoder(data.Body).Decode(&githubPRs)
	if err != nil {
		return []PullRequest{}, false, errors.New(fmt.Sprintf("unable to decode body: %s", err.Error()))
	}

	var prs []PullRequest
	for _, pr := range githubPRs {
		updatedAt := pr.GetUpdatedAt()
		if updatedAt.After(ONE_WEEK_AGO) {
			if pr.GetUser().GetLogin() == ctx_values.Get(ctx, "user") {
				prs = append(prs, mapPullRequest(pr))
			}
		} else {
			return prs, false, nil
		}
	}
	return prs, true, nil
}

func mapPullRequest(gpr *github.PullRequest) PullRequest {
	title := strings.TrimSpace(strings.Split(gpr.GetTitle(), "|")[1])
	label := strings.Join(getSigLabels(gpr.Labels), ", ")
	description := strings.TrimSpace(strings.Split(strings.Split(gpr.GetBody(), "Description")[1], "#")[0])
	status := gpr.GetState()
	assignee := gpr.GetAssignee()

	return PullRequest{
		Title:       title,
		Label:       label,
		Description: description,
		Status:      status,
		User:        assignee.GetName(),
		Date:        gpr.GetUpdatedAt(),
	}
}

func getSigLabels(labels []*github.Label) []string {
	var filtered []string
	m := make(map[string]struct{})
	for _, label := range labels {
		name := label.GetName()
		if name != "auto-merge" && name != "blocked" {
			if _, ok := m[name]; !ok {
				m[name] = struct{}{}
				filtered = append(filtered, "Spare "+strings.Title(name))
			}
		}
	}
	return filtered
}
