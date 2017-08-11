/*
Wajibu is an online web app that collects,analyses and aggregates sentiments from the public
pertaining the government of a nation. This tool allows citizens to contribute to the
governance talk by airing out their honest views about the state of the nation and in
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

import { Component } from "@angular/core";
import { Router } from "@angular/router"
import { SessionChecker } from "../../../../../services/session.checker";

@Component({
  moduleId:module.id,
  template: `
    <div class="dashboard">
      <div class="top-container">
        <h1 class="top-item">Wajibu</h1>
        <button class="top-item" md-raised-button (tap)="logout()">LOGOUT</button>
      </div>
      <div class="body-container">
          <div class="side-bar body-item">
              <side-bar></side-bar>
          </div>
          <div class="main-bar body-item">
              <router-outlet></router-outlet>
          </div>
      </div>      
    </div>    
  `,
  styleUrls:["dash.center.home.css"]
})
export class DashCenterHomeComponent {
  constructor(private router:Router,private session:SessionChecker){
  }
  logout(){
    this.session.destroySession()
    this.router.navigate(["init"])
  }
}