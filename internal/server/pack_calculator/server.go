package pack_calculator

import (
	"context"
	"fmt"

	"github.com/Yapanyushin/tabeo-challenge/api/proto"
	"github.com/Yapanyushin/tabeo-challenge/internal/app"
)

func NewServer(packCalculator app.PackCalculator) proto.PackCalculatorServer {
	return &packsCalculator{
		packCalculator: packCalculator,
	}
}

type packsCalculator struct {
	packCalculator app.PackCalculator
	proto.UnimplementedPackCalculatorServer
}

// CalculatePack handles CalculatePacksAmountRequest
func (p packsCalculator) CalculatePack(ctx context.Context, request *proto.CalculatePacksAmountRequest) (*proto.CalculatePacksAmountResponse, error) {
	if request.Items <= 0 {
		return nil, fmt.Errorf("items must be more than 0")
	}

	result := p.packCalculator.CalculatePacksAmounts(request.Items)

	packs := make([]*proto.PacksAmount, len(result))

	for i, pack := range result {
		packs[i] = &proto.PacksAmount{
			Size:   pack.Size,
			Amount: pack.Amount,
		}
	}

	return &proto.CalculatePacksAmountResponse{
		Packs: packs,
	}, nil

}
