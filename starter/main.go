package main

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
	"signalscheck"

	"time"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		fmt.Println("Failed to create client", err)
		return
	}
	defer c.Close()

	// Start a workflow execution

	we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        signalscheck.WorkflowID,
		TaskQueue: signalscheck.TaskQueue,
	}, signalscheck.LoanApplicationWorkflow)
	if err != nil {
		fmt.Println("Failed to execute workflow", err)
		return
	}

	// Send a signal to the workflow after some time
	time.Sleep(10 * time.Second)
	err = c.SignalWorkflow(context.Background(), signalscheck.WorkflowID, "", signalscheck.KYCSignal, true)
	if err != nil {
		fmt.Println("Failed to signal workflow", err)
		return
	}

	// Wait for the workflow to complete
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		fmt.Println("Failed to get workflow result", err)
		return
	}

	// Print the workflow result
	fmt.Println("Workflow result:", result)
}
