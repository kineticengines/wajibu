import { Component,OnInit } from "@angular/core";
import { trigger, style,animate,transition,keyframes} from "@angular/animations";
import { Initializer } from "../../../../services/init.service"

@Component({
    moduleId:module.id,
    selector:"sentiment-box",
    template:`        
        <div class="container sentiment-box" [@sentimentBoxAnim]>
            <div class="container-item">
                <h4>Total sentiments : </h4>
            </div>
            <div class="container-item">
                <h2>other items here</h2>
            </div>
        </div>
    `,
    styleUrls:["./sentiment.box.component.css"],
    animations:[
        trigger('sentimentBoxAnim',[
            transition(':enter',[                           
                animate('0.6s ease-in',keyframes([
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
export class SentimentBoxComponent implements OnInit{
    constructor(private init:Initializer){}
    ngOnInit(){
        this.init.getSentiments().subscribe(d =>{
            console.log(d)
        })
    }
}