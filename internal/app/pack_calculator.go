//go:generate mockery --all --with-expecter=true
package app

import (
	"fmt"
	"slices"
	"sort"
)

type PackCalculator interface {
	CalculatePacksAmounts(items int32) []PacksAmount
}

// NewPackCalculator returns service interface
func NewPackCalculator(packSizes []int32) (PackCalculator, error) {
	if len(packSizes) == 0 {
		return nil, fmt.Errorf("empty pack sizes")
	}
	filtered := make(map[int32]bool)

	for _, packSize := range packSizes {
		filtered[packSize] = true
	}

	sorted := make([]int32, 0, len(filtered))
	for packSize := range filtered {
		sorted = append(sorted, packSize)
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i] < sorted[j] {
			return false
		}
		return true
	})

	return &packCalculator{
		sizes: sorted,
	}, nil
}

type packCalculator struct {
	sizes []int32
}

type PacksAmount struct {
	Size   int32
	Amount int32
}

// CalculatePacksAmounts calculates the optimal number of packs to ship to customers based on their
// order quantity and available pack sizes.
// It adheres to the following rules:
//
// Whole Packs Only: Only complete packs can be shipped; no breaking packs open.
// Minimize Items: Fulfill the order with the least number of items possible.
// Minimize Packs: Use the fewest number of packs.
func (pc packCalculator) CalculatePacksAmounts(items int32) []PacksAmount {
	result := doCalculatePacksAmounts(pc.sizes, packsAmounts{amounts: []PacksAmount{}}, items)

	return result.amounts
}

type packsAmounts struct {
	amounts []PacksAmount
	total   int32
	packs   int32
}

func doCalculatePacksAmounts(sizes []int32, a packsAmounts, items int32) packsAmounts {
	if len(sizes) == 0 || a.total >= items {
		return a
	}

	var closest packsAmounts

	i := int32(0)
	if len(sizes) == 1 {
		i = (items - a.total) / sizes[0]
	}

	for ; i <= (items-a.total)/sizes[0]+1; i++ {
		passed := packsAmounts{
			amounts: slices.Clone(a.amounts),
			total:   a.total + sizes[0]*i,
			packs:   a.packs + i,
		}

		if i > 0 {
			passed.amounts = append(passed.amounts, PacksAmount{
				Size:   sizes[0],
				Amount: i,
			})
		}

		next := doCalculatePacksAmounts(sizes[1:], passed, items)
		if next.total < items {
			continue
		}

		if closest.total == 0 || next.total < closest.total || (next.total == closest.total && next.packs < closest.packs) {
			closest = next
		}

	}

	return closest

}
