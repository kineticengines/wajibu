import { Component,OnInit } from "@angular/core"
import { trigger, style,animate,transition,keyframes} from "@angular/animations";
import { Initializer } from "../../../../services/init.service";

@Component({
    moduleId:module.id,
    selector:"filter-box",
    template:` 
                
        <div class="container list-box" [@filterBoxAnim]>            
            <div class="container-item">   
                 <h5>Titles</h5>              
                <ul>
                    <li *ngFor="let title of titles"><a	[routerLink]="['fan/',title]">{{title | titlecase}}</a></li>                    
                </ul>
            </div>
            <div class="container-item">  
                <h5>Pillars</h5>              
                <ul>
                    <li *ngFor="let pillar of pillars"><a [routerLink]="['fan/',pillar]">{{pillar | titlecase}}</a></li>                    
                </ul>
            </div>
        </div> 
        
            
    `,
    styleUrls:["./filter.box.component.css"],
    animations:[
        trigger('filterBoxAnim',[
            transition(':enter',[                           
                animate('0.9s ease-in',keyframes([
                    style({opacity: 0,offset:0}),
                    style({opacity: 0.3,offset:0.5}),
                    style({opacity: 0.6,offset:0.7}),
                    style({opacity: 1,offset:1}),
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
    pillars:Array<string> = [];
    constructor(private init:Initializer){}
    ngOnInit(){
        this.init.getTitles().subscribe(d => {
            this.titles.splice(0,this.titles.length)
            d.titles.forEach(element => {
                this.titles.push(element)
            });
        })

        this.init.getPillars().subscribe(d =>{   
            this.pillars.splice(0,this.pillars.length)         
            d.pillars.forEach(e => {
                this.pillars.push(e.pillar)
            });
        })
    }

    
}