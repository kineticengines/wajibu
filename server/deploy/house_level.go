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

func processHouseLevel() *bool {
	wg.Add(2)
	var houseData []types.HousePosition
	go func() {
		defer wg.Done()
		c := dbase.CreateHouseLevelTable()
		if c != true {
			e := errors.New("Could not create house level table")
			report.ErrLogger(e)
		}
	}()
	go func() {
		defer wg.Done()
		for index := 0; index < NumOfHouses; index++ {
			//get title and slot designation
			key := "hdetails:" + radix.BuildOneData + ":" + strconv.Itoa(index)
			res, err := radix.RDB.Cmd("HGETALL", key).Map()
			report.ErrLogger(err)
			title := res["Reptitle"]
			slotdesignation := res["Repslot"]
			setKey := radix.BuildThree + ":house:" + strconv.Itoa(index)
			//get the members
			members, err := radix.RDB.Cmd("SMEMBERS", setKey).List()
			report.ErrLogger(err)
			for _, member := range members {
				//get the data for each
				res, err := radix.RDB.Cmd("HGETALL", member).Map()
				report.ErrLogger(err)
				var b types.BuildThree
				var data types.HousePosition
				//map to struct
				if err := mp.Decode(res, &b); err != nil {
					report.ErrLogger(err)
				}
				data.HouseName = strings.ToLower(b.HouseName)
				data.Title = strings.ToLower(title)
				data.SlotDesignation = strings.ToLower(slotdesignation)
				data.Name = strings.ToLower(b.RepName)
				data.Image = strings.ToLower(b.RepImage)
				data.SlotName = strings.ToLower(b.RepSlot)
				data.Term = strings.ToLower(b.RepTerm)
				data.Gender = strings.ToLower(b.RepGender)
				data.API = strings.ToLower(houseLevelPrefix + "/" + strings.TrimSpace(b.HouseName) + "/" + title + "/" + slotdesignation + "/" + b.RepSlot + "/" + b.RepName)
				houseData = append(houseData, data)
			}
		}
	}()
	wg.Wait()
	return dbase.SaveHouseLevel(&houseData, DeploySpan)
}
