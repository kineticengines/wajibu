import { Component,OnInit } from "@angular/core"

@Component({
    selector:"not-found",
    template:`
        <div>
            <h1>Resource not found</h1>
        </div>
    `
})

export class DashNotFoundComponent implements OnInit{
    constructor(){}
    ngOnInit(){
        
    }
}