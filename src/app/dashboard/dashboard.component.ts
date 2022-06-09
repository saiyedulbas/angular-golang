import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import jwt_decode from 'jwt-decode';
@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss'],
})
export class DashboardComponent implements OnInit {
  public mona: any;
  public zinda: any;
  public poka: any;
  constructor(private router: Router) {
    let token = localStorage.getItem('token');
    if (token == null) {
      this.router.navigate(['/home']);
    } else {
      this.mona = localStorage.getItem('token');
      this.zinda = jwt_decode(this.mona);
      this.poka = this.zinda.name;
    }
  }

  ngOnInit(): void {}
}
