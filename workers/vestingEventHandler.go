package workers

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/nicknikandish/vesting_program/api"
	"github.com/nicknikandish/vesting_program/models"
	"io"
	"log"
	"math"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var Events []models.VestingEvent

const (
	KILOBYTE   = 1024
	BufferSize = 32 * KILOBYTE
)

func ReadCSVFile(filepath string, precision int, targetDate time.Time) error {
	numberOfCPUs := runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	// CSV file read all at once
	lines := csv.NewReader(f)
	num, err := lineCounter(f)
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	// create the pair of input/output channels for the controller=>workers com.
	src := make(chan []models.VestingEvent)
	dest := make(chan string)

	// use a wait group to manage synchronization
	var wg sync.WaitGroup
	// create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// that cancels at ctrl+C
	go onSignal(os.Interrupt, cancel)

	// declare the workers
	declareWorker(ctx, &wg, src, targetDate, precision, numberOfCPUs)
	// reading the csv and write events to the channel
	go readCSVFileAndPutInAChannel(src, lines, num, numberOfCPUs, precision)
	// wait for worker group to finish and close out
	go func() {
		wg.Wait()
		close(dest)
	}()

	// drain the output
	for res := range dest {
		fmt.Println(res)
	}

	return nil
}

func worker(ctx context.Context, src chan []models.VestingEvent, precision int, targetDate time.Time) {
	select {

	case events, ok := <-src: // checking for readable state of the channel.
		if !ok {
			return
		}
		mu.Lock()
		Process(events, targetDate, precision)
		mu.Unlock()
	case <-ctx.Done(): // if the context is cancelled, quit.
		return
	}
}

func onSignal(s os.Signal, h func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, s)
	<-c
	h()
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, BufferSize)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func declareWorker(ctx context.Context, wg *sync.WaitGroup, src chan []models.VestingEvent, targetDate time.Time, precision, numberOfCPUs int) {
	for i := 0; i < numberOfCPUs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, src, precision, targetDate)
		}()
	}
}

func readCSVFileAndPutInAChannel(src chan []models.VestingEvent, lines *csv.Reader, num int, numberOfCPUs int, precision int) {
	i := 0
	for {
		if i == int(math.Ceil(float64(num)/float64(numberOfCPUs))) {
			src <- Events
			i = 0
			Events = nil
		}
		record, err := lines.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		date, err := time.Parse("2006-01-02", record[4])
		if err != nil {
			fmt.Println(err)
		}
		quantity, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			fmt.Println(err)
		}
		vestingEvent := models.VestingEvent{
			VestType: record[0],
			Award:    api.Award{EmployeeId: record[1], EmployeeName: record[2], AwardId: record[3]},
			Date:     date,
			Quantity: RoundFloat(quantity, precision),
		}
		Events = append(Events, vestingEvent)
		i++
	}
	src <- Events
	close(src) // close src to signal workers that no more job are incoming.
}
