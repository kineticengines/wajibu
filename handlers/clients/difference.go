/*
Wajibu is an online web app that collects,analyses and aggregates sentiments from the public
pertaining the government of a nation. This tool allows citizens to contribute to the
governance talk by airing out their honest views about the state of the nation and in
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

package clients

import (
	"sync"

	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/server/dbase"
)

func determineTheQueryType(item string) (*[]types.QueryType, *bool) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var rAll []types.QueryType
	var falseCount int
	var v bool
	wg.Add(3)
	go func() {
		//check if type is location
		defer wg.Done()
		var r types.QueryType
		n := *dbase.CheckIfLocation(item)
		switch n {
		case true:
			r.Type = "location"
			r.IsTrue = true
			r.Value = item
			break
		case false:
			r.Type = "location"
			r.IsTrue = false
			r.Value = item
			mutex.Lock()
			falseCount++
			mutex.Unlock()
			break
		}
		mutex.Lock()
		rAll = append(rAll, r)
		mutex.Unlock()
	}()

	go func() {
		//check if type is title
		defer wg.Done()
		var r types.QueryType
		n := *dbase.CheckIfTitle(item)
		switch n {
		case true:
			r.Type = "title"
			r.IsTrue = true
			r.Value = item
			break
		case false:
			r.Type = "title"
			r.IsTrue = false
			r.Value = item
			mutex.Lock()
			falseCount++
			mutex.Unlock()
			break
		}
		mutex.Lock()
		rAll = append(rAll, r)
		mutex.Unlock()
	}()

	go func() {
		//check if type is pillar
		defer wg.Done()
		var r types.QueryType
		n := *dbase.CheckIfPillar(item)
		switch n {
		case true:
			r.Type = "pillar"
			r.IsTrue = true
			r.Value = item
			break
		case false:
			r.Type = "pillar"
			r.IsTrue = false
			r.Value = item
			mutex.Lock()
			falseCount++
			mutex.Unlock()
			break
		}
		mutex.Lock()
		rAll = append(rAll, r)
		mutex.Unlock()
	}()
	wg.Wait()
	if falseCount == 3 {
		v = true
	}
	return &rAll, &v
}

func querySwitcher(s []types.QueryType, t string) *types.ContentDataAll {
	//log.Println(s)
	var r types.ContentDataAll
	switch t {
	case "main-content":
		for _, v := range s {
			switch v.Type {
			case "location":
				if v.IsTrue == true {
					r = *queryForLocationContent(v.Value)
				}
				break
			case "title":
				if v.IsTrue == true {
					r = *queryForTitleContent(v.Value)
				}
				break
			case "pillar":
				if v.IsTrue == true {
					r = *queryForPillarContent(v.Value)
				}
				break
			}
		}
	case "details-content":
		for _, v := range s {
			switch v.Type {
			case "location":
				if v.IsTrue == true {
					queryForLocationDetail(v.Value)
				}
				break
			case "title":
				if v.IsTrue == true {
					queryForTitleDetail(v.Value)
				}
				break
			case "pillar":
				if v.IsTrue == true {
					queryForPillarDetail(v.Value)
				}
				break
			}
		}
	}

	return &r
}
