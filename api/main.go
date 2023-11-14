package main

import (
	"fmt"
	"io"
	"net/http"
	"signalscheck"
)

func main() {
	// Create a HTTP client
	client := &http.Client{}

	// Send a POST request to suspend the workflow
	suspendURL := fmt.Sprintf("%s/api/workflows/%s/%s/%s/signal/%s", signalscheck.TemporalServerURL, signalscheck.Namespace, signalscheck.WorkflowID, signalscheck.RunID, signalscheck.SuspendSignal)
	resp, err := client.Post(suspendURL, "", nil)
	if err != nil {
		fmt.Println("Failed to send suspend request", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read suspend response", err)
		return
	}
	fmt.Println("Sent suspend request, response:", string(body))

	//// Wait for some time
	//time.Sleep(10 * time.Second)
	//
	//// Send a POST request to resume the workflow
	//resumeURL := fmt.Sprintf("%s/api/workflows/%s/%s/%s/signal/%s", signalscheck.TemporalServerURL, signalscheck.Namespace, signalscheck.WorkflowID, signalscheck.RunID, signalscheck.ResumeSignal)
	//resp, err = client.Post(resumeURL, "", nil)
	//if err != nil {
	//	fmt.Println("Failed to send resume request", err)
	//	return
	//}
	//defer resp.Body.Close()
	//body, err = io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Failed to read resume response", err)
	//	return
	//}
	//fmt.Println("Sent resume request, response:", string(body))
}
