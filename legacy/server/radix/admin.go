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
	"errors"

	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
	"github.com/daviddexter/wajibu/server/dbase"
)

func AllTitles() *[]string {
	var rAll []string
	res, err := RDB.Cmd("EXISTS", TITLES).Int()
	report.ErrLogger(err)
	switch res {
	case 0:
		//does not exist in redis
		p := *dbase.GetTitlesFromDB()
		if p.Error == nil {
			for _, v := range p.Title {
				rAll = append(rAll, v)
			}
		}
	case 1:
		//get data from redis : smembers
		res, err := RDB.Cmd("SMEMBERS", TITLES).List()
		report.ErrLogger(err)
		for _, v := range res {
			rAll = append(rAll, v)
		}
	}
	return &rAll
}

func AllHouses() *[]string {
	var rAll []string
	res, err := RDB.Cmd("EXISTS", HOUSES).Int()
	report.ErrLogger(err)
	switch res {
	case 0:
		//does not exist in redis
		p := *dbase.GetHousesFromDB()
		if p.Error == nil {
			for _, v := range p.House {
				rAll = append(rAll, v)
			}
		}
	case 1:
		//get data from redis : smembers
		res, err := RDB.Cmd("SMEMBERS", HOUSES).List()
		report.ErrLogger(err)
		for _, v := range res {
			rAll = append(rAll, v)
		}
	}
	return &rAll
}

func AllSubGovs() *struct {
	Designation string   `json:"designation"`
	Govs        []string `json:"govs"`
} {
	var rGovs []string
	var designation string
	res, err := RDB.Cmd("EXISTS", SUBGOVS).Int()
	report.ErrLogger(err)
	switch res {
	case 0:
		//does not exist in redis
		p := *dbase.GetSubGovsFromDB()
		if p.Data.Error == nil {
			designation = p.Designation
			for _, v := range p.Data.House {
				rGovs = append(rGovs, v)
			}
		}
	case 1:
		//get data from redis : smembers
		res, err := RDB.Cmd("SMEMBERS", SUBGOVS).List()
		report.ErrLogger(err)
		for _, v := range res {
			rGovs = append(rGovs, v)
		}
	}
	var rAll struct {
		Designation string   `json:"designation"`
		Govs        []string `json:"govs"`
	}
	rAll.Designation = designation
	rAll.Govs = rGovs
	return &rAll
}

func AllPillars() (*[]types.Pillar, error) {
	var rAll []types.Pillar
	var pillarErr error
	res, err := RDB.Cmd("EXISTS", PILLARS).Int()
	report.ErrLogger(err)
	switch res {
	case 0:
		//does not exist in redis
		p := *dbase.GetPillarsFromDB()
		if p.Error == nil {
			for _, v := range p.Pillars {
				rAll = append(rAll, v)
			}
			pillarErr = nil
		} else {
			pillarErr = errors.New("Error getting pillars")
		}
	case 1:
		//get data from redis : smembers
		/*res, err := RDB.Cmd("SMEMBERS", PILLARS).List()
		report.ErrLogger(err)
		for _, v := range res {
			rAll = append(rAll, v)
		}*/
	}
	return &rAll, pillarErr
}

func AddSentimentsTableFields(d map[string]interface{}) {
	for key, _ := range d {
		err := RDB.Cmd("SADD", SENTIMENT, key).Err
		report.ErrLogger(err)
	}
}
