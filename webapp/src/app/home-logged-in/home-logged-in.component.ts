import {Component} from '@angular/core';

import { Trip } from '../trip';
import { TripsService } from '../trips.service';

import {FormControl} from '@angular/forms';
// Added the Angular Router so that if the user has successfully logged in, we can route them to the trips
// page for the example implementation.
import { Router } from '@angular/router';

/**
 * @title Basic grid-list
 */
@Component({
  selector: 'app-home-logged-in',
  styleUrls: ['home-logged-in.component.scss'],
  templateUrl: 'home-logged-in.component.html',
})
export class HomeLoggedInComponent {

  // Injecting Router to be used within the Home Component
  constructor(private router: Router) {}

  // Creating the tripsIfLoggedIn function that is called when the home button is clicked in the home component.
  // It creates a new variable based off of the sessionStorage value which is set on a successful login, and not ever
  // set if the user has not logged in. Therefore we're checking if it's null or not.
  logOut() {
    const loggedIn = sessionStorage.getItem('loggedIn')
    // Deleting the loggedIn val and routing back home
    sessionStorage.removeItem('loggedIn')
    this.router.navigate([''])
  }

}