package autoscaling

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func TerminateInstance(id string, sess client.ConfigProvider) {
	svc := autoscaling.New(sess, &aws.Config{
		Region: aws.String("us-west-2"),
	})
	input := &autoscaling.TerminateInstanceInAutoScalingGroupInput{
		InstanceId:                     aws.String(id),
		ShouldDecrementDesiredCapacity: aws.Bool(false),
	}

	resp, err := svc.TerminateInstanceInAutoScalingGroup(input)
	if err != nil {
		log.Println(err.Error())
	}

	log.Printf("%+v", resp)
}

func CollectInstances(resp *ec2.DescribeInstancesOutput) []string {
	var ids []string
	for idx, res := range resp.Reservations {
		fmt.Println("  > Number of instances: ", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {
			ids = append(ids, *inst.InstanceId)
		}
	}
	return ids
}
