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

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
)

type HouseBuilt struct {
	Key        string                 `json:"key"`
	BuiltData  types.BuidlOneHDetails `json:"builtdata"`
	NumOfSeats string                 `json:"numofseats"`
	complete   bool
}

func BuildThreeInitializer() *HouseBuilt {
	var house HouseBuilt
	builds := FetchBuildLevel()
	for _, v := range builds {
		switch v.Level {
		case BuildOne:
			n := checkHousesAndBuildKeys(v.Data)
			h := housePicker(n)
			house.BuiltData = h.BuiltData
			house.complete = h.complete
			house.Key = h.Key
			house.NumOfSeats = h.NumOfSeats
			break
		}
	}
	return &house
}

func housePicker(h []HouseBuilt) HouseBuilt {
	buildInComplete := 0
	var house HouseBuilt
	//check if both have complete as FALSE
	for _, v := range h {
		if v.complete == false {
			buildInComplete++
		}
	}
	//if all houses are incomplete,randomly pick one
	if buildInComplete == len(h) && len(h) > 1 {
		index := randomdata.Number(len(h) - 1)
		house.BuiltData = h[index].BuiltData
		house.complete = h[index].complete
		house.Key = h[index].Key
		house.NumOfSeats = h[index].NumOfSeats
	} else if buildInComplete == len(h) && len(h) == 1 {
		if h[0].complete == false {
			house.BuiltData = h[0].BuiltData
			house.complete = h[0].complete
			house.Key = h[0].Key
			house.NumOfSeats = h[0].NumOfSeats
		}
	} else {
		//loop and return house which is not complete
		for _, v := range h {
			if v.complete == false {
				house.BuiltData = v.BuiltData
				house.complete = v.complete
				house.Key = v.Key
				house.NumOfSeats = v.NumOfSeats
			}
		}
	}
	return house
}

func checkHousesAndBuildKeys(data interface{}) []HouseBuilt {
	var built HouseBuilt
	var builtSlice = make([]HouseBuilt, 0)
	switch dd := data.(type) {
	case types.BuildOneAll:
		for i := 0; i < len(dd.HDetails); i++ {
			houseKey := BuildThree + ":house:" + strconv.Itoa(i)
			res, err := RDB.Cmd("EXISTS", houseKey).Int() //check if the SET exists
			report.ErrLogger(err)
			switch res {
			case 0:
				//build SET does not exist
				//create a SET of the houseKey. Each member of the set should be a key to a Hash
				ln, _ := strconv.Atoi(dd.HDetails[i].Numofseats)
				var hk []string
				for n := 1; n <= ln; n++ {
					mKey := BuildThree + ":house:" + strconv.Itoa(i) + ":" + strconv.Itoa(n)
					err := RDB.Cmd("SADD", houseKey, mKey).Err
					report.ErrLogger(err)
					hk = append(hk, mKey)
				}
				//check if all HASH keys for this house exist.
				c := houseHashChecker(hk, ln)
				built.Key = houseKey
				built.BuiltData = dd.HDetails[i]
				built.NumOfSeats = dd.HDetails[i].Numofseats
				built.complete = c
				builtSlice = append(builtSlice, built)
			case 1:
				//build SET exists
				//get all members of the SET
				ln, _ := strconv.Atoi(dd.HDetails[i].Numofseats)
				res, err := RDB.Cmd("SMEMBERS", houseKey).List()
				report.ErrLogger(err)
				c := houseHashChecker(res, ln)
				built.Key = houseKey
				built.BuiltData = dd.HDetails[i]
				built.NumOfSeats = dd.HDetails[i].Numofseats
				built.complete = c
				builtSlice = append(builtSlice, built)
			}
		}
		break
	}
	return builtSlice
}

func houseHashChecker(keys []string, numOfSeats int) bool {
	x := 0
	for _, key := range keys {
		//check if the key exists
		res, err := RDB.Cmd("EXISTS", key).Int()
		report.ErrLogger(err)
		switch res {
		case 1:
			//key exists
			x++
		}
	}
	if x == numOfSeats {
		return true
	}
	return false

}
