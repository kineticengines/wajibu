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


import { Injectable } from "@angular/core";
import {Initializer} from "./init.service";

@Injectable()
export class SessionChecker{
     

   constructor(private init:Initializer){
       this.deployChecker()
   }    

   createSession(p:string){
      window.localStorage.setItem("session",p)
   }  

   destroySession(){
      //window.localStorage.clear() 
      window.localStorage.removeItem("session")     
   }
   
   deployChecker(){
    this.init.getInit().subscribe(d =>{
        window.localStorage.setItem("deployed",d.Deployed)
    })
   }

   setDeployed(d:any){
       window.localStorage.removeItem("deployed")
       window.localStorage.setItem("deployed",d)
   }

   getSession():Session{
       let d = new Session() 
       let s = window.localStorage.getItem("session") 
       let dp =  window.localStorage.getItem("deployed")      
       switch (s && dp) {
           case null:
               d.isSet = false
               break;  
           case undefined:
                d.isSet = false
                break;     
           default:
                d.isSet = true
               break;
       }     

       return d
   }  

    getDeploy():Status{
       let d = new Status() 
       let s = window.localStorage.getItem("deployed")       
       switch (s) {
           case "true":
               d.isDeployed = true
               break;       
           default:
                d.isDeployed = false
               break;
       }

       return d
       
    }

    getNotDeployed():Status{
       let d = new Status() 
       let s = window.localStorage.getItem("deployed")       
       switch (s) {
           case "true":
               d.isDeployed = false
               break;       
           default:
                d.isDeployed = true
               break;
       }

       return d 
    }
    
}

class Session{
    isSet:boolean
}

class Status{
    isDeployed:boolean
}