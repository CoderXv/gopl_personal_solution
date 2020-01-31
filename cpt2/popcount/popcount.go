package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	ret := byte(0)
	for i := uint8(0); i < 8; i++ {
		ret += pc[byte(x>>(i*8))]
	}
	return int(ret)
}

// 2.4 写一个用于统计位的PopCount,它在其实际参数的64位上执行移位操作,每次判断最右边的位,进而实现统计功能
func PopCount2(x uint64) int {
	ret := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			ret += 1
		}
		x = x >> 1
	}
	return ret
}

// 2.5 使用x&(x-1)可以清除x最右边的非零位,利用该特点写一个PopCount
func PopCount3(x uint64) int {
	ret := 0
	for x != 0 {
		x = x&(x-1)
		ret += 1
	}
	return ret
}
