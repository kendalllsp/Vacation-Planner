import {Component, OnInit} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Component({
  selector: 'trips-deck-component',
  templateUrl: 'deck.component.html',
  styleUrls: ['deck.component.scss'],
})
export class TripsCardDeckComponent{

  cards: any[] = [];

  constructor(private http: HttpClient) {}

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
  
      // Calling the backend with the post request to create new destination in the list with the location and users email
      this.http.get<any[]>('http://localhost:8181/updateDestination', { params: { Email } }).subscribe((data: any[]) => {
        this.cards = data
        console.log(data)
      });
  }

}