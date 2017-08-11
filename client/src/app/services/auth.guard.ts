import { Injectable } from "@angular/core";
import { CanActivate,CanActivateChild,
     Router,ActivatedRouteSnapshot ,RouterStateSnapshot} from "@angular/router";
import { Initializer } from "./init.service"     

import { Observable} from "rxjs/Rx";
import "rxjs/add/operator/do";
import "rxjs/add/operator/map";
import "rxjs/add/operator/take";

@Injectable()
export class AuthGuard implements CanActivate{
    constructor(private router: Router,private init:Initializer){} 

    canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot):Observable<boolean>{
        return Observable.of(this.init.checker())
                        .take(1)
                        .map(d => d.IsDeployed)
                        .do(IsDeployed => {                                                       
                            if (!IsDeployed) this.router.navigate(["/init"])
                        })
     }
}