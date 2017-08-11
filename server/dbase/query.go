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
	"database/sql"
	"log"
	"strconv"
	"time"

	"sync"

	cfg "github.com/daviddexter/wajibu/configure"
	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
)

var wg sync.WaitGroup

const DefaultRole string = "SUPER"

func DefaultToDB(table string, username string, password string, email string) bool {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS ` + table + `(
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		role VARCHAR(255) NOT NULL,
		createdDate VARCHAR(255) NOT NULL,
		createdBy VARCHAR(255) NOT NULL,
		PRIMARY KEY (id)
	)`)
	report.ErrLogger(err)
	stmt, err := DB.Prepare(`INSERT INTO ` + table + ` SET username=?,
	password=?,email=?,role=?,createdDate=?,createdBy=?`)
	report.ErrLogger(err)
	res, err := stmt.Exec(username, password, email, DefaultRole, time.Now(), "Auto Generated")
	report.ErrLogger(err)
	_, err = res.LastInsertId()
	report.ErrLogger(err)
	return true
}

type lQueryRes struct {
	Param string
	pass  string
	Exist bool
}

func CheckLoginCred(nameoremail string, password string, table string) lQueryRes {
	wg.Add(2)
	q := lQueryRes{Param: nameoremail, pass: password}
	var uBox lQueryRes
	var eBox lQueryRes
	var vBox lQueryRes
	go func(db *sql.DB, table string, i lQueryRes) {
		defer wg.Done()
		res := db.QueryRow(`SELECT COUNT(username) FROM `+table+` WHERE username = ? AND password = ? `, i.Param, i.pass)
		var n string
		res.Scan(&n)
		switch j, _ := strconv.Atoi(n); j {
		case 0:
			uBox.Param = i.Param
			uBox.Exist = false
			break
		default:
			uBox.Param = i.Param
			uBox.Exist = true
			break
		}

	}(DB, table, q)

	go func(db *sql.DB, table string, i lQueryRes) {
		defer wg.Done()
		res := db.QueryRow(`SELECT COUNT(email) FROM `+table+` WHERE email = ? AND password = ? `, i.Param, i.pass)
		var n string
		res.Scan(&n)
		switch j, _ := strconv.Atoi(n); j {
		case 0:
			eBox.Param = i.Param
			eBox.Exist = false
			break
		default:
			eBox.Param = i.Param
			eBox.Exist = true
			break
		}

	}(DB, table, q)
	wg.Wait()
	switch true {
	case uBox.Exist:
		vBox = uBox
	case eBox.Exist:
		vBox = eBox
	}
	return vBox
}

func GetTitlesFromDB() *types.TitleData {
	var r types.TitleData
	_, err := DB.Exec(`DESCRIBE ` + cfg.Loader().TitlesTable)
	switch err {
	case nil:
		//get titles
		rows, err := DB.Query(`SELECT titlename FROM ` + cfg.Loader().TitlesTable)
		if err == nil {
			for rows.Next() {
				var title string
				err = rows.Scan(&title)
				if err == nil {
					r.Title = append(r.Title, title)
				} else {
					r.Error = err
				}
			}
		} else {
			r.Error = err
		}
	default:
		r.Error = err
	}
	return &r
}

func GetPillarsFromDB() *types.PillarData {
	var r types.PillarData
	_, err := DB.Exec(`DESCRIBE ` + cfg.Loader().PillarsTable)
	switch err {
	case nil:
		//get pillars
		rows, err := DB.Query(`SELECT pillar,fortitle FROM ` + cfg.Loader().PillarsTable)
		if err == nil {
			for rows.Next() {
				var pillar string
				var fortitle string
				var p types.Pillar

				err := rows.Scan(&pillar, &fortitle)

				p.Pillar = pillar
				p.Fortitle = fortitle

				if err == nil {
					r.Pillars = append(r.Pillars, p)
				} else {
					r.Error = err
				}
			}
		} else {
			r.Error = err
		}
	default:
		r.Error = err
	}
	return &r
}

func GetPillarsFor(level string) *[]string {
	var r []string
	all := "All"
	switch level {
	case "president":
		_, err := DB.Exec(`DESCRIBE ` + cfg.Loader().PillarsTable)
		switch err {
		case nil:
			//get pillars
			rows, err := DB.Query(`SELECT pillar FROM `+cfg.Loader().PillarsTable+` WHERE fortitle=? OR fortitle=?`, level, all)
			if err == nil {
				for rows.Next() {
					var pillar string
					err := rows.Scan(&pillar)
					if err == nil {
						r = append(r, pillar)
					}
				}
			} else {
				log.Println(err)
			}
		}
	case "deputy president":
		_, err := DB.Exec(`DESCRIBE ` + cfg.Loader().PillarsTable)
		switch err {
		case nil:
			//get pillars
			rows, err := DB.Query(`SELECT pillar FROM `+cfg.Loader().PillarsTable+` WHERE fortitle=? OR fortitle=?`, level, all)
			if err == nil {
				for rows.Next() {
					var pillar string
					err := rows.Scan(&pillar)
					if err == nil {
						r = append(r, pillar)
					}
				}
			}
		}
	}
	return &r
}

func NewPillar(d string, e string) *bool {
	var r bool
	table := cfg.Loader().PillarsTable
	_, errDB := DB.Exec(`CREATE TABLE IF NOT EXISTS ` + table + `(
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		pillar VARCHAR(255) NOT NULL,
		fortitle VARCHAR(255) NOT NULL,
		createdAt VARCHAR(255) NOT NULL,
		PRIMARY KEY (id),
		UNIQUE KEY(pillar)
	)`)
	report.ErrLogger(errDB)
	// insert
	stmt, errI := DB.Prepare(`INSERT IGNORE INTO ` + table + ` (pillar,fortitle,createdAt) values(?,?,?)`)
	report.ErrLogger(errI)
	res, _ := stmt.Exec(d, e, time.Now())
	_, errL := res.LastInsertId()
	if errL == nil {
		r = true
	} else {
		r = false
	}
	return &r
}

func RemovePillar(d string) *bool {
	var r bool
	table := cfg.Loader().PillarsTable
	stmt, err := DB.Prepare(`DELETE FROM ` + table + ` WHERE pillar=? `)
	report.ErrLogger(err)
	res, _ := stmt.Exec(d)
	_, errL := res.LastInsertId()
	if errL == nil {
		r = true
	} else {
		r = false
	}
	return &r

}

func GetRepSlots() *map[string][]string {
	//table := cfg.Loader().SlotsTable
	var d []string
	Slots := make(map[string][]string)
	rows, err := DB.Query(`SELECT DISTINCT designation FROM slotstable` /*+ table*/)
	if err == nil {
		for rows.Next() {
			var slot types.Slot
			err := rows.Scan(&slot.Designation)
			if err == nil {
				d = append(d, slot.Designation)
			}
		}
	}
	for i := range d {
		key := d[i]
		var slotsOfKey []string
		//get the slotnames for the key
		rows, err := DB.Query(`SELECT DISTINCT slotname FROM slotstable WHERE designation=?`, key)
		if err == nil {
			for rows.Next() {
				var slotname string
				err := rows.Scan(&slotname)
				if err == nil {
					slotsOfKey = append(slotsOfKey, slotname)
				}
			}
		}
		Slots[key] = slotsOfKey
	}
	return &Slots
}

func GetAPIofForTopLevel(d string) *string {
	var theAPI string
	table := cfg.Loader().TopLevelTable
	rows, err := DB.Query(`SELECT api FROM `+table+` WHERE position=?`, d)
	if err == nil {
		for rows.Next() {
			err := rows.Scan(&theAPI)
			report.ErrLogger(err)
		}
	} else {
		log.Println(err)
	}
	return &theAPI
}
