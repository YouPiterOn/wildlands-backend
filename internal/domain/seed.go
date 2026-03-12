package domain

import (
	"crypto/rand"
	"encoding/binary"
	"hash/fnv"
)

func GenerateSeed() (int64, error) {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(b[:])), nil
}

func DeriveSeedFromInt(seed int64, i int) int64 {
	h := fnv.New64a()

	var buf [16]byte
	binary.LittleEndian.PutUint64(buf[0:], uint64(seed))
	binary.LittleEndian.PutUint64(buf[8:], uint64(i))

	h.Write(buf[:])
	return int64(h.Sum64())
}

func DeriveSeedFromString(seed int64, s string) int64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return int64(h.Sum64())
}
