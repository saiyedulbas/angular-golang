import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { RegistrationService } from '../registration.service';
@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.scss'],
})
export class RegistrationComponent implements OnInit {
  public successMessage: any;
  public shimali: any;
  public loaderStatus: any = false;
  public dataToBeSent: any;
  public data: any;
  public val: any;
  constructor(
    private fb: FormBuilder,
    private _registrationService: RegistrationService,
    private titleService: Title,
    private location: Location
  ) {
    this.data = location.onUrlChange((val) => {});
  }
  registrationForm = this.fb.group({
    name: ['', Validators.required],
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required, Validators.minLength(6)]],
  });
  onSubmit() {
    this.dataToBeSent = {
      name: this.registrationForm.value.name,
      email: this.registrationForm.value.email,
      password: this.registrationForm.value.password,
    };
    this.loaderStatus = true;
    this._registrationService.register(this.dataToBeSent).subscribe(
      (response) => {
        this.successMessage = response.message;
        this.loaderStatus = false;
      },
      (error) => console.log(error)
    );
  }

  ngOnInit(): void {}
}
