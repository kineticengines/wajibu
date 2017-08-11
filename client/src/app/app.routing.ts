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

import { RouterModule, Routes } from "@angular/router";
import { RootComponent } from './pages/root/root.component';
import { IndexCompenent } from "./pages/index/index.component";
import { GeneralComponent} from "./pages/index/general/general.component"
import { ChartBoxComponent } from "./pages/index/general/chart-box/chart.box.component";
import { FilterBoxComponent } from "./pages/index/general/filter-box/filter.box.component";
import { SentimentBoxComponent } from "./pages/index/general/sentiment-box/sentiment.box.component";
import { TopBoxComponent } from "./pages/index/general/top-box/top.box.component";
import { FilterComponent } from "./pages/index/filter/filter.component";

import { NotFoundComponent } from "./pages/not-found/not.found.component";

import { AuthGuard } from "./services/auth.guard";

const routes:Routes = [    
    { path: '', redirectTo: 'init', pathMatch: 'full' },
    { path: 'init', component: RootComponent },
    {path:'home',canActivate:[AuthGuard],component:IndexCompenent,children:[
        {path:"",component:GeneralComponent},
        {path:"fan/:who",component:FilterComponent},
        {path:"**",component:NotFoundComponent} 
    ]},
    {path:"**",component:NotFoundComponent}    
]

export const routing = RouterModule.forRoot(routes);

export const ServerURL:string = 'http://localhost:5678'

export const appComponents = [
    RootComponent,IndexCompenent,GeneralComponent,ChartBoxComponent,FilterBoxComponent,SentimentBoxComponent,
    TopBoxComponent,FilterComponent,NotFoundComponent
]