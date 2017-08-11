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
	"strconv"

	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
	mp "github.com/mitchellh/mapstructure"
)

func mapCast(toCast map[string]string) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range toCast {
		switch k {
		case "Subgovhasleg":
			b, _ := strconv.ParseBool(v)
			m[k] = b
		default:
			m[k] = v
		}
	}
	return m
}

func populateBuildOne(res map[string]string, dataLevel string) types.BuildOneAll {
	var rMain types.BuildOneMain
	var rHDetails types.BuidlOneHDetails
	var rAll types.BuildOneAll
	mainKey := "main:" + dataLevel

	ex, err := RDB.Cmd("EXISTS", mainKey).Int()
	report.ErrLogger(err)
	switch ex {
	case 1:
		resMain, err := RDB.Cmd("HGETALL", mainKey).Map()
		report.ErrLogger(err)
		//convert map[string]string to map[string]interface{}
		castMain := mapCast(resMain)
		//convert map[string]interface{} to struct
		if err := mp.Decode(castMain, &rMain); err != nil {
			report.ErrLogger(err)
		}
		count, _ := strconv.Atoi(res["hlen"])
		d := make([]types.BuidlOneHDetails, 0)
		for index := 0; index < count; index++ {
			key := "hdetails:" + dataLevel + ":" + strconv.Itoa(index)
			res, err := RDB.Cmd("HGETALL", key).Map()
			report.ErrLogger(err)
			if err := mp.Decode(res, &rHDetails); err != nil {
				report.ErrLogger(err)
			}
			d = append(d, rHDetails)

		}
		rAll.Main = rMain
		rAll.HDetails = d
	}

	return rAll
}

func populateBuildTwo(dataLevel string) types.BuildTwoAll {
	var rAll types.BuildTwoAll
	mainKey := "main:" + dataLevel
	ex, err := RDB.Cmd("EXISTS", mainKey).Int()
	report.ErrLogger(err)
	switch ex {
	case 1:
		resMain, err := RDB.Cmd("HGETALL", mainKey).Map()
		report.ErrLogger(err)
		//convert map[string]string to map[string]interface{}
		castMain := mapCast(resMain)
		//convert map[string]interface{} to struct
		if err := mp.Decode(castMain, &rAll); err != nil {
			report.ErrLogger(err)
		}
	}
	return rAll
}

func populateBuildThree(dataLevel string) types.BuildThreeAll {
	var buildType types.BuildThree
	var rAll types.BuildThreeAll
	mainKey := "main:" + dataLevel
	ex, err := RDB.Cmd("EXISTS", mainKey).Int()
	report.ErrLogger(err)
	switch ex {
	case 1:
		houseCount, err := RDB.Cmd("HGET", mainKey, "Housecount").Int()
		report.ErrLogger(err)
		for index := 1; index <= houseCount; index++ {
			i := index - 1
			setKey := "build:three:house:" + strconv.Itoa(i)
			//get members of set key
			sm, err := RDB.Cmd("SMEMBERS", setKey).List()
			report.ErrLogger(err)
			//get data of each key and append to slice
			for _, key := range sm {
				res, err := RDB.Cmd("HGETALL", key).Map()
				report.ErrLogger(err)
				//convert map[string]string to map[string]interface{}
				castMain := mapCast(res)
				if err := mp.Decode(castMain, &buildType); err != nil {
					report.ErrLogger(err)
				}
				rAll.Housesdata = append(rAll.Housesdata, buildType)
			}
		}
	}
	return rAll
}

func populateBuildFour(dataLevel string) types.BuildFourAll {
	var rAll types.BuildFourAll
	mainKey := "main:" + dataLevel

	ex, err := RDB.Cmd("EXISTS", mainKey).Int()
	report.ErrLogger(err)
	switch ex {
	case 1:
		//get the number of subgovs
		res, err := RDB.Cmd("HGETALL", mainKey).Map()
		report.ErrLogger(err)
		num, _ := strconv.Atoi(res["NumOfSubGovs"])
		for i := 0; i < num; i++ {
			var b types.BuildFour
			govKey := BuildFour + ":gov:" + strconv.Itoa(i)
			res, err := RDB.Cmd("HGETALL", govKey).Map()
			report.ErrLogger(err)
			//map[string]inteface to struct
			if err := mp.Decode(res, &b); err != nil {
				report.ErrLogger(err)
			}
			rAll.SubgovData = append(rAll.SubgovData, b)
		}
	}
	return rAll
}

func populateBuildFive() types.BuildFiveAll {
	var rAll types.BuildFiveAll
	//first get the number of sub governments
	//key of main build:four data
	mainKey := "main:" + BuildFourData

	ex, err := RDB.Cmd("EXISTS", mainKey).Int()
	report.ErrLogger(err)
	switch ex {
	case 1:
		res, err := RDB.Cmd("HGET", mainKey, "NumOfSubGovs").Str()
		report.ErrLogger(err)
		numofsubgovs, _ := strconv.Atoi(res)

		//for each subgov get the number of leg seats slotted
		for index := 0; index < numofsubgovs; index++ {
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
			subgovname := target.SlotName

			//get subgov block designation
			block, err := RDB.Cmd("HGET", "main:build:one:data", "Subgovname").Str()
			report.ErrLogger(err)

			//get subgov block title
			blocktitle, err := RDB.Cmd("HGET", "main:build:one:data", "Subgovreptitle").Str()
			report.ErrLogger(err)

			//get subgov block slot eg ward
			blockslot, err := RDB.Cmd("HGET", "main:build:one:data", "Subgovhouserepslot").Str()
			report.ErrLogger(err)

			var box []types.BuildFiveRep
			//for each seat,get its data fields
			for index := 1; index <= numofseats; index++ {
				seatKey := govKey + ":" + strconv.Itoa(index)
				res, err := RDB.Cmd("HGETALL", seatKey).Map()
				report.ErrLogger(err)
				//convert map[string]interface{} to struct of build four
				var target types.BuildFive
				if err := mp.Decode(res, &target); err != nil {
					report.ErrLogger(err)
				}
				var container types.BuildFiveRep //box for each individual rep seat
				container.From = subgovname
				container.Block = block
				container.BlockTitle = blocktitle
				container.BlockSlot = blockslot
				container.RepData = target
				//append to box
				box = append(box, container)
			}
			if len(box) == numofseats {
				for _, v := range box {
					rAll.RepsData = append(rAll.RepsData, v)
				}
			}

		}
	}
	return rAll
}
