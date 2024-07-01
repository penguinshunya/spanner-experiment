package common

import (
	"context"
	"os"

	"cloud.google.com/go/spanner"
)

func NewSpannerClient(ctx context.Context) (*spanner.Client, error) {
	projectID := os.Getenv("GCP_PROJECT_ID")
	instance := os.Getenv("SPANNER_INSTANCE")
	database := os.Getenv("SPANNER_DATABASE")
	return spanner.NewClient(ctx, "projects/"+projectID+"/instances/"+instance+"/databases/"+database)
}
