package bitbucketv1

import (
	"context"

	stash "github.com/gfleury/go-bitbucket-v1"
)

// WIP

func NewClient(ctx context.Context) (*stash.APIClient, error) {
	cfg := stash.NewConfiguration()

	client := stash.NewAPIClient(ctx, cfg)

	return client, nil
}

func CreateRepository() error {
	ctx := context.TODO()

	client, _ := NewClient(ctx)

	_, err := client.DefaultApi.CreateRepository("stash-project", stash.Repository{Name: "repo-name"})
	if err != nil {
		return err
	}

}
