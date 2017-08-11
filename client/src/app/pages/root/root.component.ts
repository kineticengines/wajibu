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

import { Component, OnInit} from '@angular/core';
import { Router } from "@angular/router";
import { Initializer } from "../../services/init.service"

import { Observable } from 'rxjs/Rx';
import { AnonymousSubscription } from "rxjs/Subscription"
 
@Component({
  selector: 'app-root',
  template:`
    <Section>
      <md-card *ngIf="notDeployed">
        <md-card-content>
          <h3 align="center">Wajibu Has Not Yet Been Deployed</h3>
        </md-card-content>
      </md-card>
    <Section>
  `,
  styleUrls: ['./root.component.css'],
  providers:[Initializer]
})
export class RootComponent implements OnInit {
  private timerSubscription: AnonymousSubscription;

  public notDeployed:boolean;
  constructor(private initializer:Initializer,private router:Router){}
  ngOnInit(){
       this.startDeployChecker() 
  }

  private startDeployChecker(){
    this.initializer.getInit().subscribe(d => {      
      if(d.Deployed === false){
        this.notDeployed = true
        this.subscribeToDeployCheck()
      }else{
        //change route to index
        this.router.navigate(["home"])
      } 
    }) 
  }

  private subscribeToDeployCheck(): void {    
    this.timerSubscription = Observable.timer(1000)
      .subscribe(() => this.startDeployChecker());
   }

  
}
