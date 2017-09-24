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

package types

type BuildOneMain struct {
	Deployname         string `json:"deployname"`
	Deploycountry      string `json:"deploycountry"`
	Deployspan         string `json:"deployspan"`
	Governmenttype     string `json:"governmenttype"`
	Numofhouses        string `json:"numofhouses"`
	Subgovname         string `json:"subgovname"`
	Subgovtitle        string `json:"subgovtitle"`
	Numofsubgov        string `json:"numofsubgov"`
	Subgovhasleg       bool   `json:"subgovhasleg"`
	Subgovhousename    string `json:"subgovhousename"`
	Subgovhouserepslot string `json:"subgovhouserepslot"`
	Subgovreptitle     string `json:"subgovreptitle"`
}

type BuidlOneHDetails struct {
	Housename  string `json:"housename"`
	Repslot    string `json:"repslot"`
	Reptitle   string `json:"reptitle"`
	Numofseats string `json:"numofseats"`
}

type BuildOneAll struct {
	Main     BuildOneMain       `json:"main"`
	HDetails []BuidlOneHDetails `json:"hdetails"`
}

type BuildTwoAll struct {
	PresidentName  string `json:"presidentName"`
	PresidentTerm  string `json:"presidentTerm"`
	Pgender        string `json:"pgender"`
	PnthPosition   string `json:"pnthPosition"`
	PImage         string `json:"pImage"`
	DpresidentName string `json:"dpresidentName"`
	DpresidentTerm string `json:"dpresidentTerm"`
	Dgender        string `json:"dgender"`
	DnthPosition   string `json:"dnthPosition"`
	DImage         string `json:"dImage"`
}

type BuildThree struct {
	RepName   string `json:"repName"`
	RepTerm   string `json:"repTerm"`
	RepGender string `json:"repGender"`
	RepSlot   string `json:"repSlot"`
	RepImage  string `json:"repImage"`
	HouseName string `json:"houseName"`
	SlotName  string `json:"slotName"`
}

type BuildThreeAll struct {
	Housesdata []BuildThree `json:"housesdata"`
}

type BuildFour struct {
	HeadName        string `json:"headname"`
	HeadTerm        string `json:"headterm"`
	HeadGender      string `json:"headgender"`
	HeadnthPosition string `json:"headnthposition"`
	HeadImage       string `json:"headimage"`
	SlotName        string `json:"slotname"`      //name of county or state
	NumOfLegSeats   string `json:"numoflegseats"` //number of seats in county assembly
}

type BuildFourAll struct {
	IsCentral   bool        `json:"iscentral,omitempty"`
	HasLeg      bool        `json:"hasleg,omitempty"`
	Numofsubgov string      `json:"numofsubgov,omitempty"`
	Complete    bool        `json:"complete,omitempty"`
	OfficeTitle string      `json:"officetitle,omitempty"` // governor or whatever
	Subgovname  string      `json:"subgovname,omitempty"`  //whether county or state or region or province
	SubgovData  []BuildFour `json:"subgovdata,omitempty"`
}

type BuildFive struct {
	RepName        string `json:"repname"`
	RepTerm        string `json:"repterm"`
	RepGender      string `json:"repgender"`
	RepnthPosition string `json:"repnthposition"`
	RepImage       string `json:"repimage"`
	SlotName       string `json:"slotname"`
}

type BuildFiveCache struct {
	Key     string
	TheData []BuildFive
}

type BuildFiveRep struct {
	RepData    BuildFive `json:"repdata"`
	Block      string    `json:"block"` //county,state or province or region
	BlockTitle string    `json:"blocktitle"`
	BlockSlot  string    `json:"blockslot"`
	From       string    `json:"from"`
}

type BuildFiveAll struct {
	RepsData []BuildFiveRep `json:"repsdata"`
}

type Slot struct {
	SlotName    string `json:"slotname"`
	Designation string `json:"designation"`
}

type TopPosition struct { //be used by top level and subgovernment
	Name            string
	Position        string //president,vp or governor
	Term            string
	Gender          string
	NthPosition     string
	SlotDesignation string //county or state or province
	SlotName        string //name
	Image           string
	API             string
}

type HousePosition struct { // be used for house level and subgov leg house
	HouseName       string //senate,national assembly or county assembly
	Name            string
	Title           string //mp,senator or mca
	Term            string
	Gender          string
	NthPosition     string
	SlotDesignation string //county or constituency or ward
	LegOf           string // county
	SlotName        string //name
	Image           string
	API             string
}

type ConfigUpdater struct {
	Path  string
	Value string
}

type DeployStatus struct {
	Percent  int  `json:"percent"`
	Complete bool `json:"complete"`
}

type Pillar struct {
	Pillar   string `json:"pillar"`
	Fortitle string `json:"fortitle"`
}
type PillarData struct {
	Error   error
	Pillars []Pillar
}

type TitleData struct {
	Error error
	Title []string
}

type HouseData struct {
	Error error
	House []string
}

type FormConfig struct {
	Type        string        `json:"type"`
	Name        string        `json:"name"`
	Label       string        `json:"label"`
	Options     []interface{} `json:"options,omitempty"`
	Placeholder string        `json:"placeholder"`
	Value       string        `json:"value"`
	//Disabled    bool          `json:"disabled"`
	//Validation  func()        `json:"validation"`
}

type ConfigAll struct {
	Config []FormConfig `json:"config"`
}

type NewSentiment struct {
	API   string
	Image string
	Data  map[string]interface{}
}

type SentimentRow struct {
	Key    string
	Date   string
	Image  string
	Fields []map[string]string
}

type GetHouseSlotsData struct {
	Designation string   `json:"designation"`
	Slots       []string `json:"slots"`
}

type QueryType struct {
	Type   string
	IsTrue bool
	Value  string
}

type LevelType struct {
	IsTrue bool
	API    string
	Level  string
}

type BioData struct {
	Name     string
	Position string
	API      string
}

type ContentData struct {
	Title string              `json:"title"`
	Name  string              `json:"name"`
	Data  []map[string]string `json:"data"`
}

type ContentDataAll struct {
	Length  int           `json:"length"`
	Content []ContentData `json:"content"`
}
