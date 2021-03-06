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

package main

import (
	"log"
	"net/http"

	cfg "github.com/daviddexter/wajibu/configure"
)

const (
	PUBLIC = "dist"
)

func main() {
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir(PUBLIC))
	router.Handle("/", fs)
	serve := &http.Server{
		Addr:    cfg.Loader().Adminport,
		Handler: router,
	}
	log.Println("Admin started at port ", cfg.Loader().Adminport)
	serve.ListenAndServe()

}
