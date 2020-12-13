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
    selector:"dash-step-intermidiate",
    template:`
        <h2>Step Three</h2>
        <h4 *ngIf="allHouses.length < 1">Seats remaining : {{numOfSeats - houseSlots.length}}</h4>        
        <Section align="center">
             <form [formGroup]="interForm" novalidate class="general-form" *ngIf="houseSlots.length < numOfSeats">
                <h3>{{houseName | titlecase}}</h3>
                <md-input-container class="general-form-input">
                    <input  mdInput type="text" formControlName="RepName" placeholder="{{houseRepTitle | titlecase}} Name" class="the-input">                    
                </md-input-container><br><br>
                <md-input-container class="general-form-input">
                    <input  mdInput type="text" formControlName="RepTerm" placeholder="{{houseRepTitle | titlecase}} Term" class="the-input">                    
                </md-input-container><br><br>
                <md-select class="general-form-input" placeholder="{{houseRepTitle | titlecase}} Gender" formControlName="RepGender">
                    <md-option *ngFor="let type of genderTypes"  [value]="type.typeName">{{type.typeName}}</md-option>
                </md-select><br><br> 
                <md-input-container class="general-form-input">
                    <input  mdInput type="text" formControlName="RepSlot" placeholder="{{houseRepSlot | titlecase}} Name" class="the-input">                    
                </md-input-container><br><br>
               <md-input-container class="general-form-input">
                    <input  mdInput type="url" formControlName="RepImage" placeholder="{{houseRepTitle | titlecase}} Image URL" class="the-input-extra">                    
                </md-input-container><br><br> 
                <button type="button" md-raised-button class="save-button" (click)="addHouseRep(interForm)" *ngIf="interForm.valid">ADD {{houseRepTitle.toUpperCase()}}</button>
             </form><br><br>          
             <table align="center" *ngIf="houseSlots.length >= 1">
                <thead>                
                    <tr>
                        <th>{{houseRepTitle | titlecase}} Name</th>
                        <th>{{houseRepTitle | titlecase}} Term</th>
                        <th>{{houseRepTitle | titlecase}} Gender</th>
                        <th>{{houseRepSlot | titlecase}} Name</th>
                        <th>{{houseRepTitle | titlecase}} Image URL</th>
                        <th>Action</th>
                    <tr>
                </thead>
                <tbody>
                    <tr *ngFor="let house of houseSlots; let i = index" [attr.data-index]="i" >
                        <td>{{house.RepName | titlecase}}</td>
                        <td>{{house.RepTerm | titlecase}}</td>
                        <td>{{house.RepGender | titlecase }}</td>
                        <td>{{house.RepSlot | titlecase}}</td>
                        <td>{{house.RepImage | titlecase}}</td>
                        <td><button md-raised-button type="button" class="remove-button" (click)="removeHouseRep(i)">REMOVE</button></td>
                    </tr>
                </tbody>
             </table> 
             <button type="button" md-raised-button  *ngIf="houseSlots.length >= 1 && houseSlots.length == numOfSeats" class="save-button" (click)="saveAll()">SAVE ALL {{houseRepTitle.toUpperCase()}}S</button>   
             <div>
                <table align="center" *ngIf="allHouses.length >= 1">
                    <thead>                
                        <tr>
                            <th>House Name</th>
                            <th>Name</th>
                            <th>Term</th>
                            <th>Gender</th>
                            <th>Represented Slot</th>
                            <th>Image URL</th>                        
                        <tr>
                    </thead>
                    <tbody>
                        <tr *ngFor="let house of allHouses">
                            <td>{{house.houseName | titlecase}}</td>
                            <td>{{house.repName | titlecase }}</td>
                            <td>{{house.repTerm | titlecase}}</td>
                            <td>{{house.repGender | titlecase}}</td>
                            <td>{{house.repSlot | titlecase }} {{house.slotName | titlecase}}</td>
                            <td>{{house.repImage | lowercase}}</td>
                        </tr>
                        <br>
                    <button type="button" md-raised-button  class="remove-button" (click)="resetHouses()">RESET HOUSE LEVEL</button> 
                    </tbody>
                </table><br><br>
                 <Section class="deploy-section" align="center" *ngIf="allHouses.length >= 1 && !notCentralGovernment">                    
                    <button type="button" class="deploy-button" (click)="goToDeploy()">DEPLOY WAJIBU</button><br>
                    <span>please confirm the inputs are correct before deploying</span>
                </Section>
             </div>             
        </Section>
    `,
    styleUrls:["../dash.install.root.css"] 
})
export class DashInstallStepHouse implements OnInit{
    private timerSubscription: AnonymousSubscription;
    numOfSeats:number;
    houseSlots:Array<Object> = [];
    houseName:string
    houseRepSlot:string
    houseRepTitle:string
    houseKey:string
    notCentralGovernment:boolean;

