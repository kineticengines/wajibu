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

import { Observable } from 'rxjs/Rx';
import { AnonymousSubscription } from "rxjs/Subscription"


class StepOneData{
    main : any
    hDetails : Array<Object>    
}

@Component({
    moduleId:module.id,
    selector:"dash-step-overview",
    template:`  
        <h2>Step One</h2>  
        <Section align="center">            
            <form [formGroup]="overviewForm" novalidate class="general-form" (ngSubmit)="processInputs(overviewForm)">
                <md-input-container class="general-form-input">
                    <input  mdInput type="text" formControlName="deployname" placeholder="Deployment Name" class="the-input">                    
                </md-input-container><br><br>                
                <!--<md-select class="general-form-input" formControlName="deploycountry" placeholder="Country">
                    <md-option *ngFor="let c of countries" [value]="c.name.common"  class="the-input">{{c.name.common}}</md-option>
                </md-select><br><br>-->
                <md-input-container class="general-form-input">
                    <input  mdInput type="text" formControlName="deploycountry" placeholder="Deployment Country" class="the-input">                    
                </md-input-container><br><br> 
                <md-input-container class="general-form-input">
                    <input  mdInput mdTooltip="The number of years the government holds office" type="text" formControlName="deployspan" placeholder="Deployment Span " class="the-input">                                    
                </md-input-container><br><br>                
                <md-select class="general-form-input" placeholder="Government Type" formControlName="governmenttype">
                    <md-option *ngFor="let type of govTypes" [value]="type.typeName">{{type.typeName}}</md-option>
                </md-select><br><br>              
                <md-input-container class="general-form-input">
                    <input  mdInput type="text" mdTooltip="The number of legislative houses" formControlName="numofhouses" (keyup)="updateCount(overviewForm.get('numofhouses').value)" placeholder="Number of Houses" class="the-input">                                                                   
                </md-input-container><br><br>            
                <Form [formGroup]="addForm" *ngIf="overviewForm.get('numofhouses').value >= 1" [hidden]="houseDetails.length == allHouseCount">
                    <fieldset>
                        <legend class="the-label">House Details</legend><br>                         
                        <md-input-container class="general-form-input">
                            <input type="text" mdInput formControlName="housename" placeholder="House Name" class="the-input">
                        </md-input-container><br>
                        <md-input-container class="general-form-input">
                             <input type="text" mdInput formControlName="repslot" placeholder="Represented Slot" class="the-input">
                        </md-input-container><br><br>
                        <md-input-container class="general-form-input">
                             <input type="text" mdInput formControlName="reptitle" placeholder="Representive Title" class="the-input">
                        </md-input-container><br><br>
                        <md-input-container class="general-form-input">
                            <input type="text" mdInput formControlName="numofseats" placeholder="Number of Seats" class="the-input"> 
                        </md-input-container><br><br>
                        <md-hint align="end" *ngIf="addForm.invalid">All fields required</md-hint>
                        <br><md-hint align="end">Total Houses Added : {{houseDetails.length}}</md-hint><br>                          
                        <button type="button" md-raised-button (click)="addHouseDetails(addForm)">ADD</button>                                                                    
                    </fieldset> 
                </Form><br><br>
                <Section>
                    <md-card class="house-container" *ngFor="let house of houseDetails; let i = index" [attr.data-index]="i">
                        <div class="house-item">
                            <h4 class="detail-header">House Name</h4>
                            <p>{{house.housename | titlecase}}</p>
                        </div>
                        <div class="house-item">
                            <h4 class="detail-header">Represented Slot</h4>
                            <p>{{house.repslot | titlecase}}</p>
                        </div>
                        <div class="house-item">
                            <h4 class="detail-header">Representive Title</h4>
                            <p>{{house.reptitle | titlecase}}</p>
                        </div>
                        <div class="house-item">
                            <h4 class="detail-header">Number of Seats</h4>
                            <p>{{house.numofseats}}</p>
                        </div> 
                        <div class="house-item">
                            <h4 class="detail-header">Action</h4>                            
                            <button md-raised-button type="button" class="remove-button" (click)="removeHouse(i)">REMOVE</button>
                        </div>                       
                    </md-card>
                </Section>
                <br>                            
                <fieldset *ngIf="overviewForm.get('governmenttype').value === 'Federal Government' || overviewForm.get('governmenttype').value === 'Devolved Government'  ">  
                    <legend class="the-label">Subgovernment Details</legend><br>                 
                    <md-input-container class="general-form-input">
                        <input  mdInput type="text" formControlName="numofsubgov" placeholder="Number of Subgovernments" class="the-input">                    
                    </md-input-container><br><br>
                    <md-input-container class="general-form-input">
                        <input  mdInput type="text" formControlName="subgovtitle" placeholder="Office Title" class="the-input">                    
                    </md-input-container><br><br>
                    <md-select class="general-form-input" placeholder="Subgovernment name" formControlName="subgovname">
                        <md-option *ngFor="let type of govSubTypes"  [value]="type.typeName">{{type.typeName}}</md-option>
                    </md-select><br><br> 
                    <md-slide-toggle [color]="primary" formControlName="subgovhasleg" >Does the Subgovernment have a legistlative arm?</md-slide-toggle>
                    <Section *ngIf="overviewForm.get('subgovhasleg').value">
                       <md-input-container class="general-form-input">
                            <input type="text" mdInput formControlName="subgovhousename" placeholder="House Name" class="the-input">
                        </md-input-container>
                        <md-input-container class="general-form-input">
                             <input type="text" mdInput formControlName="subgovhouserepslot" placeholder="Represented Slot" class="the-input">
                        </md-input-container>
                        <md-input-container class="general-form-input">
                             <input type="text" mdInput formControlName="subgovreptitle" placeholder="Representive Title" class="the-input">
                        </md-input-container>                        
                    </Section>
                </fieldset><br><br>                      
                <button type="submit" md-raised-button class="save-button" *ngIf="overviewForm.valid">SAVE</button> 
            </form>
        </Section>
    `,
    styleUrls:["../dash.install.root.css"]   
})
export class DashInstallStepOverview implements OnInit{
    overviewForm:FormGroup;
    addForm:FormGroup;
    govTypes = [{typeName:"Central Government"},{typeName:"Federal Government"},{typeName:"Devolved Government"}]
    govSubTypes = [{typeName:"State"},{typeName:"County"},{typeName:"Region/Province"}]
    houseDetails:Array<Object> = []
    allHouseCount:number = 0;
    showAddForm:boolean = false;    

