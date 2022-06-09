import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import jwt_decode from 'jwt-decode';
import { LoginService } from '../login.service';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit {
  public successMessage: any;
  public shimali: any;
  public loaderStatus: any = false;
  public dataToBeSent: any;
  public data: any;
  constructor(
    private fb: FormBuilder,
    private _loginService: LoginService,
    private location: Location,
    private router: Router
  ) {}
  loginForm = this.fb.group({
    name: ['', Validators.required],
    password: ['', Validators.required],
  });
  onSubmit() {
    this.dataToBeSent = {
      name: this.loginForm.value.name,
      password: this.loginForm.value.password,
    };
    this.loaderStatus = true;
    this._loginService.login(this.dataToBeSent).subscribe(
      (response) => {
        this.successMessage = response.message;
        this.loaderStatus = false;
        if (response.message == 'Logged In') {
          let mona = jwt_decode(response.token);
          console.log(mona);
          localStorage.setItem('token', response.token);
          this.router.navigate(['/dashboard']);
        }
      },
      (error) => console.log(error)
    );
  }

  ngOnInit(): void {}
}
