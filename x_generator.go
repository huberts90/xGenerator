package xGenerator

import (
	"math/rand"
	"strconv"
	"time"
)

type XGenerator struct{}

func NewXGenerator() *XGenerator {
	initRand()
	return &XGenerator{}
}

func (g *XGenerator) Get() string {
	ts := uint64(time.Now().Unix())
	ts32 := strconv.FormatUint(ts, 32)
	r := randomBetween(1299, 46600)
	r36 := strconv.FormatUint(r, 36)

	id := r36 + ts32
	return id
}

// InitRand func
func initRand() {
	rand.Seed(time.Now().UnixNano())
}

func randomBetween(min, max int) uint64 {
	return uint64(rand.Intn(max-min) + min)
}
