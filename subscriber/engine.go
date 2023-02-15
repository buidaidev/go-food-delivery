package subscriber

import (
	"context"
	component "go-food-delivery/component/appctx"
	"go-food-delivery/component/asyncjob"
	"go-food-delivery/pubsub"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx component.AppContext
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func NewEngine(appCtx component.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appCtx}
}

// func (engine *consumerEngine) Start() error {
// }

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("Setup comsumer for:", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ".Value", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		msg := <-c

		jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

		for i := range consumerJobs {
			jobHdl := getJobHandler(&consumerJobs[i], msg)
			jobHdlArr[i] = asyncjob.NewJob(jobHdl)
		}

		group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

		if err := group.Run(context.Background()); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