    interForm:FormGroup;
    genderTypes = [{typeName:"Male"},{typeName:"Female"}]

    allHouses:Array<Object> = []

     constructor(private init:Initializer,private formBuilder:FormBuilder,
        private snackbar:MdSnackBar,private router:Router){     
    }
    ngOnInit(){
       this.stepInterInitializer()
       this.startFetcher()
       this.createForm()
    }
    private stepInterInitializer(){  
        this.houseSlots.splice(0,this.houseSlots.length) 
        this.allHouses.splice(0,this.allHouses.length)     
        this.init.buildThreeInitializer().subscribe(d => {                                      
            this.numOfSeats = d.numofseats;
            this.houseName = d.builtdata.housename;
            this.houseRepSlot = d.builtdata.repslot;
            this.houseRepTitle = d.builtdata.reptitle;
            this.houseKey = d.key;
        })
    }
    private createForm(){
        this.interForm = this.formBuilder.group({        
            RepName : ['',[Validators.required]],
            RepTerm : ['',[Validators.required,isNumber,isNotZero]],
            RepGender : ['',[Validators.required]], 
            RepSlot : ['',[Validators.required]],            
            RepImage : ['',[Validators.required,CustomValidators.url]], 
                       
        })
    }
    private startFetcher(){
        this.init.checkBuildLevel().subscribe(d => this.processBuild(d))
    }
    private addHouseRep(inForm:any){        
        if(inForm.valid){
            this.houseSlots.push(inForm.value)
            this.interForm.reset()
        }
    }
    private removeHouseRep(index:any){
        if(index > -1){
            this.houseSlots.splice(index,1)
        }
    }
    private saveAll(){ 
        let h = new BuildThreeData()
        h.key = this.houseKey
        h.housename = this.houseName
        h.slotname = this.houseRepSlot
        this.houseSlots.forEach(el => {
            h.main.push(el)
        });        
        this.init.goProcessStepThree(h).subscribe(d =>{
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
    private resetHouses(){
        this.init.resetBuildThree().subscribe(d => {            
            if (d.Status === true){                
                this.ngOnInit()
            }
        })
    }

    private processBuild(d:any){        
        for (var index = 0; index < d.levels.length; index++) {
            let element = d.levels[index]; 
            if(element.Level === "build:one"){     
                this.checkIfCentral(element.Data.main.governmenttype)
            }else if(element.Level === "build:three"){
                if(element.Exist === false && element.Complete === false){                    
                    this.subscribeToData()
                }else{                                      
                    element.Data.housesdata.forEach(element => {                        
                        this.allHouses.push(element)
                    });                        
                } 
            }
        }
    }
    private subscribeToData(): void {
        this.timerSubscription = Observable.timer(2000)
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
   private goToDeploy(){
        this.router.navigateByUrl("dash/deploy")
    }
}

class BuildThreeData{
    key:string
    housename:string
    slotname:string
    main:Array<Object> = []
}