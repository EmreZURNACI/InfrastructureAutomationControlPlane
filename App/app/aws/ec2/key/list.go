package key

import (
	"context"
	"errors"
	"time"

	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/pkg/ptr"
	"github.com/EmreZURNACI/InfrastructureAutomationControlPlane/ports"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ListRequest struct {
}
type ListResponse struct {
	Keys    []Key   `json:"keys"`
	Message *string `json:"message"`
}
type ListHandler struct {
	client ports.KeyClient
}

func NewListHandler(client ports.KeyClient) *ListHandler {
	return &ListHandler{
		client: client,
	}
}
func (h *ListHandler) Handle(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	out, err := h.client.DescribeKeys(ctx, &ec2.DescribeKeyPairsInput{})
	if err != nil {
		return nil, err
	}

	if len(out.KeyPairs) == 0 {
		return nil, errors.New("no item found")
	}

	var keys []Key
	for _, value := range out.KeyPairs {
		keys = append(keys, Key{
			KeyName:        value.KeyName,
			KeyType:        value.KeyType,
			KeyPairID:      value.KeyPairId,
			PublicKey:      value.PublicKey,
			KeyFingerPrint: value.KeyFingerprint,
			CreateTime:     value.CreateTime,
		})
	}

	return &ListResponse{
		Message: ptr.String("keypairs listed successfully"),
		Keys:    keys,
	}, nil
}

type Key struct {
	KeyName        *string       `json:"key_name"`
	KeyType        types.KeyType `json:"key_type"`
	KeyPairID      *string       `json:"key_pair_id"`
	PublicKey      *string       `json:"public_key"`
	KeyFingerPrint *string       `json:"key_finger_print"`
	CreateTime     *time.Time    `json:"create_time"`
}
