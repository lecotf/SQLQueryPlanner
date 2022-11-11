package scheduler

import (
	"fmt"
	"time"
)

func QueryScheduler() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	duration := 60 - time.Duration(currentTime.Second())
	time.Sleep(duration * time.Second)
	for {
		currentTime = time.Now()
		// DO THE JOB
		fmt.Println("Do the job")
		// END THE JOB
		currentTime = time.Now()
		fmt.Println(currentTime)
		duration = 60 - time.Duration(currentTime.Second())
		time.Sleep(duration * time.Second)
	}
}
