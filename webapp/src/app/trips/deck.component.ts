import {Component,Inject, OnInit} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {MatDialog, MAT_DIALOG_DATA} from '@angular/material/dialog';
import { Trip } from '../trip';
import { TripsService } from '../trips.service';
import { faBookmark as faBookmarkSolid  } from '@fortawesome/free-solid-svg-icons';
import { faBookmark  } from '@fortawesome/free-regular-svg-icons';

import { Destination } from '../destination';
@Component({
  selector: 'trips-deck-component',
  templateUrl: 'deck.component.html',
  styleUrls: ['deck.component.scss'],
})
export class TripsCardDeckComponent{

  cards: any[] = [];

  constructor(private http: HttpClient, private tripsService: TripsService, public dialog: MatDialog) {}
  
  ngOnInit() {

    var Email = sessionStorage.getItem('loggedIn')
    if (Email == null) {
      Email = "";
    }

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    };
  
      // Calling the backend with the post request to new destination in the list with the location and users email
      this.http.get<any[]>('http://localhost:8181/updateDestination', { params: { Email } }).subscribe((data: any[]) => {
        this.cards = data
        console.log(data)
      });
  }

}