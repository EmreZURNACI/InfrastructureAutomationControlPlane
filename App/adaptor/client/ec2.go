package client

import (
	"context"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/config"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Ec2Client struct {
	client *ec2.Client
}

func Connect() (*Ec2Client, error) {

	cfg := aws.Config{
		Region: config.AppConfig.AwsConfig.Region,
		Credentials: aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(
				config.AppConfig.AwsConfig.AccessKey,
				config.AppConfig.AwsConfig.SecretAccessKey,
				"",
			),
		),
	}

	client := ec2.NewFromConfig(cfg)

	// awsConfig, err := config.LoadDefaultConfig(
	// 	context.Background(),
	// 	config.WithRegion(cfg.AppConfig.AwsConfig.Region),
	// 	config.WithCredentialsProvider(
	// 		credentials.NewStaticCredentialsProvider(
	// 			cfg.AppConfig.AwsConfig.AccessKey,
	// 			cfg.AppConfig.AwsConfig.SecretAccessKey,
	// 			"",
	// 		),
	// 	),
	// )

	// if err != nil {
	// 	return nil, err
	// }

	// client := ec2.NewFromConfig(awsConfig)
	return &Ec2Client{
		client: client,
	}, nil

}

var _ ports.InstanceClient = (*Ec2Client)(nil)
var _ ports.VolumeClient = (*Ec2Client)(nil)
var _ ports.NetworkClient = (*Ec2Client)(nil)
var _ ports.ImageClient = (*Ec2Client)(nil)
var _ ports.KeyClient = (*Ec2Client)(nil)

// ///////////////////
func (c *Ec2Client) DescribeImages(ctx context.Context, params *ec2.DescribeImagesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error) {
	return c.client.DescribeImages(ctx, params, optFns...)
}

// ///////////////////
func (c *Ec2Client) DescribeKeys(ctx context.Context, params *ec2.DescribeKeyPairsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error) {
	return c.client.DescribeKeyPairs(ctx, params, optFns...)
}

// ///////////////////
func (c *Ec2Client) DescribeVpcs(ctx context.Context, params *ec2.DescribeVpcsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcsOutput, error) {
	return c.client.DescribeVpcs(ctx, params, optFns...)
}
func (c *Ec2Client) DescribeSubnets(ctx context.Context, params *ec2.DescribeSubnetsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error) {
	return c.client.DescribeSubnets(ctx, params, optFns...)
}
func (c *Ec2Client) DescribeSecurityGroups(ctx context.Context, params *ec2.DescribeSecurityGroupsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error) {
	return c.client.DescribeSecurityGroups(ctx, params, optFns...)
}

// ///////////////////
func (c *Ec2Client) RunInstances(ctx context.Context, params *ec2.RunInstancesInput, optFns ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error) {
	return c.client.RunInstances(ctx, params, optFns...)
}
func (c *Ec2Client) DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	return c.client.DescribeInstances(ctx, params, optFns...)
}
func (c *Ec2Client) EditInstances(ctx context.Context, params *ec2.ModifyInstanceAttributeInput, optFns ...func(*ec2.Options)) (*ec2.ModifyInstanceAttributeOutput, error) {
	return c.client.ModifyInstanceAttribute(ctx, params, optFns...)
}
func (c *Ec2Client) StartInstances(ctx context.Context, params *ec2.StartInstancesInput, optFns ...func(*ec2.Options)) (*ec2.StartInstancesOutput, error) {
	return c.client.StartInstances(ctx, params, optFns...)
}
func (c *Ec2Client) StopInstances(ctx context.Context, params *ec2.StopInstancesInput, optFns ...func(*ec2.Options)) (*ec2.StopInstancesOutput, error) {
	return c.client.StopInstances(ctx, params, optFns...)
}
func (c *Ec2Client) TerminateInstances(ctx context.Context, params *ec2.TerminateInstancesInput, optFns ...func(*ec2.Options)) (*ec2.TerminateInstancesOutput, error) {
	return c.client.TerminateInstances(ctx, params, optFns...)
}
func (c *Ec2Client) RestartInstances(ctx context.Context, params *ec2.RebootInstancesInput, optFns ...func(*ec2.Options)) (*ec2.RebootInstancesOutput, error) {
	return c.client.RebootInstances(ctx, params, optFns...)
}
func (c *Ec2Client) ListInstanceTypes(ctx context.Context, params *ec2.DescribeInstanceTypesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstanceTypesOutput, error) {
	return c.client.DescribeInstanceTypes(ctx, params, optFns...)
}

// ///////////////////
func (c *Ec2Client) AttachVolume(ctx context.Context, params *ec2.AttachVolumeInput, optFns ...func(*ec2.Options)) (*ec2.AttachVolumeOutput, error) {
	return c.client.AttachVolume(ctx, params, optFns...)
}

func (c *Ec2Client) DetachVolume(ctx context.Context, params *ec2.DetachVolumeInput, optFns ...func(*ec2.Options)) (*ec2.DetachVolumeOutput, error) {
	return c.client.DetachVolume(ctx, params, optFns...)
}

func (c *Ec2Client) CreateVolume(ctx context.Context, params *ec2.CreateVolumeInput, optFns ...func(*ec2.Options)) (*ec2.CreateVolumeOutput, error) {
	return c.client.CreateVolume(ctx, params, optFns...)
}

func (c *Ec2Client) DeleteVolume(ctx context.Context, params *ec2.DeleteVolumeInput, optFns ...func(*ec2.Options)) (*ec2.DeleteVolumeOutput, error) {
	return c.client.DeleteVolume(ctx, params, optFns...)
}

func (c *Ec2Client) EditVolume(ctx context.Context, params *ec2.ModifyVolumeInput, optFns ...func(*ec2.Options)) (*ec2.ModifyVolumeOutput, error) {
	return c.client.ModifyVolume(ctx, params, optFns...)
}

func (c *Ec2Client) DescribeVolumes(ctx context.Context, params *ec2.DescribeVolumesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVolumesOutput, error) {
	return c.client.DescribeVolumes(ctx, params, optFns...)
}

func (c *Ec2Client) CreateSnapshot(ctx context.Context, params *ec2.CreateSnapshotInput, optFns ...func(*ec2.Options)) (*ec2.CreateSnapshotOutput, error) {
	return c.client.CreateSnapshot(ctx, params, optFns...)
}

func (c *Ec2Client) DeleteSnapshot(ctx context.Context, params *ec2.DeleteSnapshotInput, optFns ...func(*ec2.Options)) (*ec2.DeleteSnapshotOutput, error) {
	return c.client.DeleteSnapshot(ctx, params, optFns...)
}
