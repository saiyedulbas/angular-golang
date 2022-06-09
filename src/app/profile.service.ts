import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ProfileService {
  _url = 'http://localhost:8091/api/user/profile/27';
  constructor(private _http: HttpClient) {}
  getProfileInfo(userId: any, token: any) {
    return this._http.get<any>(
      'http://localhost:8091/api/user/profile/' + userId,
      {
        headers: new HttpHeaders({
          Token: token,
        }),
      }
    );
  }
}
