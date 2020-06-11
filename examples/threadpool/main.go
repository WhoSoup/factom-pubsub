package main

import (
	"fmt"
	"log"
	"math/big"
	"reflect"
	"sync"

	pubsub "github.com/WhoSoup/factom-pubsub"
	"github.com/WhoSoup/factom-pubsub/localchannel"
)

const buffer int = 100

func main() {
	max := int64(1e5)

	source := localchannel.New(buffer)
	agg := localchannel.New(buffer)

	pubsub.GlobalRegistry().Register("/source", source)
	pubsub.GlobalRegistry().Register("/aggregate", agg)

	workers := 5

	rrr := pubsub.NewRoundRobinReader(source, workers)

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go PrimeWorker(rrr.GetListener(i), &wg)
	}

	// publish numbers
	go func() {
		for i := int64(0); i < max; i++ {
			if err := source.GetWriter().Write(i); err != nil {
				log.Fatal(err)
			}
		}
		source.Close()
	}()

	go func() {
		wg.Wait()
		agg.Close()
	}()

	primes := 0
	for range agg.GetReader().Channel() {
		primes++
	}
	fmt.Println(primes, "primes found")
}

func PrimeWorker(reader <-chan interface{}, wg *sync.WaitGroup) {
	pub, _ := pubsub.GlobalRegistry().Get("/aggregate")
	for v := range reader {
		if i, ok := v.(int64); !ok {
			log.Println("invalid type arrived:", reflect.TypeOf(v))
			continue
		} else if big.NewInt(i).ProbablyPrime(0) {
			pub.GetWriter().Write(true)
		}
	}
	wg.Done()
}
