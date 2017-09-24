import { Component,OnInit } from "@angular/core";
import { FormBuilder, FormGroup} from "@angular/forms";
import {MdSnackBar,MdDialog} from "@angular/material";
import { SearchResultDialogComponent } from "../search-result-dialog/search.result.dialog";
import { trigger, style,animate,transition,keyframes} from "@angular/animations";

import { Initializer } from "../../../../services/init.service"
import { FilterData,CacheItem } from "../../../../services/util";
import { Observable } from 'rxjs/Rx';
import "rxjs/add/operator/startWith";
import "rxjs/add/operator/map";

@Component({
    selector:"top-box",
    template:`
        <nav class="container" [@topBoxAnim]>
            <div class="container-item">
                <!--<h2>Wajibu</h2>-->
                <a [routerLink]="['/home/dash']"><img src="./assets/logo.png" class="logo"></a> 
            </div>
            <div class="container-item">
                <form [formGroup]="searchForm" novalidate>
                    <md-input-container class="search">
                        <input type="search" mdInput formControlName="searchitem"
                         placeholder="Search Wajibu by title, pillar or location" [mdAutocomplete]="auto">                        
                    </md-input-container>   
                    <md-autocomplete #auto="mdAutocomplete">
                        <md-option *ngFor="let option of filteredOptions | async" [value]="option" style="color:teal;font-weight:bold;">
                         {{option | titlecase }}
                        </md-option>
                    </md-autocomplete>              
                </form>                                      
            </div>
            <div class="container-item">
                <div class="container-item-social">
                    <a class="container-item-social-item link" [routerLink]="['stats']">Statistical</a>
                    <a href="https://facebook.com" target="_blank" class="container-item-social-item link">Facebook</a>
                    <a href="https://twitter.com" target="_blank" class="container-item-social-item link">Twitter</a>
                </div>                
            </div>
        </nav>
        
    `,
    styleUrls:["./top.box.component.css"],
    animations:[
        trigger('topBoxAnim',[
            transition(':enter',[                           
                animate('1.1s ease-in',keyframes([
                    style({opacity: 0,offset:0}),
                    style({opacity: 0.3,offset:0.5}),
                    style({opacity: 0.6,offset:0.7}),
                    style({opacity: 1,offset:1}),
                ]))
            ]),
            transition(':leave',[
                style({transform:'translateX(150px,25px)'}),
                animate(350)
            ])
        ])
    ]
})

export class TopBoxComponent implements OnInit{
    searchForm:FormGroup
    
    filteredOptions:Observable<string[]>
    constructor(private formBuilder:FormBuilder,private init:Initializer,
        private snackbar:MdSnackBar,private resultDialog:MdDialog){ }

    ngOnInit(){
        
        this.searchForm = this.formBuilder.group({
            searchitem : ['']           
        })                 
        
        this.filteredOptions = this.searchForm.controls.searchitem.valueChanges  
             .startWith(null)          
             .switchMap(val => this.init.getCachedItems())
             .map(val =>{
                 let s = []
                 val.forEach((el:string) => {
                     if(this.searchForm.controls.searchitem.value.length >= 1){
                        if(el.indexOf(this.searchForm.controls.searchitem.value) === 0){
                            s.push(el)
                        }
                     }
                                   
                 });
                 return s
             })                       
        
        let s = this.searchForm.controls.searchitem.valueChanges
            .debounceTime(300)  
            .map(val => {
                let f = new FilterData()
                f.Item = val.toLowerCase()
                f.Type = "main-content"
                return f
            })          
            .switchMap(val => this.init.getFilterContentData(val))
            .subscribe(val => {                
                if (val.status === true){
                    this.cacheQuery(this.searchForm.controls.searchitem.value)
                    let adj:string = val.content.length <= 1 ? "result" : "results";
                    let msg:string = `${val.content.length} ${adj} found`
                    this.snackbar.open(msg,"View")
                    this.snackbar._openedSnackBarRef.onAction().subscribe(() =>{
                        this.resultDialog.open(SearchResultDialogComponent,{
                            height:'35rem',width:'65rem',data:val.content.content
                        })                        
                    })
                }else if (val.status === false){
                    this.snackbar.open("Nothing found","Close",{duration:1000})                    
                }                
            })
            
    
    }

    cacheQuery(item:string){
        let f = new CacheItem()
        f.Item = item
        this.init.cacheTheQuery(f)
                  .subscribe(val => {return val})
    }
    
}
