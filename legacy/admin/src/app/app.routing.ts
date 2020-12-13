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


import {NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";
import { RootComponent } from './pages/root/root.component';
import { DashModule } from "./pages/dash/dash.module";
import { DashNotFoundComponent } from "./pages/dash/outlet/center/not-found/dash.not.found.component";

import { AuthGuard } from "./services/auth.guard";
import { Initializer } from "./services/init.service"
import { SessionChecker } from "./services/session.checker";

export const ServerURL:string = 'http://localhost:5678'
export const appComponents = [RootComponent]
export const appModules = [DashModule]
export const appProviders = [Initializer,SessionChecker,AuthGuard ]

const routes:Routes = [    
    { path: '', redirectTo: 'init', pathMatch: 'full' },
    { path: 'init', component: RootComponent, },//test if session active. redirect to dash if true ?  
    {path : '**',component:DashNotFoundComponent}  
]

@NgModule({
    imports:[RouterModule.forRoot(routes) ],
    exports:[RouterModule]
})
export class AppRoutingModule{}
