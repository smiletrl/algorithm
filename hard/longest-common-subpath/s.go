package solution

// https://leetcode.com/problems/longest-common-subpath/description/
// rolling hash

func longestCommonSubpath(n int, paths [][]int) int {
	if len(paths) == 0 {
		return 0
	}
	p := &Path{
		base: n + 1,
		// mod:  1e9 + 7,
		// largest prime number under 2^32
		mod:   2147483647,
		paths: paths,
	}
	min := p.getMinimLength()

	// pre-calculte base pow mod value.
	p.basePowMod = make(map[int]int, min)
	p.basePowMod[0] = 1
	for i := 1; i < min; i++ {
		p.basePowMod[i] = p.basePowMod[i-1] * p.base % p.mod
	}

	p.binarySearch(1, min)
	return p.currentCommonLen
}

type Path struct {
	currentCommonLength int
	paths               [][]int
	base                int
	mod                 int

	// to keep value to avoid collision
	// initRollingHashVal map[int]struct{}
	initRollingHashVal     map[int]map[int]struct{}
	initRollingHashPattern int

	currentCommonLen int

	currentLenLastPara int

	// this map's key is x, value is math.Pow(b, x) % m
	// these values are precaculated for key x, ranges from 1 to max-possible-length
	basePowMod map[int]int

	miniStartIndex int
}

func (p *Path) getMinimLength() int {
	minLength := len(p.paths[0])
	for i := 1; i < len(p.paths); i++ {
		if minLength > len(p.paths[i]) {
			minLength = len(p.paths[i])
			p.miniStartIndex = i
		}
	}
	return minLength
}

func (p *Path) binarySearch(low, high int) {
	// if low has outcome high, end search
	if low > high {
		return
	}
	mid := (low + high) / 2
	if mid != 0 {
		newLow, newHigh := p.processLength(mid, low, high)
		p.binarySearch(newLow, newHigh)
	}
}

func (p *Path) processLength(mid, low, high int) (newLow, newHigh int) {
	// if we find one match, we go high, otherwise go low.
	if p.checkRound(mid) {
		p.currentCommonLength = mid
		// go to high
		return mid + 1, high
	} else {
		// go to low
		return low, mid - 1
	}
}

func (p *Path) rollingHash(lastVal, lastStart, newEnd, checkLen int) int {
	// remove last start value
	result := (lastVal + p.mod - p.basePowMod[checkLen-1]*lastStart%p.mod) % p.mod
	// add the new end
	result = (result*p.base + newEnd) % p.mod
	return result
}

func (p *Path) initFirstLineHash(checkLen int) {
	data := p.paths[p.miniStartIndex]
	if len(data) < checkLen {
		panic("check len should be larger than data length")
	}
	p.initRollingHashVal = make(map[int]map[int]struct{}, len(data)-checkLen)
	p.initRollingHashPattern = 0

	// init val
	result := 0
	for i := 0; i < checkLen; i++ {
		result = result + (p.basePowMod[checkLen-1-i]*data[i])%p.mod
	}
	result = result % p.mod

	// add hash value's start index
	p.initRollingHashVal[result] = map[int]struct{}{
		0: {},
	}
	p.initRollingHashPattern++

	for i := 1; i <= len(data)-checkLen; i++ {
		hval := p.rollingHash(result, data[i-1], data[i+checkLen-1], checkLen)
		// same hash value is found, this may or may not be a hash collision
		existingIndexes, ok := p.initRollingHashVal[hval]
		if !ok {
			// insert this new hash value to local map
			p.initRollingHashVal[hval] = map[int]struct{}{
				i: {},
			}
			p.initRollingHashPattern++
		} else {
			// pre hash value found
			// firstly check if new value matches existing values
			matched := false
			for in := range existingIndexes {
				if equal := p.verifyEquality(data[in:in+checkLen], data[i:i+checkLen]); equal {
					// find a match, then we break
					matched = true
					break
				}
			}
			// if no match found, hash collision happend, add this new index
			if !matched {
				p.initRollingHashVal[hval][i] = struct{}{}
				p.initRollingHashPattern++
			}
		}
		result = hval
	}
}

func (p *Path) checkLine(data []int, checkLen int) {
	reservedHash := make(map[int]map[int]struct{})
	reservedHashPattern := 0

	// init val
	result := 0
	for i := 0; i < checkLen; i++ {
		result = result + (p.basePowMod[checkLen-1-i]*data[i])%p.mod
	}
	result = result % p.mod

	// if we find one match, keep this result in reserved hashing list
	existingIndexes, ok := p.initRollingHashVal[result]
	if ok {
		// see if current value belongs to existing values
		for in := range existingIndexes {
			if p.verifyEquality(p.paths[p.miniStartIndex][in:in+checkLen], data[:checkLen]) {
				// record first line's index
				reservedHash[result] = map[int]struct{}{
					in: {},
				}
				reservedHashPattern++
				break
			}
		}
	}

	for i := 1; i <= len(data)-checkLen; i++ {
		hval := p.rollingHash(result, data[i-1], data[i+checkLen-1], checkLen)
		result = hval
		existingIndexes, ok := p.initRollingHashVal[result]
		if ok {
			for in := range existingIndexes {
				if p.verifyEquality(p.paths[p.miniStartIndex][in:in+checkLen], data[i:i+checkLen]) {
					// record first line's index
					if _, ok2 := reservedHash[result]; !ok2 {
						reservedHash[result] = map[int]struct{}{
							in: {},
						}
						reservedHashPattern++
					} else {
						if _, ok2 := reservedHash[result][in]; !ok2 {
							reservedHash[result][in] = struct{}{}
							reservedHashPattern++
						}
					}
					break
				}
			}
		}
		// if we have found same number of pattern just like the start line, no need to proceed with
		// this line anymore
		if reservedHashPattern == p.initRollingHashPattern {
			break
		}
	}

	// if reserved pattern and init rolling pattern is the same, no need to do replace
	if reservedHashPattern != p.initRollingHashPattern {
		p.initRollingHashVal = reservedHash
	}
}

func (p *Path) checkRound(le int) (res bool) {
	p.initLastValParam(le)
	p.initFirstLineHash(le)
	for i := 0; i < len(p.paths); i++ {
		// skip start index
		if i == p.miniStartIndex {
			continue
		}
		p.checkLine(p.paths[i], le)
		// if no common length left already, terminate this round
		if len(p.initRollingHashVal) == 0 {
			return false
		}
	}
	p.currentCommonLen = le
	return true
}

func (p *Path) initLastValParam(le int) {
	p.currentLenLastPara = 0
	for i := 0; i < le-1; i++ {
		if p.currentLenLastPara == 0 {
			p.currentLenLastPara = p.base % p.mod
		} else {
			p.currentLenLastPara = p.currentCommonLen * p.base % p.mod
		}
	}
}

func (p *Path) verifyEquality(s, end []int) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != end[i] {
			return false
		}
	}
	return true
}