    private timerSubscription: AnonymousSubscription;
   

    constructor(private init:Initializer,private formBuilder:FormBuilder,
    private snackbar:MdSnackBar,private router:Router){ 
                     
    }

    ngOnInit(){ 
        this.startFetcher()           
        this.createForm()
        this.createAddHouseForm() 
             
    }
    createForm(){
        this.overviewForm = this.formBuilder.group({            
            deployname : ['',[Validators.required]], 
            deploycountry : ['',[Validators.required]],
            deployspan : ['',[Validators.required,isNumber,isNotZero]],            
            governmenttype : ['',[Validators.required]], 
            numofhouses : ['',[Validators.required,isNumber,isNotZero]],  
            subgovname : ['',[]], 
            subgovtitle : ['',[]],
            numofsubgov : ['',[]], 
            subgovhasleg : [false],
            subgovhousename : [''], 
            subgovhouserepslot: [''],
            subgovreptitle : [''],                    
        })
        
    }
    createAddHouseForm(){
        this.addForm = this.formBuilder.group({
            housename : ['',[Validators.required]],
            repslot : ['',[Validators.required]],
            reptitle : ['',[Validators.required]],
            numofseats : ['',[Validators.required,isNumber,isNotZero]]
        })
    }    

    updateCount(newCount:number){
        this.allHouseCount = newCount;     
    }
    addHouseDetails(){
        if(this.addForm.valid){
            this.houseDetails.push(this.addForm.value)
            this.addForm.reset()
        }
    }
    removeHouse(index:any){
        if(index > -1){
            this.houseDetails.splice(index,1)
        }
    }
   
    processInputs(ovForm:any){
        //TODO: Create service to validate input and return errors        
        let sOne = new StepOneData()
        sOne.main = ovForm.value;
        sOne.hDetails = this.houseDetails;                
        this.init.goProcessStepOne(sOne).subscribe(d => {
            switch (d.Status) {
                case true:
                    this.snackbar.open("Save Successfull","Close",{duration:2500})
                    this.router.navigate(["dash/install/toplevel"])
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
            if(element.Level === "build:one"){        
                if(element.Exist === false && element.Complete === false){                    
                    this.subscribeToData()
                }else{                                  
                    this.overviewForm.setValue({
                        deployname : element.Data.main.deployname,
                        deploycountry : element.Data.main.deploycountry,
                        deployspan : element.Data.main.deployspan,            
                        governmenttype : element.Data.main.governmenttype, 
                        numofhouses : element.Data.main.numofhouses, 
                        subgovname : element.Data.main.subgovname,
                        subgovtitle : element.Data.main.subgovtitle,
                        numofsubgov : element.Data.main.numofsubgov, 
                        subgovhasleg : element.Data.main.subgovhasleg,
                        subgovhousename : element.Data.main.subgovhousename,
                        subgovhouserepslot: element.Data.main.subgovhouserepslot,
                        subgovreptitle : element.Data.main.subgovreptitle
                    })                                                          
                    element.Data.hdetails.forEach(elem => {
                        this.houseDetails.push(elem)
                    }); 
                    this.updateCount(element.Data.hdetails.length)
                                 
                } 
            }
        }    
        
    }

   private subscribeToData(): void {
    this.timerSubscription = Observable.timer(2000)
      .subscribe(() => this.startFetcher());
   }
    
}

