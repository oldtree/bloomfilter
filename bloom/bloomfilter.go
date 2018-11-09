package bloom

// bloom fitler impl

func defaultBloom(value interface{}) uint64 {
	return 0
}

type BloomHash func(value interface{}) uint64

type BloomFilter struct {
	Hashfunclist []BloomHash
	Length       uint64
	bitset       []uint64
}

func NewBloomFilter(length uint64, bloomfunc []BloomHash) *BloomFilter {
	if len(bloomfunc) <= 0 {
		bloomfunc = append(bloomfunc, defaultBloom)
	}
	return &BloomFilter{
		Hashfunclist: bloomfunc,
		Length:       length,
		bitset:       make([]uint64, length, length),
	}
}

func (b *BloomFilter) maskbit() {

}

func (b *BloomFilter) AddSync(value interface{}) {
	var hashVal uint64
	for index, valCall := range b.Hashfunclist {
		println("call [%d] bloom hash func", index)
		hashVal = valCall(value)
		idx := hashVal / b.Length
		b.bitset = b.bitset[idx] && uint64(0x1<<hashVal)
	}
}

func (b *BloomFilter) AddAsync(value interface{}) {

}

func (b *BloomFilter) IsExist() bool {
	return false
}
