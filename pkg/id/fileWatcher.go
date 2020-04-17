// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

package id

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
)

// ListFileWatcher checks id list file for changes
// taken from https://medium.com/@skdomino/watch-this-file-watching-in-go-5b5a247cf71f
func (p *ListT) FileWatcher() {

	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	//
	done := make(chan bool)

	//
	go func() {
		var cnt int
		var now, last time.Time
		for {
			select {
			// watch for events
			case <-watcher.Events:
				//fmt.Printf("EVENT! %#v\n", event)

				now = time.Now()
				diff := now.Sub(last)
				if diff > 1000*time.Millisecond {
					fmt.Println("renew id.List")
					cnt++
					renewIDList()
					last = time.Now()
				}

			// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add(FnJSON); err != nil {
		fmt.Println("ERROR", err)
	}
	fmt.Println(FnJSON, " watched now for changes")
	<-done
}

func renewIDList() {
	if "none" != FnJSON {
		List = List[:0]
		List.Read(FnJSON)
	}
	/*
		newList := make(ListT, 0, 65536)
		newList.Read(FnJSON)
		for i := range newList {
			it := newList[i]
			List.appendIfMissing(it, false)
		}
	*/
}