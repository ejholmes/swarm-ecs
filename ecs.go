package ecs

import (
	"github.com/ejholmes/ecs/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws"
	"github.com/ejholmes/ecs/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/ejholmes/ecs/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/ec2"
	"github.com/ejholmes/ecs/Godeps/_workspace/src/github.com/aws/aws-sdk-go/service/ecs"
)

var DefaultClient *Client

func init() {
	config := defaults.DefaultConfig
	DefaultClient = &Client{
		ecs: ecs.New(config),
		ec2: ec2.New(config),
	}
}

type Client struct {
	ecs ecsClient
	ec2 ec2Client
}

// Returns the ec2 Instances in this ecs cluster.
func (c *Client) Instances(cluster string) ([]*ec2.Instance, error) {
	var instances []*ec2.Instance

	var instanceArns []*string
	if err := c.ecs.ListContainerInstancesPages(&ecs.ListContainerInstancesInput{
		Cluster: aws.String(cluster),
	}, func(p *ecs.ListContainerInstancesOutput, lastPage bool) bool {
		instanceArns = append(instanceArns, p.ContainerInstanceArns...)
		return true
	}); err != nil {
		return instances, err
	}

	if len(instanceArns) == 0 {
		return instances, nil
	}

	ecsDesc, err := c.ecs.DescribeContainerInstances(&ecs.DescribeContainerInstancesInput{
		Cluster:            aws.String(cluster),
		ContainerInstances: instanceArns,
	})
	if err != nil {
		return instances, err
	}

	var containerInstances []*string
	for _, instance := range ecsDesc.ContainerInstances {
		containerInstances = append(containerInstances, instance.Ec2InstanceId)
	}

	ec2Desc, err := c.ec2.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-id"),
				Values: containerInstances,
			},
		},
	})
	if err != nil {
		return instances, err
	}

	for _, reservation := range ec2Desc.Reservations {
		instances = append(instances, reservation.Instances...)
	}

	return instances, nil
}

func Instances(cluster string) ([]*ec2.Instance, error) {
	return DefaultClient.Instances(cluster)
}

type ecsClient interface {
	ListContainerInstancesPages(*ecs.ListContainerInstancesInput, func(*ecs.ListContainerInstancesOutput, bool) bool) error
	DescribeContainerInstances(*ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error)
}

type ec2Client interface {
	DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error)
}
