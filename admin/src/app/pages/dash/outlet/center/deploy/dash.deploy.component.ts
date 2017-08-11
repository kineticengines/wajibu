import { Component,OnInit } from "@angular/core";
import { Router } from "@angular/router";
import {Location} from "@angular/common";
import { Initializer } from "../../../../../services/init.service";
import { SessionChecker } from "../../../../../services/session.checker";

import { Observable } from 'rxjs/Rx';
import { AnonymousSubscription } from "rxjs/Subscription"


@Component({
    selector:"dash-deploy",
    template:`
        <div align="center">
          <h3>Deploying Wajibu</h3>          
        </div>
        <div>
            <h5>Progress...</h5>
            <md-progress-bar mode="determinate" [value]="percent"></md-progress-bar> 
        </div>           
        <!--<div class="deploy-page-container">            
            <div>
               <h2>step here</h2> 
            </div>
        </div>-->

        
    `,
    styleUrls:["./dash.deploy.component.css"]
})
export class DashDeploy implements OnInit{
     private timerSubscription: AnonymousSubscription;
     public percent:number;
    constructor(private router:Router,private location: Location,private init:Initializer,
    private session:SessionChecker){
    }

    ngOnInit(){
      this.location.subscribe((v)=>{
        if(v.pop){
            this.location.forward()
        }
      }) 
      this.deploy() 
      this.checkdeploy()    
      
    }

    private deploy(){
        this.init.goStartDeployProcess().subscribe(d => d)        
    }
    private checkdeploy(){
        this.init.goCheckDeployProcess().subscribe(d =>{  
            this.percent = d.Status.percent               
            if(d.Status.complete === false){
                //call subscription to check deploy
                this.subscribeToData()
            }else if(d.Status.complete === true){                
                this.init.updateConfig().subscribe(d => {                    
                    if (d.Status === true){
                        this.session.setDeployed(true)
                        setTimeout(()=>{
                            this.router.navigate(["/dash"])
                        },1000)                        
                    }
                })
            }
        })
    }

    private subscribeToData(): void {
        this.timerSubscription = Observable.timer(50)
            .subscribe(() => this.checkdeploy());
    }
    
}