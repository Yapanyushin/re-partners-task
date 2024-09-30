package pack_calculator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Yapanyushin/tabeo-challenge/internal/app"
	"github.com/Yapanyushin/tabeo-challenge/internal/app/mocks"
)

func TestCalculatePack(t *testing.T) {

	testCases := []struct {
		name            string
		request         *CalculatePacksAmountRequest
		expected        *CalculatePacksAmountResponse
		expectErr       bool
		mockExpectation func(m *mocks.PackCalculator_Expecter)
	}{
		{
			name: "Valid request",
			request: &CalculatePacksAmountRequest{
				Items: 501,
			},
			mockExpectation: func(m *mocks.PackCalculator_Expecter) {
				m.CalculatePacksAmounts(int32(501)).Return([]app.PacksAmount{
					{Size: 500, Amount: 1}, {Size: 250, Amount: 1},
				})
			},
			expected: &CalculatePacksAmountResponse{
				Packs: []*PacksAmount{{Size: 500, Amount: 1}, {Size: 250, Amount: 1}},
			},
		},
		{
			name: "Zero items",
			request: &CalculatePacksAmountRequest{
				Items: 0,
			},
			expectErr: true,
		},
		{
			name: "Negative items",
			request: &CalculatePacksAmountRequest{
				Items: -10,
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			m := mocks.NewPackCalculator(t)
			p := packsCalculator{packCalculator: m}
			if tc.mockExpectation != nil {
				tc.mockExpectation(m.EXPECT())
			}

			response, err := p.CalculatePack(context.Background(), tc.request)

			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, response)
			}
		})
	}
}
