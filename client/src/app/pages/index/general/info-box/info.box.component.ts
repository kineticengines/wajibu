import { Component} from "@angular/core";
import { trigger, style,animate,transition,keyframes} from "@angular/animations";

@Component({
    moduleId:module.id,
    selector:"info-box",
    template:`
        <div class="container info-box" [@infoBoxAnim]>
            <div class="container-item">                
                <p>
                    Wajibu is an online tool that collects and agggerates sentiments from the
                    public regarding the governance of a nation in an effort to provide a 
                    statistical analysis on the kind of leaders a country has and how such 
                    leaders contribute to political,economical and social development of a nation.
                </p>
            <div>
            <div class="container-item">
               <p>Support the efforts of Wajibu by sending your donation via Mpesa Paybill number
               12345</p>
               <p>Help us shape a better Kenya for present and future generations.</p>
            <div>
        </div> 
    `,
    styleUrls:["./info.box.component.css"],
    animations:[
        trigger('infoBoxAnim',[
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
export class InfoBoxComponent{

}