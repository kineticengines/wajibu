import { Component,OnInit,OnDestroy} from "@angular/core";
import {ActivatedRoute}	from	'@angular/router';

@Component({
    selector:"filter",
    template:`
        <h3>filtered</h3>
    `
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