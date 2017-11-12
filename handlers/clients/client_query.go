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
	"github.com/daviddexter/wajibu/server/radix"
)

func queryForLocationContent(s string) *types.ContentDataAll {
	return fromAnyLevelLocation(s)
}

func queryForTitleContent(s string) *types.ContentDataAll {
	return fromAnyLevelTitle(s)
}

func queryForPillarContent(s string) *types.ContentDataAll {
	return fromAnyLevelPillar(s)
}

func queryForLocationDetail(s string) {

}

func queryForTitleDetail(s string) {

}

func queryForPillarDetail(s string) {

}

func fromAnyLevelLocation(item string) *types.ContentDataAll {
	var r types.ContentDataAll
	l := *radix.GetSentimentFields()
	s := *dbase.DetailForRepSlot(item)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		//get the sentiments for the slot
		r = *dbase.GetSentimentForLocation(s, l)
	}()
	wg.Wait()
	return &r
}

func fromAnyLevelTitle(item string) *types.ContentDataAll {
	var r types.ContentDataAll
	l := *radix.GetSentimentFields()
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		//get api if from toplevel
		defer wg.Done()
		n := *dbase.GetAPIForLevel(item, "toplevel")
		r = *dbase.GetSentimentForLevelTitle(n, l)
	}()
	go func() {
		//get api if from houselevel
		defer wg.Done()
		n := *dbase.GetAPIForLevel(item, "houselevel")
		r = *dbase.GetSentimentForLevelTitle(n, l)
	}()
	go func() {
		//get api if from houselevel
		defer wg.Done()
		n := *dbase.GetAPIForLevel(item, "subgovlevel")
		r = *dbase.GetSentimentForLevelTitle(n, l)
	}()
	go func() {
		//get api if from grasslevel
		defer wg.Done()
		n := *dbase.GetAPIForLevel(item, "grasslevel")
		r = *dbase.GetSentimentForLevelTitle(n, l)
	}()
	wg.Wait()
	//log.Println(r)
	return &r
}

func fromAnyLevelPillar(item string) *types.ContentDataAll {
	var r types.ContentDataAll
	l := *radix.GetSentimentFields()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		//get sentiments of pillar
		r = *dbase.GetSentimentForPillar(item, l)
	}()
	wg.Wait()
	return &r
}
