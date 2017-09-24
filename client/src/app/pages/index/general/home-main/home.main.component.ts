import { Component } from "@angular/core";

@Component({
    moduleId:module.id,
    selector:"home-main",
    template:`        
        <div class="container">
            <div class="container-item oneth">
                <sentiment-box></sentiment-box>           
                <locale-box></locale-box>           
            </div>
            <div class="container-item twoth">
                <chart-box></chart-box>                
            </div>
        </div>
       
    `,
    styleUrls:["./home.main.component.css"]
})
export class HomeMainComponent{
    constructor(){

    }
}