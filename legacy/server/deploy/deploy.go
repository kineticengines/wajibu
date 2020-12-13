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
	"strconv"

	"github.com/daviddexter/wajibu/report"
	"github.com/daviddexter/wajibu/server/radix"
)

var (
	DeploySteps         int
	DeploySpan          int
	NumOfHouses         int
	isCentralGovernment bool
	centralBuildKey     string = "central:build"
	nonCentralBuildKey  string = "non:central:build"
	topLevelPrefix      string = "api/top/level"
	houseLevelPrefix    string = "api/house/level"
	subGovLevelPrefix   string = "api/subgov/level"
	grassLevelPrefix    string = "api/grass/level"
)

func Initilizer() {
	radix.ConnectToRDB()
	//check the type of government
	c := CheckIfCentralGov()
	switch c {
	case true:
		isCentralGovernment = true
		DeploySteps = 4
	case false:
		isCentralGovernment = false
		//check if sub gov has leg
		key := "main:" + radix.BuildOneData
		res, err := radix.RDB.Cmd("HGET", key, "Subgovhasleg").Str()
		report.ErrLogger(err)
		b, _ := strconv.ParseBool(res)
		switch b {
		case true:
			DeploySteps = 6
		case false:
			DeploySteps = 5
		}

	}
}

func CheckIfCentralGov() bool {
	key := "main:" + radix.BuildOneData
	res, err := radix.RDB.Cmd("HGETALL", key).Map()
	report.ErrLogger(err)
	var verdict bool
	n, _ := strconv.Atoi(res["Deployspan"])
	m, _ := strconv.Atoi(res["Numofhouses"])
	DeploySpan = n
	NumOfHouses = m
	switch res["Governmenttype"] {
	case "Central Government":
		verdict = true
	default:
		verdict = false
	}
	return verdict
}
