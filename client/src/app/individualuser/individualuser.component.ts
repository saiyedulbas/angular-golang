import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ProfileService } from '../profile.service';
@Component({
  selector: 'app-individualuser',
  templateUrl: './individualuser.component.html',
  styleUrls: ['./individualuser.component.scss'],
})
export class IndividualuserComponent implements OnInit {
  public id: any;
  public mona: any;
  public zinda: any;
  public poka: any;
  public name: any;
  public email: any;
  public password: any;
  constructor(
    private route: ActivatedRoute,
    private _profileService: ProfileService
  ) {
    let id = this.route.snapshot.paramMap.get('id');
    this.mona = localStorage.getItem('token');

    this._profileService.getProfileInfo(id, this.mona).subscribe(
      (response) => {
        this.name = response.message[0].name;
        this.email = response.message[0].email;
        this.password = response.message[0].password;
      },
      (error) => console.log(error)
    );
  }

  ngOnInit(): void {}
}
