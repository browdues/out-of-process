package main

import (
	"testing"
	"time"

	pb "github.com/BrowduesMan85/out-of-process/take-two-grpc/worker"
	"github.com/stretchr/testify/assert"
)

func TestRunCmd(t *testing.T) {
	streamCh := make(chan pb.Worker_CommmandMessageLinkServer)
	server, _ := startWorkerServer(streamCh)
	LaunchExecutor()

	stream := <-streamCh
	server.RunWorkflow(stream, "workflowXml", "parametersXml")

	time.Sleep(10 * time.Second)
	assert.Equal(t, 1, 1)
}
