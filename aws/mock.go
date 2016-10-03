package autoscaling

import (
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
)

type mockAutoScalingClient struct {
	autoscalingiface.AutoScalingAPI
}

func (m *mockAutoScalingClient) AttachInstances(input *autoscaling.AttachInstancesInput) (*autoscaling.AttachInstancesOutput, error) {
	return &autoscaling.AttachInstancesOutput{}, nil
}

func (m *mockAutoScalingClient) TerminateInstanceInAutoScalingGroup(*autoscaling.TerminateInstanceInAutoScalingGroupInput) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error) {
	return &autoscaling.TerminateInstanceInAutoScalingGroupOutput{}, nil
}
