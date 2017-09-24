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
	"sync"
	tmpl "text/template"

	cfg "github.com/daviddexter/wajibu/configure"
	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
)

func GetSentimentForLocation(item map[string]string, f []string) *types.ContentDataAll {
	var rALL types.ContentDataAll
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var key string
		var value string
		for k, v := range item {
			key = tmpl.JSEscapeString(tmpl.HTMLEscapeString(k))
			value = tmpl.JSEscapeString(tmpl.HTMLEscapeString(v))
		}
		table := cfg.Loader().SentimentsTable
		var content types.ContentData
		content.Title = key
		content.Name = value
		switch len(f) {
		case 1:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') FROM `+table+` WHERE `+key+`=? ORDER BY id DESC`, value)
			defer row.Close()
			if err == nil {
				for row.Next() {
					var f0 string
					err := row.Scan(&f0)
					report.ErrLogger(err)
					field := make(map[string]string)
					field[f[0]] = f0
					content.Data = append(content.Data, field)
				}
			}
		case 2:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,''),COALESCE(`+f[1]+`,'') FROM `+table+` WHERE `+key+`=? ORDER BY id DESC`, value)
			defer row.Close()
			if err == nil {
				for row.Next() {
					var f0 string
					var f1 string
					err := row.Scan(&f0, &f1)
					report.ErrLogger(err)
					field := make(map[string]string)
					field[f[0]] = f0
					field[f[1]] = f1
					content.Data = append(content.Data, field)
				}
			}
		case 3:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,''),COALESCE(`+f[1]+`,''),COALESCE(`+f[2]+`,'') FROM `+table+` WHERE `+key+`=? ORDER BY id DESC`, value)
			defer row.Close()
			if err == nil {
				for row.Next() {
					var f0 string
					var f1 string
					var f2 string
					err := row.Scan(&f0, &f1, &f2)
					report.ErrLogger(err)
					field := make(map[string]string)
					field[f[0]] = f0
					field[f[1]] = f1
					field[f[2]] = f2
					content.Data = append(content.Data, field)
				}
			}
		case 4:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,''),COALESCE(`+f[1]+`,''),COALESCE(`+f[2]+`,''),COALESCE(`+f[3]+`,'') FROM `+table+` WHERE `+key+`=? ORDER BY id DESC`, value)
			defer row.Close()
			if err == nil {
				for row.Next() {
					var f0 string
					var f1 string
					var f2 string
					var f3 string
					err := row.Scan(&f0, &f1, &f2, &f3)
					report.ErrLogger(err)
					field := make(map[string]string)
					field[f[0]] = f0
					field[f[1]] = f1
					field[f[2]] = f2
					field[f[3]] = f3
					content.Data = append(content.Data, field)
				}
			}
		case 5:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,''),COALESCE(`+f[1]+`,''),COALESCE(`+f[2]+`,''),COALESCE(`+f[3]+`,''),COALESCE(`+f[4]+`,'') FROM `+table+` WHERE `+key+`=? ORDER BY id DESC`, value)
			defer row.Close()
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
					content.Data = append(content.Data, field)
				}
			}
		case 6:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,''),COALESCE(`+f[1]+`,''),COALESCE(`+f[2]+`,''),COALESCE(`+f[3]+`,''),COALESCE(`+f[4]+`,''),COALESCE(`+f[5]+`,'') FROM `+table+` WHERE `+key+`=? ORDER BY id DESC`, value)
			defer row.Close()
			if err == nil {
				for row.Next() {
					var f0 string
					var f1 string
					var f2 string
					var f3 string
					var f4 string
					var f5 string
					err := row.Scan(&f0, &f1, &f2, &f3, &f4, &f5)
					report.ErrLogger(err)
					field := make(map[string]string)
					field[f[0]] = f0
					field[f[1]] = f1
					field[f[2]] = f2
					field[f[3]] = f3
					field[f[4]] = f4
					field[f[5]] = f5
					content.Data = append(content.Data, field)
				}
			}
		case 7:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,''),COALESCE(`+f[1]+`,''),COALESCE(`+f[2]+`,''),COALESCE(`+f[3]+`,''),COALESCE(`+f[4]+`,''),COALESCE(`+f[5]+`,''),COALESCE(`+f[6]+`,'') FROM `+table+` WHERE `+key+`=? ORDER BY id DESC`, value)
			defer row.Close()
			if err == nil {
				for row.Next() {
					var f0 string
					var f1 string
					var f2 string
					var f3 string
					var f4 string
					var f5 string
					var f6 string
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
					content.Data = append(content.Data, field)
				}
			}
		case 8:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,''),COALESCE(`+f[1]+`,''),COALESCE(`+f[2]+`,''),COALESCE(`+f[3]+`,''),COALESCE(`+f[4]+`,''),COALESCE(`+f[5]+`,''),COALESCE(`+f[6]+`,''),COALESCE(`+f[7]+`,'') FROM `+table+` WHERE `+key+`=? ORDER BY id DESC`, value)
			defer row.Close()
			if err == nil {
				for row.Next() {
					var f0 string
					var f1 string
					var f2 string
					var f3 string
					var f4 string
					var f5 string
					var f6 string
					var f7 string
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
					content.Data = append(content.Data, field)
				}
			}
		case 9:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,''),COALESCE(`+f[1]+`,''),COALESCE(`+f[2]+`,''),COALESCE(`+f[3]+`,''),COALESCE(`+f[4]+`,''),COALESCE(`+f[5]+`,''),COALESCE(`+f[6]+`,''),COALESCE(`+f[7]+`,''),COALESCE(`+f[8]+`,'') FROM `+table+` WHERE `+key+`=? ORDER BY id DESC`, value)
			defer row.Close()
			if err == nil {
				for row.Next() {
					var f0 string
					var f1 string
					var f2 string
					var f3 string
					var f4 string
					var f5 string
					var f6 string
					var f7 string
					var f8 string
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
					content.Data = append(content.Data, field)
				}
			}
		case 10:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,''),COALESCE(`+f[1]+`,''),COALESCE(`+f[2]+`,''),COALESCE(`+f[3]+`,''),COALESCE(`+f[4]+`,''),COALESCE(`+f[5]+`,''),COALESCE(`+f[6]+`,''),COALESCE(`+f[7]+`,''),COALESCE(`+f[8]+`,''),COALESCE(`+f[9]+`,'') FROM `+table+` WHERE `+key+`=? ORDER BY id DESC`, value)
			defer row.Close()
			if err == nil {
				for row.Next() {
					var f0 string
					var f1 string
					var f2 string
					var f3 string
					var f4 string
					var f5 string
					var f6 string
					var f7 string
					var f8 string
					var f9 string
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
					content.Data = append(content.Data, field)
				}
			}
		}

		rALL.Length = len(content.Data)
		rALL.Content = append(rALL.Content, content)
	}()
	wg.Wait()
	return &rALL
}
