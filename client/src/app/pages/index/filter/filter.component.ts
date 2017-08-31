import { Component,OnInit,OnDestroy} from "@angular/core";
import {ActivatedRoute}	from	'@angular/router';

@Component({
    moduleId:module.id,
    selector:"filter",
    template:`
        <div class="container">
            <div class="container-item">
                <filter-top-box></filter-top-box>
            </div>
            <div class="container-item">
                <h3>main content</h3>
            </div>
        </div>
    `,
    styleUrls:["./filter.component.css"]
})
export class FilterComponent implements OnInit, OnDestroy{
   sub:any; 
   who:string
   constructor(private route:ActivatedRoute){
   } 

   ngOnInit(){
       this.sub = this.route.params.subscribe(params => {
            this.who = params["who"]
       })  
        
       console.log(this.who)
   }
   ngOnDestroy(){
     this.sub.unsubscribe();
   }
}