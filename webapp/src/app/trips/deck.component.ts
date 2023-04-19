import {Component,Inject, OnInit} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {MatDialog, MAT_DIALOG_DATA} from '@angular/material/dialog';
import { Trip } from '../trip';
import { TripsService } from '../trips.service';
import { faBookmark as faBookmarkSolid  } from '@fortawesome/free-solid-svg-icons';
import { faBookmark  } from '@fortawesome/free-regular-svg-icons';
import { DatePipe } from '@angular/common';
import { Destination } from '../destination';


@Component({
  selector: 'trips-deck-component',
  templateUrl: 'deck.component.html',
  styleUrls: ['deck.component.scss'],
  providers: [DatePipe],
})
export class TripsCardDeckComponent{

  cards: any[] = [];

  constructor(private http: HttpClient, private tripsService: TripsService, public dialog: MatDialog, private datePipe: DatePipe) {}
  
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
        this.cards = data.map((item) => {
          return {
            ...item,
            Start: this.datePipe.transform(item.Start, 'MM/dd/yyyy'),
            End: this.datePipe.transform(item.End, 'MM/dd/yyyy')
          };
        });
        console.log(this.cards);
      });
  }

  deleteTrip(location: string) {

    var email = sessionStorage.getItem("loggedIn")

    const options = {
      headers: { 'Content-Type': 'application/json' },
      body: { "Email": email,
              "Location": location }
    };

    this.http.delete('http://localhost:8181/updateDestination', options).subscribe(() => {
      window.location.reload()
      console.log("test")
    });

  }

}