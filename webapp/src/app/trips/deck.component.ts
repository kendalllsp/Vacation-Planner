import {Component, Inject, OnInit} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import {MatDialog, MAT_DIALOG_DATA} from '@angular/material/dialog';
import { Trip } from '../trip';
import { map, catchError, tap } from 'rxjs/operators';
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
  destinationResults: Destination = new Destination();
  
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

  openDialog(): void {
    this.dialog.open(DialogDataExampleDialog, {
      data: {results: this.destinationResults},
    });
  }

  viewTrip(location: string){

    var tripsUrl = 'http://localhost:8181/';

    console.log(tripsUrl.concat("newDestination/", location) );
    
    this.http.get<any>(tripsUrl.concat("newDestination/", location)).pipe(
      map(response => {
      // Print request response to JS console
      console.log(response)
      this.destinationResults = response;
    }))
    .subscribe(response => {
      this.openDialog();
    });;

  }

}

@Component({
  selector: 'deck-dialog',
  styleUrls: ['deck-dialog.component.scss'],
  templateUrl: 'deck-dialog.component.html',
})
export class DialogDataExampleDialog {
  constructor(@Inject(MAT_DIALOG_DATA) public data: any, private http: HttpClient) {}

  faBookmark = faBookmark;
  faBookmarkSolid = faBookmarkSolid;

  isBookmarked = faBookmarkSolid;

  // Created function for when the save destination button is clicked on within the dialog window
  saveDestination() {
    if(this.isBookmarked == faBookmark){
      this.isBookmarked = faBookmarkSolid;
       // Checking if the loggedIn sessionStorage value has been set
      if (sessionStorage.getItem('loggedIn') != null)
      {
        // Setting the location from the trip component to be passed to the backend
        // Using the email from the loggedIn variable
        var location = this.data.results.Location[0] + ", " + this.data.results.Location[1]
        const params = { Email: sessionStorage.getItem('loggedIn'), Location: location }

        const httpOptions = {
          headers: new HttpHeaders({
            'Content-Type': 'application/json'
          })
        };
    
        // Calling the backend with the post request to create new destination in the list with the location and users email
        this.http.post("http://localhost:8181/updateDestination", params, httpOptions)
        .subscribe(response => {
            // Print request response to JS console
            console.log(response)
        });
      }
      else
      {
        // Logging error message if the user is not logged in.
        console.log("You have to log in, dork.")
      }
    }
    else{
      this.isBookmarked = faBookmark;
       // Checking if the loggedIn sessionStorage value has been set
       if (sessionStorage.getItem('loggedIn') != null)
       {
         // Setting the location from the trip component to be passed to the backend
         // Using the email from the loggedIn variable
         var location = this.data.results.Location[0] + ", " + this.data.results.Location[1]
         var email = sessionStorage.getItem("loggedIn")
         
         const httpOptions = {
           headers: new HttpHeaders({ 'Content-Type': 'application/json'}),
           body: { "Email": email,
                  "Location": location }
         };
     
         // Calling the backend with the post request to create new destination in the list with the location and users email
         this.http.delete("http://localhost:8181/updateDestination", httpOptions)
         .subscribe(response => {
             // Print request response to JS console
             window.location.reload()
         });
       }
       else
       {
         // Logging error message if the user is not logged in.
         console.log("You have to log in, dork.")
       }
      }
    }
}