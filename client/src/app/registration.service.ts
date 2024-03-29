import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class RegistrationService {
  _url = 'http://localhost:8091/api/auth/register';
  constructor(private _http: HttpClient) {}
  register(userData: any) {
    return this._http.post<any>(this._url, userData);
  }
}
