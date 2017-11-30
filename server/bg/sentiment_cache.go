/*
Wajibu is an online web app that collects,analyses, aggregates and visualizes sentiments
from the public pertaining the government of a nation. This tool allows citizens to contribute
to the governance talk by airing out their honest views about the state of the nation and in
particular the people placed in government or leadership positions.

Copyright (C) 2017
David 'Dexter' Mwangi
dmwangimail@gmail.com
https://github.com/daviddexter
https://github.com/daviddexter/wajibu

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package bg

import (
	"strconv"
	"sync"
	"time"

	"github.com/daviddexter/wajibu/report"
	"github.com/daviddexter/wajibu/server/dbase"
	"github.com/daviddexter/wajibu/server/radix"
)

func cacheProcessor() {
	caching()
}

func caching() {
	result := make(chan bool)
	go worker(result)
	for {
		select {
		case r := <-result:
			switch r {
			case false:
				time.Sleep(700 * time.Duration(time.Millisecond))
				caching()
			case true:
				time.Sleep(100 * time.Duration(time.Millisecond))
				r := *workerCache()
				if r == true {
					time.Sleep(700 * time.Duration(time.Millisecond))
					caching()
				}
			}
		}
	}
}

func worker(r chan bool) {
	var wg sync.WaitGroup
	wg.Add(1)
	var sentimentsTableExist bool
	go func() {
		defer wg.Done()
		//check if sentiments table exists
		switch i := *dbase.IfSentimentsExist(); i {
		case true:
			sentimentsTableExist = true
		case false:
			removeSentimentsHash()
			removeSentimentsZSet()
			sentimentsTableExist = false
		}
	}()
	wg.Wait()
	switch sentimentsTableExist {
	case true:
		r <- true
	case false:
		r <- false
	}
}

func workerCache() *bool {
	//get the sentiment:field
	var r bool
	fields, err := radix.RDB.Cmd("SMEMBERS", radix.SENTIMENT).List()
	report.ErrLogger(err)
	s := *dbase.GetCurrentSentiments(fields)
	//log.Println(s)
	var counter int
	if len(s) >= 1 {
		removeSentimentsHash()
		removeSentimentsZSet()
		for k, val := range s {
			//PROCESS NEW SENTIMENTS
			index := k + 1
			key := radix.SENTIMENTKEYPREFIX + ":" + strconv.Itoa(index)
			err = radix.RDB.Cmd("HMSET", key, "api", val.Key, "date", val.Date, "image", val.Image, val.Fields).Err
			if err == nil {
				err = radix.RDB.Cmd("ZADD", radix.SENTIMENTLIST, "NX", k, key).Err
				if err == nil {
					counter++
				}
			}
		}
		if len(s) == counter {
			r = true
		}
	} else {
		r = true
	}
	return &r
}

func removeSentimentsHash() {
	//get current sentiment hash a delete them
	s, err := radix.RDB.Cmd("KEYS", radix.SENTIMENTKEYPREFIX+":*").List()
	report.ErrLogger(err)
	for _, sen := range s {
		_ = radix.RDB.Cmd("DEL", sen).Err
	}
}

func removeSentimentsZSet() {
	//delete zset key
	_ = radix.RDB.Cmd("DEL", radix.SENTIMENTLIST).Err
}
