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
import { CanActivate,CanActivateChild,
     Router,ActivatedRouteSnapshot ,RouterStateSnapshot} from "@angular/router";
import { SessionChecker } from "./session.checker";

import { Observable} from "rxjs/Rx";
import "rxjs/add/operator/do";
import "rxjs/add/operator/map";
import "rxjs/add/operator/take";

@Injectable()
export class AuthGuard implements CanActivate,CanActivateChild{    
      
      constructor(private router: Router, private session:SessionChecker){}            
      
      canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot):Observable<boolean>{           
            return Observable.of(this.session.getSession())
                             .take(1)
                             .map(d => d.isSet)
                             .do(isSet => {                                 
                                 if (!isSet) this.router.navigate(["init"])                                 
                             })
          
          
      }  
       canActivateChild(route: ActivatedRouteSnapshot, state: RouterStateSnapshot):Observable<boolean> {           
            return Observable.of(this.session.getDeploy())                             
                             .take(1)
                             .map(d => d.isDeployed)                             
                             .do(isDeployed => {                                 
                                 if (!isDeployed) this.router.navigate(["dash/install"])   
                                 //console.log(isDeployed)                                               
                             })                     
                             
       }
    
       
}

@Injectable()
export class DeployedGuard implements CanActivateChild{
    constructor(private router: Router, private session:SessionChecker){} 
    canActivateChild(route: ActivatedRouteSnapshot, state: RouterStateSnapshot):Observable<boolean>{        
            return Observable.of(this.session.getNotDeployed())
                             .take(1)
                             .map(d => d.isDeployed)
                             .do(isDeployed => {                                                       
                                 if (!isDeployed) this.router.navigate(["dash"])
                             })
    }
}






