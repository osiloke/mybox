// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/osiloke/mybox/services"
	"github.com/spf13/cobra"
)

var (
	api, key, host, port string
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Runs a box worker service",
	Long:  `Runs a box worker service`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("service called")
		services.Run(
			map[string]string{
				"url": api,
				"key": key,
			},
			"mybox",
			host, port)
	},
}

func init() {
	RootCmd.AddCommand(serviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")
	serviceCmd.Flags().StringVarP(&api, "api", "a", "http://localhost:3001/v1", "api")
	serviceCmd.Flags().StringVarP(&key, "key", "k", "8b69af27-4fae-46a2-a1ad-f74744ee9de2", "key")
	serviceCmd.Flags().StringVarP(&host, "redis-host", "e", "localhost", "redis host")
	serviceCmd.Flags().StringVarP(&port, "redis-port", "p", "6379", "redis port")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
