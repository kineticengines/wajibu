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

package configure

import (
	"encoding/json"
	"io"
	"os"
	"strconv"

	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
)

var configFile string = "../configure/app.json"

type ConfigStruct struct {
	Adminport           string
	Publicport          string
	Serverport          string
	Deployed            bool
	DNS                 string
	Driver              string
	AdminTable          string
	MailGunDomain       string
	MailGunAPIKey       string
	MailGunPublic       string
	RedisNetwork        string
	RedisAddress        string
	RedisPoolSize       int
	SlotsTable          string
	TitlesTable         string
	TopLevelTable       string
	HouseLevelTable     string
	SubGovLevelTable    string
	GrassRootLevelTable string
	PillarsTable        string
	SentimentsTable     string
}

func Loader() ConfigStruct {
	var config ConfigStruct
	file, err := os.Open(configFile)
	report.ErrLogger(err)
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.Decode(&config)
	return config
}

func Updater(u *types.ConfigUpdater) bool {
	newDeployed, err := strconv.ParseBool(u.Value)
	report.ErrLogger(err)
	var config ConfigStruct
	oData := Loader()
	config.Adminport = oData.Adminport
	config.AdminTable = oData.AdminTable
	config.Deployed = newDeployed
	config.DNS = oData.DNS
	config.Driver = oData.Driver
	config.GrassRootLevelTable = oData.GrassRootLevelTable
	config.HouseLevelTable = oData.HouseLevelTable
	config.MailGunAPIKey = oData.MailGunAPIKey
	config.MailGunDomain = oData.MailGunDomain
	config.MailGunPublic = oData.MailGunPublic
	config.Publicport = oData.Publicport
	config.RedisAddress = oData.RedisAddress
	config.RedisNetwork = oData.RedisNetwork
	config.RedisPoolSize = oData.RedisPoolSize
	config.Serverport = oData.Serverport
	config.SlotsTable = oData.SlotsTable
	config.SubGovLevelTable = oData.SubGovLevelTable
	config.TitlesTable = oData.TitlesTable
	config.TopLevelTable = oData.TopLevelTable
	config.PillarsTable = oData.PillarsTable
	config.SentimentsTable = oData.SentimentsTable

	jsonFile, err := os.Create(configFile)
	report.ErrLogger(err)
	jsonWriter := io.Writer(jsonFile)
	encoder := json.NewEncoder(jsonWriter)
	err = encoder.Encode(&config)
	if err != nil {
		return false
	}
	return true
}
