import { Component, ViewChild, OnInit } from '@angular/core';
import {MdSnackBar } from "@angular/material";
import { FormBuilder, FormGroup, Validators,FormControl} from "@angular/forms";
import{ DashService } from "../../../../../../services/dash.service";

import { FieldConfig } from '../../../../../../dynamic-form/models/field-config.interface';
import { DynamicFormComponent } from '../../../../../../dynamic-form/containers/dynamic-form/dynamic-form.component';
@Component({
    moduleId:module.id,
    selector:"sentiment",
    template:`
        <div>
            <Section>
                <h3>Choose Sentiment Level</h3>
                <md-radio-group class="top-radio-group" (change)="choiceChange($event.value)">                   
                    <md-radio-button class="top-radio-group-item" *ngFor="let level of levels" [value]="level">
                        Show {{level}}
                    </md-radio-button>                    
                </md-radio-group><br>
                <md-radio-group class="top-radio-group" *ngIf="showTopForm" (change)="choiceTopChange($event.value)">
                    <md-radio-button class="top-radio-group-item" *ngFor="let level of toplevels" [value]="level">
                        Add Sentiment For  {{level}}
                    </md-radio-button>  
                </md-radio-group>
            </Section>
            <Section>
                <md-chip-list>
                    <md-chip *ngIf="presConfig.length <= 0"> Error : Could not initilize president level data</md-chip>
                    <md-chip *ngIf="dpresConfig.length <= 0"> Error : Could not initilize deputy president level data</md-chip>
                    <md-chip>Error : could not initilize house level data</md-chip>
                    <md-chip>Error : Could not initilize subgovernment level data</md-chip>
                     <md-chip>Error : Could not initilize grassroot level data</md-chip>
                </md-chip-list> 
            </Section><br><br> 
            <Section><md-spinner *ngIf="processingSubmit"></md-spinner></Section><br>                                   
            <Section *ngIf="showPres">
                <dynamic-form [config]="presConfig"  #form="dynamicForm" (submit)="submit($event)"></dynamic-form>            
            </Section>
            <Section *ngIf="showDPres">
                <dynamic-form [config]="dpresConfig"  #form="dynamicForm" (submit)="submit($event)"></dynamic-form>                        
            </Section>                 
            
        </div>
    `,
    styleUrls:["./dash.sentiment.css"]
})
export class SentimentComponent implements OnInit{
     @ViewChild(DynamicFormComponent) form: DynamicFormComponent;

    topForm:FormGroup;
    houseForm:FormGroup
    subForm:FormGroup
    grassForm:FormGroup
    processingSubmit:boolean = false;
   
    showTopForm:boolean
    showHouseForm:boolean
    showSubForm:boolean
    showGrassForm:boolean

    showPres:boolean
    showDPres:boolean

    levels = ['Top Level','House Level','Subgovernment Level','Grassroot Level',];
    toplevels = ['President','Deputy President']

    presConfig:Array<FieldConfig> = []
    dpresConfig:Array<FieldConfig> = []
    

    constructor(private formBuilder:FormBuilder,private dash:DashService,private snackbar:MdSnackBar){
    }
    
    ngOnInit(){
        this.dash.presidentLevelConfigure().subscribe(d =>{            
            d.config.config.forEach(element => {
                this.presConfig.push(element)                
            });      
        })

        this.dash.dpresidentLevelConfigure().subscribe(d =>{
            d.config.config.forEach(element => {
                this.dpresConfig.push(element)                
            });
        })
    }    

    choiceChange(value:any){
        switch (value) {
            case this.levels[0]:
                    this.showTopForm = true
                    this.showPres = false
                    this.showDPres = false
                break;
            case this.levels[1]:
                    this.showTopForm = false
                    this.showPres = false
                    this.showDPres = false
                break;
            case this.levels[1]:
                    this.showTopForm = false
                    this.showPres = false
                    this.showDPres = false
                break;
            case this.levels[3]:
                    this.showTopForm = false;
                    this.showPres = false
                    this.showDPres = false
                break;        
            default:
                this.showTopForm = false;
                this.showPres = false
                this.showDPres = false
                break;
        }
    }  
    
    choiceTopChange(value:any){        
        switch (value) {
            case "President":
                this.showPres = true
                this.showDPres = false
                break;        
            case "Deputy President":
                this.showPres = false
                this.showDPres = true
                break;
        }
    }    
    submit(value: {[name: string]: any}) {  
        this.processingSubmit = true; 
        let count:number = 0;    
        let keys = Object.keys(value.thedata).filter((v)=> { return v !== ""})
        keys.forEach(key => {
            let val = value.thedata[key]           
            switch (val) {
                case undefined:
                    let msg:string = "Error: required inputs not entered";
                    this.snackbar.open(msg,"Close",{duration:2500})
                    break;            
                default:
                    count++
                    break;
            }
        });
        if (count === keys.length){            
            Object.getOwnPropertyNames(value.thedata).forEach(element => {
                if (element === ""){                    
                    delete value.thedata[element]
                }
            });
            let s = new Sentiment()
            s.api = value.forwho
            s.data = value.thedata
            this.dash.addSentiment(s).subscribe(d => {                
                if(d.status === true){
                    this.processingSubmit = false;                   
                    this.snackbar.open("Sentiment added.","Close",{duration:2500})
                }else if (d.status === false){
                   this.processingSubmit = false;
                   this.snackbar.open("Error occured.","Close",{duration:2500}) 
                }
            })
            
        }
    }
}

class Sentiment{
    api:string
    data:Object
}