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

package deploy

import (
	"strconv"

	"sync"

	"strings"

	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
	"github.com/daviddexter/wajibu/server/dbase"
	"github.com/daviddexter/wajibu/server/radix"
)

var (
	wg   sync.WaitGroup
	mutx sync.Mutex
)

func generateRepSlots(t bool) { // key central:build:1
	switch t {
	case true: //central government
		slotsForCentralGov()
	case false: //not central government
		slotsForNotCentralGov()
	}
}

func generateRepTitles(t bool) { // key central:build:2
	switch t {
	case true: //central government
		titlesForCentralGov()
	case false: //not central government
		titlesForNotCentralGov()
	}
}

func generateTopLevelAPI() { // key central:build:3	/
	success := *processTopLevel()
	if success == true {
		//add central:build:1 key into redis set
		k := centralBuildKey + ":" + strconv.Itoa(3)
		err := radix.RDB.Cmd("SADD", centralBuildKey, k).Err
		report.ErrLogger(err)
	}
}

func generateHouseLevelAPI() { // key central:build:4
	success := *processHouseLevel()
	if success == true {
		//add central:build:1 key into redis set
		k := centralBuildKey + ":" + strconv.Itoa(4)
		err := radix.RDB.Cmd("SADD", centralBuildKey, k).Err
		report.ErrLogger(err)
	}
}

func generateSubGovLevelAPI() { // key central:build:5
	success := *processSubGovLevel()
	if success == true {
		//add central:build:1 key into redis set
		k := centralBuildKey + ":" + strconv.Itoa(5)
		err := radix.RDB.Cmd("SADD", centralBuildKey, k).Err
		report.ErrLogger(err)
	}
}

func generateGrassRootLevelAPI() { // key central:build:6
	success := *processGrassRootGovLevel()
	if success == true {
		//add central:build:1 key into redis set
		k := centralBuildKey + ":" + strconv.Itoa(6)
		err := radix.RDB.Cmd("SADD", centralBuildKey, k).Err
		report.ErrLogger(err)
	}
}

func slotsForCentralGov() {
	h := commonSlots()
	success := *dbase.AddSlotsFromDeploy(*h)
	if success == true {
		//add central:build:1 key into redis set
		k := centralBuildKey + ":" + strconv.Itoa(1)
		err := radix.RDB.Cmd("SADD", centralBuildKey, k).Err
		report.ErrLogger(err)
	}
}

func slotsForNotCentralGov() {
	a := commonSlots()
	b := notCommonSlots()
	h := combineS(a, b)
	success := *dbase.AddSlotsFromDeploy(*h)
	if success == true {
		//add central:build:1 key into redis set
		k := centralBuildKey + ":" + strconv.Itoa(1)
		err := radix.RDB.Cmd("SADD", centralBuildKey, k).Err
		report.ErrLogger(err)
	}
}

func titlesForCentralGov() {
	h := commonTitles()
	v := dbase.AddTitlesFromDeploy(*h)
	success := *v
	if success == true {
		//add central:build:2 key into redis set
		k := centralBuildKey + ":" + strconv.Itoa(2)
		err := radix.RDB.Cmd("SADD", centralBuildKey, k).Err
		report.ErrLogger(err)
	}
}

func titlesForNotCentralGov() {
	a := commonTitles()
	b := notCommonTitles()
	h := combine(a, b)
	v := dbase.AddTitlesFromDeploy(*h)
	success := *v
	if success == true {
		//add central:build:2 key into redis set
		k := centralBuildKey + ":" + strconv.Itoa(2)
		err := radix.RDB.Cmd("SADD", centralBuildKey, k).Err
		report.ErrLogger(err)
	}
}

func combine(a *[]string, b *[]string) *[]string {
	c := *a
	d := *b
	var s []string
	wg.Add(2)
	go func(list []string) {
		defer wg.Done()
		for _, v := range list {
			mutx.Lock()
			s = append(s, v)
			mutx.Unlock()
		}
	}(c)

	go func(list []string) {
		defer wg.Done()
		for _, v := range list {
			mutx.Lock()
			s = append(s, v)
			mutx.Unlock()
		}
	}(d)
	wg.Wait()
	return &s
}

func combineS(a *[]types.Slot, b *[]types.Slot) *[]types.Slot {
	c := *a
	d := *b
	var s []types.Slot
	wg.Add(2)
	go func(list []types.Slot) {
		defer wg.Done()
		for _, v := range list {
			mutx.Lock()
			s = append(s, v)
			mutx.Unlock()
		}
	}(c)
	go func(list []types.Slot) {
		defer wg.Done()
		for _, v := range list {
			mutx.Lock()
			s = append(s, v)
			mutx.Unlock()
		}
	}(d)
	wg.Wait()
	return &s

}

