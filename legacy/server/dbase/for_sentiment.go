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
	"strconv"
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
			case "image":
				v := reflect.ValueOf(value)
				dataDB.Image = v.Interface().(string)
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
	switch err {
	case nil:
		//alter the table
		for key, _ := range d {
			_, err := DB.Query(`SELECT ` + key + ` FROM ` + table)
			//defer rows.Close()
			if err != nil {
				_, err := DB.Exec(`ALTER TABLE ` + table + ` ADD ` + key + ` TEXT`)
				report.ErrLogger(err)
			}
		}
		break
	default:
		//create the table
		for key, _ := range d {
			//add the keys to a set
			_, err := DB.Exec(`CREATE TABLE ` + table + `(
				id INT NOT NULL AUTO_INCREMENT, api  VARCHAR(255) NOT NULL,
				` + key + ` TEXT, image VARCHAR(255) NOT NULL,createdDate VARCHAR(255) NOT NULL,
				PRIMARY KEY (id))`)
			if err != nil {
				CreateOrAlterTable(d)
			}
		}
		break
	}
}

func AddSentimentToDB(d types.NewSentiment) *bool {
	var r bool
	table := cfg.Loader().SentimentsTable
	db, _ := DB.Begin()
	stmt, errI := db.Prepare(`INSERT INTO ` + table + ` (api,createdDate,image) values(?,?,?)`)
	report.ErrLogger(errI)
	t := time.Now()
	res, _ := stmt.Exec(d.API, t.Format("Jan 2, 2006 at 3:04pm MST"), d.Image)
	lastID, errL := res.LastInsertId()
	if errL != nil {
		db.Rollback()
	}
	db.Commit()
	var counter int
	var notEmpty int
	for key, val := range d.Data {
		switch dd := val.(type) {
		case string:
			if len(dd) != 0 {
				notEmpty++
				db, _ := DB.Begin()
				updateStmt, errU := DB.Prepare(`UPDATE ` + table + ` SET ` + key + `=? WHERE id=?`)
				report.ErrLogger(errU)
				_, err := updateStmt.Exec(dd, lastID)
				if err != nil {
					db.Rollback()
				}
				db.Commit()
				counter++
			}

		}

	}

	switch counter {
	case notEmpty:
		r = true
		break
	default:
		r = false
		break
	}
	return &r
}

func IfSentimentsExist() *bool {
	var r bool
	table := cfg.Loader().SentimentsTable
	_, err := DB.Exec(`DESCRIBE ` + table)
	switch err {
	case nil:
		//check the number of rows
		res := DB.QueryRow(`SELECT COUNT(*) FROM ` + table)
		var n string
		res.Scan(&n)
		switch j, _ := strconv.Atoi(n); j {
		case 0:
			r = false
			break
		default:
			r = true
			break
		}
	default:
		r = false
	}
	return &r
}

