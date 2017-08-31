import { Component,OnInit } from "@angular/core";
import { Initializer } from "../../../../services/init.service";

@Component({
    selector:"filter-top-box",
    template:`
        <div class="container">
            <div class="container-item">
                <img src="./assets/logo.png" class="logo"> 
            </div>
            <div class="container-item">
                <ul class="tags" *ngFor="let item of items" fxLayout="row" fxLayoutWrap="wrap">
                    <li><a>Binary planet</a></li>  
                    <li><a>Binary planet</a></li>                                 
                </ul>
            </div>
        </div>
    `,
    styleUrls:["./filter.top.box.component.css"]
})
export class FilterTopBoxComponent implements OnInit{
    items:Array<string> = [];    

    constructor(private init:Initializer){
        this.init.getTitles().subscribe(d => {
            d.titles.forEach(element => {
                this.items.push(element)
            });
        })

        this.init.getPillars().subscribe(d =>{            
            d.pillars.forEach(e => {
                this.items.push(e.pillar)
            });
        })
    }

    ngOnInit(){

    }
}