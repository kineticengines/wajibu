/*
Wajibu is an online web app that collects,analyses, aggregates and visualizes sentiments
from the public pertaining the government of a nation. This tool allows citizens to contribute
to the governance talk by airing out their honest views about the state of the nation and in
particular the people placed in government or leadership positions.

Copyright (C) 2017
David 'Dexter' Mwangi
dmwangimail@gmail.com
https://github.com/daviddexter
https://github.com/daviddexter/wajibu

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/


import { Component, OnInit, ChangeDetectionStrategy, ChangeDetectorRef} from '@angular/core';
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { Router } from "@angular/router";
import * as EmailValidator from 'email-validator';
import { Initializer } from "../../services/init.service"
import { SessionChecker } from "../../services/session.checker";
 
@Component({
  selector: 'app-root',
  templateUrl: './root.component.html',
  styleUrls: ['./root.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class RootComponent implements OnInit {  
  processing:boolean = false;
  processingLogin:boolean = false;
  showLoginForm:boolean = false;
  showProvideEmailForm:boolean = false;
  createDefault:boolean = false;
  defaultCred:string;
  disable:boolean = false;
  statusString:string = "press Enter to send";
  loginStatusString:string;

  loginForm:FormGroup;

  constructor(private initializer:Initializer,private ref: ChangeDetectorRef,
    private router:Router,private formBuilder:FormBuilder,private sessionChecker:SessionChecker){
  }

  ngOnInit(){
    this.createForm()
    this.processing = true;
    this.ref.detectChanges();
    this.initializer.getDefaultCred().subscribe(d => this.checkDefault(d))     
  }
  createForm(){
    this.loginForm = this.formBuilder.group({
      nameoremail : ['',[Validators.required]],
      password : ['',[Validators.required,Validators.minLength(6)]]
    })
  }
  

  checkDefault(d:any){    
    if(d.Exists === false){
      this.processing = false;
      this.showProvideEmailForm = true;
      this.ref.detectChanges();      
    }else{
      switch (this.sessionChecker.getSession().isSet) {
        case false:
          this.processing = false;
          this.showLoginForm = d.Exists;
          this.ref.detectChanges(); 
          break;     
        case true:
          this.processing = false;
          this.ref.detectChanges();          
          this.router.navigateByUrl('dash')
          break;
      }
      
    }
  }   

  submitEmail(email:string){  
    this.disable = true; 
    this.statusString = "Checking..."
    if (EmailValidator.validate(email) === false){
      this.disable = false;
       this.statusString = "Email Address Invalid."
    }else{
      this.statusString = "Sending Credentials to Email Address..."
      let obj = { Email : email.toLowerCase()}
      this.initializer.createDefaultCred(obj).subscribe(d => this.processEmailSend(d))
    }
  } 

  processEmailSend(d:any){
    if(d.Sent === false){
      this.statusString = "Error occured. Try again later. :("
    }else{
      this.statusString = "Credentials sent.Please check your inbox."
      setTimeout(()=>{
        this.showProvideEmailForm = false;
        this.showLoginForm = d.Sent;       
        this.ref.detectChanges(); 
      },1500)
    }
  }  

  proceedToLogin(){
    this.loginStatusString = "";
    this.processingLogin = true; 
    this.ref.detectChanges();     
    if(this.loginForm.valid === true ){
      this.initializer.createLogin(this.loginForm.value).subscribe(d => this.processLogin(d))
    }    
  }

  processLogin(d:any){
    switch (d.Accurate) {
      case false:
          this.processingLogin = false; 
          this.loginStatusString = "Credentials not accurate";
          this.ref.detectChanges();
        break;
      case true:
        this.processingLogin = false;
        this.ref.detectChanges();        
        this.sessionChecker.createSession(d.Cred)        
        this.router.navigate(['dash'])
        break;
    }
  }

  
}
