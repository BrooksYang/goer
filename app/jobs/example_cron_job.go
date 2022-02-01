package jobs

import (
	"log"
)

type ExampleCronJob struct{}

// Run Job
func (f ExampleCronJob) Run() {
	log.Println("example job running...")
}
