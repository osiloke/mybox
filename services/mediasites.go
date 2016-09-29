package services

import (
	"encoding/json"
	"github.com/codeskyblue/go-sh"
	"github.com/gocraft/work"
)

func (c *Context) DownloadUrl(job *work.Job) error {
	//uses youtube-dl to download file
	url := job.ArgString("url")
	if err := job.ArgError(); err != nil {
		return err
	}
	cmd := sh.Command("youtube-dl", "--no-progress", "--print-json", url)
	job.Checkin("started download")
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	var info map[string]interface{}
	json.Unmarshal(out, &info)
	println(info)
	job.Checkin("finished download")
	return nil
}
