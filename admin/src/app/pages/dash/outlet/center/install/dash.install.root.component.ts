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


import { Component,OnInit, ViewEncapsulation } from "@angular/core";
import { Router } from "@angular/router";
import { Initializer } from "../../../../../services/init.service";
import { SessionChecker } from "../../../../../services/session.checker";

import { Observable } from 'rxjs/Rx';
import { AnonymousSubscription } from "rxjs/Subscription"

@Component({
  moduleId:module.id,
  template: `
    <Section align="center">
      <!--<h1>Wajibu</h1>-->
      <img src="./assets/logo.png" style="width:150px;height:150px;">
      <h3>Building and deployement</h3>
    </Section> 
    <div class="logout-container">
      <button md-button (click)="logOut()" style="background-color:teal;color: white;">LOGOUT</button> 
    </div>     
    <nav id="crouton" align="center">
        <ul>
            <li><a routerLink="../install">Overview</a></li>
            <li><a routerLink="../install/toplevel" routerLinkActive="active" *ngIf="stepOneComplete">Top Level</a></li>
            <li><a routerLink="../install/houselevel" routerLinkActive="active" *ngIf="stepTwoComplete">House Level</a></li>
            <li><a routerLink="../install/subgovlevel" routerLinkActive="active" *ngIf="stepThreeComplete && notCentralGovernment">Subgovernment Level</a></li>
            <li><a routerLink="../install/rootlevel" routerLinkActive="active" *ngIf="stepFourComplete && notCentralGovernment">Grassroot Level</a></li>
        </ul>
    </nav> 
    <br>    
    <router-outlet></router-outlet>
  `,
  encapsulation : ViewEncapsulation.None,
  styleUrls:["dash.install.root.css"]
})
export class DashInstallRootComponent {
  stepOneComplete:boolean;
  stepTwoComplete:boolean;
  stepThreeComplete:boolean;
  stepFourComplete:boolean;
  notCentralGovernment:boolean;
  
  private timerSubscription: AnonymousSubscription;

  constructor(private init:Initializer,private session:SessionChecker,private router:Router){}
  ngOnInit(){
    this.startFetcher()    
  }
 
  startFetcher(){
    this.init.checkBuildLevel().subscribe(d => this.processBuild(d))
  }
  processBuild(d:any){
    for (var index = 0; index < d.levels.length; index++) {
      let element = d.levels[index];       
      if(element.Level === "build:one"){        
        if(element.Data !== null){
          this.checkIfCentral(element.Data.main.governmenttype)
        }         
        switch (element.Exist) {
          case false:   
            this.stepOneComplete = element.Exist;         
            break;        
          case true:
            this.stepOneComplete = element.Complete;
            break;
        }     
      }else if(element.Level === "build:two"){        
        switch (element.Exist) {
          case false:  
            this.stepTwoComplete = element.Exist;          
            break;        
          case true:
            this.stepTwoComplete = element.Complete;
            break;
        } 
      }else if(element.Level === "build:three"){      
        switch (element.Exist) {
          case false: 
            this.stepThreeComplete = element.Exist;         
            break;        
          case true:
            this.stepThreeComplete = element.Complete;
            break;
        } 
      }else if(element.Level === "build:four"){        
        switch (element.Exist) {
          case false: 
            this.stepFourComplete = element.Exist;         
            break;        
          case true:
            this.stepFourComplete = element.Complete;
            break;
        } 
      }
    }    
    this.subscribeToData();
  }

   private subscribeToData(): void {
    this.timerSubscription = Observable.timer(500)
      .subscribe(() => this.startFetcher());
   }

   private checkIfCentral(type:string){
      switch (type) {
        case "Central Government":
            this.notCentralGovernment = false;
          break;      
        default:
            this.notCentralGovernment = true;
          break;
      }
   }

  logOut(){
    this.session.destroySession()
    this.router.navigateByUrl("init")
  }
  


}

