import { Component,OnInit,OnDestroy,ChangeDetectorRef} from "@angular/core";
import { trigger, style,animate,transition,keyframes} from "@angular/animations";
import {ActivatedRoute,Router}	from	'@angular/router';
import { Initializer } from "../../../services/init.service";
import { FilterData } from "../../../services/util";

import { Observable } from "rxjs/Observable";
import { Subject } from "rxjs/Subject";

@Component({
    moduleId:module.id,
    selector:"filter",
    template:`
        <div class="container" [@filterBoxAnim]>   
            <div class="container-item container-top">
                <p class="container-top-item" >Filter by: {{filterBy | titlecase }}</p>
                <a class="container-top-item" [routerLink]="['../..']">Back</a>                
            </div>         
            <div class="container-item">
                <div class="main-content-box">
                    <div class="main-content-box-item oneth content-box">
                        <!--<div class="content-box-item content-box-oneth">
                            details box
                        </div>-->
                        <div class="content-box-item content-box-twoth">
                            <div *ngIf="dataNotFound">                                
                                <md-chip-list>
                                    <md-chip><span>Data not available</span></md-chip>                    
                                </md-chip-list> 
                            </div>
                            <div *ngIf="!dataNotFound" class="main-data-content-box">
                                <div class="main-data-content-box-item" [@filterBoxAnim]>
                                    <small *ngIf="responsesFound >= 1">{{responsesFound}}  {{responsesFound <= 1 ? "result found" : "results found"}}</small>
                                    <md-chip-list *ngIf="responsesFound == 0">
                                        <md-chip><span>Error occured.Please try again later</span></md-chip>                    
                                    </md-chip-list>
                                </div>
                                <div class="main-data-content-box-item">
                                    <div *ngFor="let dd of contentData | async" [@filterBoxAnim]>
                                        <div class="content-data">
                                            <div class="content-data-item content-data-item-top">
                                                <div>                                                    
                                                    <small><em style="color:teal;">Name</em> : {{dd.name | titlecase}}</small>
                                                 </div>
                                                 <div>                                                    
                                                    <small><em style="color:teal;">Title</em> : {{dd.title | titlecase}}</small>
                                                 </div>                                                 
                                            </div>
                                            <div class="content-data-item content-data-item-bottom">
                                                <div *ngFor="let data of dd.data" class="sentiment-box">
                                                    <div class="sentiment-box-item the-sentiment">
                                                        <p>{{data.sentiment}}</p>
                                                    </div>
                                                    <div class="sentiment-box-item sentiment-box-item-bottom">
                                                        <div class="sentiment-box-item-bottom-item">                                                            
                                                            <small><em style="color:teal;">Pillar</em> : {{data.pillars | titlecase}}</small>
                                                        </div>
                                                        <div class="sentiment-box-item-bottom-item" *ngIf="data.county">                                                            
                                                            <small><em style="color:teal;">County</em> : {{data.county | titlecase}}</small> 
                                                        </div>
                                                        <div class="sentiment-box-item-bottom-item" *ngIf="data.constituency">                                                            
                                                            <small><em style="color:teal;">Constituency</em> : {{data.constituency | titlecase}}</small>
                                                        </div>
                                                        <div class="sentiment-box-item-bottom-item" *ngIf="data.ward">                                                            
                                                             <small><em style="color:teal;">Ward</em> : {{data.ward | titlecase}}</small> 
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>                                        
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="main-content-box-item twoth">
                     charts
                    </div>
                </div>                
            </div>
        </div>
    `,
    styleUrls:["./filter.component.css"],
    animations:[
        trigger('filterBoxAnim',[
            transition(':enter',[                           
                animate('0.9s ease-in',keyframes([
                    style({opacity: 0,offset:0}),
                    style({opacity: 0.3,offset:0.5}),
                    style({opacity: 0.6,offset:0.7}),
                    style({opacity: 1,offset:1}),
                ]))
            ])            
        ])
    ]
})
export class FilterComponent implements OnInit, OnDestroy{
   sub:any;
   dataNotFound:boolean;
   responsesFound:number;
   filterBy:any;
   private contentStream:Observable<any>; 
   private contentData:Subject<any>; 
    
   constructor(private aroute:ActivatedRoute,private init:Initializer,private router:Router,
    private chng:ChangeDetectorRef){
        this.contentData = new Subject()
   } 

   ngOnInit(){
       this.sub = this.aroute.params.subscribe(params => {            
            this.getTheData(params["who"]) 
            this.filterBy = params["who"];
       })        
   }

   ngOnDestroy(){
     this.sub.unsubscribe();
   }

   getTheData(val:string){
      this.mainContent(val)
      this.detailsContent(val)  
   }

   mainContent(val:string){
       let f = new FilterData()
       f.Item = val.toLowerCase()
       f.Type = "main-content"       
       this.init.getFilterContentData(f).subscribe(res =>{                              
           switch (res.status) {
               case false:
                   this.dataNotFound = !res.status;
                   this.chng.detectChanges()
                   break;           
               case true:
                   this.dataNotFound = !res.status; 
                   this.chng.detectChanges()
                   break;
           }
           this.responsesFound = res.content.length;
           this.chng.detectChanges()
           this.contentData.next(res.content.content) 
       })       
       
   }

   detailsContent(val:string){
       let f = new FilterData()
       f.Item = val.toLowerCase()
       f.Type = "details-content"
   }
   

}

