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
  selector: 'app-home',
  styleUrls: ['home.component.scss'],
  templateUrl: 'home.component.html',
})
export class HomeComponent {

  // Injecting Router to be used within the Home Component
  constructor(private router: Router) {}

  // Creating the tripsIfLoggedIn function that is called when the home button is clicked in the home component.
  // It creates a new variable based off of the sessionStorage value which is set on a successful login, and not ever
  // set if the user has not logged in. Therefore we're checking if it's null or not.
  tripsIfLoggedIn() {
    const loggedIn = sessionStorage.getItem('loggedIn')
    if (loggedIn != null)
    {
      // If the loggedIn variable has a value(email) then the router will take us to the trips page.
      this.router.navigate(['/trips'])
    }
  }

}