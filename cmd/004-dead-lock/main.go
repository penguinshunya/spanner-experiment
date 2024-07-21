package main

import (
	"context"
	"errors"
	"fmt"
	"spanner-experiment/common"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()
	client, err := common.NewSpannerClient(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	go func() {
		_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
			println("start tx1")
			id, num, err := read(ctx, tx, "table1")
			if err != nil {
				return err
			}
			time.Sleep(3 * time.Second)
			if err := write(ctx, tx, "table2", id, num+1); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			println(err.Error())
		}
		println("end tx1")
	}()

	go func() {
		_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, tx *spanner.ReadWriteTransaction) error {
			time.Sleep(1 * time.Second)
			println("start tx2")
			id, num, err := read(ctx, tx, "table2")
			if err != nil {
				return err
			}
			time.Sleep(1 * time.Second)
			if err := write(ctx, tx, "table1", id, num+1); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			println(err.Error())
		}
		println("end tx2")
	}()

	time.Sleep(32 * time.Second)
}

func read(ctx context.Context, tx *spanner.ReadWriteTransaction, name string) (int64, int64, error) {
	sql := fmt.Sprintf("SELECT id, num FROM %s", name)
	stmt := spanner.Statement{SQL: sql}
	iter := tx.Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, 0, err
		}
		var id int64
		var num int64
		if err := row.Columns(&id, &num); err != nil {
			return 0, 0, err
		}
		return id, num, nil
	}
	return 0, 0, errors.New("not found")
}

func write(ctx context.Context, tx *spanner.ReadWriteTransaction, name string, id int64, num int64) error {
	sql := fmt.Sprintf("UPDATE %s SET num = %d WHERE id = %d", name, num, id)
	stmt := spanner.Statement{SQL: sql}
	if _, err := tx.Update(ctx, stmt); err != nil {
		return err
	}
	return nil
}
