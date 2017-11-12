/*
Wajibu is an online web app that collects,analyses and aggregates sentiments from the public
pertaining the government of a nation. This tool allows citizens to contribute to the
governance talk by airing out their honest views about the state of the nation and in
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

package clients

import (
	"encoding/json"
	"net/http"
	tmpl "text/template"

	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
	"github.com/daviddexter/wajibu/server/radix"
)

func FetchCurrentSentiments(w http.ResponseWriter, r *http.Request) {
	n := radix.GetCachedSentiments()
	res, err := json.Marshal(struct {
		All []map[string]string `json:"all"`
	}{All: *n})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func FilterByQuery(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Item string
		Type string
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)

	n, fc := determineTheQueryType(tmpl.JSEscapeString(tmpl.HTMLEscapeString(data.Item)))

	if *fc == true { // true means no match was found.Status should be false
		res, err := json.Marshal(struct {
			Status bool `json:"status"`
		}{Status: !*fc})
		report.ErrLogger(err)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		return
	}
	m := querySwitcher(*n, data.Type)
	res, err := json.Marshal(struct {
		Status bool                 `json:"status"`
		Data   types.ContentDataAll `json:"content"`
	}{Status: true, Data: *m})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func CacheTheQuery(w http.ResponseWriter, r *http.Request) {
	var data struct{ Item string }
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	n := *radix.CacheQuery(tmpl.JSEscapeString(tmpl.HTMLEscapeString(data.Item)))
	res, err := json.Marshal(struct {
		Status bool `json:"status"`
	}{Status: n})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func GetCachedQuery(w http.ResponseWriter, r *http.Request) {
	n := *radix.GetCachedQueryItems()
	res, err := json.Marshal(struct {
		Items []string `json:"items"`
	}{Items: n})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
