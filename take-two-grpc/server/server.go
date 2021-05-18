package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"sync"
	"time"

	pb "github.com/BrowduesMan85/out-of-process/take-two-grpc/worker"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type workerServer struct {
	pb.UnimplementedWorkerServer
	executors          sync.Map // executors is a concurrent map that maps executorID:executor
	executorSubscribed chan bool
}

type executor struct {
	finished chan bool                           // finished signals that an executor is finished executing work
	stream   pb.Worker_CommmandMessageLinkServer // stream is the server side of the RPC stream
}

func (s *workerServer) CommmandMessageLink(stream pb.Worker_CommmandMessageLinkServer) error {
	ctx := stream.Context()

	registration, err := stream.Recv()
	if err != nil {
		return err
	}

	fin := make(chan bool)
	s.executors.Store(registration.ExecutorId, executor{stream: stream, finished: fin})
	fmt.Printf("Worker received Executor ID %q registration request\n", registration.ExecutorId)

	s.executorSubscribed <- true

	select {
	case <-fin:
		log.Printf("Closing stream for executor ID %q", registration.ExecutorId)
		return nil
	case <-ctx.Done():
		log.Printf("Executor ID %q disconnected", registration.ExecutorId)
		return nil
	}
}

func (s *workerServer) RunWorkflow(executorID string, workflowXML string, parametersXML string) error {
	v, ok := s.executors.Load(executorID)
	if !ok {
		log.Fatal("Error loading executor")
	}

	e, ok := v.(executor)
	if !ok {
		log.Fatal("Error type casting executor")
	}

	workerMsg := pb.WorkerMessage{
		Command:       "run",
		WorkflowXml:   []byte(workflowXML),
		ParametersXml: []byte(parametersXML),
	}

	return e.stream.Send(&workerMsg)
}

func newServer(streamCh chan bool) *workerServer {
	return &workerServer{
		executorSubscribed: streamCh,
	}
}

func startWorkerServer(streamCh chan bool, socket string) (*workerServer, error) {
	if err := os.RemoveAll(socket); err != nil {
		return nil, err
	}

	lis, err := net.Listen("unix", socket)
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer()

	server := newServer(streamCh)
	pb.RegisterWorkerServer(grpcServer, server)

	// TODO: Handle serve error
	go grpcServer.Serve(lis)

	return server, nil
}

func LaunchExecutor(executorID string, socket string) {
	cmdPath := "/Users/nbrowdues/Workspace/github.com/BrowduesMan85/out-of-process/take-two-grpc/client/client"
	cmd := exec.Command(cmdPath, executorID, socket)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Start()
}

func StartWorkerExecutor(executorID string, socket string) {
	executorSubscribed := make(chan bool)
	workerServer, _ := startWorkerServer(executorSubscribed, socket)

	LaunchExecutor(executorID, socket)
	<-executorSubscribed

	fmt.Println("Worker executing RunWorkflowCommand")
	workerServer.RunWorkflow(executorID, "workflowXML", "parametersXML")
}

func main() {
	go StartWorkerExecutor("fake-id-1", fmt.Sprintf("/tmp/%s", uuid.New()))
	go StartWorkerExecutor("fake-id-2", fmt.Sprintf("/tmp/%s", uuid.New()))
	time.Sleep(15 * time.Second)
}
