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

package admin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/daviddexter/wajibu/report"
	"github.com/daviddexter/wajibu/server/dbase"
	"github.com/daviddexter/wajibu/server/radix"
)

func AddSentiment(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	d := *dbase.AddNewSentiment(data)
	dbase.CreateOrAlterTable(d.Data)
	radix.AddSentimentsTableFields(d.Data)
	time.Sleep(time.Duration(100 * time.Microsecond))
	n := dbase.AddSentimentToDB(d)
	res, err := json.Marshal(struct {
		Status bool `json:"status"`
	}{Status: *n})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
