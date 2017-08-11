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
    selector:"dash-sub-gov",
    template:`
        <h2>Step Four</h2>
        <h4 *ngIf="govsBuild.length < numOfGovs">Subgovernments remaining : {{numOfGovs - govsBuild.length}}</h4>
        <Section align="center">
            <form [formGroup]="govForm" novalidate class="general-form" *ngIf="govsBuild.length < numOfGovs">
                <md-input-container class="general-form-input">
                    <input  mdInput type="text"  formControlName="HeadName" placeholder="{{officeTitle | titlecase}} Name" class="the-input">                    
                </md-input-container><br><br>
                <md-input-container class="general-form-input">
                    <input  mdInput type="text" formControlName="HeadTerm" placeholder="{{officeTitle | titlecase}} Term" class="the-input">                    
                </md-input-container><br><br>
                <md-select class="general-form-input" placeholder="{{officeTitle | titlecase}} Gender" formControlName="HeadGender">
                    <md-option *ngFor="let type of genderTypes"  [value]="type.typeName">{{type.typeName}}</md-option>
                </md-select><br><br> 
                <md-input-container class="general-form-input">
                    <input  mdInput type="text" formControlName="HeadnthPosition" placeholder="{{officeTitle | titlecase}} nth Position" class="the-input">                    
                </md-input-container><br><br>
                <md-input-container class="general-form-input">
                    <input  mdInput type="url" formControlName="HeadImage" placeholder="{{officeTitle | titlecase}} Image URL" class="the-input-extra">                    
                </md-input-container><br><br>
                <md-input-container class="general-form-input">
                        <input  mdInput type="text" formControlName="SlotName" placeholder="{{subGovName | titlecase}} Name" class="the-input">                    
                    </md-input-container><br><br>
                <section *ngIf="hasLeg">                    
                    <md-input-container class="general-form-input">
                        <input  mdInput type="text"  formControlName="NumOfLegSeats" placeholder="{{subGovName}} House Number of seats" class="the-input">                    
                    </md-input-container><br><br> 
                </section>
                
                <button type="button" md-raised-button class="save-button" (click)="addHouseRep(govForm)" *ngIf="govForm.valid">ADD {{officeTitle.toUpperCase()}}</button>
            </form>
            <div>
                <table align="center" *ngIf="govsBuild.length >= 1">
                    <thead>                
                        <tr>
                            <th>{{officeTitle | titlecase}} Name</th>
                            <th>{{officeTitle | titlecase}} Term</th>
                            <th>{{officeTitle | titlecase}} Gender</th>
                            <th>{{officeTitle | titlecase}} nth Position</th>                         
                            <th>{{officeTitle | titlecase}} Image URL</th>
                            <th>{{subGovName}} Name</th>
                            <th *ngIf="hasLeg">{{subGovName}} seats</th>
                            <th>Action</th>
                        <tr>
                    </thead>
                    <tbody>
                        <tr *ngFor="let gov of govsBuild; let i = index" [attr.data-index]="i" >
                            <td>{{gov.HeadName | titlecase}}</td>
                            <td>{{gov.HeadTerm | titlecase }}</td>
                            <td>{{gov.HeadGender | titlecase}}</td>
                             <td>{{gov.HeadnthPosition | titlecase}}</td>
                            <td>{{gov.HeadImage | lowercase}}</td>
                            <td>{{gov.SlotName | titlecase}}</td>
                            <td *ngIf="hasLeg">{{gov.NumOfLegSeats | titlecase}}</td>
                            <td><button md-raised-button type="button" class="remove-button" (click)="removeGovBuild(i)">REMOVE</button></td>
                        </tr>
                    </tbody>
                </table>
                <button type="button" md-raised-button  *ngIf="govsBuild.length >= 1 && govsBuild.length == numOfGovs" class="save-button" (click)="saveAll()">SAVE ALL {{officeTitle | uppercase}}S</button> 
             </div>
             <div>
                <table align="center" *ngIf="allGov.length >= 1">
                    <thead>                
                        <tr>
                            <th>{{officeTitle | titlecase}} Name</th>
                            <th>{{officeTitle | titlecase}} Term</th>
                            <th>{{officeTitle | titlecase}} Gender</th>  
                            <th>{{officeTitle | titlecase}} nth Position</th>                      
                            <th>{{officeTitle | titlecase}} Image URL</th>
                            <th>{{subGovName | titlecase}} Name</th>                        
                            <th>{{subGovName | titlecase}} Seats</th>
                        <tr>
                    </thead>
                    <tbody>
                        <tr *ngFor="let gov of allGov; let i = index" [attr.data-index]="i" >
                            <td>{{gov.headname | titlecase}}</td>
                            <td>{{gov.headterm | titlecase}}</td>
                            <td>{{gov.headgender | titlecase}}</td>
                            <td>{{gov.headnthposition | titlecase}}</td>
                            <td>{{gov.headimage | lowercase}}</td>
                            <td>{{gov.slotname | titlecase}}</td>
                            <td>{{gov.numoflegseats | titlecase}}</td>                            
                        </tr>
                        <button type="button" md-raised-button  class="remove-button" (click)="resetSubGov()">RESET SUBGOVERNMENT LEVEL</button> 
                    </tbody>
                </table>
             </div>
        </Section>
    `,
    styleUrls:["../dash.install.root.css"] 
})
export class DashInstallStepSubGov implements OnInit{
    private timerSubscription: AnonymousSubscription;
    public numOfGovs:number;
    public officeTitle:string;
    public subGovName:string;
    public hasLeg:boolean;
    public govsBuild:Array<Object> = []
    public allGov:Array<Object> = []

