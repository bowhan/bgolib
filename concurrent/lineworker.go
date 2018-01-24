package concurrent

import (
	"os"
	"sync"
	"bufio"
	"io"
)

type LineWorker interface {
	Process(line []byte)
}

type LineWorkerGroup struct {
	FileHandler *os.File
	Workers     []LineWorker
	reader      *bufio.Reader
	buffer      []byte
	ch          chan []byte
	wg          sync.WaitGroup
}

func (this *LineWorkerGroup) AddWorker(newWorker LineWorker) {
	this.Workers = append(this.Workers, newWorker)
}

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

func (this *LineWorkerGroup) ReadLine() (err error) {
	this.buffer, err = this.reader.ReadBytes('\n')
	return
}

func (this *LineWorkerGroup) Feed() {
	this.ch <- this.buffer
}

func (this *LineWorkerGroup) ReadAndFeed() (err error) {
	for {
		if err = this.ReadLine(); err == nil {
			this.Feed()
		} else if err == io.EOF {
			return nil
		} else {
			return
		}
	}
}

func (this *LineWorkerGroup) Wait() {
	close(this.ch)
	this.wg.Wait()
}
