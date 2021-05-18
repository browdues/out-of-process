package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/BrowduesMan85/out-of-process/take-two-grpc/worker"
	"google.golang.org/grpc"
)

func runRegisterExecutor(ctx context.Context, stream pb.Worker_CommmandMessageLinkClient, executorID string, timeout time.Duration) {
	message := pb.ExecutorMessage{
		ExecutorId: executorID,
	}

	_, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := stream.Send(&message); err != nil {
		log.Fatalf("Failed to send a message: %v", err)
	}
	stream.CloseSend()
}

func handleRunnerCommand(ctx context.Context, cancel context.CancelFunc, stream pb.Worker_CommmandMessageLinkClient, executorID string) {

	getCommand := func() string {
		response, err := stream.Recv()
		if err == io.EOF {
			return "EOF"
		}
		if err != nil {
			log.Fatalf("Failed to receive a command : %v", err)
		}
		return response.Command
	}

	for {
		c := getCommand()
		if c == "run" {
			fmt.Printf("Executor ID %q received run command\n", executorID)
			go func() {
				time.Sleep(5000 * time.Millisecond)
				cancel()
			}()
		} else if c == "cancel" {
			fmt.Println("Received cancel command, cancelling now!")
			cancel()
			break
		}
	}
}

func newStream(client pb.WorkerClient, timeout time.Duration) (context.Context, context.CancelFunc, pb.Worker_CommmandMessageLinkClient) {
	streamCtx, streamCancel := context.WithCancel(context.Background())
	stream, err := client.CommmandMessageLink(streamCtx)
	if err != nil {
		log.Fatalf("%v.CommandMessageLink(_) = _, %v", client, err)
	}
	return streamCtx, streamCancel, stream
}

func newClient(socket string) (pb.WorkerClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(fmt.Sprintf("unix:%v", socket), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return pb.NewWorkerClient(conn), conn
}

func main() {
	executorID := os.Args[1]
	socket := os.Args[2]

	client, conn := newClient(socket)
	defer conn.Close()

	streamCtx, streamCancel, stream := newStream(client, 10*time.Second)
	done := streamCtx.Done()
	go handleRunnerCommand(streamCtx, streamCancel, stream, executorID)
	runRegisterExecutor(streamCtx, stream, executorID, 10*time.Second)
	<-done
	fmt.Printf("Executor ID %q work is finished\n", executorID)
}
