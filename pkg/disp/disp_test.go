// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

package disp

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/rokath/trice/pkg/lgf"
	"github.com/rokath/trice/pkg/lib"
)

func lineGenerator(t *testing.T, s string, len, count int, wg *sync.WaitGroup) {
	var ss []string
	var i int
	// create line slice
	for i = 0; i < len; i++ { // line length
		var c string
		switch i & 11 { // 11 color channels
		case 0:
			c = "err:"
		case 1:
			c = "wrn:"
		case 2:
			c = "rd_:"
		case 3:
			c = "wr_:"
		case 4:
			c = "tim:"
		case 5:
			c = "att:"
		case 6:
			c = "dbg:"
		case 7:
			c = "dia:"
		case 8:
			c = "isr:"
		case 9:
			c = "sig:"
		case 10:
			c = "tst:"
		}
		su := fmt.Sprintf(c+s+"%02x ", i) // line position
		ss = append(ss, su)
	}
	ss = append(ss, "\n") // line end

	wg.Add(1)
	i = 0
	go func() {
		defer wg.Done()
		for ; i < count; i++ { // line count
			err := Out(ss)
			lib.Ok(t, err)
			//time.Sleep(100 * time.Microsecond)
		}
	}()
}

func TestServerMutex(t *testing.T) {
	//lgf.Name = "C:/GitRepos/trice/pkg/disp/testdata/serverMutexTest.log"
	//uniqName := "C:/GitRepos/trice/pkg/disp/testdata/serverMutexUniq.txt"
	lgf.Name = "./testdata/serverMutexTest.log"
	uniqName := "./testdata/serverMutexUniq.txt"
	os.Remove(lgf.Name)
	os.Remove(uniqName)

	var wg sync.WaitGroup
	StartServer()
	err := Connect()
	lib.Ok(t, err)

	var result int64
	err = PtrRPC.Call("Server.ColorPalette", []string{"off"}, &result)
	lib.Ok(t, err)

	ss := []string{"first line", "\n"}
	err = Out(ss)
	lib.Ok(t, err)

	ll := 21
	lc := 10
	lineGenerator(t, "r", ll, lc, &wg)
	lineGenerator(t, "s", ll, lc, &wg)
	lineGenerator(t, "t", ll, lc, &wg)
	lineGenerator(t, "u", ll, lc, &wg)
	lineGenerator(t, "v", ll, lc, &wg)
	lineGenerator(t, "w", ll, lc, &wg)
	lineGenerator(t, "x", ll, lc, &wg)
	lineGenerator(t, "y", ll, lc, &wg)
	lineGenerator(t, "z", ll, lc, &wg)
	wg.Wait()
	time.Sleep(3000 * time.Millisecond)
	StopServer()
	n := lib.UniqLines(t, lgf.Name, uniqName)
	lib.Equals(t, n, 11) // first line + 9 lines + last empty line
	os.Remove(lgf.Name)
	os.Remove(uniqName)
}