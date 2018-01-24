package main

import (
	"github.com/bowhan/bgolib/concurrent"
	"os"
	"fmt"
)

type CharCounter struct {
	N int
}

func (this *CharCounter) Process(line []byte) {
	this.N += len(line)
}

func main() {
	numWorker := 8
	fh, err := os.Open("/Users/bo.han/annotation/hg38.fa")
	if err != nil {
		os.Exit(1)
	}
	workGroup := concurrent.LineWorkerGroup{FileHandler: fh}
	for i := 0; i < numWorker; i ++ {
		workGroup.AddWorker(&CharCounter{N: 0})
	}
	workGroup.StartWorkers()
	workGroup.ReadAndFeed()
	workGroup.Wait()
	totalChar := 0
	for i := 0; i < numWorker; i++ {
		totalChar += workGroup.Workers[i].(*CharCounter).N
	}
	fmt.Printf("Total char:%d\n", totalChar)
}
