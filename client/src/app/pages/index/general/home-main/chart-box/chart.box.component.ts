import { Component } from "@angular/core"
import { trigger, style,animate,transition,keyframes} from "@angular/animations";

@Component({
    moduleId:module.id,
    selector:"chart-box",
    template:`
        <div class="container chart-box" [@chatBoxAnim]>
            <div class="container-item">
                <div class="chart-item">
                    <canvas baseChart
                    [datasets]="barChartData"
                    [labels]="barChartLabels"
                    [options]="barChartOptions"
                    [legend]="barChartLegend"
                    [chartType]="barChartType"></canvas>
                </div>                           
            </div>
            <div class="container-item">
                <div class="chart-item">
                    <canvas baseChart
                        [data]="pieChartData"
                        [labels]="pieChartLabels"
                        [chartType]="pieChartType"></canvas>
                </div>               
            </div>
        </div>
        
    `,
    styleUrls:["./chart.box.component.css"],
    animations:[
        trigger('chatBoxAnim',[
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

export class ChartBoxComponent{
    public barChartOptions:any = {
        scaleShowVerticalLines: false,        
        responsive: true
    };
    public barChartLabels:string[] = ['2006', '2007', '2008', '2009', '2010', '2011', '2012'];
    public barChartType:string = 'bar';
    public barChartLegend:boolean = true;
    
    public barChartData:any[] = [
        {data: [65, 59, 80, 81, 56, 55, 40], label: 'Series A'},
        {data: [28, 48, 40, 19, 86, 27, 90], label: 'Series B'}
    ];



    public pieChartLabels:string[] = ['Download Sales', 'In-Store Sales','Others' ,'Mail Sales','Main Sales'];
    public pieChartData:number[] = [300, 500,200, 100,50];
    public pieChartType:string = 'pie';
    
    
}