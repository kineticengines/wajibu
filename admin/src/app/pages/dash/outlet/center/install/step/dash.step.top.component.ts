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


import { Component ,OnInit} from "@angular/core";
import {MdSnackBar } from "@angular/material";
import { Router } from "@angular/router";
import { FormBuilder, FormGroup, Validators,FormControl} from "@angular/forms";
import { Initializer } from "../../../../../../services/init.service";
import { isNumber,isNotZero } from "../../../../../../services/util.service";
import { CustomValidators } from 'ng2-validation';

import { Observable } from 'rxjs/Rx';
import { AnonymousSubscription } from "rxjs/Subscription"

@Component({
    selector:"dash-step-top",
    template:`
        <h2>Step Two</h2>  
        <Section align="center">
            <form [formGroup]="topLevelForm" novalidate class="general-form" (ngSubmit)="processInputs(topLevelForm)">
                <Section>
                    <h3>President</h3>
                    <md-input-container class="general-form-input">
                        <input  mdInput type="text" formControlName="presidentName" placeholder="President Name" class="the-input">                    
                    </md-input-container><br><br>
                    <md-input-container class="general-form-input">
                        <input  mdInput type="text" formControlName="presidentTerm" placeholder="President Term" class="the-input">                    
                    </md-input-container><br><br>
                    <md-select class="general-form-input" placeholder="President Gender" formControlName="pgender">
                        <md-option *ngFor="let type of genderTypes"  [value]="type.typeName">{{type.typeName}}</md-option>
                    </md-select><br><br> 
                    <md-input-container class="general-form-input">
                        <input  mdInput type="text" formControlName="pnthPosition" placeholder="President nTH Position" class="the-input">                    
                    </md-input-container><br><br>
                    <md-input-container class="general-form-input">
                        <input  mdInput type="text" formControlName="pImage" placeholder="President Image URL" class="the-input-extra">                    
                    </md-input-container><br><br>
                </Section>
                <Section>
                    <h3>Deputy President</h3>
                    <md-input-container class="general-form-input">
                        <input  mdInput type="text" formControlName="dpresidentName" placeholder="Deputy President Name" class="the-input">                    
                    </md-input-container><br><br>
                    <md-input-container class="general-form-input">
                        <input  mdInput type="text" formControlName="dpresidentTerm" placeholder="Deputy President Term" class="the-input">                    
                    </md-input-container><br><br>
                    <md-select class="general-form-input" placeholder="Deputy President Gender" formControlName="dgender">
                        <md-option *ngFor="let type of genderTypes"  [value]="type.typeName">{{type.typeName}}</md-option>
                    </md-select><br><br> 
                    <md-input-container class="general-form-input">
                        <input  mdInput type="text" formControlName="dnthPosition" placeholder="Deputy President nTH Position" class="the-input">                    
                    </md-input-container><br><br>
                    <md-input-container class="general-form-input">
                        <input  mdInput type="text" formControlName="dImage" placeholder="Deputy President Image URL" class="the-input-extra">                    
                    </md-input-container><br><br>
                </Section>
                 <button type="submit" md-raised-button class="save-button" *ngIf="topLevelForm.valid">SAVE</button> 
            </form>
        </Section> 
    `,
    styleUrls:["../dash.install.root.css"]   
})
export class DashInstallStepTop implements OnInit{
    topLevelForm:FormGroup
    genderTypes = [{typeName:"Male"},{typeName:"Female"}]

   timerSubscription: AnonymousSubscription; 

    constructor(private init:Initializer,private formBuilder:FormBuilder,
        private snackbar:MdSnackBar,private router:Router){                  
    }

    ngOnInit(){
       this.startFetcher()   
       this.createForm() 
    }
    private createForm(){
        this.topLevelForm = this.formBuilder.group({
            presidentName : ['',[Validators.required]],
            presidentTerm : ['',[Validators.required,isNumber,isNotZero]],
            pgender : ['',[Validators.required]],
            pnthPosition : ['',[Validators.required,isNumber,isNotZero]],
            pImage : ['',[Validators.required,CustomValidators.url]],
            dpresidentName : ['',[Validators.required]],
            dpresidentTerm : ['',[Validators.required,isNumber,isNotZero]],
            dgender : ['',[Validators.required]],
            dnthPosition : ['',[Validators.required,isNumber,isNotZero]],
            dImage : ['',[Validators.required,CustomValidators.url]]
        })
    }

    processInputs(tpForm:any){
        //TODO: Create service to validate input and return errors       
        this.init.goProcessStepTwo(tpForm.value).subscribe(d => {
            switch (d.Status) {
                case true:
                    this.snackbar.open("Save Successfull","Close",{duration:2500})
                    this.router.navigateByUrl("dash/install/houselevel")
                    break;            
                case false:
                    this.snackbar.open("Save Unsuccessfull","",{duration:5000})
                    break;
            }
        })
    }

    private startFetcher(){
        this.init.checkBuildLevel().subscribe(d => this.processBuild(d))
    }

    private processBuild(d:any){     
        for (var index = 0; index < d.levels.length; index++) {
            let element = d.levels[index];  
            //console.log(element)      
            if(element.Level === "build:two"){        
                if(element.Exist === false && element.Complete === false){                    
                    this.subscribeToData()
                }else{                                  
                    this.topLevelForm.setValue({
                        presidentName : element.Data.presidentName,
                        presidentTerm : element.Data.presidentTerm,
                        pgender : element.Data.pgender,
                        pnthPosition : element.Data.pnthPosition,
                        pImage : element.Data.pImage,
                        dpresidentName : element.Data.dpresidentName,
                        dpresidentTerm : element.Data.dpresidentTerm,
                        dgender : element.Data.dgender,
                        dnthPosition : element.Data.dnthPosition,
                        dImage : element.Data.dImage
                    })                           
                } 
            }
        }       
    }

    private subscribeToData(): void {
        this.timerSubscription = Observable.timer(2000)
        .subscribe(() => this.startFetcher());
    }
}