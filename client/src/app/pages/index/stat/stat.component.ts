import { Component,OnInit } from "@angular/core";
import { trigger, style,animate,transition,keyframes} from "@angular/animations";

@Component({
    moduleId:module.id,
    selector:"stat",
    template:`
        <div class="container stats-box" [@statBoxAnim]>
            <div class="container-item header">
                <a class="header-item" [routerLink]="['.']">Pillar</a>
                <a class="header-item" [routerLink]="['./gender']">Gender</a>               
                <a class="header-item" [routerLink]="['./nth']">nTh Comparison</a>                
            </div>
            <div class="container-item body">
                <router-outlet></router-outlet>
            </div>
        </div>
    `,
    styleUrls:["./stat.component.css"],
    animations:[
        trigger('statBoxAnim',[
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
export class StatComponent implements OnInit{
    constructor(){}
    ngOnInit(){}
}