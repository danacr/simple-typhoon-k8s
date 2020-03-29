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
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use: "encrypt",

	Run: func(cmd *cobra.Command, args []string) {
		err := uploadcfg()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

}

func uploadcfg() error {
	if err := cfgencrypt(); err != nil {
		return err
	}
	if err := serviceAccount(); err != nil {
		return err
	}
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	bucket := client.Bucket(os.Getenv("TF_VAR_cluster_id"))

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	f, err := os.Open("cluster-config.gpg")
	if err != nil {
		return err
	}
	defer f.Close()

	wc := bucket.Object("cluster-config.gpg").NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	acl := bucket.Object("cluster-config.gpg").ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return err
	}

	return nil
}

func cfgencrypt() error {
	key, err := ioutil.ReadFile("pubkey.b64")
	if err != nil {
		log.Fatal(err)
	}
	// Decode public key
	decoded, err := base64.StdEncoding.DecodeString(string(key))
	if err != nil {
		return err
	}
	// Read in public key
	recipient, err := readEntity(decoded)
	if err != nil {
		return err
	}

	f, err := os.Open("cluster-config")
	if err != nil {
		return err
	}
	defer f.Close()

	dst, err := os.Create("cluster-config.gpg")
	if err != nil {
		return err
	}
	defer dst.Close()
	encrypt([]*openpgp.Entity{recipient}, nil, f, dst)
	return nil
}
func encrypt(recip []*openpgp.Entity, signer *openpgp.Entity, r io.Reader, w io.Writer) error {
	wc, err := openpgp.Encrypt(w, recip, signer, &openpgp.FileHints{IsBinary: true}, nil)
	if err != nil {
		return err
	}
	if _, err := io.Copy(wc, r); err != nil {
		return err
	}
	return wc.Close()
}

func readEntity(key []byte) (*openpgp.Entity, error) {
	r := bytes.NewReader(key)

	block, err := armor.Decode(r)
	if err != nil {
		return nil, err
	}
	return openpgp.ReadEntity(packet.NewReader(block.Body))
}
