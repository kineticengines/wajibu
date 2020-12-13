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
	"log"
	"strconv"

	"github.com/daviddexter/wajibu/report"
)

func ResetBuildThree() *bool {
	r := true
	//delete all build three keys
	errA := RDB.Cmd("DEL", BuildThree).Err
	if errA != nil {
		r = false
	}
	keys1, err1 := RDB.Cmd("keys", "build:three:*").List()
	report.ErrLogger(err1)
	for _, key := range keys1 {
		err := RDB.Cmd("DEL", key).Err
		if err != nil {
			r = false
		}

	}
	keys2, err2 := RDB.Cmd("keys", "main:build:three:*").List()
	report.ErrLogger(err2)
	for _, key := range keys2 {
		err := RDB.Cmd("DEL", key).Err
		if err != nil {
			r = false
		}
	}

	return &r
}

func ResetBuildFour() *bool {
	r := true
	//delete all build four keys
	errA := RDB.Cmd("DEL", BuildFour).Err
	if errA != nil {
		r = false
	}
	keys1, err1 := RDB.Cmd("keys", "build:four:*").List()
	report.ErrLogger(err1)
	for _, key := range keys1 {
		err := RDB.Cmd("DEL", key).Err
		if err != nil {
			r = false
		}
	}
	keys2, err2 := RDB.Cmd("keys", "main:build:four:*").List()
	report.ErrLogger(err2)
	for _, key := range keys2 {
		err := RDB.Cmd("DEL", key).Err
		if err != nil {
			r = false
		}
	}

	return &r
}

func ResetBuildFive() *bool {
	r := true
	//delete all build five keys
	errA := RDB.Cmd("DEL", BuildFive).Err
	if errA != nil {
		r = false
	}

	mainKey := "main:" + BuildFourData
	res, err := RDB.Cmd("HGET", mainKey, "NumOfSubGovs").Str()
	log.Println(err) //check later
	//log.Println("reset error")
	//report.ErrLogger(err)
	numofsubgovs, _ := strconv.Atoi(res)

	for index := 0; index < numofsubgovs; index++ {
		govKey := BuildFour + ":gov:" + strconv.Itoa(index)
		res, err := RDB.Cmd("HGET", govKey, "NumOfLegSeats").Str()
		report.ErrLogger(err)
		num, _ := strconv.Atoi(res)
		for index := 1; index <= num; index++ {
			seatKey := govKey + ":" + strconv.Itoa(index)
			err := RDB.Cmd("DEL", seatKey).Err
			if err != nil {
				r = false
			}
		}
	}

	return &r
}
