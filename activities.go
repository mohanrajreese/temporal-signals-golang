package signalscheck

import (
	"fmt"
	"time"
)

func GreetingActivity() (res string, err error) {

	res = "Workflow Started . Welcome to Loanwiser...!"

	time.Sleep(2 * time.Second)
	fmt.Println("Now in Greeting for the Applicant activity", res)
	return res, nil
}

func FillApplicationActivity() (res string, err error) {

	res = "Application Filled Successfully"

	time.Sleep(2 * time.Second)
	fmt.Println("Now in Fill Application activity", res)
	return res, nil
}

func SubmitApplicationActivity() (res string, err error) {

	res = "Submitted"

	time.Sleep(2 * time.Second)
	fmt.Println("Now in Submit Application activity", res)
	return res, nil
}
func ApprovalApplicationActivity() (res string, err error) {
	res = "Waiting For KYC Approval"
	time.Sleep(2 * time.Second)
	fmt.Println("Now in Approval Application activity", res)
	return res, nil
}

func SanctionLoanActivity() (res string, err error) {
	res = "20L Loan Sanctioned Successfully. Please Accept to Disburse the Loan Amount"
	time.Sleep(2 * time.Second)
	fmt.Println("Now in Approval Application activity", res)
	return res, nil
}

func DisbursalActivity() (res string, err error) {
	res = "Loan Disbursed Successfully"
	time.Sleep(2 * time.Second)
	fmt.Println("Now in Loan Disbursed activity", res)
	return res, nil
}

func CompletionActivity() (res string, err error) {
	res = "Workflow Completed"
	time.Sleep(2 * time.Second)
	fmt.Println("Now in Completion activity", res)
	return res, nil
}
