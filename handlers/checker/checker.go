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

package checker

import (
	"encoding/json"
	"math/rand"
	"net/http"

	mailgun "gopkg.in/mailgun/mailgun-go.v1"

	"strings"

	"time"

	"github.com/Pallinder/go-randomdata"
	cfg "github.com/daviddexter/wajibu/configure"
	"github.com/daviddexter/wajibu/handlers/types"
	"github.com/daviddexter/wajibu/report"
	"github.com/daviddexter/wajibu/server/dbase"
	"github.com/daviddexter/wajibu/server/deploy"
	"github.com/daviddexter/wajibu/server/radix"
)

func CheckDeployed() bool {
	return *radix.DeployChecker()
}

func CheckInitHandler(w http.ResponseWriter, r *http.Request) {
	if d := CheckDeployed(); d != true {
		res, err := json.Marshal(
			struct {
				Deployed bool
			}{
				Deployed: cfg.Loader().Deployed,
			})
		report.ErrLogger(err)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		return
	}
	res, err := json.Marshal(
		struct {
			Deployed bool
		}{
			Deployed: cfg.Loader().Deployed,
		})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}

func CheckDefaultCred(w http.ResponseWriter, r *http.Request) {
	d := &dbase.PingTable{Table: cfg.Loader().AdminTable}
	if e := d.TableExists(); e != nil {
		res, err := json.Marshal(struct{ Exists bool }{Exists: false})
		report.ErrLogger(err)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		return
	} else {
		res, err := json.Marshal(struct{ Exists bool }{Exists: true})
		report.ErrLogger(err)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		return
	}
}

