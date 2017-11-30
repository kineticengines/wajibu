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

func DeployBuild() {
	switchBuildSteps()
}

func switchBuildSteps() {
	//depends if isCentralGovernment or not
	switch isCentralGovernment {
	case true:
		switchForCentral()
	case false:
		switchForNonCentral()
	}
}

func switchForCentral() {
	k := getInCompleteSteps(centralBuildKey, DeploySteps)
	for _, v := range k {
		switch v {
		case 1:
			generateRepSlots(isCentralGovernment)
		case 2:
			generateRepTitles(isCentralGovernment)
		case 3:
			generateTopLevelAPI()
		case 4:
			generateHouseLevelAPI()
		}
	}
}

func switchForNonCentral() {
	k := getInCompleteSteps(centralBuildKey, DeploySteps)
	for _, v := range k {
		switch v {
		case 1:
			generateRepSlots(isCentralGovernment)
		case 2:
			generateRepTitles(isCentralGovernment)
		case 3:
			generateTopLevelAPI()
		case 4:
			generateHouseLevelAPI()
		case 5:
			generateSubGovLevelAPI()
		case 6:
			generateGrassRootLevelAPI()
		}
	}
}

func getInCompleteSteps(key string, steps int) []int {
	var inComplete []int
	for index := 1; index <= steps; index++ {
		m := key + ":" + strconv.Itoa(index)
		res, err := radix.RDB.Cmd("SISMEMBER", key, m).Int()
		report.ErrLogger(err)
		switch res {
		case 0:
			inComplete = append(inComplete, index)
			break
		case 1:
			break //deploy is complete.
		}
	}
	return inComplete
}
