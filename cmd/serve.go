package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"goer/bootstrap"
	"goer/global"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start API server",
	Run:   runServe,
}

func init() {
	rootCmd.AddCommand(cmdServe)
}

func runServe(cmd *cobra.Command, args []string) {
	// Gin mode: debug, release, test
	if global.Config.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := bootstrap.SetupRouter()

	// http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.Config.App.Port),
		Handler: router,
	}

	// Listen and Server
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// log
	fmt.Printf("\r\n")
	fmt.Println("Server run at:")
	fmt.Printf("- Local: http://localhost:%d/ \r\n", global.Config.App.Port)
	fmt.Printf("\r\n")
	log.Printf("Enter Control + C Shutdown Server \r\n")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// log
	fmt.Printf("\r\n")
	log.Println("Shutdown Server ...")

	// Timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// log
	log.Println("Server exiting")
}
