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
	"reflect"
	"time"

	cfg "github.com/daviddexter/wajibu/configure"
	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
)

func AddNewSentiment(data interface{}) *types.NewSentiment {
	var dataDB types.NewSentiment
	switch dd := data.(type) {
	case map[string]interface{}:
		for key, value := range dd {
			switch key {
			case "api":
				v := reflect.ValueOf(value)
				dataDB.API = v.Interface().(string)
			case "data":
				v := reflect.ValueOf(value)
				dataDB.Data = v.Interface().(map[string]interface{})
			}
		}
	}
	return &dataDB
}

func CreateOrAlterTable(d map[string]interface{}) {
	table := cfg.Loader().SentimentsTable
	_, err := DB.Exec(`DESCRIBE ` + table)
	if err != nil {
		//create the table
		for key, _ := range d {
			//add the keys to a set
			_, err := DB.Exec(`CREATE TABLE ` + table + `(
				id INT UNSIGNED NOT NULL AUTO_INCREMENT, api  VARCHAR(255) NOT NULL,
				` + key + ` VARCHAR(255) NULL,createdDate VARCHAR(255) NOT NULL,				
				PRIMARY KEY (id))`)
			if err != nil {
				CreateOrAlterTable(d)
			}
		}
	} else {
		//alter the table
		for key, _ := range d {
			_, err := DB.Query(`SELECT ` + key + ` FROM ` + table)
			if err != nil {
				_, err := DB.Exec(`ALTER TABLE ` + table + ` ADD ` + key + ` VARCHAR(255) NULL`)
				report.ErrLogger(err)
			}
		}

	}
}

func AddSentimentToDB(d types.NewSentiment) *bool {
	table := cfg.Loader().SentimentsTable
	stmt, errI := DB.Prepare(`INSERT INTO ` + table + ` (api,createdDate) values(?,?)`)
	report.ErrLogger(errI)
	res, _ := stmt.Exec(d.API, time.Now())
	lastID, _ := res.LastInsertId()
	var counter int
	var r bool
	for key, val := range d.Data {
		updateStmt, errU := DB.Prepare(`UPDATE ` + table + ` SET ` + key + `=? WHERE id=?`)
		report.ErrLogger(errU)
		_, err := updateStmt.Exec(val, lastID)
		if err == nil {
			counter++
		}
	}
	switch len(d.Data) {
	case counter:
		r = true
		break
	default:
		r = false
	}
	return &r
}

func IfSentimentsExist() *bool {
	var r bool
	table := cfg.Loader().SentimentsTable
	_, err := DB.Exec(`DESCRIBE ` + table)
	switch err {
	case nil:
		r = true
	default:
		r = false
	}
	return &r
}

func GetCurrentSentiments(f []string) *[]types.SentimentRow {
	var data []types.SentimentRow
	table := cfg.Loader().SentimentsTable
	rows, err := DB.Query(`SELECT id,api,createdDate FROM ` + table + ` ORDER BY id`)
	if err == nil {
		for rows.Next() {
			var dataRow types.SentimentRow
			var id string
			var api string
			var createdDate string
			err := rows.Scan(&id, &api, &createdDate)
			report.ErrLogger(err)
			dataRow.Key = api
			dataRow.Date = createdDate

			switch len(f) {
			case 1:
			case 2:
			case 3:
			case 4:
			case 5:
				row, err := DB.Query(`SELECT `+f[0]+` ,`+f[1]+` ,`+f[2]+` ,`+f[3]+` ,`+f[4]+` FROM `+table+` WHERE api=? && id=?`, api, id)
				if err == nil {
					for row.Next() {
						var f0 string
						var f1 string
						var f2 string
						var f3 string
						var f4 string
						err := row.Scan(&f0, &f1, &f2, &f3, &f4)
						report.ErrLogger(err)
						field := make(map[string]string)
						field[f[0]] = f0
						field[f[1]] = f1
						field[f[2]] = f2
						field[f[3]] = f3
						field[f[4]] = f4
						dataRow.Fields = append(dataRow.Fields, field)
					}
				}
			case 6:
			case 7:
			case 8:
			case 9:
			case 10:

			}
			data = append(data, dataRow)
		}
	}
	return &data
}
