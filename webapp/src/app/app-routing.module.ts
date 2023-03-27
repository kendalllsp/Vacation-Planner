import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { TripsComponent } from './trips/trips.component';
import { SignupComponent } from './signup/signup.component';
import { HomeLoggedInComponent } from './home-logged-in/home-logged-in.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'login', component: LoginComponent },
  { path: 'trips', component: TripsComponent },
  { path: 'signup', component: SignupComponent},
  { path: 'homeLoggedIn', component: HomeLoggedInComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
