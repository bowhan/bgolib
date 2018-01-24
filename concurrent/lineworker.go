package concurrent

import (
	"os"
	"sync"
	"bufio"
	"io"
)

/// An interface to be implemented by any worker doing real work
type LineWorker interface {
	Process(line []byte) /// the function to process the line
}

/// A class wrapping a bunch of workers using channels to process lines in parallel
type LineWorkerGroup struct {
	FileHandler *os.File     /// file handler
	Workers     []LineWorker /// workers
	reader      *bufio.Reader
	buffer      []byte
	ch          chan []byte
	wg          sync.WaitGroup
}

/// Add new worker which has implemented the LineWorker interface
func (this *LineWorkerGroup) AddWorker(newWorker LineWorker) {
	this.Workers = append(this.Workers, newWorker)
}

/// Prepare workers
func (this *LineWorkerGroup) StartWorkers() {
	numWorker := len(this.Workers)
	this.wg.Add(numWorker)
	this.ch = make(chan []byte, numWorker*10)
	this.reader = bufio.NewReader(this.FileHandler)
	for i := 0; i < numWorker; i++ {
		go func(worker LineWorker) {
			defer this.wg.Done()
			for {
				if line, ok := <-this.ch; ok {
					worker.Process(line)
				} else {
					break
				}
			}
		}(this.Workers[i])
	}
}

func (this *LineWorkerGroup) readline() (err error) {
	this.buffer, err = this.reader.ReadBytes('\n')
	return
}

func (this *LineWorkerGroup) feed() {
	this.ch <- this.buffer
}

/// Read input file line by line and feed them to the channel
func (this *LineWorkerGroup) ReadAndFeed() (err error) {
	for {
		if err = this.readline(); err == nil {
			this.feed()
		} else if err == io.EOF {
			return nil
		} else {
			return
		}
	}
}

/// wait for all workers to finish job
func (this *LineWorkerGroup) Wait() {
	close(this.ch)
	this.wg.Wait()
}
