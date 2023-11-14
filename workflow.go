package signalscheck

import (
	"go.temporal.io/sdk/workflow"
	"time"
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

type MySignal struct {
	Message string
}

func LoanApplicationWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)

	selector := workflow.NewSelector(ctx)

	var kycResult bool

	selector.AddReceive(workflow.GetSignalChannel(ctx, KYCSignal), func(c workflow.ReceiveChannel, more bool) {
		// Receive the signal value and assign it to the kycResult variable
		c.Receive(ctx, &kycResult)
		logger.Info("Received KYC signal: ", kycResult)
	})
	//var signal bool
	//signalChan := workflow.GetSignalChannel(ctx, KYCSignal)
	//signalChan.Receive(ctx, &signal)
	//if !signal {
	//	return errors.New("signal")
	//}
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 96 * time.Hour,
	}
	var res string
	ctx = workflow.WithActivityOptions(ctx, ao)
	err := workflow.ExecuteActivity(ctx, SubmitFormActivity, ao).Get(ctx, &res)
	if err != nil {
		return err
	}
	logger.Info("Submitted the form to the bank")

	logger.Info("Waiting for KYC signal or timeout")
	selector.Select(ctx)
	if !kycResult {
		logger.Info("KYC verification failed or timed out")
		return nil
	}

	logger.Info("KYC verification succeeded, continuing the workflow")
	//time.Sleep(20 * time.Second)
	logger.Info("other processes completed")
	print("Result : ", &res)

	//c, err := client.Dial(client.Options{
	//	HostPort: client.DefaultHostPort,
	//})
	//if err != nil {
	//	log.Fatalln("Unable to create client", err)
	//}
	//defer c.Close()
	//
	//err = c.CancelWorkflow(context.Background(), WorkflowID, "")
	//if err != nil {
	//	log.Fatalln("Unable to cancel Workflow Execution", err)
	//}
	return nil
}
