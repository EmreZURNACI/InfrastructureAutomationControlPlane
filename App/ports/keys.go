package ports

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type KeyClient interface {
	DescribeKeys(ctx context.Context, params *ec2.DescribeKeyPairsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error)
}
