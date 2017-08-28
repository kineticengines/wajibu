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

package radix

import (
	"strconv"

	cfg "github.com/daviddexter/wajibu/configure"
	"github.com/daviddexter/wajibu/report"
	"github.com/fatih/structs"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

var (
	RDB        *pool.Pool
	keyMap     map[string]string
	levels     []string
	dataLevels map[string]string
)

const (
	ConfigKey  string = "config"
	BuildOne   string = "build:one"
	BuildTwo   string = "build:two"
	BuildThree string = "build:three"
	BuildFour  string = "build:four"
	BuildFive  string = "build:five"

	BuildOneData   string = "build:one:data"
	BuildTwoData   string = "build:two:data"
	BuildThreeData string = "build:three:data"
	BuildFourData  string = "build:four:data"
	BuildFiveData  string = "build:five:data"

	PILLARS              string = "pillars"          //key to a set
	TITLES               string = "titles"           //key to a set
	HOUSES               string = "houses"           //key to a set
	SENTIMENT            string = "sentiments:field" //key to a set
	SENTIMENT_KEY_PREFIX string = "sentiment:key"    //key to a set
	SENTIMENT_LIST       string = "sentiment:list"   //key to a list
)

//Common in all builds is a `complete` field which by default of false

func init() {
	keyMap = make(map[string]string)
	keyMap["level"] = "level"
	keyMap["complete"] = "complete"
	keyMap["data"] = "data"

	levels = []string{BuildOne, BuildTwo, BuildThree, BuildFour, BuildFive}

	dataLevels = make(map[string]string)
	dataLevels["one"] = BuildOneData
	dataLevels["two"] = BuildTwoData
	dataLevels["three"] = BuildThreeData
	dataLevels["four"] = BuildFourData
	dataLevels["five"] = BuildFiveData
}

func ConnectToRDB() {
	var err error
	RDB, err = pool.New(cfg.Loader().RedisNetwork, cfg.Loader().RedisAddress, cfg.Loader().RedisPoolSize)
	report.ErrLogger(err)
}

func ConfigDefaulter() {
	data := cfg.Loader()
	//into map[string]interface{}
	mainData := structs.Map(data)
	//add into redis with key config
	err := RDB.Cmd("HMSET", ConfigKey, mainData).Err
	report.ErrLogger(err)
}

func DeployChecker() *bool {
	var b bool
	res, err := RDB.Cmd("EXISTS", ConfigKey).Int()
	report.ErrLogger(err)
	switch res {
	case 1:
		res, err := RDB.Cmd("HGET", ConfigKey, "Deployed").Str()
		report.ErrLogger(err)
		b, _ = strconv.ParseBool(res)

	}
	return &b
}

type LevelsRes struct {
	Level    string      `json:level`
	Complete bool        `json:complete,omitempty`
	Exist    bool        `json:exist`
	Data     interface{} `json:data,omitempty`
}

func FetchBuildLevel() []LevelsRes {
	build := make([]LevelsRes, 0)
	lvRes := make(chan LevelsRes)

	for _, v := range levels {
		go checkBuild(lvRes, v)
	}

	for {
		select {
		case res := <-lvRes:
			build = append(build, res)
			if len(build) == 5 {
				close(lvRes)
				return build
			}
		}
	}

}

func checkBuild(lRes chan<- LevelsRes, level string) {
	//check if build exists and if is complete.
	//Return false if build one does not exist.
	//if exists,check if is complete.
	//check if data exists and retrieve
	data := LevelsRes{}
	data.Level = level
	conn, err := RDB.Get()
	report.ErrLogger(err)
	defer RDB.Put(conn)

	res, _ := conn.Cmd("EXISTS", level).Int()
	switch res {
	case 0:
		//build level does not exist
		data.Exist = false
		lRes <- data
		break
	case 1:
		//build level exists
		data.Exist = true
		ifComplete(conn, level)
		ex, v := ifComplete(conn, level)
		if ex == true {
			data.Complete = v
		} else if ex == false {
			data.Complete = v
		}
		e, d := ifData(conn, level)
		if e == true {
			data.Data = d
		}
		lRes <- data
		break
	}

}

func ifComplete(conn *redis.Client, level string) (bool, bool) {
	res, err := conn.Cmd("HEXISTS", level, keyMap["complete"]).Int()
	report.ErrLogger(err)
	var exist bool
	var value bool
	switch res {
	case 0:
		//complete does not exist
		exist = false
		value = false
		break
	case 1:
		//complete key has been found
		r, err := conn.Cmd("HGET", level, keyMap["complete"]).Str()
		report.ErrLogger(err)
		v, _ := strconv.ParseBool(r)
		exist = true
		value = v
		break
	}
	return exist, value
}

func ifData(conn *redis.Client, level string) (bool, interface{}) {
	res, err := conn.Cmd("HEXISTS", level, keyMap["data"]).Int()
	report.ErrLogger(err)
	var exist bool
	var value interface{}
	switch res {
	case 0:
		exist = false
		value = false
		break
	case 1:
		//data key has been found
		switch level {
		case levels[0]:
			value = getTheData(conn, dataLevels["one"])
		case levels[1]:
			value = getTheData(conn, dataLevels["two"])
		case levels[2]:
			value = getTheData(conn, dataLevels["three"])
		case levels[3]:
			value = getTheData(conn, dataLevels["four"])
		case levels[4]:
			value = getTheData(conn, dataLevels["five"])
		}
		exist = true
		break
	}
	return exist, value
}

func getTheData(conn *redis.Client, dataLevel string) interface{} {
	var dataRes interface{}
	switch dataLevel {
	case dataLevels["one"]:
		res, err := conn.Cmd("HGETALL", dataLevel).Map()
		report.ErrLogger(err)
		dataRes = populateBuildOne(res, dataLevel)
	case dataLevels["two"]:
		dataRes = populateBuildTwo(dataLevel)
	case dataLevels["three"]:
		dataRes = populateBuildThree(dataLevel)
	case dataLevels["four"]:
		dataRes = populateBuildFour(dataLevel)
	case dataLevels["five"]:
		dataRes = populateBuildFive()
	}
	return dataRes
}

func HousesCompleteness() bool {
	//get build:three:house keys
	k := BuildThree + ":house:*"
	houseskeys, err := RDB.Cmd("KEYS", k).List()
	report.ErrLogger(err)
	x := make([]bool, 0)
	for _, v := range houseskeys {
		//check if the key is a set
		t, e := RDB.Cmd("TYPE", v).Str()
		report.ErrLogger(e)
		switch t {
		case "set":
			//get the members and test if members exist
			memberKeys, err := RDB.Cmd("SMEMBERS", v).List()
			report.ErrLogger(err)
			//test existence
			for _, memberKey := range memberKeys {
				exist, err := RDB.Cmd("EXISTS", memberKey).Int()
				report.ErrLogger(err)
				switch exist {
				case 1:
					//hash key exists
					//append true
					x = append(x, true)
				case 0:
					//hash key does not exists
					//append false
					x = append(x, false)

				}
			}

		}

	}
	size := len(x)
	t := 0
	for _, v := range x {
		if v == true {
			t++
		}
	}
	if size == t {
		return true
	} else {
		return false
	}
}
