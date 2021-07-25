import {Injectable} from '@angular/core';
import {AuthHttpClient} from "../http/auth-http.client";
import {RegisterDto} from "../../interfaces/RegisterDto";
import {Observable} from "rxjs";
import {TokenDto} from "../../interfaces/TokenDto";
import {LoginDto} from "../../interfaces/LoginDto";
import {JwtHelperService} from "@auth0/angular-jwt";

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(
    private authHttpClient: AuthHttpClient,
    private jwtHelper: JwtHelperService
  ) {}

  loginUser(loginDto: LoginDto): Observable<TokenDto> {
    return this.authHttpClient.loginUser(loginDto);
  }

  registerUser(registerDto: RegisterDto): Observable<TokenDto> {
    return this.authHttpClient.registerUser(registerDto)
  }

  successfulLogin(response: TokenDto): void {
    if(response.access_token) {
      localStorage.setItem("token", response.access_token);
    }
  }

  isLogged(): boolean {
    try {
      return !this.jwtHelper.isTokenExpired(AuthService.getToken() as string);
    } catch (e) {
      return false;
    }
  }

  static getToken(): string | null {
    return localStorage.getItem("token");
  }
}

