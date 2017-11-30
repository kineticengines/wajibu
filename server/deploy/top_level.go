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
	"strings"

	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
	"github.com/daviddexter/wajibu/server/dbase"
	"github.com/daviddexter/wajibu/server/radix"
	mp "github.com/mitchellh/mapstructure"
)

func processTopLevel() *bool {
	wg.Add(2)
	var r []types.TopPosition
	go func() {
		defer wg.Done()
		c := dbase.CreateTopLevelTable()
		if c != true {
			e := errors.New("Could not create top level table")
			report.ErrLogger(e)
		}
	}()
	go func() {
		defer wg.Done()
		var b types.BuildTwoAll
		var pd types.TopPosition
		var dpp types.TopPosition
		k := "main:" + radix.BuildTwoData
		res, err := radix.RDB.Cmd("HGETALL", k).Map()
		report.ErrLogger(err)
		//map to struct
		if err := mp.Decode(res, &b); err != nil {
			report.ErrLogger(err)
		}
		pd.Name = strings.ToLower(b.PresidentName)
		pd.Position = "president"
		pd.Term = strings.ToLower(b.PresidentTerm)
		pd.NthPosition = strings.ToLower(b.PnthPosition)
		pd.Gender = strings.ToLower(b.Pgender)
		pd.Image = strings.ToLower(b.PImage)
		pd.API = strings.ToLower(topLevelPrefix + "/president/" + b.PresidentName)
		r = append(r, pd)
		dpp.Name = strings.ToLower(b.DpresidentName)
		dpp.Position = "deputy president"
		dpp.Term = strings.ToLower(b.DpresidentTerm)
		dpp.NthPosition = strings.ToLower(b.DnthPosition)
		dpp.Gender = strings.ToLower(b.Dgender)
		dpp.Image = strings.ToLower(b.DImage)
		dpp.API = strings.ToLower(topLevelPrefix + "/dpresident/" + b.DpresidentName)
		r = append(r, dpp)
	}()
	wg.Wait()
	return dbase.SaveTopLevel(&r, DeploySpan)

}
