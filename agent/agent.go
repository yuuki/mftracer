package agent

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yuuki/mkr-flow-tracer/collector"
	"github.com/yuuki/mkr-flow-tracer/datastore"
)

// Start starts agent.
func Start(interval time.Duration) {
	go Watch(interval)

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGTERM, syscall.SIGINT)
	sig := <-sigch
	log.Printf("Received %s gracefully shutdown...\n", sig)

	time.Sleep(3 * time.Second)
}

// Watch watches host flows for localhost.
func Watch(interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	errChan := make(chan error, 1)
	for {
		select {
		case err := <-errChan:
			if err != nil {
				return err
			}
		case <-ticker.C:
			go CollectAndPostHostFlows(errChan)
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

// CollectAndPostHostFlows collect host flows and
// post it to the data store.
func CollectAndPostHostFlows(errChan chan error) {
	flows, err := collector.CollectHostFlows()
	if err != nil {
		errChan <- err
		return
	}
	errChan <- datastore.PostHostFlows(flows)
	log.Printf("Post host flows (%d)", len(flows))
}
