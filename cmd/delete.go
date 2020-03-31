/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serviceAccount(); err != nil {
			log.Fatal(err)
		}
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Fatal(err)
		}
		bucket := client.Bucket(os.Getenv("TF_VAR_cluster_id"))
		err = deletegcs(bucket, ctx)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Deleted bucket")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deletegcs(bucket *storage.BucketHandle, ctx context.Context) error {

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	it := bucket.Objects(ctx, nil)

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		err = bucket.Object(attrs.Name).Delete(ctx)
		if err != nil {
			return err
		}

	}
	if err := bucket.Delete(ctx); err != nil {
		return err
	}
	return nil
}
