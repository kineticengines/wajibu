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

func GetSentimentForPillar(item string, v []string) *types.ContentDataAll {
	var rALL types.ContentDataAll
	n := *sentimentsFetcherFromFilterPillar(v, tmpl.JSEscapeString(tmpl.HTMLEscapeString(item)))

	switch len(n.Data) {
	case 0:
		rALL.Length = 0
	default:
		rALL.Length = len(n.Data)
		var bios []types.BioData
		for _, val := range n.Data {
			api := val["api"]
			var wg sync.WaitGroup
			wg.Add(4)
			var mutex sync.Mutex
			var r types.BioData
			go func() {
				defer wg.Done()
				//get bio data for top level
				table := cfg.Loader().TopLevelTable
				rows, err := DB.Query(`SELECT  name,position FROM `+table+` WHERE api=?`, api)
				if err == nil {
					for rows.Next() {
						var n string
						var p string
						err := rows.Scan(&n, &p)
						if err == nil {
							mutex.Lock()
							r.Name = n
							r.Position = p
							r.API = api
							mutex.Unlock()
						}
					}
				}
			}()
			go func() {
				defer wg.Done()
				//get bio data for house level
				table := cfg.Loader().HouseLevelTable
				rows, err := DB.Query(`SELECT  name,title FROM `+table+` WHERE api=?`, api)
				if err == nil {
					for rows.Next() {
						var n string
						var p string
						err := rows.Scan(&n, &p)
						if err == nil {
							mutex.Lock()
							r.Name = n
							r.Position = p
							r.API = api
							mutex.Unlock()
						}
					}
				}
			}()
			go func() {
				defer wg.Done()
				//get bio data for subgov level
				table := cfg.Loader().SubGovLevelTable
				rows, err := DB.Query(`SELECT  name,position FROM `+table+` WHERE api=?`, api)
				if err == nil {
					for rows.Next() {
						var n string
						var p string
						err := rows.Scan(&n, &p)
						if err == nil {
							mutex.Lock()
							r.Name = n
							r.Position = p
							r.API = api
							mutex.Unlock()
						}
					}
				}
			}()
			go func() {
				defer wg.Done()
				//get bio data for grass level
				table := cfg.Loader().GrassRootLevelTable
				rows, err := DB.Query(`SELECT  name,title FROM `+table+` WHERE api=?`, api)
				if err == nil {
					for rows.Next() {
						var n string
						var p string
						err := rows.Scan(&n, &p)
						if err == nil {
							mutex.Lock()
							r.Name = n
							r.Position = p
							r.API = api
							mutex.Unlock()
						}
					}
				}
			}()
			wg.Wait()
			bios = append(bios, r)
			//var newContent types.ContentData
			//newContent.Name = r.Name
			//newContent.Title = r.Position
			//newContent.Data = n.Data
			//dataForWhichAPI(r, n.Data)

			//rALL.Content = append(rALL.Content, newContent)
		}
		dataForWhichAPI(bios, n.Data)
	}
	return &rALL
}

func sentimentsFetcherFromFilterPillar(f []string, item string) *types.ContentData {
	var content types.ContentData
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		table := cfg.Loader().SentimentsTable
		switch len(f) {
		case 1:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') FROM `+table+` WHERE pillars=? ORDER BY id DESC`, item)
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
			break
		case 2:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') FROM `+table+` WHERE pillars=? ORDER BY id DESC`, item)
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
					content.Data = append(content.Data, field)
				}
			}
			break
		case 3:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') FROM `+table+` WHERE pillars=? ORDER BY id DESC`, item)
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
					content.Data = append(content.Data, field)
				}
			}
			break
		case 4:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') FROM `+table+` WHERE pillars=? ORDER BY id DESC`, item)
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
					content.Data = append(content.Data, field)
				}
			}
			break
		case 5:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,''),COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') ,COALESCE(`+f[4]+`,'') FROM `+table+` WHERE pillars=? ORDER BY id DESC`, item)
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
					content.Data = append(content.Data, field)
				}
			}
			break
		case 6:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') ,COALESCE(`+f[4]+`,'') ,COALESCE(`+f[5]+`,'') FROM `+table+` WHERE pillars=? ORDER BY id DESC`, item)
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
					content.Data = append(content.Data, field)
				}
			}
			break
		case 7:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') ,COALESCE(`+f[4]+`,'') ,COALESCE(`+f[5]+`,'') ,COALESCE(`+f[6]+`,'') FROM `+table+` WHERE pillars=? ORDER BY id DESC`, item)
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
					content.Data = append(content.Data, field)
				}
			}
			break
		case 8:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') ,COALESCE(`+f[4]+`,'') ,COALESCE(`+f[5]+`,'') ,COALESCE(`+f[6]+`,'') ,COALESCE(`+f[7]+`,'') FROM `+table+` WHERE pillars=? ORDER BY id DESC`, item)
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
					content.Data = append(content.Data, field)
				}
			}
			break
		case 9:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') ,COALESCE(`+f[4]+`,'') ,COALESCE(`+f[5]+`,'') ,COALESCE(`+f[6]+`,'') ,COALESCE(`+f[7]+`,'') ,COALESCE(`+f[8]+`,'') FROM `+table+` WHERE pillars=? ORDER BY id DESC`, item)
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
					content.Data = append(content.Data, field)
				}
			}
			break
		case 10:
			row, err := DB.Query(`SELECT COALESCE(`+f[0]+`,'') ,COALESCE(`+f[1]+`,'') ,COALESCE(`+f[2]+`,'') ,COALESCE(`+f[3]+`,'') ,COALESCE(`+f[4]+`,'') ,COALESCE(`+f[5]+`,'') ,COALESCE(`+f[6]+`,'') ,COALESCE(`+f[7]+`,'') ,COALESCE(`+f[8]+`,'') ,COALESCE(`+f[9]+`,'') FROM `+table+` WHERE pillars=? ORDER BY id DESC`, item)
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
					content.Data = append(content.Data, field)
				}
			}
			break
		}
	}()
	wg.Wait()
	return &content
}
