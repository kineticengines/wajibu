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


import { Injectable } from "@angular/core"
import { Http,Response,Headers,RequestOptions } from "@angular/http"


import { ServerURL } from "../app.routing"

import { Observable } from "rxjs/Observable";
import "rxjs/add/operator/map";
import "rxjs/add/operator/catch";

@Injectable()
export class DashService{
    private headers = new Headers({'Content-Type':'application/json'})
    private options = new RequestOptions({headers:this.headers})

    constructor(private http:Http){}
    private extractData(res: Response) {
        let body = res.json();       
        return body || { };    
    }

    presidentLevelConfigure(){
        return this.http
                .get(`${ServerURL}/api/dash/president/level/configure`)
                .catch(err => Observable.throw(err))
                .map(this.extractData) 
    }

    dpresidentLevelConfigure(){
        return this.http
                .get(`${ServerURL}/api/dash/dpresident/level/configure`)
                .catch(err => Observable.throw(err))
                .map(this.extractData) 
    }   
    houseLevelConfigure(){
        return this.http
                .get(`${ServerURL}/api/dash/house/level/configure`)
                .catch(err => Observable.throw(err))
                .map(this.extractData)
    }

    fetchTitles(){
        return this.http
                .get(`${ServerURL}/api/dash/fetch/titles`)
                .catch(err => Observable.throw(err))
                .map(this.extractData)  
    }

    fetchPillars(){
       return this.http
                .get(`${ServerURL}/api/dash/fetch/pillars`)
                .catch(err => Observable.throw(err))
                .map(this.extractData)  
    }

    fetchHouses(){        
        return this.http
                .get(`${ServerURL}/api/dash/fetch/houses`)
                .map(this.extractData) 
    }

    savePillar(obj:Object){
        let OBJ = JSON.stringify(obj)        
        return this.http.post(`${ServerURL}/api/dash/save/pillar`,OBJ,this.options)
                .catch(err => Observable.throw(err))
                .map(this.extractData)  
                        
    }

    removePillar(obj:Object){
        let OBJ = JSON.stringify(obj)        
        return this.http.post(`${ServerURL}/api/dash/remove/pillar`,OBJ,this.options)
                .catch(err => Observable.throw(err))
                .map(this.extractData)  
                        
    }

    addSentiment(obj:Object){
       let OBJ = JSON.stringify(obj)        
       return this.http.post(`${ServerURL}/api/dash/sentiment/add`,OBJ,this.options)
                .catch(err => Observable.throw(err))
                .map(this.extractData)  
    }
    
    getRepSlotsForHouse(obj:Object){
       let OBJ = JSON.stringify(obj)        
       return this.http.post(`${ServerURL}/api/dash/sentiment/get/house/reps`,OBJ,this.options)
                .catch(err => Observable.throw(err))
                .map(this.extractData)
    }

    configureHouseLevel(obj:Object){
        let OBJ = JSON.stringify(obj)        
        return this.http.post(`${ServerURL}/api/dash/house/level/configure`,OBJ,this.options)
                .catch(err => Observable.throw(err))
                .map(this.extractData)
    }
}