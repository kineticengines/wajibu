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
import { Observable} from "rxjs/Rx";
import "rxjs/Rx";


@Injectable()
export class Initializer{  
     private headers = new Headers({'Content-Type':'application/json'})
    private options = new RequestOptions({headers:this.headers})

    constructor(private http:Http){}

    
    getInit(){
        return this.http
                .get(`${ServerURL}/api/check/init`)
                .map(this.extractData)                
    }
    private extractData(res: Response) {
        let body = res.json();
        return body || { };    
    }
    checker():Status{        
        this.getInit().subscribe(d => {                 
            window.localStorage.clear()
            window.localStorage.setItem("get",d.Deployed)
        })
        
        let dd = new Status()
        
        switch (this.getter()) {
            case "true":              
                dd.IsDeployed = true;   
                break;     
            case "false":                
               dd.IsDeployed = false 
               break;
        }      
        
        return dd
    }  
    
    getter():string{        
        return window.localStorage.getItem("get")
    }

    getTitles(){        
        return Observable.interval(2000)
                        .startWith(0) 
                        .switchMap(() => 
                            this.http
                                .get(`${ServerURL}/api/init/get/titles`)
                                .map(this.extractData)  
                    )         
    }   
            
    getPillars(){
       return Observable.interval(2000)
                        .startWith(0) 
                        .switchMap(() => 
                        this.http
                            .get(`${ServerURL}/api/dash/fetch/pillars`)
                            .map(this.extractData) 
                    )
        
    }

    getSentiments(){       
        return Observable.interval(1000)
                         .startWith(0)
                         .switchMap(() => 
                            this.http
                                .get(`${ServerURL}/api/init/get/seintiments`)            
                                .map((res: Response) => {
                                    let body = res.json();
                                    return body.all.reverse()
                                })
                                ._catch(e => Observable.of(e))                                  
                         )                  
                         
                                          
    }

    getFilterContentData(obj:Object){
        let OBJ = JSON.stringify(obj) 
        return this.http.post(`${ServerURL}/api/init/filter/query`,OBJ,this.options)
                 .map(this.extractData)                                
                 ._catch(e => Observable.of(e))            
    }

    cacheTheQuery(obj:Object){
        let OBJ = JSON.stringify(obj) 
        return this.http.post(`${ServerURL}/api/init/filter/cache/query`,OBJ,this.options)
                 .map(this.extractData)                                
                 ._catch(e => Observable.of(e))
    }

    getCachedItems(){
        return this.http
                .get(`${ServerURL}/api/init/filter/get/cache/query`)
                .map(this.extractData) 
                .map(val => {
                    return val.items
                })
                
    }
    
}

class Status{
    IsDeployed:boolean
}