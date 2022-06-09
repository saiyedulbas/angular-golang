import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
@Injectable({
  providedIn: 'root',
})
export class UsereditService {
  constructor(private _http: HttpClient) {}
  userEdit(userId: any, data: any, token: any) {
    let headers = new HttpHeaders({
      Token: token,
    });
    let options = { headers: headers };
    return this._http.post(
      'http://localhost:8091/api/user/useredit/' + userId,
      data,
      options
    );
  }
}
