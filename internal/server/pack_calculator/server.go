//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=../../../api/proto/ pack_calculator.proto
package pack_calculator

import (
	"context"
	"fmt"

	"github.com/Yapanyushin/tabeo-challenge/internal/app"
)

func NewServer(packCalculator app.PackCalculator) PackCalculatorServer {
	return &packsCalculator{
		packCalculator: packCalculator,
	}
}

type packsCalculator struct {
	packCalculator app.PackCalculator
	UnimplementedPackCalculatorServer
}

// CalculatePack handles CalculatePacksAmountRequest
func (p packsCalculator) CalculatePack(ctx context.Context, request *CalculatePacksAmountRequest) (*CalculatePacksAmountResponse, error) {
	if request.Items <= 0 {
		return nil, fmt.Errorf("items must be more than 0")
	}

	result := p.packCalculator.CalculatePacksAmounts(request.Items)

	packs := make([]*PacksAmount, len(result))

	for i, pack := range result {
		packs[i] = &PacksAmount{
			Size:   pack.Size,
			Amount: pack.Amount,
		}
	}

	return &CalculatePacksAmountResponse{
		Packs: packs,
	}, nil

}
