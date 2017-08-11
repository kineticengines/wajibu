import { Component } from "@angular/core";

@Component({
    moduleId:module.id,
    selector:"side-bar",
    template:`
        <ul>
           <li class="Sidebar-navItem">                    
              <a routerLink="/dash">General</a>
           </li>
           <li class="Sidebar-navItem">                    
             <a routerLink="/dash/sentiment" routerLinkActive="active">Sentiments</a>
           </li>
           <li class="Sidebar-navItem">                    
             <a routerLink="/dash/amend" routerLinkActive="active">Amendments</a>
           </li>
           <li class="Sidebar-navItem">                   
            <a routerLink="/dash/settings" routerLinkActive="active">Settings</a>
           </li>
        </ul>        

    `,
    styleUrls:["./dash.sidebar.css"]
})
export class SideBarComponent{

}