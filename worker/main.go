package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"signalscheck"
)

const (
	TemporalServerURL = "http://localhost:8233"
	Namespace         = "default"
	WorkflowID        = "loan-application-workflow"
	RunID             = "runID"
	SuspendSignal     = "suspend-signal"
	ResumeSignal      = "resume-signal"
	TaskQueue         = "sample_queue"
	KYCSignal         = "kyc-signal"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, signalscheck.TaskQueue, worker.Options{})
	//ao := workflow.ActivityOptions{
	//	StartToCloseTimeout: 96 * time.Hour,
	//}
	//ctx = workflow.WithActivityOptions(ctx, ao)

	w.RegisterWorkflow(signalscheck.LoanApplicationWorkflow)
	w.RegisterActivity(signalscheck.SubmitFormActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
