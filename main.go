package potato

import "github.com/attilaolah/intperm.go"

const seed = 42

func Perm() interface {
	MapTo(uint64) uint64
	MapFrom(uint64) uint64
} {
	return intperm.New(seed)
}
