import { Location } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { Router } from '@angular/router';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
  public data: any;
  public val: any;
  public val2: any;
  constructor(
    private location: Location,
    private titleService: Title,
    private router: Router
  ) {
    this.data = location.onUrlChange((val) => {
      let token = localStorage.getItem('token');
      if (token == null) {
        this.val = true;
        this.val2 = false;
      } else {
        this.val = false;
        this.val2 = true;
      }
      if (val === '/register') {
        this.titleService.setTitle('Register');
      } else if (val === '/login') {
        this.titleService.setTitle('Login');
      } else if (val === '/home') {
        this.titleService.setTitle('Home');
      } else if (val === '/dashboard') {
        this.titleService.setTitle('Dashboard');
      } else if (val === '/profile') {
        this.titleService.setTitle('Profile');
      } else if (val === '/allusers') {
        this.titleService.setTitle('All Users');
      }
    });
  }
  ngOnInit(): void {}
}
