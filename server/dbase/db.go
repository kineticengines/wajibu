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

	cfg "github.com/daviddexter/wajibu/configure"
	"github.com/daviddexter/wajibu/report"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectToDB() {
	var err error
	DB, err = sql.Open(cfg.Loader().Driver, cfg.Loader().DNS)
	report.ErrLogger(err)
}

type PingTable struct {
	Table string
}

func (p *PingTable) TableExists() error {
	var r error
	_, err := DB.Exec(`DESCRIBE ` + cfg.Loader().AdminTable)
	switch err {
	case nil:
		//check if table has data
		_, err := DB.Query(`SELECT * FROM ` + cfg.Loader().AdminTable)
		if err != nil {
			r = err
		}
	default:
		r = err
	}
	return r
}
