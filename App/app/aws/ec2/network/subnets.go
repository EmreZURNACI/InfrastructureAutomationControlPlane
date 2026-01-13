package network

import (
	"context"
	"errors"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/ptr"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ListSubnetRequest struct {
	VpcID string `json:"vpc_id" validate:"required"`
}
type ListSubnetResponse struct {
	Subnets []Subnet `json:"subnets"`
	Message *string  `json:"message"`
}
type ListSubnetHandler struct {
	client ports.NetworkClient
}

func NewListSubnetHandler(client ports.NetworkClient) *ListSubnetHandler {
	return &ListSubnetHandler{
		client: client,
	}
}

func (h *ListSubnetHandler) Handle(ctx context.Context, req *ListSubnetRequest) (*ListSubnetResponse, error) {

	out, err := h.client.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{
		Filters: []types.Filter{
			{
				Name:   ptr.String("vpc-id"),
				Values: []string{req.VpcID},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(out.Subnets) == 0 {
		return nil, errors.New("no data found")
	}

	var subnets []Subnet
	for _, value := range out.Subnets {

		ipv6s := make([]ipv6Association, 0, len(value.Ipv6CidrBlockAssociationSet))
		for _, ipv6 := range value.Ipv6CidrBlockAssociationSet {
			ipv6s = append(ipv6s, ipv6Association{
				AssociationId:    ipv6.AssociationId,
				CidrBlock:        ipv6.Ipv6CidrBlock,
				Status:           ipv6.Ipv6CidrBlockState.StatusMessage,
				AddressAttribute: ipv6.Ipv6AddressAttribute,
			})
		}

		subnets = append(subnets, Subnet{
			VpcID:                       value.VpcId,
			AvailabilityZone:            value.AvailabilityZone,
			AvailableIPAddressCount:     value.AvailableIpAddressCount,
			CidrBlock:                   value.CidrBlock,
			MapCustomerOwnedIPOnLaunch:  value.MapCustomerOwnedIpOnLaunch,
			MapPublicIPOnLaunch:         value.MapPublicIpOnLaunch,
			State:                       value.State,
			SubnetID:                    value.SubnetId,
			OwnerID:                     value.OwnerId,
			AssignIpv6AddressOnCreation: value.AssignIpv6AddressOnCreation,
			Ipv6CidrBlockAssociationSet: ipv6s,
			SubnetArn:                   value.SubnetArn,
			EnableDNS:                   value.EnableDns64,
			Ipv6Native:                  value.Ipv6Native,
			PrivateDNSNameOptionsOnLaunch: PrivateDNSNameOptionsOnLaunch{
				HostnameType:                    value.PrivateDnsNameOptionsOnLaunch.HostnameType,
				EnableResourceNameDNSARecord:    value.PrivateDnsNameOptionsOnLaunch.EnableResourceNameDnsARecord,
				EnableResourceNameDNSAAAARecord: value.PrivateDnsNameOptionsOnLaunch.EnableResourceNameDnsAAAARecord,
			},
		})
	}

	return &ListSubnetResponse{
		Subnets: subnets,
		Message: ptr.String("subnets listed successfully"),
	}, nil
}

type Subnet struct {
	VpcID                         *string                       `json:"vpc_id"`
	AvailabilityZone              *string                       `json:"availability_zone"`
	AvailableIPAddressCount       *int32                        `json:"availability_ip_address_count"`
	CidrBlock                     *string                       `json:"cidr_block"`
	MapPublicIPOnLaunch           *bool                         `json:"map_public_ip_on_launch"`
	MapCustomerOwnedIPOnLaunch    *bool                         `json:"map_customer_owned_ip_on_launch"`
	State                         types.SubnetState             `json:"state"`
	SubnetID                      *string                       `json:"subnet_id"`
	OwnerID                       *string                       `json:"owner_id"`
	AssignIpv6AddressOnCreation   *bool                         `json:"assign_ipv6_adress_on_creation"`
	Ipv6CidrBlockAssociationSet   []ipv6Association             `json:"ipv6_cidr_block_association_set"`
	SubnetArn                     *string                       `json:"subnet_arn"`
	EnableDNS                     *bool                         `json:"enable_dns"`
	Ipv6Native                    *bool                         `json:"ipv6_dns"`
	PrivateDNSNameOptionsOnLaunch PrivateDNSNameOptionsOnLaunch `json:"private_dns_name_options_on_launch"`
}

type PrivateDNSNameOptionsOnLaunch struct {
	HostnameType                    types.HostnameType `json:"hostname_type"`
	EnableResourceNameDNSARecord    *bool              `json:"enable_resource_name_dns_record"`
	EnableResourceNameDNSAAAARecord *bool              `json:"enable_resource_name_dns_aa_record"`
}
type ipv6Association struct {
	AssociationId    *string
	AddressAttribute types.Ipv6AddressAttribute
	CidrBlock        *string
	Status           *string
}
