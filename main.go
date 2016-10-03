package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/chazsmi/dragon/aws"
)

func main() {
	sess := session.New()
	//log.Println(sess)
	svc := ec2.New(sess, &aws.Config{
		Region: aws.String("us-west-2"),
	})

	// Call the DescribeInstances Operation
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		log.Printf("%+v", err)
	}

	// resp has all of the response data, pull out instance IDs:
	fmt.Println("> Number of reservation sets: ", len(resp.Reservations))
	ins := autoscaling.CollectInstances(resp)

	log.Println("%+v", ins)
	autoscaling.TerminateInstance(ins[rand.Intn(len(ins))], sess)
}