func commonSlots() *[]types.Slot { //houselevel
	var slots []types.Slot
	//get the number of houses
	key := "main:" + radix.BuildOneData
	res, err := radix.RDB.Cmd("HGET", key, "Numofhouses").Str()
	report.ErrLogger(err)
	numH, _ := strconv.Atoi(res)

	for indexA := 0; indexA < numH; indexA++ {
		setKey := radix.BuildThree + ":house:" + strconv.Itoa(indexA)
		//get members
		members, err := radix.RDB.Cmd("SMEMBERS", setKey).List()
		report.ErrLogger(err)
		for _, member := range members {
			resB, err := radix.RDB.Cmd("HGETALL", member).Map()
			report.ErrLogger(err)
			var slot types.Slot
			slot.SlotName = strings.ToLower(resB["RepSlot"])
			slot.Designation = strings.ToLower(resB["SlotName"])
			slots = append(slots, slot)
		}

	}
	return &slots
}

func notCommonSlots() *[]types.Slot { // subgov level ang grass root level
	var slots []types.Slot
	wg.Add(2)
	key := "main:" + radix.BuildOneData
	go func() {
		//get subgovname
		defer wg.Done()
		res, err := radix.RDB.Cmd("HGETALL", key).Map()
		report.ErrLogger(err)
		designation := strings.ToLower(res["Subgovname"])
		num, _ := strconv.Atoi(res["Numofsubgov"])
		for index := 0; index < num; index++ {
			govKey := radix.BuildFour + ":gov:" + strconv.Itoa(index)
			//get the rep slot name of each subgov
			res, err := radix.RDB.Cmd("HGET", govKey, "SlotName").Str()
			report.ErrLogger(err)
			var slot types.Slot
			slot.Designation = designation
			slot.SlotName = strings.ToLower(res)
			mutx.Lock()
			slots = append(slots, slot)
			mutx.Unlock()
		}
	}()

	go func() {
		//get slots from grassroot
		defer wg.Done()
		//check if subgov has legislative arm
		res, err := radix.RDB.Cmd("HGETALL", key).Map()
		report.ErrLogger(err)
		hasLeg, _ := strconv.ParseBool(res["Subgovhasleg"])
		num, _ := strconv.Atoi(res["Numofsubgov"])
		switch hasLeg {
		case true:
			//leg slot designation
			k := "main:" + radix.BuildOneData
			resD, err := radix.RDB.Cmd("HGET", k, "Subgovhouserepslot").Str()
			report.ErrLogger(err)
			for index := 0; index < num; index++ {
				govKey := radix.BuildFour + ":gov:" + strconv.Itoa(index)
				//get leg seats
				res, err := radix.RDB.Cmd("HGET", govKey, "NumOfLegSeats").Str()
				report.ErrLogger(err)
				num, err := strconv.Atoi(res)
				report.ErrLogger(err)
				for index := 1; index <= num; index++ {
					repKey := govKey + ":" + strconv.Itoa(index)
					res, err := radix.RDB.Cmd("HGET", repKey, "SlotName").Str()
					report.ErrLogger(err)
					var slot types.Slot
					slot.Designation = resD
					slot.SlotName = strings.ToLower(res)
					mutx.Lock()
					slots = append(slots, slot)
					mutx.Unlock()
				}
			}

		}
	}()
	wg.Wait()
	return &slots
}

func commonTitles() *[]string {
	var titles []string
	//add the top level titles;president and deputy president
	titles = append(titles, "president")
	titles = append(titles, "deputy president")
	//houses titles
	key := "main:" + radix.BuildOneData
	res, err := radix.RDB.Cmd("HGET", key, "Numofhouses").Str()
	report.ErrLogger(err)
	num, _ := strconv.Atoi(res)
	//get the representative slot of every house
	for index := 0; index < num; index++ {
		houseKey := "hdetails:" + radix.BuildOneData + ":" + strconv.Itoa(index)
		res, err := radix.RDB.Cmd("HGET", houseKey, "Reptitle").Str()
		report.ErrLogger(err)
		titles = append(titles, strings.ToLower(res))
	}
	return &titles
}

func notCommonTitles() *[]string { // subgov level ang grass root level
	var titles []string
	wg.Add(2)
	key := "main:" + radix.BuildOneData
	go func() {
		//get subgovname
		defer wg.Done()
		res, err := radix.RDB.Cmd("HGET", key, "Subgovtitle").Str()
		report.ErrLogger(err)
		mutx.Lock()
		titles = append(titles, strings.ToLower(res))
		mutx.Unlock()
	}()

	go func() {
		//get slots from grassroot
		defer wg.Done()
		//check if subgov has legislative arm
		res, err := radix.RDB.Cmd("HGET", key, "Subgovhasleg").Str()
		report.ErrLogger(err)
		hasLeg, err := strconv.ParseBool(res)
		report.ErrLogger(err)
		switch hasLeg {
		case true:
			res, err := radix.RDB.Cmd("HGET", key, "Subgovreptitle").Str()
			report.ErrLogger(err)
			mutx.Lock()
			titles = append(titles, strings.ToLower(res))
			mutx.Unlock()
		}
	}()
	wg.Wait()
	return &titles
}
