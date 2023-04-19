import { Component } from '@angular/core';
import { faSearch  } from '@fortawesome/free-solid-svg-icons';
import { Routes } from "@angular/router"
import { LoginComponent } from '../login/login.component';
import { TripsCardDeckComponent } from '../trips/deck.component';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent {
  faSearch = faSearch;
}

export const routes: Routes = [
  {path:'login', component: LoginComponent}, 
  {path: 'trips', component: TripsCardDeckComponent}
];