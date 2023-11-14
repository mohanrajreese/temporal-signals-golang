package signalscheck

import (
	"context"
	"fmt"
	"time"
)

func SubmitFormActivity(ctx context.Context) error {
	res := "Submitted"
	time.Sleep(25 * time.Second)
	fmt.Println("Now in Submit form activity", res)
	return nil
}