type CreateDefault struct {
	Email string
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CreateDefaultCred(w http.ResponseWriter, r *http.Request) {
	var createDefault CreateDefault
	err := json.NewDecoder(r.Body).Decode(&createDefault)
	report.ErrLogger(err)
	var toEmail string = createDefault.Email
	var username string = strings.ToLower(randomdata.SillyName())
	var password string = randSeq(10)

	HTMLtoSend :=
		`<html>
		<body>
			<p>Hello,</p>
			<p>Welcome to <b>Wajibu</b></p>
			<p>These are default credentials randomly generated to allow you to accces Wajibu.</p>
			<h4>Username : </h4>` + username +
			` 
			<h4>Password : </h4> ` + password +
			`
			<p><b>NOTE:</b> Please change your credentials to secure your deployment.</p>
			<p>Thank you for promoting good governance.</p>
			<h3> :) </h3>
		</body>
	</html>	
	`

	mg := mailgun.NewMailgun(cfg.Loader().MailGunDomain, cfg.Loader().MailGunAPIKey, cfg.Loader().MailGunPublic)
	msg := mailgun.NewMessage(
		"wajibu@wajibu.com",
		"Administrator Default Credentials",
		"Setting up Wajibu Default Administrator Credentials",
		toEmail)
	msg.SetHtml(HTMLtoSend)
	_, id, err := mg.Send(msg)
	report.ErrLogger(err)
	if len(id) == 0 {
		//Empty response
		res, _ := json.Marshal(struct {
			Sent bool
			Hint string
		}{Sent: false, Hint: "Server error"})
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		return
	} else {
		//Insert into database then send response
		if s := dbase.DefaultToDB(cfg.Loader().AdminTable, username, password, toEmail); s == true {
			//Insert not successfull
			res, _ := json.Marshal(struct {
				Sent bool
				Hint string
			}{Sent: s, Hint: "Default credentials created"})
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			return
		}
	}
}

type LoginCreds struct {
	NameOrEmail string
	Password    string
}

func CheckLoginCredsThenLogin(w http.ResponseWriter, r *http.Request) {
	var loginCreds LoginCreds
	err := json.NewDecoder(r.Body).Decode(&loginCreds)
	report.ErrLogger(err)
	chkRes := dbase.CheckLoginCred(loginCreds.NameOrEmail, loginCreds.Password, cfg.Loader().AdminTable)
	res, err := json.Marshal(struct {
		Cred     string
		Accurate bool
	}{
		Cred:     chkRes.Param,
		Accurate: chkRes.Exist,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

type BuildRes struct {
	Levels []radix.LevelsRes `json:"levels"`
}

func CheckBuildLevel(w http.ResponseWriter, r *http.Request) {
	n := radix.FetchBuildLevel()
	time.Sleep(time.Duration(10) * time.Millisecond)
	i := BuildRes{Levels: n}
	j, _ := json.Marshal(i)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

/******************** Intialize Builds *******************************/
/******************************************************************/

func IntializeBuildThree(w http.ResponseWriter, r *http.Request) { //house level
	h := radix.BuildThreeInitializer()
	res, err := json.Marshal(*h)
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func IntializeBuildFour(w http.ResponseWriter, r *http.Request) { //subgovernment level
	h := radix.BuildFourInitializer()
	res, err := json.Marshal(*h)
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func IntializeBuildFive(w http.ResponseWriter, r *http.Request) { //grassroot level
	h := radix.BuildFiveInitializer()
	res, err := json.Marshal(*h)
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

/*******************************************************************/
/*******************************************************************/

/******************** Process Builds *******************************/
/******************************************************************/
func GoProcessStepOne(w http.ResponseWriter, r *http.Request) {
	var data types.BuildOneAll
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	result := radix.AddCache(radix.BuildOne, radix.BuildOneData, data)
	res, err := json.Marshal(struct{ Status bool }{Status: *result})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func GoProcessStepTwo(w http.ResponseWriter, r *http.Request) {
	var data types.BuildTwoAll
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	result := radix.AddCache(radix.BuildTwo, radix.BuildTwoData, data)
	res, err := json.Marshal(struct{ Status bool }{Status: *result})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
func GoProcessStepThree(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	result := radix.AddCache(radix.BuildThree, radix.BuildThreeData, data)
	res, err := json.Marshal(struct{ Status bool }{Status: *result})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
func GoProcessStepFour(w http.ResponseWriter, r *http.Request) {
	var data []types.BuildFour
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	result := radix.AddCache(radix.BuildFour, radix.BuildFourData, data)
	res, err := json.Marshal(struct{ Status bool }{Status: *result})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
func GoProcessStepFive(w http.ResponseWriter, r *http.Request) {
	var data types.BuildFiveCache
	err := json.NewDecoder(r.Body).Decode(&data)
	report.ErrLogger(err)
	result := radix.AddCache(radix.BuildFive, radix.BuildFiveData, data)
	res, err := json.Marshal(struct{ Status bool }{Status: *result})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

/*******************************************************************/
/*******************************************************************/

/******************** Reset Builds *******************************/
/******************************************************************/
func ResetBuildThree(w http.ResponseWriter, r *http.Request) {
	//var bb bool
	x := radix.ResetBuildThree()
	if *x == true {
		y := radix.ResetBuildFour()
		if *y == true {
			z := radix.ResetBuildFive()
			if *z == true {
				res, err := json.Marshal(struct{ Status bool }{Status: true})
				report.ErrLogger(err)
				w.Header().Set("Content-Type", "application/json")
				w.Write(res)
				return
			}
		}
	}

}

func ResetBuildFour(w http.ResponseWriter, r *http.Request) {
	x := radix.ResetBuildFour()
	if *x == true {
		y := radix.ResetBuildFive()
		if *y == true {
			res, err := json.Marshal(struct{ Status bool }{Status: true})
			report.ErrLogger(err)
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			return
		}
	}

}

func ResetBuildFive(w http.ResponseWriter, r *http.Request) {
	n := radix.ResetBuildFive()
	res, err := json.Marshal(struct{ Status bool }{Status: *n})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

/*******************************************************************/
/*******************************************************************/

/******************** Deploy Process *******************************/
/******************************************************************/
func StartDeployProcess(w http.ResponseWriter, r *http.Request) {
	deploy.Initilizer()
	time.Sleep(time.Duration(500) * time.Millisecond)
	deploy.DeployBuild()
	res, err := json.Marshal(struct{ Status bool }{Status: true})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func CheckDeployProcess(w http.ResponseWriter, r *http.Request) {
	n := deploy.CheckDeploy()
	res, err := json.Marshal(struct{ Status types.DeployStatus }{Status: *n})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func UpdateConfig(w http.ResponseWriter, r *http.Request) {
	n := deploy.Update()
	res, err := json.Marshal(struct{ Status bool }{Status: n})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

/******************************************************************/
/******************************************************************/

/******************** Common Process *******************************/
/******************************************************************/

func FetchTitles(w http.ResponseWriter, r *http.Request) {
	p := *allTitles()
	res, err := json.Marshal(struct {
		Titles []string `json:"titles"`
	}{Titles: p})
	report.ErrLogger(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}

func allTitles() *[]string {
	return radix.AllTitles()
}

/******************************************************************/
/******************************************************************/
