import {HttpService, Injectable} from "@nestjs/common";
import { AxiosResponse } from "axios";
import {Observable} from "rxjs";
import {User} from "../domain/User";

@Injectable()
export class UserProvider {
    private url: string = process.env.AUTH_SERVER_URL;

    constructor(private http: HttpService) {}

    getUser(id: string): Observable<AxiosResponse<User>> {
        return this.http.get(`${this.url}/user/${id}`);
    }
}