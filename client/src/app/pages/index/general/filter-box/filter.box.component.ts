import { Component,OnInit } from "@angular/core"
import { trigger, style,animate,transition,keyframes} from "@angular/animations";
import { Initializer } from "../../../../services/init.service";

@Component({
    moduleId:module.id,
    selector:"filter-box",
    template:`
        <div class="container list-box" [@filterBoxAnim]>
            <div class="container-item">
                <h4>Filter by;</h4>
            <div>
            <div class="container-item">
                <ul *ngFor="let title of titles">
                    <li><a 	[routerLink]="['../home/fan/',title]">{{title | titlecase}}</a></li>                    
                </ul>
            <div>
        </div>    
            
    `,
    styleUrls:["./filter.box.component.css"],
    animations:[
        trigger('filterBoxAnim',[
            transition(':enter',[                           
                animate('0.9s ease-in',keyframes([
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

export class FilterBoxComponent implements OnInit{
    titles:Array<string> = [];
    constructor(private init:Initializer){}
    ngOnInit(){
        this.init.getTitles().subscribe(d => {
            d.titles.forEach(element => {
                this.titles.push(element)
            });
        })
    }

    
}