package lazy

import (
	"sync"
	"testing"
	"time"
)

type spinupClient struct{}

func TestFetcherReturnsSamePointer(t *testing.T) {
	timesCalled := 0
	newSpinupClient := func() (*spinupClient, error) {
		timesCalled++
		time.Sleep(200 * time.Millisecond)
		return &spinupClient{}, nil
	}
	fetcher := Fetcher(newSpinupClient, nil)
	var wg sync.WaitGroup
	wg.Add(2)
	var a, b *spinupClient
	go func() {
		a = fetcher()
		wg.Done()
	}()
	go func() {
		b = fetcher()
		wg.Done()
	}()
	wg.Wait()
	if a != b {
		t.Fatalf("returned pointers are not equal")
	}
	if timesCalled != 1 {
		t.Fatalf("initialiation happened more than once")
	}
}
