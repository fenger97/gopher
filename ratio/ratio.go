package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type AccelInfo struct {
	reqErr string
}

type PmWorker struct {
	ratio      uint64
	index      uint64
	recvCount  uint64
	ratioCount uint64
}

const (
	MaxRatio       = 100
	AccelInfoTopic = "hopMsg"
)

func (pm *PmWorker) needToSend(topic string, msg *AccelInfo) bool {
	if pm.ratio == MaxRatio {
		return true
	}

	if topic == AccelInfoTopic && msg.reqErr == "" {
		index := atomic.AddUint64(&pm.index, 1)
		if index%MaxRatio >= pm.ratio {
			return false
		}
	}

	return true
}

func (pm *PmWorker) SendToKafka(topic string, msg *AccelInfo) {
	if pm.needToSend(topic, msg) {
		atomic.AddUint64(&pm.ratioCount, 1)
	}
	atomic.AddUint64(&pm.recvCount, 1)
}

func testSend(pm *PmWorker, id int) {
	count := 0
	for i := 0; i < 10000; i++ {
		pm.SendToKafka(AccelInfoTopic, &AccelInfo{})
		count++
	}
	fmt.Printf("id: %d send: %d\n", id, count)
}

func main() {
	ratio := 1

	numWorkers := 10
	pm := &PmWorker{ratio: uint64(ratio)}
	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			defer wg.Done()
			testSend(pm, id)
		}(i)
	}
	wg.Wait()

	// wait for pm status finish
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("ratio: %d / recv: %d = %0.2f\n", pm.ratioCount, pm.recvCount, float64(pm.ratioCount)/float64(pm.recvCount))
}
