package cmd

import (
	"time"

	"goer/global"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

var CmdCron = &cobra.Command{
	Use:   "cron",
	Short: "Start cron job",
	Run:   runCron,
}

func init() {
	rootCmd.AddCommand(CmdCron)
}

func runCron(cmd *cobra.Command, args []string) {
	// Load location
	location, _ := time.LoadLocation(global.Config.App.Timezone)

	// New Cron
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(cron.DelayIfStillRunning(cron.DefaultLogger)),
		cron.WithLocation(location),
	)

	// Schedule
	schedule(c)

	// Start
	c.Start()

	// Keep awake
	select {}
}

// Schedule
func schedule(c *cron.Cron) {
	// Example cron job
	// _, _ = c.AddJob("@every 1s", &jobs.ExampleCronJob{})
}
