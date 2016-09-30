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
	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"
	"github.com/jinzhu/now"
	"github.com/spf13/cobra"
	"time"
)

var secondsLater float64
var startAt string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download from a url/torrent/ftp etc",
	Long:  `Download from a url/torrent/ftp etc`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		if len(args) != 1 {
			return
		}
		if args[0] == "" {
			return
		}
		var redisPool = &redis.Pool{
			MaxActive: 5,
			MaxIdle:   5,
			Wait:      true,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", ":6379")
			},
		}
		enqueuer := work.NewEnqueuer("mybox", redisPool)
		var (
			job interface{}
			err error
		)
		if startAt != "" {
			t, err := now.Parse(startAt)
			if err != nil {
				println(err.Error())
				return
			}
			ct := time.Now()
			if t.After(ct) {
				diff := t.Sub(ct)
				secondsLater = diff.Seconds()
			} else {
				println(t.String())
				println("can't schedule download in the past")
				return
			}
			job, err = enqueuer.EnqueueUniqueIn("download_url", int64(secondsLater), work.Q{"startAt": t, "url": args[0]}) // job returned
		} else {
			job, err = enqueuer.EnqueueUnique("download_url", work.Q{"url": args[0]})
		}
		if err != nil {
			println(err.Error())
		} else {
			println(job)
		}

	},
}

func init() {
	RootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	downloadCmd.Flags().StringVarP(&startAt, "startAt", "s", "", "Start download at future time")

}
