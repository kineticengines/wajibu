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
	"reflect"
	"strconv"

	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
	"github.com/fatih/structs"
)

func AddCache(level string, dataLevel string, data interface{}) *bool {
	var r bool
	switch dd := data.(type) {
	case types.BuildOneAll: //cache for buildone

		/************* IMPORTANT ***********************/
		/****************************************************/
		err := RDB.Cmd("FLUSHALL").Err
		report.ErrLogger(err)
		/*******************************************************/
		/*******************************************************/

		for index := 0; index < len(dd.HDetails); index++ {
			key := "hdetails:" + dataLevel + ":" + strconv.Itoa(index)
			m := structs.Map(dd.HDetails[index])
			err := RDB.Cmd("HMSET", key, m).Err
			report.ErrLogger(err)
		}
		mainKey := "main:" + dataLevel
		//convert the struct to map[string]interface{}
		mainData := structs.Map(dd.Main)
		err1 := RDB.Cmd("HMSET", mainKey, mainData).Err
		report.ErrLogger(err1)

		err2 := RDB.Cmd("HMSET", dataLevel, "hlen", len(dd.HDetails)).Err
		report.ErrLogger(err2)

		if err1 == nil && err2 == nil {
			e := RDB.Cmd("HMSET", level, keyMap["complete"], "true", keyMap["data"], "true").Err
			report.ErrLogger(e)
			r = true
		} else {
			r = false
		}
	case types.BuildTwoAll: //build thwo cache
		//cache for buildtwo
		mainKey := "main:" + dataLevel
		//convert the struct to map[string]interface{}
		mainData := structs.Map(dd)
		err := RDB.Cmd("HMSET", mainKey, mainData).Err
		//report.ErrLogger(err)
		if err == nil {
			e := RDB.Cmd("HMSET", level, keyMap["complete"], "true", keyMap["data"], "true").Err
			report.ErrLogger(e)
			r = true
		} else {
			r = false
		}
	case map[string]interface{}: //build three cache
		//cache for buildthree
		type mainReturn struct{ Housecount int }
		var theMainReturn mainReturn
		theMainReturn.Housecount = 0
		mainKey := "main:" + dataLevel
		houseLevelKey := dd["key"]

		switch v := dd["main"].(type) {
		case []interface{}:
			sm, err := RDB.Cmd("SMEMBERS", houseLevelKey).List()
			report.ErrLogger(err)
			mp := make(map[string]interface{})
			for i, _ := range v {
				mp[sm[i]] = v[i]
			}
			for k, v := range mp {
				err := RDB.Cmd("HMSET", k, v, "HouseName", dd["housename"], "SlotName", dd["slotname"]).Err
				report.ErrLogger(err)

			}
			hs, _ := RDB.Cmd("HGET", mainKey, "Housecount").Int()
			hs++
			theMainReturn.Housecount = hs
			//build:three:data
			mainData := structs.Map(theMainReturn)
			er := RDB.Cmd("HMSET", mainKey, mainData).Err
			report.ErrLogger(er)
			//check if all houses have being filled.Return false if not
			c := HousesCompleteness()
			if c == true {
				//add complete to build three
				e := RDB.Cmd("HMSET", level, keyMap["complete"], "true", keyMap["data"], "true").Err
				report.ErrLogger(e)
				r = true
			} else {
				r = true
			}
		}
	case []types.BuildFour:
		switch reflect.TypeOf(data).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(data)
			size := s.Len()
			var ss int
			type mainReturn struct{ NumOfSubGovs int }
			var theMainReturn mainReturn
			mainKey := "main:" + dataLevel

			theMainReturn.NumOfSubGovs = size
			//build:three:data
			mainData := structs.Map(theMainReturn)
			er := RDB.Cmd("HMSET", mainKey, mainData).Err
			report.ErrLogger(er)

			for i := 0; i < size; i++ {
				//typecast to type.BuildFour type
				v := s.Index(i)
				x := v.Interface().(types.BuildFour)
				govKey := BuildFour + ":gov:" + strconv.Itoa(i)
				//convert to map[string]interface{}
				var data = structs.Map(x)
				err := RDB.Cmd("HMSET", govKey, data).Err
				report.ErrLogger(err)
				if err == nil {
					ss++
				}
			}

			if size == ss {
				e := RDB.Cmd("HMSET", level, keyMap["complete"], "true", keyMap["data"], "true").Err
				report.ErrLogger(e)
				if e == nil {
					r = true
				}
			} else {
				r = false
			}

		}
	case types.BuildFiveCache:
		switch reflect.TypeOf(data).Kind() {
		case reflect.Struct:
			s := reflect.ValueOf(data)
			v := s.Interface().(types.BuildFiveCache)
			gkey := v.Key
			for k, v := range v.TheData {
				i := k + 1
				repKey := gkey + ":" + strconv.Itoa(i)
				var data = structs.Map(v)
				err := RDB.Cmd("HMSET", repKey, data).Err
				report.ErrLogger(err)
			}
			c, _ := checkIfLevelIsComplete()
			if c == true {
				e := RDB.Cmd("HMSET", level, keyMap["complete"], "true", keyMap["data"], "true").Err
				report.ErrLogger(e)
				if e == nil {
					r = true
				}
			}
			r = true
		}
	}
	return &r
}
