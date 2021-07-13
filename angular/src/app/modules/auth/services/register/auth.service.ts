import {Injectable} from '@angular/core';
import {AuthHttpClient} from "../http/auth-http.client";
import {RegisterDto} from "../../interfaces/RegisterDto";
import {Observable} from "rxjs";
import {TokenDto} from "../../interfaces/TokenDto";
import {LoginDto} from "../../interfaces/LoginDto";

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(
    private authHttpClient: AuthHttpClient,
  ) {}

  loginUser(loginDto: LoginDto): Observable<TokenDto> {
    return this.authHttpClient.loginUser(loginDto);
  }

  registerUser(registerDto: RegisterDto): Observable<TokenDto> {
    return this.authHttpClient.registerUser(registerDto)
  }

  successfulLogin(response: TokenDto): void {
    if(response.token) {
      localStorage.setItem("token", response.token);
    }
  }
}

