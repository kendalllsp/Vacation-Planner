import { Component } from '@angular/core';
import { faSearch  } from '@fortawesome/free-solid-svg-icons';
import { Router } from '@angular/router';
import { Routes } from "@angular/router"
import { LoginComponent } from '../login/login.component';
import { TripsCardDeckComponent } from '../trips/deck.component';

@Component({
  selector: 'app-navbar-logged-in',
  templateUrl: './navbar-logged-in.component.html',
  styleUrls: ['./navbar-logged-in.component.scss']
})
export class NavbarLoggedInComponent {
  faSearch = faSearch;
  constructor(private router: Router) {}

  // Creating the tripsIfLoggedIn function that is called when the home button is clicked in the home component.
  // It creates a new variable based off of the sessionStorage value which is set on a successful login, and not ever
  // set if the user has not logged in. Therefore we're checking if it's null or not.
  logOut() {
    const loggedIn = sessionStorage.getItem('loggedIn')
    sessionStorage.removeItem('loggedIn')
    this.router.navigate([''])
  }
}

export const routes: Routes = [
  {path:'login', component: LoginComponent}, 
  {path: 'trips', component: TripsCardDeckComponent}
];