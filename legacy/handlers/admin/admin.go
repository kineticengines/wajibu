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

	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
	"github.com/daviddexter/wajibu/server/dbase"
)

func FetchPillars(w http.ResponseWriter, r *http.Request) {
	p, errA := allPillars()
	var marshalRes []byte

	if errA != nil {
		marshalRes, _ = json.Marshal(struct {
			Pillars []types.Pillar `json:"pillars"`
			Error   string         `json:"error"`
		}{Error: errA.Error()})
	} else {
		marshalRes, _ = json.Marshal(struct {
			Pillars []types.Pillar `json:"pillars"`
			Error   string         `json:"error"`
		}{Pillars: *p})
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalRes)
	return
}

type NewPillarData struct {
	Pillar string
	For    string
}

func AddNewPillar(w http.ResponseWriter, r *http.Request) {
	var data NewPillarData
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	n := dbase.NewPillar(data.Pillar, data.For)
	res, err := json.Marshal(struct {
		Status bool `json:"status"`
	}{Status: *n})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return

}

func RemovePillar(w http.ResponseWriter, r *http.Request) {
	var data NewPillarData
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	n := dbase.RemovePillar(data.Pillar)
	res, err := json.Marshal(struct {
		Status bool `json:"status"`
	}{Status: *n})

	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func PresidentLevelCongfigure(w http.ResponseWriter, r *http.Request) {
	d := presidentLevelConfigurer()
	res, err := json.Marshal(struct {
		Config types.ConfigAll `json:"config"`
	}{Config: *d})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func DPresidentLevelCongfigure(w http.ResponseWriter, r *http.Request) {
	d := dpresidentLevelConfigurer()
	res, err := json.Marshal(struct {
		Config types.ConfigAll `json:"config"`
	}{Config: *d})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func GetHouseRepSlots(w http.ResponseWriter, r *http.Request) {
	var data struct{ HouseName string }
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	d := getHouseLevelRepSlots(data)
	res, err := json.Marshal(d)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func HouseLevelConfigure(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Designation string
		Type        string
		Data        struct {
			SlotName string
		}
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	d := houseLevelConfigurer(data)
	res, err := json.Marshal(struct {
		Config types.ConfigAll `json:"config"`
	}{Config: *d})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func SubGovLevelConfigure(w http.ResponseWriter, r *http.Request) {
	var data struct{ GovName string }

	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	d := subGovLevelConfigurer(data)
	res, err := json.Marshal(struct {
		Config types.ConfigAll `json:"config"`
	}{Config: *d})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func GetRootLevelReps(w http.ResponseWriter, r *http.Request) {
	var data struct{ GovName string }
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	d := getRootLevelReps(data)
	res, err := json.Marshal(d)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func RootLevelConfigure(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Designation string
		Type        string
		Data        struct {
			SlotName string
		}
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	d := rootLevelConfigurer(data)
	res, err := json.Marshal(struct {
		Config types.ConfigAll `json:"config"`
	}{Config: *d})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
