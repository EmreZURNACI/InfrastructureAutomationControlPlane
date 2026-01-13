package network

import (
	"context"
	"errors"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/ptr"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ListSecurityGroupRequest struct {
}
type ListSecurityGroupResponse struct {
	SecurityGroups []SecurityGroup `json:"security_group"`
	Message        *string         `json:"message"`
}
type ListSecurityGroupHandler struct {
	client ports.NetworkClient
}

func NewListSecurityGroupHandler(client ports.NetworkClient) *ListSecurityGroupHandler {
	return &ListSecurityGroupHandler{
		client: client,
	}
}

func (h *ListSecurityGroupHandler) Handle(ctx context.Context, req *ListSecurityGroupRequest) (*ListSecurityGroupResponse, error) {

	out, err := h.client.DescribeSecurityGroups(ctx, &ec2.DescribeSecurityGroupsInput{})
	if err != nil {
		return nil, err
	}

	if len(out.SecurityGroups) == 0 {
		return nil, errors.New("no data found")
	}

	var security_groups []SecurityGroup
	for _, value := range out.SecurityGroups {
		security_groups = append(security_groups, SecurityGroup{
			ID:                  value.VpcId,
			Description:         value.Description,
			GroupID:             value.GroupId,
			GroupName:           value.GroupName,
			IpPermissions:       value.IpPermissions,
			IpPermissionsEgress: value.IpPermissionsEgress,
			OwnerID:             value.OwnerId,
			Arn:                 value.SecurityGroupArn,
		})
	}

	return &ListSecurityGroupResponse{
		SecurityGroups: security_groups,
		Message:        ptr.String("security groups listed successfully"),
	}, nil
}

type SecurityGroup struct {
	ID *string `json:"id"`

	Description *string `json:"description"`

	GroupID *string `json:"group_id"`

	GroupName *string `json:"group_name"`

	IpPermissions []types.IpPermission `json:"ip_permissions"`

	IpPermissionsEgress []types.IpPermission `json:"ip_permissions_egress"`

	OwnerID *string `json:"owner_id"`

	Arn *string `json:"Arn"`
}
