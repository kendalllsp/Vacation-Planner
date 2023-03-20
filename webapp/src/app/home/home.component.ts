import {Component} from '@angular/core';

import { Trip } from '../trip';
import { TripsService } from '../trips.service';

import {FormControl} from '@angular/forms';
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

  constructor(private router: Router) {}

  homeIfLoggedIn() {
    const loggedIn = sessionStorage.getItem('loggedIn')
    if (loggedIn)
    {
      this.router.navigate(['/'])
    }
  }

}