import { Component } from "@angular/core";
import { trigger, style,animate,transition,keyframes} from "@angular/animations";

@Component({
    selector:"nth",
    template:`
        <div [@statNthBoxAnim]>
             <h3>nth</h3>
        </div>
    `,
    animations:[
        trigger('statNthBoxAnim',[
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
export class StatNthComponent{

}