/*
Wajibu is an online web app that collects,analyses, aggregates and visualizes sentiments
from the public pertaining the government of a nation. This tool allows citizens to contribute
to the governance talk by airing out their honest views about the state of the nation and in
particular the people placed in government or leadership positions..

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


import { Component ,OnInit,OnDestroy} from "@angular/core";
import {MdSnackBar } from "@angular/material";
import { Router } from "@angular/router";
import { FormBuilder, FormGroup, Validators,FormControl} from "@angular/forms";
import { Initializer } from "../../../../../../services/init.service";
import { isNumber,isNotZero } from "../../../../../../services/util.service";
import { CustomValidators } from 'ng2-validation';


import { Observable } from 'rxjs/Rx';
import { AnonymousSubscription } from "rxjs/Subscription"

@Component({
    selector:"dash-step-grassroot",
    template:`
        <h2>Step Four</h2>
        <h4 *ngIf="repsBuild.length < numofseats">{{houseName | titlecase}} of {{subgovname | titlecase}} {{govname | titlecase}}</h4>
        <h4 *ngIf="repsBuild.length < numofseats">Number of seats remaining : {{numofseats - repsBuild.length}}</h4>
        <Section align="center">
            <form [formGroup]="repsForm" novalidate class="general-form" *ngIf="repsBuild.length < numofseats">
                <md-input-container class="general-form-input">
                    <input  mdInput type="text"  formControlName="RepName" placeholder="{{repTitle | titlecase}} Name" class="the-input">                    
                </md-input-container><br><br>
                <md-input-container class="general-form-input">
                    <input  mdInput type="text" formControlName="RepTerm" placeholder="{{repTitle | titlecase}} Term" class="the-input">                    
                </md-input-container><br><br>
                <md-select class="general-form-input" placeholder="{{repTitle | titlecase}} Gender" formControlName="RepGender">
                    <md-option *ngFor="let type of genderTypes"  [value]="type.typeName">{{type.typeName}}</md-option>
                </md-select><br><br> 
                <md-input-container class="general-form-input">
                    <input  mdInput type="text" formControlName="RepnthPosition" placeholder="{{repTitle | titlecase}} nth Position" class="the-input">                    
                </md-input-container><br><br>
                <md-input-container class="general-form-input">
                    <input  mdInput type="url" formControlName="RepImage" placeholder="{{repTitle | titlecase}} Image URL" class="the-input-extra">                    
                </md-input-container><br><br>
                <md-input-container class="general-form-input">
                        <input  mdInput type="text" formControlName="SlotName" placeholder="{{repSlot | titlecase}} Name" class="the-input">                    
                </md-input-container><br><br>               
                <button type="button" md-raised-button class="save-button" (click)="addRep(repsForm)" *ngIf="repsForm.valid">ADD {{repTitle | uppercase}}</button>
            </form>
            <div>
                <table align="center" *ngIf="repsBuild.length >= 1">
                    <thead>                
                        <tr>
                            <th>{{repTitle | titlecase}} Name</th>
                            <th>{{repTitle | titlecase}} Term</th>
                            <th>{{repTitle | titlecase}} Gender</th>
                            <th>{{repTitle | titlecase}} nth Position</th>                         
                            <th>{{repTitle | titlecase}} Image URL</th>
                            <th>{{repSlot | titlecase}} Name</th>                            
                            <th>Action</th>
                        <tr>
                    </thead>
                    <tbody>
                        <tr *ngFor="let rep of repsBuild; let i = index" [attr.data-index]="i" >
                            <td>{{rep.RepName | titlecase}}</td>
                            <td>{{rep.RepTerm | titlecase}}</td>
                            <td>{{rep.RepGender | titlecase}}</td>
                             <td>{{rep.RepnthPosition | lowercase}}</td>
                            <td>{{rep.RepImage}}</td>
                            <td>{{rep.SlotName | titlecase}}</td>                           
                            <td><button md-raised-button type="button" class="remove-button" (click)="removeRepBuild(i)">REMOVE</button></td>
                        </tr>
                    </tbody>
                </table>
                <button type="button" md-raised-button  *ngIf="repsBuild.length >= 1 && repsBuild.length == numofseats" class="save-button" (click)="saveAll()">SAVE ALL {{repTitle | uppercase}}s</button> 
             </div>
             <div>
                <table align="center" *ngIf="levelData.length >= 1">
                    <thead>                
                        <tr>
                            <th>{{block | titlecase}}</th>
                            <th>{{blocktitle | titlecase}} Name</th>
                            <th>{{blockslot | titlecase}} Name</th>  
                            <th>{{blocktitle | titlecase}} Gender</th> 
                            <th>{{blocktitle | titlecase}} Term</th> 
                            <th>{{blocktitle | titlecase}} nth Position</th>                      
                            <th>{{blocktitle | titlecase}} Image URL</th>
                            
                        <tr>
                    </thead>
                    <tbody>
                        <tr *ngFor="let rep of levelData; let i = index" [attr.data-index]="i" >
                            <td>{{rep.from | titlecase}}</td>
                            <td>{{rep.repname | titlecase}}</td>
                            <td>{{rep.slotname | titlecase}}</td>
                            <td>{{rep.repgender | titlecase}}</td>
                            <td>{{rep.repterm | titlecase}}</td>
                            <td>{{rep.repnthposition | lowercase}}</td>
                            <td>{{rep.repimage | lowercase}}</td>
                                                        
                        </tr>
                        <button type="button" md-raised-button  class="remove-button" (click)="resetGrass()">RESET GRASSROOT LEVEL</button> 
                    </tbody>
                </table><br><br><br>
                <Section class="deploy-section" align="center" *ngIf="levelData.length >= 1">                    
                    <button type="button" class="deploy-button" (click)="goToDeploy()">DEPLOY WAJIBU</button><br>
                    <span>please confirm the inputs are correct before deploying</span>
                </Section>
                
             </div>
        </Section>
    `,
    styleUrls:["../dash.install.root.css"] 
})
export class DashInstallStepGrass implements OnInit{    
    private timerSubscription: AnonymousSubscription;
    public houseName:string;
    public repSlot:string;
    public repTitle:string;
    public govname:string;
    public numofseats:number;
    public subgovname:string;
    public key:string

    repsForm:FormGroup;
    genderTypes = [{typeName:"Male"},{typeName:"Female"}]
    repsBuild:Array<Object> = [];

    levelData:Array<AllReps> = []
    public block:string;
    public blocktitle:string;
    public blockslot:string;   
    

    constructor(private init:Initializer,private formBuilder:FormBuilder,
        private snackbar:MdSnackBar,private router:Router){ 
              
    }

    ngOnInit(){
       this.stepGrassRootInitializer()  
       //this.createForm()     
    }
    
    private stepGrassRootInitializer(){
        this.repsBuild.splice(0,this.repsBuild.length)
        this.levelData.splice(0,this.levelData.length) 
        this.init.buildFiveInitializer().subscribe(d => {                      
            if (d.iscomplete === false){
                this.houseName = d.housename;
                this.repSlot = d.repslot;
                this.repTitle = d.reptitle;
                this.govname = d.govname;
                this.key = d.govdata.govinkey;
                this.numofseats = d.govdata.build.numoflegseats;
                this.subgovname = d.govdata.build.slotname;
                this.createForm()
            }else if (d.iscomplete === true){
                //call populate data
                this.startFetcher()
            }
        })
    }
    startFetcher(){        
        this.init.checkBuildLevel().subscribe(d => this.processBuild(d))
    }

    private processBuild(d:any){                 
        for (var index = 0; index < d.levels.length; index++) {
            let element = d.levels[index]; 
            if(element.Level === "build:five"){                
                if(element.Exist === false && element.Complete === false){                    
                    this.subscribeToData()
                }else{                                                                                
                    element.Data.repsdata.forEach(e => { 
                        this.block = e.block  
                        this.blocktitle = e.blocktitle
                        this.blockslot = e.blockslot

                        let d = new AllReps()
                        d.from = e.from                        
                        d.repgender = e.repdata.repgender
                        d.repimage = e.repdata.repimage
                        d.repname = e.repdata.repname
                        d.repnthposition = e.repdata.repnthposition
                        d.repterm = e.repdata.repterm
                        d.slotname = e.repdata.slotname

                        this.levelData.push(d)
                    });        
                                       
                } 
            }
        }
    }

    private subscribeToData(): void {
        this.timerSubscription = Observable.timer(2000)
        .subscribe(() => this.startFetcher());
    }

    private createForm(){
        this.repsForm = this.formBuilder.group({
            RepName : ['',[Validators.required]],
            RepTerm : ['',[Validators.required,isNumber,isNotZero]],
            RepGender : ['',[Validators.required]], 
            RepnthPosition : ['',[Validators.required,isNumber,isNotZero]],
            RepImage : ['',[Validators.required,CustomValidators.url]], 
            SlotName :  ['',[Validators.required]] 
        })
    }

    private addRep(inForm:any){
        if(inForm.valid){
            this.repsBuild.push(inForm.value)
            this.repsForm.reset()
        }
    }

    private removeRepBuild(index:any){
        if(index > -1){
            this.repsBuild.splice(index,1)
        }
    }

    private saveAll(){
        let d = new LegData()
        d.Key = this.key
        d.TheData = this.repsBuild        
        this.init.goProcessStepFive(d).subscribe(d => {                  
            if(d.Status === true){
                window.location.reload() //TODO => review logic
            }
        })
    }

    private resetGrass(){
        this.init.resetBuildFive().subscribe(d =>{
            if (d.Status === true){                
                this.ngOnInit()
            }
        })
    }

    private goToDeploy(){
        this.init.checkBuildLevel().subscribe().unsubscribe()
        this.router.navigateByUrl("dash/deploy")
    }
}

class LegData{
    Key:string
    TheData:any
}

class AllReps{
    from:string    
    repgender:string
    repimage:string
    repname:string
    repnthposition:string
    repterm:string
    slotname:string

}