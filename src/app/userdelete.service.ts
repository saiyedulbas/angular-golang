import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
@Injectable({
  providedIn: 'root',
})
export class UserdeleteService {
  constructor(private _http: HttpClient) {}
  deleteUser(userId: any, token: any) {
    let headers = new HttpHeaders({
      Token: token,
    });
    let options = { headers: headers };
    return this._http.post(
      'http://localhost:8091/api/user/userdelete/' + userId,
      null,
      options
    );
  }
}
