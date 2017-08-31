import { Component , OnInit} from "@angular/core";
import { DashService } from "../../../../../../services/dash.service";
import { FormBuilder,FormGroup,Validators } from "@angular/forms";
import {MdSnackBar} from '@angular/material';

import { Observable } from 'rxjs/Rx';
import { AnonymousSubscription } from "rxjs/Subscription"

@Component({
    moduleId:module.id,
    selector:"amend-pillars",
    template:`
        <div class="pillars-container">                       
            <div class="pillars-item">
                <Section *ngIf="!isPillars">
                    <md-chip-list>
                        <md-chip><span>Error : pillars not found</span></md-chip>                    
                    </md-chip-list> 
                </Section><br> 
                <Section>
                    <form [formGroup]="addForm" novalidate class="general-form" (submit)="savePillar(addForm.value)">
                        <md-input-container class="general-form-input">
                            <input mdInput placeholder="Pillar" value="" formControlName="pillar" class="the-input">
                        </md-input-container>
                        <md-select class="general-form-input" formControlName="for" placeholder="For">
                            <md-option *ngFor="let title of titles" [value]="title"  class="the-input">{{title }}</md-option>
                        </md-select><br><br>
                         <button type="submit" md-raised-button class="save-button" *ngIf="addForm.valid">SAVE</button><br>
                    </form>
                </Section>
            </div>
            <Section class="pillars-item" *ngIf="isPillars">
                <div>
                    <h4 align="right">Numbers of current pillars : {{numOfPillars}}</h4>
                </div>
                <div class="all-pillars-container">                    
                    <md-card *ngFor="let pillar of pillars" class="all-pillars-container-item">
                        <md-card-content>
                            <h3>{{pillar.pillar | titlecase}}</h3>
                            <h3>For : {{pillar.fortitle | titlecase}}</h3>
                        </md-card-content>
                        <md-card-actions>
                            <button md-button (click)="removePillar(pillar.pillar)">REMOVE</button>                            
                        </md-card-actions>
                    </md-card>
                </div>
            </Section>
        </div>      
    `,
    styleUrls:["./amend.css"]
})

export class AmendPillarsComponent implements OnInit{
    private timerSubscription: AnonymousSubscription;

    isPillars:boolean;
    addForm:FormGroup
    numOfPillars:number;
    pillars:Array<Object> = []
    titles:Array<string> = []
    constructor(private dash:DashService,private fb:FormBuilder,public snackBar: MdSnackBar){

    }
    ngOnInit(){
        this.getPillars()
        this.createForm()
    }

    getPillars(){
        this.titles.splice(0,this.titles.length)
        this.dash.fetchTitles().subscribe(d =>{   
            this.titles.push("All")         
            d.titles.forEach(element =>{
                this.titles.push(element)
            })
        })
        this.dash.fetchPillars().subscribe(d =>{                  
            if(d.error === null || d.error === undefined || d.error === ""){
                if(d.pillars === null){
                    this.isPillars = false;                
                }else{                
                    this.isPillars = true;
                    this.numOfPillars = d.pillars.length;
                    d.pillars.forEach(element => {
                        this.pillars.push(element)
                    });                                      
                }
            }else{
               this.isPillars = false; 
            }
        })
    }

    createForm(){
        this.addForm = this.fb.group({
            pillar:['',[Validators.required]],
            for:['',[Validators.required]]
        })
    }

    savePillar(p:any){
        let obj = { Pillar : p.pillar,For:p.for}        
        this.dash.savePillar(obj).subscribe(d =>{
            if(d.status === true){
                this.snackBar.open("Added Successfully", "", {
                     duration: 2000,
                });                                
                this.subscribeToData() 
                this.addForm.reset()
                
            }else{
               this.snackBar.open("Error Occured", "", {
                     duration: 2000,
                }); 
            }
        })
    }

    removePillar(p:string){
        let obj = { Pillar : p}
        this.dash.removePillar(obj).subscribe(d =>{
            if(d.status === true){
                this.snackBar.open("Removed Successfully", "", {
                     duration: 2000,
                });
                this.addForm.reset()  
                this.subscribeToData()  
            }else{
               this.snackBar.open("Error Occured", "", {
                     duration: 2000,
                }); 
            }
        }) 
    }

    private subscribeToData(): void {
        this.pillars.splice(0,this.pillars.length)
        this.timerSubscription = Observable.timer(100)
        .subscribe(() => this.getPillars());
    }
}