    govForm:FormGroup;
    genderTypes = [{typeName:"Male"},{typeName:"Female"}]

    constructor(private init:Initializer,private formBuilder:FormBuilder,
        private snackbar:MdSnackBar,private router:Router){     
    }
    ngOnInit(){
       this.stepInterInitializer()       
       this.createForm()
    }
    private stepInterInitializer(){ 
        this.govsBuild.splice(0,this.govsBuild.length) 
        this.allGov.splice(0,this.allGov.length) 
        this.numOfGovs = 0;           
        this.init.buildFourInitializer().subscribe(d => { 
            this.officeTitle = d.officetitle;
            this.subGovName = d.subgovname;                                                                  
            if (d.complete === true){
                this.startFetcher()
            }else{
                this.numOfGovs = d.numofsubgov; 
                this.hasLeg = d.hasleg;                
            }
        })
    }
    private createForm(){
        switch (this.hasLeg) {
            case false:
                this.govForm = this.formBuilder.group({        
                    HeadName : ['',[Validators.required]],
                    HeadTerm : ['',[Validators.required,isNumber,isNotZero]],
                    HeadGender : ['',[Validators.required]], 
                    HeadnthPosition : ['',[Validators.required,isNumber,isNotZero]],
                    HeadImage : ['',[Validators.required,CustomValidators.url]], 
                    SlotName :  ['',[Validators.required]]                                    
                })                
                break;        
            default:
                this.govForm = this.formBuilder.group({        
                    HeadName : ['',[Validators.required]],
                    HeadTerm : ['',[Validators.required,isNumber,isNotZero]],
                    HeadGender : ['',[Validators.required]], 
                    HeadnthPosition : ['',[Validators.required,isNumber,isNotZero]],
                    HeadImage : ['',[Validators.required,CustomValidators.url]], 
                    SlotName :  ['',[Validators.required]],           
                    NumOfLegSeats : ['',[Validators.required,isNumber,isNotZero]]                       
                })
                break;
        }
        
    }
    private addHouseRep(inForm:any){
        if(inForm.valid){
            this.govsBuild.push(inForm.value)
            this.govForm.reset()
        }
    }
    private removeGovBuild(index:any){
        if(index > -1){
            this.govsBuild.splice(index,1)
        }
    }

    private saveAll(){          
        this.init.goProcessStepFour(this.govsBuild).subscribe(d => {
            switch (d.Status) {
                case true:
                    this.stepInterInitializer()
                    break;            
                case false:
                    this.snackbar.open("Save Unsuccessfull","",{duration:5000})
                    break;
            }
        })
    }

    startFetcher(){        
        this.init.checkBuildLevel().subscribe(d => this.processBuild(d))
    }

    private resetSubGov(){
        this.init.resetBuildFour().subscribe(d => {            
            if (d.Status === true){                
                this.ngOnInit()
            }
        })
    }

    private processBuild(d:any){              
        for (var index = 0; index < d.levels.length; index++) {
            let element = d.levels[index]; 
            if(element.Level === "build:four"){                
                if(element.Exist === false && element.Complete === false){                    
                    this.subscribeToData()
                }else{                                                                    
                    element.Data.subgovdata.forEach(element => {                        
                        this.allGov.push(element)
                    });                                                    
                } 
            }
        }
    }

    private subscribeToData(): void {
        this.timerSubscription = Observable.timer(2000)
        .subscribe(() => this.startFetcher());
    }
   
}

