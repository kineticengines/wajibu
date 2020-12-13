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
export class Initializer{
    private headers = new Headers({'Content-Type':'application/json'})
    private options = new RequestOptions({headers:this.headers})

    constructor(private http:Http){}
    extractData(res: Response) {
        let body = res.json();       
        return body || { };    
    }
    
    getDefaultCred(){
        return this.http
                .get(`${ServerURL}/api/check/default`)
                .catch(err => Observable.throw(err))
                .map(this.extractData)                  
    }

    getInit(){
        return this.http
                .get(`${ServerURL}/api/check/init`)
                .catch(err => Observable.throw(err))
                .map(this.extractData)              
    }

    createDefaultCred(obj:Object){
        let OBJ = JSON.stringify(obj)        
        return this.http.post(`${ServerURL}/api/check/create`,OBJ,this.options)
                .catch(err => Observable.throw(err))
                .map(this.extractData)  
                        
    }

    createLogin(obj:Object){
        let OBJ = JSON.stringify(obj)        
        return this.http.post(`${ServerURL}/api/check/proceed/login`,OBJ,this.options)
                .catch(err => Observable.throw(err))
                .map(this.extractData)  
    }
    
    checkBuildLevel(){
        return this.http
                .get(`${ServerURL}/api/check/build/level`)
                .catch(err => Observable.throw(err))
                .map(this.extractData)  
    }

    goProcessStepOne(obj:Object){        
        let OBJ = JSON.stringify(obj)        
        return this.http.post(`${ServerURL}/api/check/build/level/one`,OBJ,this.options)
                .catch(err => Observable.throw(err))
                .map(this.extractData)
    }

    goProcessStepTwo(obj:Object){
       let OBJ = JSON.stringify(obj) 
       return this.http.post(`${ServerURL}/api/check/build/level/two`,OBJ,this.options)
                .catch(err => Observable.throw(err))
                .map(this.extractData) 
    }
    
    buildThreeInitializer(){
       return this.http
                .get(`${ServerURL}/api/check/build/level/three/init`)
                .catch(err => Observable.throw(err))
                .map(this.extractData)  
    }

    resetBuildThree(){
        return this.http
                .get(`${ServerURL}/api/check/build/level/three/reset`)
                .catch(err => Observable.throw(err))
                .map(this.extractData) 
    }

    goProcessStepThree(obj:Object){
        let OBJ = JSON.stringify(obj) 
       return this.http.post(`${ServerURL}/api/check/build/level/three`,OBJ,this.options)
                .catch(err => Observable.throw(err))
                .map(this.extractData) 
    }

    buildFourInitializer(){
       return this.http
                .get(`${ServerURL}/api/check/build/level/four/init`)
                .catch(err => Observable.throw(err))
                .map(this.extractData)  
    }

    resetBuildFour(){
        return this.http
                .get(`${ServerURL}/api/check/build/level/four/reset`)
                .catch(err => Observable.throw(err))
                .map(this.extractData) 
    }

    goProcessStepFour(obj:Object){
        let OBJ = JSON.stringify(obj)        
        return this.http.post(`${ServerURL}/api/check/build/level/four`,OBJ,this.options)
                .catch(err => Observable.throw(err))
                .map(this.extractData)
    }

    buildFiveInitializer(){
        return this.http
                .get(`${ServerURL}/api/check/build/level/five/init`)
                .catch(err => Observable.throw(err))
                .map(this.extractData) 
    }

    resetBuildFive(){
        return this.http
                .get(`${ServerURL}/api/check/build/level/five/reset`)
                .catch(err => Observable.throw(err))
                .map(this.extractData) 
    }
    goProcessStepFive(obj:Object){
        let OBJ = JSON.stringify(obj)        
        return this.http.post(`${ServerURL}/api/check/build/level/five`,OBJ,this.options)
                .catch(err => Observable.throw(err))
                .map(this.extractData)
    }
    
    goStartDeployProcess(){
       return this.http
                .get(`${ServerURL}/api/check/deploy/start`)
                .catch(err => Observable.throw(err))
                .map(this.extractData)  
    }

    goCheckDeployProcess(){
       return this.http
                .get(`${ServerURL}/api/check/deploy/check`)
                .catch(err => Observable.throw(err))
                .map(this.extractData)  
    }

    updateConfig(){
      return this.http
                .get(`${ServerURL}/api/check/deploy/update`)
                .catch(err => Observable.throw(err))
                .map(this.extractData)
                
    }
}
