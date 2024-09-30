package app

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPackCalculator(t *testing.T) {
	testCases := []struct {
		name        string
		packSizes   []int32
		expected    []int32
		expectedErr bool
	}{

		{
			name:        "Empty",
			packSizes:   []int32{},
			expectedErr: true,
		},
		{
			name:      "Single pack Size",
			packSizes: []int32{500},
			expected:  []int32{500},
		},
		{
			name:      "Already sorted pack sizes",
			packSizes: []int32{5000, 2000, 1000, 500, 250},
			expected:  []int32{5000, 2000, 1000, 500, 250},
		},
		{
			name:      "Unsorted pack sizes",
			packSizes: []int32{1000, 250, 5000, 2000, 500},
			expected:  []int32{5000, 2000, 1000, 500, 250},
		},
		{
			name:      "Duplicate pack sizes",
			packSizes: []int32{500, 1000, 250, 1000, 5000},
			expected:  []int32{5000, 1000, 500, 250},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			packSizes := slices.Clone(tc.packSizes)
			pc, err := NewPackCalculator(tc.packSizes)
			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Verify that the returned packCalculator has the expected sorted pack sizes
			assert.Equal(t, tc.expected, pc.(*packCalculator).sizes)

			// Verify that the original packSizes slice is not modified
			assert.Equal(t, tc.packSizes, packSizes)
		})
	}
}

func TestCalculatePacksAmounts(t *testing.T) {
	// Define test cases using a table-driven approach
	testCases := []struct {
		name      string
		items     int32
		packSizes []int32
		expected  []PacksAmount
	}{
		{
			name:      "Simple case",
			items:     501,
			packSizes: []int32{2000, 1000, 500, 250, 200, 100, 50},
			expected: []PacksAmount{
				{Size: 500, Amount: 1},
				{Size: 50, Amount: 1},
			},
		},
		{
			name:      "No packs needed",
			items:     0,
			packSizes: []int32{2000, 1000, 500, 250},
			expected:  []PacksAmount{},
		},
		{
			name:      "Largest pack only",
			items:     10000,
			packSizes: []int32{5000, 2000, 1000, 500, 250},
			expected: []PacksAmount{
				{Size: 5000, Amount: 2},
			},
		},
		{
			name:      "Mix of packs",
			items:     12001,
			packSizes: []int32{5000, 2000, 1000, 500, 250},
			expected: []PacksAmount{
				{Size: 5000, Amount: 2},
				{Size: 2000, Amount: 1},
				{Size: 250, Amount: 1},
			},
		},
		{
			name:      "optimal packs Size",
			items:     251,
			packSizes: []int32{2000, 1000, 500, 250},
			expected: []PacksAmount{
				{Size: 500, Amount: 1},
			},
		},
		{
			name:      "Edge case - bigger than items",
			items:     10,
			packSizes: []int32{11, 3, 2},
			expected: []PacksAmount{
				{Size: 3, Amount: 2},
				{Size: 2, Amount: 2},
			},
		},
		{
			name:      "Edge case - prime numbers",
			items:     120,
			packSizes: []int32{31, 29},
			expected: []PacksAmount{
				{Size: 31, Amount: 2},
				{Size: 29, Amount: 2},
			},
		},
		{
			name:      "Edge case - one extra item",
			items:     1999,
			packSizes: []int32{2000, 1000, 2},
			expected: []PacksAmount{
				{Size: 2000, Amount: 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t,
				tc.expected,
				packCalculator{sizes: tc.packSizes}.CalculatePacksAmounts(tc.items),
				"Unexpected pack amounts for items: %d",
				tc.items)
		})
	}
}
