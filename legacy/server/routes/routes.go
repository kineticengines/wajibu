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

package routes

import (
	"net/http"

	"github.com/daviddexter/wajibu/handlers/admin"
	chck "github.com/daviddexter/wajibu/handlers/checker"
	"github.com/daviddexter/wajibu/handlers/clients"
)

type routes []route

type route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

var Routes = routes{
	route{
		Path:    "/api/check/init",
		Method:  "GET",
		Handler: chck.CheckInitHandler,
	},
	route{
		Path:    "/api/check/default",
		Method:  "GET",
		Handler: chck.CheckDefaultCred,
	},
	route{
		Path:    "/api/check/create",
		Method:  "POST",
		Handler: chck.CreateDefaultCred,
	},
	route{
		Path:    "/api/check/proceed/login",
		Method:  "POST",
		Handler: chck.CheckLoginCredsThenLogin,
	},
	route{
		Path:    "/api/check/build/level",
		Method:  "GET",
		Handler: chck.CheckBuildLevel,
	},
	route{
		Path:    "/api/check/build/level/one",
		Method:  "POST",
		Handler: chck.GoProcessStepOne,
	},
	route{
		Path:    "/api/check/build/level/two",
		Method:  "POST",
		Handler: chck.GoProcessStepTwo,
	},
	route{
		Path:    "/api/check/build/level/three/init",
		Method:  "GET",
		Handler: chck.IntializeBuildThree,
	},
	route{
		Path:    "/api/check/build/level/three",
		Method:  "POST",
		Handler: chck.GoProcessStepThree,
	},
	route{
		Path:    "/api/check/build/level/three/reset",
		Method:  "GET",
		Handler: chck.ResetBuildThree,
	},
	route{
		Path:    "/api/check/build/level/four/init",
		Method:  "GET",
		Handler: chck.IntializeBuildFour,
	},
	route{
		Path:    "/api/check/build/level/four/reset",
		Method:  "GET",
		Handler: chck.ResetBuildFour,
	},
	route{
		Path:    "/api/check/build/level/four",
		Method:  "POST",
		Handler: chck.GoProcessStepFour,
	},
	route{
		Path:    "/api/check/build/level/five/init",
		Method:  "GET",
		Handler: chck.IntializeBuildFive,
	},
	route{
		Path:    "/api/check/build/level/five/reset",
		Method:  "GET",
		Handler: chck.ResetBuildFive,
	},
	route{
		Path:    "/api/check/build/level/five",
		Method:  "POST",
		Handler: chck.GoProcessStepFive,
	},
	route{
		Path:    "/api/check/deploy/start",
		Method:  "GET",
		Handler: chck.StartDeployProcess,
	},
	route{
		Path:    "/api/check/deploy/check",
		Method:  "GET",
		Handler: chck.CheckDeployProcess,
	},
	route{
		Path:    "/api/check/deploy/update",
		Method:  "GET",
		Handler: chck.UpdateConfig,
	},
	route{
		Path:    "/api/dash/fetch/titles",
		Method:  "GET",
		Handler: chck.FetchTitles,
	},
	route{
		Path:    "/api/dash/fetch/houses",
		Method:  "GET",
		Handler: chck.FetchHouses,
	},
	route{
		Path:    "/api/dash/fetch/subgovs",
		Method:  "GET",
		Handler: chck.FetchSubGovs,
	},
	route{
		Path:    "/api/dash/fetch/pillars",
		Method:  "GET",
		Handler: admin.FetchPillars,
	},
	route{
		Path:    "/api/dash/save/pillar",
		Method:  "POST",
		Handler: admin.AddNewPillar,
	},
	route{
		Path:    "/api/dash/remove/pillar",
		Method:  "POST",
		Handler: admin.RemovePillar,
	},
	route{
		Path:    "/api/dash/president/level/configure",
		Method:  "GET",
		Handler: admin.PresidentLevelCongfigure,
	},
	route{
		Path:    "/api/dash/dpresident/level/configure",
		Method:  "GET",
		Handler: admin.DPresidentLevelCongfigure,
	},
	route{
		Path:    "/api/init/get/titles",
		Method:  "GET",
		Handler: chck.FetchTitles,
	},
	route{
		Path:    "/api/dash/sentiment/add",
		Method:  "POST",
		Handler: admin.AddSentiment,
	},
	route{
		Path:    "/api/init/get/seintiments",
		Method:  "GET",
		Handler: clients.FetchCurrentSentiments,
	},
	route{
		Path:    "/api/init/filter/query",
		Method:  "POST",
		Handler: clients.FilterByQuery,
	},
	route{
		Path:    "/api/init/filter/cache/query",
		Method:  "POST",
		Handler: clients.CacheTheQuery,
	},
	route{
		Path:    "/api/init/filter/get/cache/query",
		Method:  "GET",
		Handler: clients.GetCachedQuery,
	},
	route{
		Path:    "/api/dash/sentiment/get/house/reps",
		Method:  "POST",
		Handler: admin.GetHouseRepSlots,
	},
	route{
		Path:    "/api/dash/house/level/configure",
		Method:  "POST",
		Handler: admin.HouseLevelConfigure,
	},
	route{
		Path:    "/api/dash/check/if/central",
		Method:  "GET",
		Handler: chck.CheckIfCentral,
	},
	route{
		Path:    "/api/dash/subgov/level/configure",
		Method:  "POST",
		Handler: admin.SubGovLevelConfigure,
	},
	route{
		Path:    "/api/dash/sentiment/get/root/reps",
		Method:  "POST",
		Handler: admin.GetRootLevelReps,
	},
	route{
		Path:    "/api/dash/root/level/configure",
		Method:  "POST",
		Handler: admin.RootLevelConfigure,
	},
}
