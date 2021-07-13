import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {environment as env} from "../../../../../environments/environment";
import {RegisterDto} from "../../interfaces/RegisterDto";
import {Observable} from "rxjs";
import {TokenDto} from "../../interfaces/TokenDto";
import {LoginDto} from "../../interfaces/LoginDto";

@Injectable({
  providedIn: 'root'
})
export class AuthHttpClient {
  constructor(private http: HttpClient) {}

  loginUser(loginDto: LoginDto): Observable<TokenDto> {
    console.log(env.AUTH_URL);
    return this.http.post<TokenDto>(`${env.AUTH_URL}/login`, loginDto);
  }

  registerUser(registerDto: RegisterDto): Observable<TokenDto> {
    return this.http.post<TokenDto>(`${env.AUTH_URL}/register`, registerDto);
  }
}
