package component

import (
	"math"
	"math/rand"
	"time"

	"../domain"
)

type priceGenerator struct {
	cache map[int]*domain.Price
}

var sharedInstance *priceGenerator = newPriceGenerator()

func newPriceGenerator() *priceGenerator {
	generator := new(priceGenerator)
	generator.cache = map[int]*domain.Price{}
	return generator
}

func GetPriceGenerator() *priceGenerator {
	return sharedInstance
}

func (generator *priceGenerator) GetLatestPrice() *domain.Price {
	return generator.cache[len(generator.cache)]
}

func (generator *priceGenerator) Generate(currency1 string, currency2 string) {
	ticker := time.NewTicker(time.Millisecond * 1000)
	count := 0
	go func() {
		for {
			select {
			case <-ticker.C:
				count++
				rand.Seed(time.Now().UnixNano())
				value := round(rand.Float64()*1000, .5, 7)
				error, price := domain.NewPrice(*domain.NewCurrency(currency1), *domain.NewCurrency(currency2), value)
				if error != nil {
					panic(error)
				}
				generator.cache[count] = price
			}
		}
	}()
	time.Sleep(time.Second * 5)
	ticker.Stop()
}

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return newVal
}
