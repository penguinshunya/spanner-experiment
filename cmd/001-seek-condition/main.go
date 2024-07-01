package main

import (
	"context"
	"spanner-experiment/common"

	"cloud.google.com/go/spanner"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	client, err := common.NewSpannerClient(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	for i := 0; i < 10; i++ {
		muts := []*spanner.Mutation{}
		for j := 0; j < 10000; j++ {
			userID := uuid.New().String()
			muts = append(muts, spanner.Insert("Users", []string{"UserID"}, []interface{}{userID}))
		}
		_, err := client.Apply(ctx, muts)
		if err != nil {
			panic(err)
		}
	}
}
