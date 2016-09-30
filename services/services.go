package services

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"
	"github.com/olebedev/emitter"
	"github.com/osiloke/dostow-contrib/store"
	"os"
	"os/signal"
)

// Make a redis pool
var redisPool *redis.Pool

type Context struct {
	stow store.Dostow
	e    *emitter.Emitter
}

func Run(dostow map[string]string, namespace, host, port, shutoffTime string) {
	stow := store.NewStore(dostow["api"], dostow["key"])
	redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host+":"+port)
		},
	}
	// Make a new pool. Arguments:
	// Context{} is a struct that will be the context for the request.
	// 10 is the max concurrency
	// "my_app_namespace" is the Redis namespace
	// redisPool is a Redis pool
	e := &emitter.Emitter{}
	pool := work.NewWorkerPool(Context{stow, e}, 10, namespace, redisPool)

	// Add middleware that will be executed for each job
	pool.Middleware((*Context).Log)

	// Map the name of jobs to handler functions
	pool.JobWithOptions("download_url", work.JobOptions{MaxFails: 1}, (*Context).DownloadUrl)

	// // Customize options:
	// pool.JobWithOptions("export", JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)

	// Start processing jobs
	pool.Start()

	stopAll := make(chan bool, 1)
	finish := make(chan bool, 1)

	scheduler.Every().Day().At(shutoffTime).Run(func() {
		<-e.Emit("action", "stop")
		stopAll <- true
	})
	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	go func() {
	OUTER:
		for {
			select {
			case <-signalChan:
				finish <- true
				break OUTER
			case <-stopAll:
				finish <- true
				break OUTER
			}
		}
	}()
	<-finish
	println("finish")
	e.Off("*")
	pool.Stop()
}
func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}
