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
)

func BuildFourInitializer() *types.BuildFourAll {
	var rAll types.BuildFourAll
	builds := FetchBuildLevel()
	for _, v := range builds {
		switch v.Level {
		case BuildOne:
			//check if gov is central in build:one
			data, cn := checkIfCentral(v.Data)
			if cn == true {
				rAll.IsCentral = true
			} else {
				buildSubGovHasKeys(data, &rAll)
			}
			break
		}
	}

	return &rAll

}
func checkIfCentral(data interface{}) (types.BuildOneAll, bool) {
	var r bool
	var b types.BuildOneAll
	switch dd := data.(type) {
	case types.BuildOneAll:
		switch dd.Main.Governmenttype {
		case "Central Government":
			r = true
			b = dd
			break
		default:
			r = false
			b = dd
			break
		}
	}
	return b, r
}

func buildSubGovHasKeys(data types.BuildOneAll, govInit *types.BuildFourAll) {
	c, err := strconv.Atoi(data.Main.Numofsubgov)
	report.ErrLogger(err)
	govInit.Numofsubgov = data.Main.Numofsubgov
	govInit.OfficeTitle = data.Main.Subgovtitle
	govInit.Subgovname = data.Main.Subgovname
	v := 0
	for i := 0; i < c; i++ {
		//sub government hash key
		govKey := BuildFour + ":gov:" + strconv.Itoa(i)
		res, err := RDB.Cmd("EXISTS", govKey).Int() //check if the HASH exists
		report.ErrLogger(err)
		switch res {
		case 0:
			//build HASH does not exist
			v++
		}
	}
	if v == c {
		govInit.Complete = false
	} else {
		govInit.Complete = true
	}

	switch data.Main.Subgovhasleg {
	case true:
		//subgov have legislative arm
		govInit.HasLeg = true
	case false:
		//subgovs have no legislative arm
		govInit.HasLeg = false

	}

}
