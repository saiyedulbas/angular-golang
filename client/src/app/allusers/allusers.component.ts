import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import jwt_decode from 'jwt-decode';
import { AllusersService } from '../allusers.service';
import { UserdeleteService } from '../userdelete.service';
@Component({
  selector: 'app-allusers',
  templateUrl: './allusers.component.html',
  styleUrls: ['./allusers.component.scss'],
})
export class AllusersComponent implements OnInit {
  public mona: any;
  public zinda: any;
  public poka: any;
  public name: any;
  public email: any;
  public password: any;
  public userId: any;
  public userList: any;
  constructor(
    private _allUsersService: AllusersService,
    private router: Router,
    private _userDeleteService: UserdeleteService
  ) {
    let token = localStorage.getItem('token');
    if (token == null) {
      this.router.navigate(['/home']);
    } else {
      this.mona = localStorage.getItem('token');
      this.zinda = jwt_decode(this.mona);
      this.poka = this.zinda.user_id;
      this._allUsersService.getAllUsers(this.mona).subscribe(
        (response) => {
          this.userList = response.data;
        },
        (error) => console.log(error)
      );
    }
  }
  onView(user: any) {
    this.router.navigate(['/user', user.id]);
  }
  onDelete(user: any) {
    let token = localStorage.getItem('token');
    this._userDeleteService.deleteUser(user.id, token).subscribe(
      (response) => {
        console.log(response);
        this._allUsersService.getAllUsers(this.mona).subscribe(
          (response) => {
            this.userList = response.data;
            alert('User Successfully Deleted');
          },
          (error) => console.log(error)
        );
      },
      (error) => console.log(error)
    );
  }
  onEdit(user: any) {
    this.router.navigate(['/edit', user.id]);
  }

  ngOnInit(): void {}
}
