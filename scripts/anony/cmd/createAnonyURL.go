/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"

	"github.com/Tatsuemon/anony/rpc"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gopkg.in/yaml.v2"
)

func init() {
	rootCmd.AddCommand(newCreateAnonyURLCmd())
}

type createAnonyURLOpts struct {
	Original string
	InActive bool
}

func newCreateAnonyURLCmd() *cobra.Command {
	opts := &createAnonyURLOpts{}
	cmd := &cobra.Command{
		Use:   "create [original url]",
		Short: "create",
		Long:  "create Anony URL from original URL.\nif you do not set 'in_active' flag, this URL is open.",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Original = args[0]
			if err := createAnonyURL(cmd, opts); err != nil {
				fmt.Println()
				return errors.Wrap(err, "failed to execute a command 'create'\n")
			}
			return nil
		},
		Args: cobra.MinimumNArgs(1),
	}
	cmd.Flags().BoolVarP(&opts.InActive, "inactive", "i", false, "set URL inactive")
	return cmd
}

func createAnonyURL(cmd *cobra.Command, opts *createAnonyURLOpts) error {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		return errors.Wrap(err, "failed to establish connection\n")
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Printf("failed to conn.Close(): \n%v", err)
		}
	}()

	cli := rpc.NewAnonyServiceClient(conn)
	req := &rpc.CreateAnonyURLRequest{
		OriginalUrl: opts.Original,
		IsActive:    !opts.InActive,
	}

	// ~/.anony/config.yamlからJWTの取得
	buf, err := ioutil.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		fmt.Println(err)
	}
	m := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(buf), &m)
	if err != nil {
		fmt.Println(err)
	}
	md := metadata.Pairs("Authorization", fmt.Sprintf("bearer %s", m["Token"].(string)))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// API call
	res, err := cli.CreateAnonyURL(ctx, req)
	if err != nil {
		return errors.Wrap(err, "failed to cli.CreateAnonyURL\n")
	}

	// TODO(Tatsuemon): 出力の調整
	fmt.Println(res)

	return nil
}
