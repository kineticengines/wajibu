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
	"strconv"
	"time"

	"sync"

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
