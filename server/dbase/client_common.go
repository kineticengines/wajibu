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

package dbase

import (
	"strconv"
	"sync"
	tmpl "text/template"

	cfg "github.com/daviddexter/wajibu/configure"
	"github.com/daviddexter/wajibu/handlers/types"
)

func CheckIfLocation(s string) *bool {
	table := cfg.Loader().SlotsTable
	var wg sync.WaitGroup
	var rAll bool
	wg.Add(1)
	go func() {
		defer wg.Done()
		res := DB.QueryRow(`SELECT COUNT(slotname) FROM `+table+` WHERE slotname=?`, tmpl.JSEscapeString(tmpl.HTMLEscapeString(s)))
		var n string
		res.Scan(&n)
		switch j, _ := strconv.Atoi(n); j {
		case 0:
			rAll = false
			break
		default:
			rAll = true
			break
		}
	}()
	wg.Wait()
	return &rAll
}

func CheckIfTitle(s string) *bool {
	table := cfg.Loader().TitlesTable
	var wg sync.WaitGroup
	var rAll bool
	wg.Add(1)
	go func() {
		defer wg.Done()
		res := DB.QueryRow(`SELECT COUNT(titlename) FROM `+table+` WHERE titlename=?`, tmpl.JSEscapeString(tmpl.HTMLEscapeString(s)))
		var n string
		res.Scan(&n)
		switch j, _ := strconv.Atoi(n); j {
		case 0:
			rAll = false
			break
		default:
			rAll = true
			break
		}
	}()
	wg.Wait()
	return &rAll
}

func CheckIfPillar(s string) *bool {
	table := cfg.Loader().PillarsTable
	var wg sync.WaitGroup
	var rAll bool
	wg.Add(1)
	go func() {
		defer wg.Done()
		res := DB.QueryRow(`SELECT COUNT(pillar) FROM `+table+` WHERE pillar=?`, tmpl.JSEscapeString(tmpl.HTMLEscapeString(s)))
		var n string
		res.Scan(&n)
		switch j, _ := strconv.Atoi(n); j {
		case 0:
			rAll = false
			break
		default:
			rAll = true
			break
		}
	}()
	wg.Wait()
	return &rAll
}

func GetAPIForLevel(item string, level string) *[]types.LevelType {
	var rAll []types.LevelType
	switch level {
	case "toplevel":
		var r types.LevelType
		table := cfg.Loader().TopLevelTable
		rows, err := DB.Query(`SELECT api FROM `+table+` WHERE position=?`, tmpl.JSEscapeString(tmpl.HTMLEscapeString(item)))
		if err == nil {
			for rows.Next() {
				var api string
				err := rows.Scan(&api)
				if err == nil {
					r.IsTrue = true
					r.API = api
					r.Level = "toplevel"
					rAll = append(rAll, r)
				}
			}
		}
		r.IsTrue = false
		rAll = append(rAll, r)
	case "houselevel":
		var r types.LevelType
		table := cfg.Loader().HouseLevelTable
		rows, err := DB.Query(`SELECT api FROM `+table+` WHERE title=?`, tmpl.JSEscapeString(tmpl.HTMLEscapeString(item)))
		if err == nil {
			for rows.Next() {
				var api string
				err := rows.Scan(&api)
				if err == nil {
					r.IsTrue = true
					r.API = api
					r.Level = "houselevel"
					rAll = append(rAll, r)
				}
			}
		}
		r.IsTrue = false
		rAll = append(rAll, r)
	case "subgovlevel":
		var r types.LevelType
		table := cfg.Loader().SubGovLevelTable
		rows, err := DB.Query(`SELECT api FROM `+table+` WHERE position=?`, tmpl.JSEscapeString(tmpl.HTMLEscapeString(item)))
		if err == nil {
			for rows.Next() {
				var api string
				err := rows.Scan(&api)
				if err == nil {
					r.IsTrue = true
					r.API = api
					r.Level = "subgovlevel"
					rAll = append(rAll, r)
				}
			}
		}
		r.IsTrue = false
		rAll = append(rAll, r)
	case "grasslevel":
		var r types.LevelType
		table := cfg.Loader().GrassRootLevelTable
		rows, err := DB.Query(`SELECT api FROM `+table+` WHERE title=?`, tmpl.JSEscapeString(tmpl.HTMLEscapeString(item)))
		if err == nil {
			for rows.Next() {
				var api string
				err := rows.Scan(&api)
				if err == nil {
					r.IsTrue = true
					r.API = api
					r.Level = "grasslevel"
					rAll = append(rAll, r)
				}
			}
		}
		r.IsTrue = false
		rAll = append(rAll, r)
	}
	return &rAll
}

func DetailForRepSlot(s string) *map[string]string {
	table := cfg.Loader().SlotsTable
	Slots := make(map[string]string)
	rows, err := DB.Query(`SELECT designation FROM `+table+` WHERE slotname=?`, tmpl.JSEscapeString(tmpl.HTMLEscapeString(s)))
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var d string
			err := rows.Scan(&d)
			if err == nil {
				Slots[d] = s
			}
		}
	}
	return &Slots
}
