package signalscheck

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

const (
	WorkflowID  = "loan-application-workflow"
	TaskQueue   = "sample_queue"
	PauseSignal = "confirmPause"
)

type MySignal struct {
	Message string
}

func LoanApplicationWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 96 * time.Hour,
	}
	var res string

	var confirmPauseResult string

	PauseSelector := workflow.NewSelector(ctx)

	PauseSignalChan := workflow.GetSignalChannel(ctx, PauseSignal)

	PauseSelector.AddReceive(PauseSignalChan, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &confirmPauseResult)
	})

	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, GreetingActivity, ao).Get(ctx, &res)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, FillApplicationActivity, ao).Get(ctx, &res)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, SubmitApplicationActivity, ao).Get(ctx, &res)
	if err != nil {
		return err
	}
	PauseSelector.Select(ctx)

	err = workflow.ExecuteActivity(ctx, ApprovalApplicationActivity, ao).Get(ctx, &res)
	if err != nil {
		return err
	}
	PauseSelector.Select(ctx)

	err = workflow.ExecuteActivity(ctx, SanctionLoanActivity, ao).Get(ctx, &res)
	if err != nil {
		return err
	}
	PauseSelector.Select(ctx)

	err = workflow.ExecuteActivity(ctx, DisbursalActivity, ao).Get(ctx, &res)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, CompletionActivity, ao).Get(ctx, &res)
	if err != nil {
		return err
	}

	logger.Info("KYC verification succeeded, continuing the workflow")

	logger.Info("other processes completed")
	print("Result : ", &res)

	return nil
}
