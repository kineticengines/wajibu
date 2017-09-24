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
	tmpl "text/template"
	"time"

	"github.com/daviddexter/moment"
	cfg "github.com/daviddexter/wajibu/configure"
	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
)

func AddSlotsFromDeploy(slots []types.Slot) *bool {
	//create slottable if it does not exist
	var verdict bool
	table := cfg.Loader().SlotsTable
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS ` + table + `(
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		slotname VARCHAR(255) NOT NULL,	
		designation VARCHAR(255) NOT NULL,
		createdDate VARCHAR(255) NOT NULL,	
		PRIMARY KEY (id),
		UNIQUE KEY (slotname)
	)`)
	report.ErrLogger(err)
	for _, v := range slots {
		var errX error
		var errY error
		stmt, err := DB.Prepare(`INSERT INTO ` + table + ` SET slotname=?,designation=?,
		createdDate=?`)
		report.ErrLogger(err)
		res, errX := stmt.Exec(tmpl.JSEscapeString(tmpl.HTMLEscapeString(v.SlotName)), tmpl.JSEscapeString(tmpl.HTMLEscapeString(v.Designation)), time.Now().Format("Jan 2, 2006 at 3:04pm MST"))
		report.ErrLogger(errX)
		_, errY = res.LastInsertId()
		report.ErrLogger(errY)
		if errX == nil && errY == nil {
			verdict = true
		} else {
			verdict = false
		}
	}
	return &verdict
}

func AddTitlesFromDeploy(slots []string) *bool {
	//create slottable if it does not exist
	var verdict bool
	table := cfg.Loader().TitlesTable
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS ` + table + `(
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		titlename VARCHAR(255) NOT NULL,	
		createdDate VARCHAR(255) NOT NULL,	
		PRIMARY KEY (id),
		UNIQUE KEY (titlename)
	)`)
	report.ErrLogger(err)
	for _, v := range slots {
		var errX error
		var errY error
		stmt, err := DB.Prepare(`INSERT INTO ` + table + ` SET titlename=?,
		createdDate=?`)
		report.ErrLogger(err)
		res, errX := stmt.Exec(tmpl.JSEscapeString(tmpl.HTMLEscapeString(v)), time.Now().Format("Jan 2, 2006 at 3:04pm MST"))
		report.ErrLogger(errX)
		_, errY = res.LastInsertId()
		report.ErrLogger(errY)
		if errX == nil && errY == nil {
			verdict = true
		} else {
			verdict = false
		}
	}
	return &verdict
}

func CreateTopLevelTable() bool {
	var r bool
	table := cfg.Loader().TopLevelTable
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS ` + table + `(
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,	
		position VARCHAR(255) NOT NULL,
		term VARCHAR(255) NOT NULL,
		gender VARCHAR(255) NOT NULL,
		nthPosition VARCHAR(255) NOT NULL,
		imageurl VARCHAR(255) NOT NULL,
		api VARCHAR(255) NOT NULL,
		termstart VARCHAR(255) NOT NULL,
		termend VARCHAR(255) NOT NULL,
		createdDate VARCHAR(255) NOT NULL,	
		PRIMARY KEY (id),
		UNIQUE KEY (name)		
	)`)
	if err == nil {
		r = true
	} else {
		r = false
	}
	return r
}

func SaveTopLevel(d *[]types.TopPosition, span int) *bool {
	var done int
	var r bool
	for _, v := range *d {
		table := cfg.Loader().TopLevelTable
		stmt, err := DB.Prepare(`INSERT INTO ` + table + ` SET name=?,position=?,term=?,gender=?,imageurl=?,
		nthPosition=?,api=?,termstart=?,termend=?,createdDate=?`)
		report.ErrLogger(err)
		termstart := time.Now().Year()
		n := moment.UtilBuilder{moment.ADDOPERATION, moment.YEARLEAP, 5}
		termend, err := n.Add(time.Now())
		report.ErrLogger(err)
		res, errX := stmt.Exec(v.Name, v.Position, v.Term, v.Gender, v.Image, v.NthPosition, v.API, termstart, termend.Year(), time.Now().Format("Jan 2, 2006 at 3:04pm MST"))
		report.ErrLogger(errX)
		_, err = res.LastInsertId()
		if err == nil {
			done++
		}
	}
	switch done {
	case len(*d):
		r = true
	default:
		r = false
	}
	return &r
}

func CreateHouseLevelTable() bool {
	var r bool
	table := cfg.Loader().HouseLevelTable
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS ` + table + `(
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		housename VARCHAR(255) NOT NULL,
		title VARCHAR(255) NOT NULL,	
		name VARCHAR(255) NOT NULL,
		term VARCHAR(255) NOT NULL,
		gender VARCHAR(255) NOT NULL,
		slotdesignation VARCHAR(255) NOT NULL,
		slotname VARCHAR(255) NOT NULL,
		imageurl VARCHAR(255) NOT NULL,
		api VARCHAR(255) NOT NULL,
		termstart VARCHAR(255) NOT NULL,
		termend VARCHAR(255) NOT NULL,
		createdDate VARCHAR(255) NOT NULL,	
		PRIMARY KEY (id),
		UNIQUE KEY (slotname)		
	)`)
	if err == nil {
		r = true
	} else {
		r = false
	}
	return r
}

