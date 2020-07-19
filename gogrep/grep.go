
package gogrep

import (
	"bytes"
	"fmt"
	"github.com/aloksguha/gogrep/utils"
	"os"
	"sync"
	"time"
)


var wg sync.WaitGroup

func NewSearch(path string, timeout int, searchstring string, noOfWorkers int) *Search {
	return &Search{
		filePath: path,
		timeout: timeout,
		searchString: searchstring,
		noOfworkers: noOfWorkers,
	}
}


func (s *Search) Search() ([]Report, error) {
	return s.search()
}

func(s *Search) search() ([]Report, error) {
	fmt.Println(utils.Info("Searching '"+s.searchString+"', in file :", s.filePath))
	file, err := os.Open(s.filePath)
	if err!= nil {
		return nil, err
	}
	check(err)
	defer file.Close()
	fileInfo, _ := file.Stat()
	fileSize := int(fileInfo.Size())
	bufSize := fileSize/s.noOfworkers

	concurrency := s.noOfworkers
	chunks := make([]chunk, concurrency)

	for i := 0; i < concurrency; i++ {
		chunks[i].bufsize = bufSize
		chunks[i].offset = int64(bufSize * i)
	}

	if remainder := fileSize % s.noOfworkers; remainder != 0 {
		lastchunk := chunks[len(chunks)-1]
		lastchunk.bufsize+= remainder
		chunks[len(chunks)-1] = lastchunk
	}

	wg.Add(concurrency)
	masterReport := make(chan Report, concurrency)


	resultsByWorker := make([]Report, 0)
	for i := 0; i < concurrency; i++ {
		chunk := chunks[i]
		go s.searchInChunk(masterReport, time.Now(), file, chunk, i)
		select {
		   case report := <- masterReport :{
			   //fmt.Println("hi",report)
			   resultsByWorker = append(resultsByWorker, report)
		   }
		   case <-time.After(time.Duration(s.timeout)*time.Second): {
				   fmt.Println("timeout")
			   resultsByWorker = append(resultsByWorker, Report{Status: STATUS(TIMEOUT)})
		   }
		}
	}
	wg.Wait() //blocks till all process finishes
	close(masterReport)
	return resultsByWorker, nil
}

func (s *Search) searchInChunk(c chan Report, start time.Time, file *os.File, chunk chunk, id int)  {
	defer wg.Done()
	buffer := make([]byte, chunk.bufsize)
	file.ReadAt(buffer, chunk.offset)
	r := Report{id, 0, 0,0, STATUS(FAILURE)}

	if bytes.Contains(buffer, [] byte(s.searchString)) {
		r.Status = STATUS(SUCEESS)
		r.Elapsed = time.Duration(time.Since(start))
		r.Remaining = (time.Duration(s.timeout)* time.Second) - (time.Since(start))
		r.ByteCnt = bytes.IndexAny(buffer, s.searchString)
		c <- r
	} else {
		c <- r
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}





