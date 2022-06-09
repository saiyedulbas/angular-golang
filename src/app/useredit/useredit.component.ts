import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { ActivatedRoute, Router } from '@angular/router';
import { ProfileService } from '../profile.service';
import { RegistrationService } from '../registration.service';
import { UsereditService } from '../useredit.service';
@Component({
  selector: 'app-useredit',
  templateUrl: './useredit.component.html',
  styleUrls: ['./useredit.component.scss'],
})
export class UsereditComponent implements OnInit {
  public successMessage: any;
  public shimali: any;
  public loaderStatus: any = false;
  public dataToBeSent: any;
  public data: any;
  public val: any;
  public mona: any;
  public name: any;
  public password: any;
  public id: any;
  userEditForm = this.fb.group({
    name: ['', Validators.required],
    password: ['', [Validators.required, Validators.minLength(6)]],
  });
  constructor(
    private fb: FormBuilder,
    private _registrationService: RegistrationService,
    private titleService: Title,
    private location: Location,
    private _userEdit: UsereditService,
    private route: ActivatedRoute,
    private _profileService: ProfileService,
    private router: Router
  ) {
    this.data = location.onUrlChange((val) => {});
    let id = this.route.snapshot.paramMap.get('id');
    this.mona = localStorage.getItem('token');
    this.id = id;
    this._profileService.getProfileInfo(id, this.mona).subscribe(
      (response) => {
        this.name = response.message[0].name;
        this.password = response.message[0].password;
        this.userEditForm.controls['name'].setValue(this.name);
        this.userEditForm.controls['password'].setValue(this.password);
      },
      (error) => console.log(error)
    );
  }

  onEditUser() {
    let token = localStorage.getItem('token');
    this.dataToBeSent = {
      name: this.userEditForm.value.name,
      password: this.userEditForm.value.password,
    };
    this.loaderStatus = true;
    this._userEdit.userEdit(this.id, this.dataToBeSent, token).subscribe(
      (response: any) => {
        alert(response.message);
        this.router.navigate(['/allusers']);
        console.log(response.message);
      },
      (error) => console.log(error)
    );
  }

  ngOnInit(): void {}
}
