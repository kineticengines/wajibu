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

package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	cfg "github.com/daviddexter/wajibu/configure"
	"github.com/daviddexter/wajibu/server/bg"
	"github.com/daviddexter/wajibu/server/dbase"
	"github.com/daviddexter/wajibu/server/radix"
	"github.com/daviddexter/wajibu/server/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var licenseNote string = `Wajibu  Copyright (C) 2017  David 'Dexter' Mwangi.
This program comes with ABSOLUTELY NO WARRANTY;This is free software, and you are welcome to redistribute it under certain conditions.`

func init() {
	dbase.ConnectToDB()
	radix.ConnectToRDB()
	radix.ConfigDefaulter()
}

func main() {
	router := mux.NewRouter()
	for _, v := range routes.Routes {
		router.HandleFunc(v.Path, v.Handler).Methods(v.Method)
	}
	serve := &http.Server{
		Addr:    cfg.Loader().Serverport,
		Handler: cors.Default().Handler(router),
	}

	fmt.Println(licenseNote)
	go bg.StartBG()
	serve.ListenAndServe()
}
