package sf_runtime

import (
	"context"
	"fmt"
	rawGRPC "google.golang.org/grpc"
	sfComponent "mosn.io/layotto/cmd/layotto_multiple_api/sf_runtime/component"
	"mosn.io/layotto/components/lock"
	"mosn.io/layotto/pkg/grpc"
	grpc_api "mosn.io/layotto/pkg/grpc"
	"mosn.io/pkg/log"
)

const (
	componentTypeTest   = "test"
	componentTypeStore  = "store"
	componentTypePubsub = "pubsub"
	componentTypeLock   = "lock"
	componentTypeFourA  = "4A"
)

const (
	testHello    = "hello"
	lockRedis    = "redis"
	pubsubPulsar = "pubsub"
	fourAXXX     = "xxx"
)

func NewSFAPI(ac *grpc_api.ApplicationContext) grpc.GrpcAPI {
	name2component := make(map[string]sfComponent.Hello)
	if len(ac.CustomComponent) != 0 {
		// we only care about those components of type "helloworld"
		name2comp, ok := ac.CustomComponent[componentTypeTest]
		if ok && len(name2comp) > 0 {
			for name, v := range name2comp {
				// convert them using type assertion
				comp, ok := v.(sfComponent.Hello)
				if !ok {
					errMsg := fmt.Sprintf("custom component %s does not implement HelloWorld interface", name)
					log.DefaultLogger.Errorf(errMsg)
				}
				name2component[name] = comp
			}
		}
	}
	return &sfServer{appId: ac.AppId}
}

type sfServer struct {
	appId            string
	test2component   map[string]sfComponent.Hello
	store2Component  map[string]lock.LockStore
	pubsub2Component map[string]lock.LockStore
	lock2Component   map[string]lock.LockStore
	fourAComponent   map[string]lock.LockStore
	sfComponent.UnimplementedGreeterServer
	AppCallbackConn *rawGRPC.ClientConn
}

func (s *sfServer) Init(conn *rawGRPC.ClientConn) error {
	// 1. set connection
	s.AppCallbackConn = conn
	return nil
}

func (s *sfServer) Register(rawGrpcServer *rawGRPC.Server) error {
	sfComponent.RegisterGreeterServer(rawGrpcServer, s)
	return nil
}

func (s *sfServer) SayHello(ctx context.Context, in *sfComponent.HelloRequest) (*sfComponent.HelloReply, error) {

	return &sfComponent.HelloReply{Message: "1233456789"}, nil
}
