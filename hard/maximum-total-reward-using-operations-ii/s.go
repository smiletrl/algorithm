package solution

import (
	"math/big"
	"sort"
)

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	// we use one bitset to hold all reward values. Each reward value represents one bit in the set.
	x := big.NewInt(int64(1))
	for _, r := range rewardValues {
		// k := x & (1<<r - 1)
		// x = x | k<<r

		// current position's allowed total reward values are smaller than current reward value's bit.
		k := big.NewInt(int64(1))
		k.And(x, k.Sub(k.Lsh(k, uint(r)), big.NewInt(int64(1))))

		// add current position's value to allowed total reward value
		x.Or(x, k.Lsh(k, uint(r)))
	}
	return x.BitLen() - 1
}
