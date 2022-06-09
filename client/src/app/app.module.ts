import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AllusersComponent } from './allusers/allusers.component';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { DepartmentDetailsComponent } from './department-details/department-details.component';
import { EmployeeDetailsComponent } from './employee-details/employee-details.component';
import { HomeComponent } from './home/home.component';
import { IndividualuserComponent } from './individualuser/individualuser.component';
import { LoginComponent } from './login/login.component';
import { LogoutComponent } from './logout/logout.component';
import { NavComponent } from './nav/nav.component';
import { ProfileComponent } from './profile/profile.component';
import { RegistrationComponent } from './registration/registration.component';
import { UsereditComponent } from './useredit/useredit.component';
import { UsernavComponent } from './usernav/usernav.component';

@NgModule({
  declarations: [
    AppComponent,
    EmployeeDetailsComponent,
    DepartmentDetailsComponent,
    RegistrationComponent,
    LoginComponent,
    HomeComponent,
    NavComponent,
    DashboardComponent,
    UsernavComponent,
    LogoutComponent,
    ProfileComponent,
    AllusersComponent,
    IndividualuserComponent,
    UsereditComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
