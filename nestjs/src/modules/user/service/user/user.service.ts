// import { Injectable } from '@nestjs/common';
// import {UserProvider} from "../../http/user.provider";
// import {Observable} from "rxjs";
// import {map} from "rxjs/operators";
//
// @Injectable()
// export class UserService {
//     constructor(private userProvider: UserProvider) {}
//
//     getUser(id: string): Observable<User>
//     {
//         return this.userProvider.getUser(id)
//             .pipe(
//                 map(res => res.data)
//             );
//     }
//
// }
