import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
@Injectable({
  providedIn: 'root',
})
export class LoginService {
  _url = 'http://localhost:8091/api/auth/login';
  constructor(private _http: HttpClient) {}
  login(userData: any) {
    return this._http.post<any>(this._url, userData);
  }
}
