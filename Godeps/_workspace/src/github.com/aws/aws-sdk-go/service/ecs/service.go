// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ecs

import (
	"github.com/ejholmes/ecs/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws"
	"github.com/ejholmes/ecs/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/ejholmes/ecs/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/request"
	"github.com/ejholmes/ecs/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/service"
	"github.com/ejholmes/ecs/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/ejholmes/ecs/Godeps/_workspace/src/github.com/aws/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/ejholmes/ecs/Godeps/_workspace/src/github.com/aws/aws-sdk-go/internal/signer/v4"
)

// Amazon EC2 Container Service (Amazon ECS) is a highly scalable, fast, container
// management service that makes it easy to run, stop, and manage Docker containers
// on a cluster of Amazon EC2 instances. Amazon ECS lets you launch and stop
// container-enabled applications with simple API calls, allows you to get the
// state of your cluster from a centralized service, and gives you access to
// many familiar Amazon EC2 features like security groups, Amazon EBS volumes,
// and IAM roles.
//
// You can use Amazon ECS to schedule the placement of containers across your
// cluster based on your resource needs, isolation policies, and availability
// requirements. Amazon EC2 Container Service eliminates the need for you to
// operate your own cluster management and configuration management systems
// or worry about scaling your management infrastructure.
type ECS struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new ECS client.
func New(config *aws.Config) *ECS {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:       defaults.DefaultConfig.Merge(config),
			ServiceName:  "ecs",
			APIVersion:   "2014-11-13",
			JSONVersion:  "1.1",
			TargetPrefix: "AmazonEC2ContainerServiceV20141113",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &ECS{service}
}

// newRequest creates a new request for a ECS operation and runs any
// custom request initialization.
func (c *ECS) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
