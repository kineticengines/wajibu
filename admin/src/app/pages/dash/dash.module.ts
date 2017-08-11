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


import { NgModule }  from '@angular/core';
import { CommonModule }   from '@angular/common';
import { FormsModule,ReactiveFormsModule } from '@angular/forms';
import { RouterModule, Routes } from '@angular/router';
import { MaterialModule } from '@angular/material';

import { AuthGuard,DeployedGuard} from "../../services/auth.guard";
import { DashRootComponent } from "./dash.root.component";
import { DashCenterHomeComponent } from "./outlet/center/home/dash.center.home.component";
import { DashInstallRootComponent  } from "./outlet/center/install/dash.install.root.component";
import { DashNotFoundComponent } from "./outlet/center/not-found/dash.not.found.component";

import { DashInstallStepOverview } from "./outlet/center/install/step/dash.step.overview.component";
import { DashInstallStepTop } from "./outlet/center/install/step/dash.step.top.component";
import { DashInstallStepHouse } from "./outlet/center/install/step/dash.step.house.component";
import { DashInstallStepSubGov } from "./outlet/center/install/step/dash.step.sub.component";
import { DashInstallStepGrass } from "./outlet/center/install/step/dash.step.grass.component";

import { DashDeploy } from "./outlet/center/deploy/dash.deploy.component";

import { GeneralComponent } from "./outlet/center/home/general/dash.general.component";
import { SideBarComponent } from "./outlet/center/home/sidebar/dash.sidebar.component";
import { AmendComponent } from "./outlet/center/home/amend/dash.amend";
import { AmendPillarsComponent } from "./outlet/center/home/amend/amend.pillars.component";
import { SentimentComponent } from "./outlet/center/home/sentiment/dash.sentiments";
import { SettingsComponent } from "./outlet/center/home/settings/dash.settings";

import { DynamicFormModule  } from "../../dynamic-form/dynamic-form.module"

import { DashService } from "../../services/dash.service";

const dashComponents = [DashRootComponent,DashCenterHomeComponent, DashNotFoundComponent,
DashInstallRootComponent,DashInstallStepOverview,DashInstallStepTop,
DashInstallStepHouse,DashInstallStepSubGov,DashInstallStepGrass,DashDeploy,GeneralComponent,
SideBarComponent,AmendComponent,SentimentComponent,SettingsComponent,AmendPillarsComponent]

const routes:Routes = [
    {
        path:"dash",component: DashRootComponent,canActivate:[AuthGuard],
        children:[
            {path:"",component: DashCenterHomeComponent,canActivateChild:[AuthGuard],
                children:[
                    { path: '', component: GeneralComponent},
                    { path: 'sentiment', component: SentimentComponent},
                    { path: 'amend', component: AmendComponent},
                    { path: 'settings', component: SettingsComponent},
                    {path : '**',component:DashNotFoundComponent}
                ]
            },
            {path:"install",component:DashInstallRootComponent,canActivateChild:[DeployedGuard],
                children:[
                    {path:"",component:DashInstallStepOverview},
                    {path:"toplevel",component:DashInstallStepTop},
                    {path:"houselevel",component:DashInstallStepHouse},
                    {path:"subgovlevel",component:DashInstallStepSubGov},
                    {path:"rootlevel",component:DashInstallStepGrass},
                    {path : '**',component:DashNotFoundComponent}
                ]
            },
            {path:"deploy",component:DashDeploy},
            {path : '**',component:DashNotFoundComponent}
        ]
    },
    
   
]


@NgModule({
    imports:[RouterModule.forChild(routes)],
    exports:[RouterModule],
    providers: [AuthGuard,DeployedGuard]     
})
export class DashRoutingModule{}

@NgModule({
    declarations: [...dashComponents],
    entryComponents:[],
    exports:[DashNotFoundComponent],
    imports: [ CommonModule, FormsModule,ReactiveFormsModule,DynamicFormModule,
    DashRoutingModule,MaterialModule],
    providers:[DashService]
      
})
export class DashModule{}

