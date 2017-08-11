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

func processSubGovLevel() *bool {
	wg.Add(2)
	var subgovdata []types.TopPosition
	go func() {
		wg.Done()
		c := dbase.CreateSubGovLevelTable()
		if c != true {
			e := errors.New("Could not create subgov level table")
			report.ErrLogger(e)
		}
	}()
	go func() {
		wg.Done()
		//get the number of subgovs,title
		mainKey := "main:" + radix.BuildOneData
		res, err := radix.RDB.Cmd("HGETALL", mainKey).Map()
		report.ErrLogger(err)
		num, _ := strconv.Atoi(res["Numofsubgov"])
		position := res["Subgovtitle"]
		slotdesignation := res["Subgovname"]
		for index := 0; index < num; index++ {
			govKey := radix.BuildFour + ":gov:" + strconv.Itoa(index)
			res, err := radix.RDB.Cmd("HGETALL", govKey).Map()
			report.ErrLogger(err)
			var b types.BuildFour
			var data types.TopPosition
			//map to struct
			if err := mp.Decode(res, &b); err != nil {
				report.ErrLogger(err)
			}
			data.Gender = strings.ToLower(b.HeadGender)
			data.Image = strings.ToLower(b.HeadImage)
			data.Name = strings.ToLower(b.HeadName)
			data.NthPosition = strings.ToLower(b.HeadnthPosition)
			data.Position = strings.ToLower(position)
			data.SlotDesignation = strings.ToLower(slotdesignation)
			data.SlotName = strings.ToLower(b.SlotName)
			data.Term = strings.ToLower(b.HeadTerm)
			data.API = strings.ToLower(subGovLevelPrefix + "/" + slotdesignation + "/" + b.SlotName + "/" + position + "/" + b.HeadName)
			subgovdata = append(subgovdata, data)
		}
	}()
	wg.Wait()
	return dbase.SaveSubGovLevel(&subgovdata, DeploySpan)
}