func SaveHouseLevel(d *[]types.HousePosition, span int) *bool {
	var done int
	var r bool
	for _, v := range *d {
		table := cfg.Loader().HouseLevelTable
		stmt, err := DB.Prepare(`INSERT INTO ` + table + ` SET housename=?,title=?,name=?,term=?,gender=?,slotdesignation=?,
		slotname=?,imageurl=?,api=?,termstart=?,termend=?,createdDate=?`)
		report.ErrLogger(err)
		termstart := time.Now().Year()
		n := moment.UtilBuilder{moment.ADDOPERATION, moment.YEARLEAP, 5}
		termend, err := n.Add(time.Now())
		report.ErrLogger(err)
		res, errX := stmt.Exec(v.HouseName, v.Title, v.Name, v.Term, v.Gender, v.SlotDesignation, v.SlotName, v.Image, v.API, termstart, termend, time.Now().Format("Jan 2, 2006 at 3:04pm MST"))
		report.ErrLogger(errX)
		_, err = res.LastInsertId()
		if err == nil {
			done++
		}
	}
	switch done {
	case len(*d):
		r = true
	default:
		r = false
	}
	return &r
}

func CreateSubGovLevelTable() bool {
	var r bool
	table := cfg.Loader().SubGovLevelTable
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS ` + table + `(
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,	
		position VARCHAR(255) NOT NULL,
		term VARCHAR(255) NOT NULL,
		gender VARCHAR(255) NOT NULL,
		slotdesignation VARCHAR(255) NOT NULL,
		slotname VARCHAR(255) NOT NULL,
		nthPosition VARCHAR(255) NOT NULL,
		imageurl VARCHAR(255) NOT NULL,
		api VARCHAR(255) NOT NULL,
		termstart VARCHAR(255) NOT NULL,
		termend VARCHAR(255) NOT NULL,
		createdDate VARCHAR(255) NOT NULL,	
		PRIMARY KEY (id)		
	)`)
	if err == nil {
		r = true
	} else {
		r = false
	}
	return r
}

func SaveSubGovLevel(d *[]types.TopPosition, span int) *bool {
	var done int
	var r bool
	for _, v := range *d {
		table := cfg.Loader().SubGovLevelTable
		stmt, err := DB.Prepare(`INSERT INTO ` + table + ` SET name=?,position=?,term=?,gender=?,slotdesignation=?,slotname=?,
		nthPosition=?,imageurl=?,api=?,termstart=?,termend=?,createdDate=?`)
		report.ErrLogger(err)
		termstart := time.Now().Year()
		n := moment.UtilBuilder{moment.ADDOPERATION, moment.YEARLEAP, 5}
		termend, err := n.Add(time.Now())
		report.ErrLogger(err)
		res, errX := stmt.Exec(v.Name, v.Position, v.Term, v.Gender, v.SlotDesignation, v.SlotName, v.NthPosition, v.Image, v.API, termstart, termend, time.Now().Format("Jan 2, 2006 at 3:04pm MST"))
		report.ErrLogger(errX)
		_, err = res.LastInsertId()
		if err == nil {
			done++
		}
	}
	switch done {
	case len(*d):
		r = true
	default:
		r = false
	}
	return &r
}

func CreateGrassRootLevelTable() bool {
	var r bool
	table := cfg.Loader().GrassRootLevelTable
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS ` + table + `(
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		housename VARCHAR(255) NOT NULL,
		title VARCHAR(255) NOT NULL,	
		name VARCHAR(255) NOT NULL,
		term VARCHAR(255) NOT NULL,
		gender VARCHAR(255) NOT NULL,
		slotdesignation VARCHAR(255) NOT NULL,
		slotname VARCHAR(255) NOT NULL,
		legof VARCHAR(255) NOT NULL,
		nthposition VARCHAR(255) NOT NULL,
		imageurl VARCHAR(255) NOT NULL,	
		api VARCHAR(255) NOT NULL,	
		termstart VARCHAR(255) NOT NULL,
		termend VARCHAR(255) NOT NULL,
		createdDate VARCHAR(255) NOT NULL,	
		PRIMARY KEY (id),
		UNIQUE KEY (slotname)		
	)`)

	if err == nil {
		r = true
	} else {
		r = false
	}
	return r
}

func SaveGrassRootGovLevel(d *[]types.HousePosition, span int) *bool {
	var done int
	var r bool
	for _, v := range *d {
		table := cfg.Loader().GrassRootLevelTable
		stmt, err := DB.Prepare(`INSERT INTO ` + table + ` SET housename=?,title=?,name=?,term=?,gender=?,
		slotdesignation=?,slotname=?,legof=?,nthPosition=?,imageurl=?,api=?,termstart=?,termend=?,createdDate=?`)
		report.ErrLogger(err)
		termstart := time.Now().Year()
		n := moment.UtilBuilder{moment.ADDOPERATION, moment.YEARLEAP, 5}
		termend, err := n.Add(time.Now())
		report.ErrLogger(err)
		res, errX := stmt.Exec(v.HouseName, v.Title, v.Name, v.Term, v.Gender, v.SlotDesignation, v.SlotName, v.LegOf, v.NthPosition, v.Image, v.API, termstart, termend, time.Now().Format("Jan 2, 2006 at 3:04pm MST"))
		report.ErrLogger(errX)
		_, err = res.LastInsertId()
		if err == nil {
			done++
		}
	}
	switch done {
	case len(*d):
		r = true
	default:
		r = false
	}

	return &r
}