func GetCurrentSentiments(f []string) *[]types.SentimentRow {
	var data []types.SentimentRow
	table := cfg.Loader().SentimentsTable
	rows, err := DB.Query(`SELECT id,api,createdDate,image FROM ` + table + ` ORDER BY id DESC`)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var dataRow types.SentimentRow
			var id string
			var api string
			var createdDate string
			var image string
			err := rows.Scan(&id, &api, &createdDate, &image)
			report.ErrLogger(err)
			dataRow.Key = api
			dataRow.Date = createdDate
			dataRow.Image = image
			switch len(f) {
			case 1:
				row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') FROM `+table+` WHERE api=? && id=?`, api, id)
				defer row.Close()
				if err == nil {
					for row.Next() {
						var f0 string
						err := row.Scan(&f0)
						report.ErrLogger(err)
						field := make(map[string]string)
						field[f[0]] = f0
						dataRow.Fields = append(dataRow.Fields, field)
					}
				}
				break
			case 2:
				row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') FROM `+table+` WHERE api=? && id=?`, api, id)
				defer row.Close()
				if err == nil {
					for row.Next() {
						var (
							f0 string
							f1 string
						)
						err := row.Scan(&f0, &f1)
						report.ErrLogger(err)
						field := make(map[string]string)
						field[f[0]] = f0
						field[f[1]] = f1
						dataRow.Fields = append(dataRow.Fields, field)
					}
				}
				break
			case 3:
				row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') FROM `+table+` WHERE api=? && id=?`, api, id)
				defer row.Close()
				if err == nil {
					for row.Next() {
						var (
							f0 string
							f1 string
							f2 string
						)
						err := row.Scan(&f0, &f1, &f2)
						report.ErrLogger(err)
						field := make(map[string]string)
						field[f[0]] = f0
						field[f[1]] = f1
						field[f[2]] = f2
						dataRow.Fields = append(dataRow.Fields, field)
					}
				}
				break
			case 4:
				row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') FROM `+table+` WHERE api=? && id=?`, api, id)
				defer row.Close()
				if err == nil {
					for row.Next() {
						var (
							f0 string
							f1 string
							f2 string
							f3 string
						)
						err := row.Scan(&f0, &f1, &f2, &f3)
						report.ErrLogger(err)
						field := make(map[string]string)
						field[f[0]] = f0
						field[f[1]] = f1
						field[f[2]] = f2
						field[f[3]] = f3
						dataRow.Fields = append(dataRow.Fields, field)
					}
				}
				break
			case 5:
				row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,''),COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') ,COALESCE(`+f[4]+`,'') FROM `+table+` WHERE api=? && id=?`, api, id)
				defer row.Close()
				if err == nil {
					for row.Next() {
						var (
							f0 string
							f1 string
							f2 string
							f3 string
							f4 string
						)
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
				break
			case 6:
				row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') ,COALESCE(`+f[4]+`,'') ,COALESCE(`+f[5]+`,'') FROM `+table+` WHERE api=? && id=?`, api, id)
				defer row.Close()
				if err == nil {
					for row.Next() {
						var (
							f0 string
							f1 string
							f2 string
							f3 string
							f4 string
							f5 string
						)
						err := row.Scan(&f0, &f1, &f2, &f3, &f4, &f5)
						report.ErrLogger(err)
						field := make(map[string]string)
						field[f[0]] = f0
						field[f[1]] = f1
						field[f[2]] = f2
						field[f[3]] = f3
						field[f[4]] = f4
						field[f[5]] = f5
						dataRow.Fields = append(dataRow.Fields, field)
					}
				}
				break
			case 7:
				row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') ,COALESCE(`+f[4]+`,'') ,COALESCE(`+f[5]+`,'') ,COALESCE(`+f[6]+`,'') FROM `+table+` WHERE api=? && id=?`, api, id)
				defer row.Close()
				if err == nil {
					for row.Next() {
						var (
							f0 string
							f1 string
							f2 string
							f3 string
							f4 string
							f5 string
							f6 string
						)
						err := row.Scan(&f0, &f1, &f2, &f3, &f4, &f5, &f6)
						report.ErrLogger(err)
						field := make(map[string]string)
						field[f[0]] = f0
						field[f[1]] = f1
						field[f[2]] = f2
						field[f[3]] = f3
						field[f[4]] = f4
						field[f[5]] = f5
						field[f[6]] = f6
						dataRow.Fields = append(dataRow.Fields, field)
					}
				}
				break
			case 8:
				row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') ,COALESCE(`+f[4]+`,'') ,COALESCE(`+f[5]+`,'') ,COALESCE(`+f[6]+`,'') ,COALESCE(`+f[7]+`,'') FROM `+table+` WHERE api=? && id=?`, api, id)
				defer row.Close()
				if err == nil {
					for row.Next() {
						var (
							f0 string
							f1 string
							f2 string
							f3 string
							f4 string
							f5 string
							f6 string
							f7 string
						)
						err := row.Scan(&f0, &f1, &f2, &f3, &f4, &f5, &f6, &f7)
						report.ErrLogger(err)
						field := make(map[string]string)
						field[f[0]] = f0
						field[f[1]] = f1
						field[f[2]] = f2
						field[f[3]] = f3
						field[f[4]] = f4
						field[f[5]] = f5
						field[f[6]] = f6
						field[f[7]] = f7
						dataRow.Fields = append(dataRow.Fields, field)
					}
				}
				break
			case 9:
				row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') ,COALESCE(`+f[4]+`,'') ,COALESCE(`+f[5]+`,'') ,COALESCE(`+f[6]+`,'') ,COALESCE(`+f[7]+`,'') ,COALESCE(`+f[8]+`,'') FROM `+table+` WHERE api=? && id=?`, api, id)
				defer row.Close()
				if err == nil {
					for row.Next() {
						var (
							f0 string
							f1 string
							f2 string
							f3 string
							f4 string
							f5 string
							f6 string
							f7 string
							f8 string
						)
						err := row.Scan(&f0, &f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8)
						report.ErrLogger(err)
						field := make(map[string]string)
						field[f[0]] = f0
						field[f[1]] = f1
						field[f[2]] = f2
						field[f[3]] = f3
						field[f[4]] = f4
						field[f[5]] = f5
						field[f[6]] = f6
						field[f[7]] = f7
						field[f[8]] = f8
						dataRow.Fields = append(dataRow.Fields, field)
					}
				}
				break
			case 10:
				row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') ,COALESCE(`+f[4]+`,'') ,COALESCE(`+f[5]+`,'') ,COALESCE(`+f[6]+`,'') ,COALESCE(`+f[7]+`,'') ,COALESCE(`+f[8]+`,'') ,COALESCE(`+f[9]+`,'') FROM `+table+` WHERE api=? && id=?`, api, id)
				defer row.Close()
				if err == nil {
					for row.Next() {
						var (
							f0 string
							f1 string
							f2 string
							f3 string
							f4 string
							f5 string
							f6 string
							f7 string
							f8 string
							f9 string
						)
						err := row.Scan(&f0, &f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9)
						report.ErrLogger(err)
						field := make(map[string]string)
						field[f[0]] = f0
						field[f[1]] = f1
						field[f[2]] = f2
						field[f[3]] = f3
						field[f[4]] = f4
						field[f[5]] = f5
						field[f[6]] = f6
						field[f[7]] = f7
						field[f[8]] = f8
						field[f[9]] = f9
						dataRow.Fields = append(dataRow.Fields, field)
					}
				}
				break
			}
			data = append(data, dataRow)
		}
	}
	return &data
}
