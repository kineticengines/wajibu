import { Component,OnInit } from "@angular/core";
import { trigger, style,animate,transition,keyframes} from "@angular/animations";
import { Initializer } from "../../../../services/init.service"

import { Observable } from 'rxjs/Rx';
import { AnonymousSubscription } from "rxjs/Subscription"

@Component({
    moduleId:module.id,
    selector:"sentiment-box",
    template:`        
        <div class="container sentiment-box" [@sentimentBoxAnim]>            
            <!--<div class="container-item">                               
                <h4>Total sentiments : {{sentiments.length}}</h4>
                <small>Live Stream</small>
            </div>-->
            <div class="container-item sentiment-container-box" >                
                <div *ngFor="let s of sentiments | async" class="sentiment-container-box-item sentiment-container">
                    <div class="sentiment-container-item">
                        <!--<img [src]="s.image">-->
                        <img src="./assets/photo.png" class="pillar-for">
                    </div>
                    <div class="sentiment-container-item">
                        <div class="sentiment-side-items">
                            <div class="sentiment-side-items-item">
                                <div class="sentiment-side-items-item-header">
                                    <div class="sentiment-side-items-item-header-item">
                                       <small><em style="color:teal;">Pillar</em> : {{s.pillars | titlecase}}</small> 
                                    </div>
                                    <div class="sentiment-side-items-item-header-item">
                                        <small><em style="color:teal;">Date</em> : {{s.date}}</small> 
                                    </div>                                                                    
                                </div>
                            </div>
                            <div class="sentiment-side-items-item">
                                <div class="sentiment-side-items-item-content">
                                    <div class="sentiment-side-items-item-content-item">
                                       <p>{{s.sentiment}} </p>                                       
                                    </div>
                                </div>
                            </div>
                            <div class="sentiment-side-items-item">
                                <div class="sentiment-side-items-item-footer">
                                    <div class="sentiment-side-items-item-footer-item" *ngIf="s.county">
                                       <small><em style="color:teal;">County</em> : {{s.county | titlecase}}</small> 
                                    </div>
                                    <div class="sentiment-side-items-item-footer-item" *ngIf="s.constituency">
                                        <small><em style="color:teal;">Constituency</em> : {{s.constituency | titlecase}}</small> 
                                    </div>
                                    <div class="sentiment-side-items-item-footer-item" *ngIf="s.ward">
                                        <small><em style="color:teal;">Ward</em> : {{s.ward | titlecase}}</small> 
                                    </div>                                    
                                </div>
                            </div>
                        </div>
                    </div>                                           
                </div>              
            </div>
        </div>        
    `,
    styleUrls:["./sentiment.box.component.css"],
    animations:[
        trigger('sentimentBoxAnim',[
            transition(':enter',[                           
                animate('0.6s ease-in',keyframes([
                    style({opacity: 0,transform:'translate3d(-50%,0,0)',offset:0}),
                    style({opacity: 0.3,transform:'translate3d(-100%,0, 0)',offset:0.5}),
                    style({opacity: 0.6,transform:'none',offset:0.7}),
                    style({opacity: 1,transform:'none',offset:1}),
                ]))
            ]),
            transition(':leave',[
                style({transform:'translateX(150px,25px)'}),
                animate(350)
            ])
        ])       
    ]
})
export class SentimentBoxComponent implements OnInit{
    private timerSubscription: AnonymousSubscription;
    sentiments:Observable<any>;
    newSentiments:Array<Object> = []
    
    constructor(private init:Initializer){}    
    ngOnInit(){
       this.startFetcher()
       
    }
    startFetcher(){     
        this.sentiments =  this.init.getSentiments()        
    }

    

   
}