import {CanActivate, Router} from "@angular/router";
import {Injectable} from "@angular/core";
import {AuthService} from "../auth/auth.service";

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(
    private authService: AuthService,
    private router: Router
  ) {
  }

  canActivate() {
    if(!this.authService.isLogged()) {
      this.router.navigate(['login']);
      return false;
    }
    return true;
  }
}
