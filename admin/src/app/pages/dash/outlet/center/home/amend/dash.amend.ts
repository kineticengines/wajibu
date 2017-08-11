import { Component,OnInit } from "@angular/core";

@Component({
    moduleId:module.id,
    selector:"amend",
    template:`
        <button md-icon-button [mdMenuTriggerFor]="menu"><md-icon>menu</md-icon></button>
            <md-menu #menu="mdMenu">
                <button md-menu-item (click)="showPillarsForm()">Pillars</button>
                <button md-menu-item (click)="showSlotsForm()">Slots</button>
        </md-menu><br><br>
        <div class="container">  
            <section *ngIf="showPillars">
                <amend-pillars class="item"></amend-pillars>
            </section> 
            <section *ngIf="showSlots">
                <h3>Slots pages</h3>
            </section>                                 
        </div>
    `,
    styleUrls:["./amend.css"]
})
export class AmendComponent implements OnInit{
    showPillars:boolean = true;
    showSlots:boolean
    constructor(){}

    ngOnInit(){

    }
    showPillarsForm(){
        this.showSlots = false;
        this.showPillars = true;
    }
    showSlotsForm(){        
        this.showPillars = false;
        this.showSlots = true;
    }
}