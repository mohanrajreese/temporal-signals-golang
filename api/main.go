//package main
//
//import (
//	"fmt"
//	"io"
//	"net/http"
//	"signalscheck"
//)
//
//func main() {
//	// Create a HTTP client
//	client := &http.Client{}
//
//	// Send a POST request to suspend the workflow
//	suspendURL := fmt.Sprintf("%s/api/workflows/%s/%s/%s/signal/%s", signalscheck.TemporalServerURL, signalscheck.Namespace, signalscheck.WorkflowID, signalscheck.RunID, signalscheck.SuspendSignal)
//	resp, err := client.Post(suspendURL, "", nil)
//	if err != nil {
//		fmt.Println("Failed to send suspend request", err)
//		return
//	}
//	defer resp.Body.Close()
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("Failed to read suspend response", err)
//		return
//	}
//	fmt.Println("Sent suspend request, response:", string(body))
//
//	//// Wait for some time
//	//time.Sleep(10 * time.Second)
//	//
//	//// Send a POST request to resume the workflow
//	//resumeURL := fmt.Sprintf("%s/api/workflows/%s/%s/%s/signal/%s", signalscheck.TemporalServerURL, signalscheck.Namespace, signalscheck.WorkflowID, signalscheck.RunID, signalscheck.ResumeSignal)
//	//resp, err = client.Post(resumeURL, "", nil)
//	//if err != nil {
//	//	fmt.Println("Failed to send resume request", err)
//	//	return
//	//}
//	//defer resp.Body.Close()
//	//body, err = io.ReadAll(resp.Body)
//	//if err != nil {
//	//	fmt.Println("Failed to read resume response", err)
//	//	return
//	//}
//	//fmt.Println("Sent resume request, response:", string(body))
//}

//package main
//
//import (
//	"bytes"
//	"context"
//	"encoding/json"
//	"fmt"
//	"log"
//	"net/http"
//
//	"go.temporal.io/sdk/client"
//)
//
//// SignalPayload Define the signal payload
//type SignalPayload struct {
//	Message string `json:"message"`
//}
//
//func main() {
//	c, err := client.Dial(client.Options{
//		HostPort: client.DefaultHostPort,
//	})
//	if err != nil {
//		log.Fatal("Unable to create client", err)
//	}
//	defer c.Close()
//
//	workflowID := "my-workflow"
//	runID := ""
//
//	payload := SignalPayload{
//		Message: "Hello",
//	}
//
//	// Marshal the payload to JSON
//	data, err := json.Marshal(payload)
//	if err != nil {
//		log.Fatal("Unable to marshal payload", err)
//	}
//
//	url := "http://localhost:8080/signal"
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
//	if err != nil {
//		log.Fatal("Unable to create request", err)
//	}
//
//	// Set the request headers
//	req.Header.Set("Content-Type", "application/json")
//
//	// Create a client and send the request
//	var client = &http.Client{}
//	res, err := client.Do(req)
//	if err != nil {
//		log.Fatal("Unable to send request", err)
//	}
//	defer res.Body.Close()
//
//	// Check the response status code
//	if res.StatusCode == http.StatusOK {
//		// Signal the workflow
//		err = c.SignalWorkflow(context.Background(), workflowID, runID, "my-signal", payload)
//		if err != nil {
//			log.Fatal("Unable to signal workflow", err)
//		}
//		fmt.Println("Signal sent successfully")
//	} else {
//		// Handle the error
//		fmt.Println("Error:", res.Status)
//	}
//}

package main

import (
	"context"
	"encoding/json"
	"go.temporal.io/sdk/client"
	"log"
)

const (
	WorkflowID  = "loan-application-workflow"
	RunID       = ""
	PauseSignal = "confirmPause"
)

type SampleSignalPayload struct {
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

func main() {
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatal("Unable to create client", err)
	}
	defer c.Close()

	samplePayload, _ := json.Marshal(SampleSignalPayload{Data: "Sample Payload", Success: true})
	err = c.SignalWorkflow(context.Background(), WorkflowID, RunID, PauseSignal, string(samplePayload))
	if err != nil {
		log.Fatal("Unable to Send Signal to the workflow", err)
	}
}
