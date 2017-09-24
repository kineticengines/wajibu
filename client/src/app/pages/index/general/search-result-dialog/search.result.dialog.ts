import { Component,Inject,OnInit } from "@angular/core";
import { MdDialogRef, MD_DIALOG_DATA} from "@angular/material"

@Component({
    moduleId:module.id,
    selector:"search-result-dialog",
    template:`
        <md-dialog-content>
            <div class="main-container">
                <div *ngFor="let d of data" class="container" >
                    <div class="container-item header-item">
                        <div class="header-item-item"><h2>Title : {{d.title | titlecase}}</h2></div>
                        <div class="header-item-item"><h2>Name :  {{d.name | titlecase}}</h2></div>
                    </div>
                    <div class="container-item body-item">
                        <div *ngFor="let dd of d.data" class="body-item-item body-content">
                            <div class="body-content-item body-content-item-image">
                                <!--{{dd.image}}-->
                                <img src="./assets/photo.png" class="body-image">
                            </div>
                            <div class="body-content-item body-content-item-body">                            
                                <div class="body-content-item-body-item middle-header">
                                    <p>{{dd.sentiment}} </p> 
                                </div>
                                <div class="body-content-item-body-item footer-header">
                                    <div class="footer-header-item" *ngIf="dd.pillars">
                                        <small><em style="color:teal;">Pillar</em> : {{dd.pillars | titlecase}}</small> 
                                    </div>
                                    <div class="footer-header-item" *ngIf="dd.county">
                                        <small><em style="color:teal;">County</em> : {{dd.county | titlecase}}</small> 
                                    </div>
                                    <div class="footer-header-item" *ngIf="dd.constituency">
                                        <small><em style="color:teal;">Constituency</em> : {{dd.constituency | titlecase}}</small> 
                                    </div>
                                    <div class="footer-header-item" *ngIf="dd.ward">
                                        <small><em style="color:teal;">Ward</em> : {{dd.ward | titlecase}}</small> 
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </md-dialog-content>
        <md-dialog-actions>
            <button md-button style="background-color:teal; color:whitesmoke;" [md-dialog-close]="true">Close</button>
        </md-dialog-actions>
    `,
    styleUrls:["./search.result.dialog.css"]
})
export class SearchResultDialogComponent implements OnInit{
    constructor(public dialogRef:MdDialogRef<SearchResultDialogComponent>,
    @Inject(MD_DIALOG_DATA) public data:any){
    }

    ngOnInit(){
        
    }
}