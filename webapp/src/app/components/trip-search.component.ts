import {Component, Inject, HostListener, ViewChild, ElementRef} from '@angular/core';
import {FormControl, FormGroup,} from '@angular/forms';
import {MatDialog, MAT_DIALOG_DATA} from '@angular/material/dialog';
import { map, catchError, tap } from 'rxjs/operators';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Trip } from '../trip';
import { TripsService } from '../trips.service';
import { faBookmark as faBookmarkSolid  } from '@fortawesome/free-solid-svg-icons';
import { faBookmark  } from '@fortawesome/free-regular-svg-icons';

import { Destination } from '../destination';

/** @title Date range picker forms integration */
@Component({
  selector: 'trip-search',
  styleUrls: ['trip-search.component.scss'],
  templateUrl: 'trip-search.component.html',
})
export class TripSearchComponent {

  public tripForm = new FormGroup({
    dest: new FormControl<string | null>(null),
    start: new FormControl<string | null>(null),
    end: new FormControl<string | null>(null),
  });

  constructor(private tripsService: TripsService, public dialog: MatDialog) {}


  destinationResults: Destination = new Destination();

  openDialog(): void {
    this.dialog.open(DestinationResultDialog, {
      data: {results: this.destinationResults},
    });
  }

  add(): void {
    var dest: string = JSON.stringify(this.tripForm.value.dest!);
    var start: string = JSON.stringify(this.tripForm.value.start!);
    var end: string = JSON.stringify(this.tripForm.value.end!);
    start = start.substring(1,11);
    end = end.substring(1,11);
    dest = dest.substring(1,dest.length-1);
    if (!dest) {return;}
    this.tripsService.getTrip({dest, start, end} as Trip).pipe(
      map((response) => {
        this.destinationResults = response;
      })
    )
    .subscribe(response => {
      this.openDialog();
    });
  }

}

@Component({
  selector: 'destination-dialog',
  styleUrls: ['destination-dialog.component.scss'],
  templateUrl: 'destination-dialog.component.html',
})
export class DestinationResultDialog {
  constructor(@Inject(MAT_DIALOG_DATA) public data: any, private http: HttpClient) {}

  faBookmark = faBookmark;
  faBookmarkSolid = faBookmarkSolid;

  isBookmarked = faBookmark;

  // Created function for when the save destination button is clicked on within the dialog window
  saveDestination() {
    if(this.isBookmarked == faBookmark){
      this.isBookmarked = faBookmarkSolid;
    }
    
    // Checking if the loggedIn sessionStorage value has been set
    if (sessionStorage.getItem('loggedIn') != null)
    {
      // Setting the location from the trip component to be passed to the backend
      // Using the email from the loggedIn variable
      var location = this.data.results.Location[0] + ", " + this.data.results.Location[1]
      const params = { Email: sessionStorage.getItem('loggedIn'), Location: location, Start: this.data.results.Start, End: this.data.results.End }

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

  
}