package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) {
	svc := ec2.New(session.New())
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String("i-1234567890abcdef0"),
		},
	}

	result, err := svc.StartInstances(input)

	// input := &ec2.StopInstancesInput{
	// 	InstanceIds: []*string{
	// 		aws.String("i-1234567890abcdef0"),
	// 	},
	// }

	// result, err := svc.StopInstances(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

func main() {
	lambda.Start(Handler)
}
