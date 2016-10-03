package autoscaling

import "testing"

func TestTerminateInstance(t *testing.T) {
	cli := mockAutoScalingClient{}
	TermainateInstance(cli)
}
