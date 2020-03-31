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
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		err := creategcs()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Created bucket")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

}

func creategcs() error {
	if err := serviceAccount(); err != nil {
		return err
	}
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	bucket := client.Bucket(os.Getenv("TF_VAR_cluster_id"))

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := bucket.Create(ctx, "k8stfw", &storage.BucketAttrs{
		StorageClass: "STANDARD",
		Location:     "us-east1",
	}); err != nil {
		return err
	}
	return nil
}

// serviceAccount shows how to use a service account to authenticate.
func serviceAccount() error {

	client, err := pubsub.NewClient(context.Background(), "k8stfw")
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	// Use the authenticated client.
	_ = client

	return nil
}
