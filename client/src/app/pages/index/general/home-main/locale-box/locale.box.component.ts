import { Component } from "@angular/core";
import { trigger, style,animate,transition,keyframes} from "@angular/animations";
@Component({
    moduleId:module.id,
    selector:"locale-box",
    template:`
        <div class="container" [@localeBoxAnim]>
            <div class="container-item left-box">  
                <small>Locale Details</small>              
                <div class="left-box-item">
                    <p>Country : <small class="detail-item">1</small></p>
                </div>
                <div class="left-box-item">
                    <p>Government Type : <small class="detail-item">1</small></p>
                </div>
                <div class="left-box-item">
                    <p>Number of Houses : <small class="detail-item">1</small></p>
                </div>
                <div class="left-box-item">
                    <p>Number of subgovernments : <small class="detail-item">1</small></p>
                </div>
            </div>
            <div class="container-item right-box">
                <div class="chart-item">
                    <small>Representatives Gender Distribution</small>
                    <canvas baseChart
                        [data]="dougData"                        
                        [labels]="dougLabels"                                       
                        [chartType]="dougType" height="140rem" width="140rem"></canvas>
                </div>
            </div>
        </div>
    `,
    styleUrls:["./locale.box.component.css"],
     animations:[
        trigger('localeBoxAnim',[
            transition(':enter',[                           
                animate('0.3s ease-in',keyframes([
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
export class LocaleBoxComponent{
    public dougLabels:string[] = ['Female', 'Male'];
    public dougData:number[] = [70,30];    
    public dougType:string = 'doughnut';    
}