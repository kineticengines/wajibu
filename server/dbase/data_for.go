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

package dbase

import (
	"log"

	"github.com/daviddexter/wajibu/handlers/types"
)

func dataForWhichAPI(bio []types.BioData, data []map[string]string) {
	distinct(bio)
}

func distinct(b []types.BioData) {
	//var d []types.BioData
	if len(b) != 0 {
		f := b[0]
		for i := 1; i < len(b); i++ {
			switch f.API {
			case b[i].API:
				j := i + 1
				log.Println(b[j:])
				distinct(b[j:])
				break
			default:
				//d = append(d, f)
				//log.Println(b[1:])
				//distinct(b[1:])
				break
				
			}
		}
	}

}
