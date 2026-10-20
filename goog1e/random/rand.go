package random

import (
	"math/rand"
	"strconv"
	"time"
)

var seed int64 = time.Now().UnixNano()
var index int64 = 0

func getSeed() int64 {
	index = index + 1
	seed = time.Now().UnixNano() + index
	return seed
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
func RandStr(num int) string {
	str := ""
	rand.Seed(getSeed())
	for i := 0; i < num; i++ {
		str = str + strconv.Itoa(rand.Intn(10))
	}
	return str
}
func RandInt(min int, max int) int {
	rand.Seed(getSeed())
	return rand.Intn(max) - min
}

var Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Ints(min int, max int, length int) []int {
	if min > max || length == 0 {
		return []int{}
	}
	var rt = make([]int, 0, length)
	for i := 0; i < length; i++ {
		rt = append(rt, Rand.Intn(max-min)+min)
	}
	return rt
}

func Float64s(length int) []float64 {
	var rt = make([]float64, 0, length)
	for i := 0; i < length; i++ {
		rt = append(rt, Rand.Float64())
	}
	return rt
}
