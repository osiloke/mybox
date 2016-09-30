package services

import (
	"encoding/json"
	"errors"
	"github.com/codeskyblue/go-sh"
	"github.com/gocraft/work"
	"os"
)

func (c *Context) DownloadUrl(job *work.Job) (err error) {
	//uses youtube-dl to download file
	url := job.ArgString("url")
	if err := job.ArgError(); err != nil {
		return err
	}
	cmd := sh.Command("youtube-dl", "--no-progress", "--print-json", url)
	outChan := make(chan []byte, 1)
	errChan := make(chan error, 1)
	job.Checkin("started download")
	go func() {
		out, err := cmd.Output()
		if string(out) == "" {
			if err != nil {
				errChan <- err
			} else {
				errChan <- errors.New("empty result")
			}
		} else {
			outChan <- out
		}

	}()
DONE:
	for {
		select {
		case <-c.e.On("action"):
			cmd.Kill(os.Kill)
			job.Checkin("killed command")
			break DONE
		case _out := <-outChan:
			var info map[string]interface{}
			json.Unmarshal(_out, &info)
			println(info)
			job.Checkin("finished download")
			break DONE
		case err = <-errChan:
			println(err)
			break DONE
		}
	}
	return
}
