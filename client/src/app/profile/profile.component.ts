import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import jwt_decode from 'jwt-decode';
import { ProfileService } from '../profile.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss'],
})
export class ProfileComponent implements OnInit {
  public mona: any;
  public zinda: any;
  public poka: any;
  public name: any;
  public email: any;
  public password: any;
  public userId: any;
  constructor(private _profileService: ProfileService, private router: Router) {
    let token = localStorage.getItem('token');
    if (token == null) {
      this.router.navigate(['/home']);
    } else {
      this.mona = localStorage.getItem('token');
      this.zinda = jwt_decode(this.mona);
      this.poka = this.zinda.user_id;
      this._profileService.getProfileInfo(this.poka, this.mona).subscribe(
        (response) => {
          this.name = response.message[0].name;
          this.email = response.message[0].email;
          this.password = response.message[0].password;
        },
        (error) => console.log(error)
      );
    }
  }

  ngOnInit(): void {}
}
