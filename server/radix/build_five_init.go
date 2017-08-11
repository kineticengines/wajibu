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

package radix

import (
	"sync"

	"strconv"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
	mp "github.com/mitchellh/mapstructure"
)

var wg sync.WaitGroup

var (
	numOfSubGovs    string
	subGovHouseName string
	subGovRepSlot   string
	subGovRepTitle  string
	subGovName      string //county or state or province or region
	starterKeys     []string
)

type InitData struct {
	IsComplete      bool   `json:"iscomplete"`
	SubGovHouseName string `json:"housename,omitempty"`
	SubGovRepSlot   string `json:"repslot,omitempty"`
	SubGovRepTitle  string `json:"reptitle,omitempty"`
	SubGovName      string `json:"govname,omitempty"`
	SubgovData      InComp `json:"govdata"`
}

type InComp struct {
	GovInKey string          `json:"govinkey,omitempty"`
	Build    types.BuildFour `json:"build,omitempty"`
}

func BuildFiveInitializer() *InitData {
	var data InitData
	//start goroutines to fetch nummber of subgovernments,subgov house name, subgov representative title and
	//subgovname
	starterKeys = []string{"numofsubgovs", "subgovhousename", "subgovrepslot", "subgovreptitle", "subgovname"}
	wg.Add(len(starterKeys))
	for _, v := range starterKeys {
		go func(key string) {
			defer wg.Done()
			switch key {
			case starterKeys[0]: // numofsubgovs
				//key of main build:four data
				mainKey := "main:" + BuildFourData
				res, err := RDB.Cmd("HGET", mainKey, "NumOfSubGovs").Str()
				report.ErrLogger(err)
				numOfSubGovs = res
			case starterKeys[1]: //subgovhousname
				//key of main build:one data
				mainKey := "main:" + BuildOneData
				res, err := RDB.Cmd("HGET", mainKey, "Subgovhousename").Str()
				report.ErrLogger(err)
				subGovHouseName = res
			case starterKeys[2]: //subgovrepslot
				mainKey := "main:" + BuildOneData
				res, err := RDB.Cmd("HGET", mainKey, "Subgovhouserepslot").Str()
				report.ErrLogger(err)
				subGovRepSlot = res
			case starterKeys[3]: //subgovreptitle
				mainKey := "main:" + BuildOneData
				res, err := RDB.Cmd("HGET", mainKey, "Subgovreptitle").Str()
				report.ErrLogger(err)
				subGovRepTitle = res
			case starterKeys[4]: //subgovname
				mainKey := "main:" + BuildOneData
				res, err := RDB.Cmd("HGET", mainKey, "Subgovname").Str()
				report.ErrLogger(err)
				subGovName = res
			}
		}(v)
	}
	wg.Wait()
	c, g := checkIfLevelIsComplete()
	if c == false {
		//pick a subgovernment and send back to be filled
		if len(g) == 0 {
			data.IsComplete = true
		} else {
			gv := pickSubgov(g)
			data.IsComplete = false
			data.SubGovHouseName = subGovHouseName
			data.SubGovRepSlot = subGovRepSlot
			data.SubGovRepTitle = subGovRepTitle
			data.SubGovName = subGovName
			data.SubgovData = *gv
		}

	} else {
		data.IsComplete = true
	}
	return &data
}

func checkIfLevelIsComplete() (bool, []InComp) {
	var govs []InComp
	num, _ := strconv.Atoi(numOfSubGovs)
	notCompleteCount := 0
	for index := 0; index < num; index++ {
		govKey := BuildFour + ":gov:" + strconv.Itoa(index)
		//get the num of seats of the sub government
		res, err := RDB.Cmd("HGETALL", govKey).Map()
		report.ErrLogger(err)
		//convert map[string]interface{} to struct of build four
		var target types.BuildFour
		if err := mp.Decode(res, &target); err != nil {
			report.ErrLogger(err)
		}
		numofseats, _ := strconv.Atoi(target.NumOfLegSeats)
		notExistCount := 0
		for index := 1; index <= numofseats; index++ {
			seatKey := govKey + ":" + strconv.Itoa(index)
			res, err := RDB.Cmd("EXISTS", seatKey).Int()
			report.ErrLogger(err)
			if res == 0 {
				notExistCount++
			}
		}
		if notExistCount == numofseats {
			notCompleteCount++
			//append into govs slice
			var m InComp
			m.GovInKey = govKey
			m.Build = target
			govs = append(govs, m)
		}
	}
	if notCompleteCount == 0 {
		return true, govs //the level is complete
	}
	return false, govs // the level is not complete
}

func pickSubgov(govs []InComp) *InComp {
	return &govs[randomdata.Number(len(govs))]
}
