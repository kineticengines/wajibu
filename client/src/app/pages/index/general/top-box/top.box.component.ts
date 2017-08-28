import { Component } from "@angular/core"
import { trigger, style,animate,transition,keyframes} from "@angular/animations";

@Component({
    selector:"top-box",
    template:`
        <div class="container" [@topBoxAnim]>
            <div class="container-item">
                <h2>Wajibu</h2>
            </div>
            <div class="container-item">
                <md-input-container class="search">
                    <input type="search" mdInput placeholder="Search Wajibu">
                </md-input-container>                
            </div>
            <div class="container-item">
                <div class="container-item-social">
                    <a href="https://facebook.com" target="_blank" class="container-item-social-item link">Facebook</a>
                    <a href="https://twitter.com" target="_blank" class="container-item-social-item link">Twitter</a>
                </div>                
            </div>
        </div>
        
    `,
    styleUrls:["./top.box.component.css"],
    animations:[
        trigger('topBoxAnim',[
            transition(':enter',[                           
                animate('1.1s ease-in',keyframes([
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

export class TopBoxComponent{

}