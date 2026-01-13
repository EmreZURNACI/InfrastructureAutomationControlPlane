package network

import (
	"context"
	"errors"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/ptr"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ListVPCRequest struct {
}
type ListVPCResponse struct {
	Vpcs    []VPC   `json:"vpcs"`
	Message *string `json:"message"`
}
type ListVPCHandler struct {
	client ports.NetworkClient
}

func NewListVPCHandler(client ports.NetworkClient) *ListVPCHandler {
	return &ListVPCHandler{
		client: client,
	}
}

func (h *ListVPCHandler) Handle(ctx context.Context, req *ListVPCRequest) (*ListVPCResponse, error) {
	out, err := h.client.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{})
	if err != nil {
		return nil, err
	}

	if len(out.Vpcs) == 0 {
		return nil, errors.New("vpc bulunamadı")
	}

	var vpcs []VPC
	for _, v := range out.Vpcs {

		ipv4s := make([]ipv4BlockAssociation, 0, len(v.CidrBlockAssociationSet))
		for _, ipv4 := range v.CidrBlockAssociationSet {
			ipv4s = append(ipv4s, ipv4BlockAssociation{
				AssociationId: ipv4.AssociationId,
				CidrBlock:     ipv4.CidrBlock,
				StatusMessage: ipv4.CidrBlockState.StatusMessage,
			})
		}

		ipv6s := make([]ipv6BlockAssociation, 0, len(v.Ipv6CidrBlockAssociationSet))
		for _, ipv6 := range v.Ipv6CidrBlockAssociationSet {
			ipv6s = append(ipv6s, ipv6BlockAssociation{
				AssociationId:        ipv6.AssociationId,
				Ipv6CidrBlock:        ipv6.Ipv6CidrBlock,
				StatusMessage:        ipv6.Ipv6CidrBlockState.StatusMessage,
				Ipv6Pool:             ipv6.Ipv6Pool,
				NetworkBorderGroup:   ipv6.NetworkBorderGroup,
				Ipv6AddressAttribute: ipv6.Ipv6AddressAttribute,
			})
		}

		vpcs = append(vpcs, VPC{
			ID:                      v.VpcId,
			State:                   v.State,
			OwnerID:                 v.OwnerId,
			IsDefault:               v.IsDefault,
			CidrBlock:               v.CidrBlock,
			IPv4:                    ipv4s,
			IPv6:                    ipv6s,
			BlockPublicAccessStates: v.BlockPublicAccessStates,
			DhcpOptionsID:           v.DhcpOptionsId,
			InstanceTenancy:         v.InstanceTenancy,
			EncryptionControl:       v.EncryptionControl,
		})
	}

	return &ListVPCResponse{
		Vpcs:    vpcs,
		Message: ptr.String("vpcs listed successfully"),
	}, nil
}

type VPC struct {
	ID                      *string                        `json:"id"`
	State                   types.VpcState                 `json:"state"`
	OwnerID                 *string                        `json:"owner_id"`
	IsDefault               *bool                          `json:"is_default"`
	CidrBlock               *string                        `json:"cidr_block"`
	IPv4                    []ipv4BlockAssociation         `json:"ipv4"`
	IPv6                    []ipv6BlockAssociation         `json:"ipv6"`
	BlockPublicAccessStates *types.BlockPublicAccessStates `json:"block_public_access_states"`
	DhcpOptionsID           *string                        `json:"dhcp_options_id"`
	InstanceTenancy         types.Tenancy                  `json:"instance_tenancy"`
	EncryptionControl       *types.VpcEncryptionControl    `json:"encryption_control"`
}
type ipv4BlockAssociation struct {
	AssociationId *string
	CidrBlock     *string
	StatusMessage *string
}
type ipv6BlockAssociation struct {
	AssociationId        *string
	Ipv6AddressAttribute types.Ipv6AddressAttribute
	Ipv6CidrBlock        *string
	StatusMessage        *string
	Ipv6Pool             *string
	NetworkBorderGroup   *string
}
