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
	"errors"
	"strconv"

	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"

	"strings"

	"github.com/daviddexter/wajibu/server/dbase"
	"github.com/daviddexter/wajibu/server/radix"
	mp "github.com/mitchellh/mapstructure"
)

func processGrassRootGovLevel() *bool {
	wg.Add(2)
	var grassData []types.HousePosition
	go func() {
		defer wg.Done()
		c := dbase.CreateGrassRootLevelTable()
		if c != true {
			e := errors.New("Could not create grassroot level table")
			report.ErrLogger(e)
		}
	}()
	go func() {
		defer wg.Done()
		mainKey := "main:" + radix.BuildOneData
		res, err := radix.RDB.Cmd("HGETALL", mainKey).Map()
		report.ErrLogger(err)
		num, _ := strconv.Atoi(res["Numofsubgov"])
		slotdesignation := res["Subgovname"]
		housename := res["Subgovhousename"]
		title := res["Subgovreptitle"]
		for index := 0; index < num; index++ {
			govKey := radix.BuildFour + ":gov:" + strconv.Itoa(index)
			//get the number of leg seats in each sub gov
			res, err := radix.RDB.Cmd("HGET", govKey, "NumOfLegSeats").Str()
			report.ErrLogger(err)
			num, _ := strconv.Atoi(res)
			for index := 1; index <= num; index++ {
				seatKey := govKey + ":" + strconv.Itoa(index)
				res, err := radix.RDB.Cmd("HGETALL", seatKey).Map()
				report.ErrLogger(err)
				var b types.BuildFive
				var data types.HousePosition
				//map to struct
				if err := mp.Decode(res, &b); err != nil {
					report.ErrLogger(err)
				}

				data.Gender = strings.ToLower(b.RepGender)
				data.HouseName = strings.ToLower(housename)
				data.Image = strings.ToLower(b.RepImage)
				data.Name = strings.ToLower(b.RepName)
				data.NthPosition = strings.ToLower(b.RepnthPosition)
				data.SlotDesignation = strings.ToLower(slotdesignation)
				data.SlotName = strings.ToLower(b.SlotName)
				data.Term = strings.ToLower(b.RepTerm)
				data.Title = strings.ToLower(title)
				data.API = strings.ToLower(grassLevelPrefix + "/" + strings.TrimSpace(housename) + "/" + title + "/" + slotdesignation + "/" + b.SlotName + "/" + b.RepName)
				grassData = append(grassData, data)
			}
		}
	}()
	wg.Wait()
	return dbase.SaveGrassRootGovLevel(&grassData, DeploySpan)
}
