package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	myTest "myTest/proto"
)

type MyTest struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *MyTest) Call(ctx context.Context, req *myTest.Request, rsp *myTest.Response) error {
	log.Info("Received MyTest.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *MyTest) Stream(ctx context.Context, req *myTest.StreamingRequest, stream myTest.MyTest_StreamStream) error {
	log.Infof("Received MyTest.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&myTest.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *MyTest) PingPong(ctx context.Context, stream myTest.MyTest_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&myTest.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
