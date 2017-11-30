import { Component } from "@angular/core"

@Component({
    moduleId:module.id,
    selector:"general",
    template:`
        <div class="container-col">
           <div class="container-item">
                <top-box></top-box>
           </div>
           <div class="container-item info-extra">
                 <p>Support the efforts of Wajibu by sending your donation via Mpesa Paybill number
               12345</p>
               <p>Help us shape a better Kenya for present and future generations.</p>
           </div>
           <div class="container-item">
                <div class="container-row">
                    <div class="container-item oneth">
                        <filter-box ></filter-box>
                        <links-box></links-box>
                        <info-box></info-box>                                           
                    </div>
                    <div class="container-item twoth">
                        <router-outlet ></router-outlet>
                    </div>                            
                </div>
           </div>
             
        </div>
    `,
    styleUrls:["./general.component.css"]
})

export class GeneralComponent{
    
}