import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { Trip } from './trip';
import { MessageService } from './message.service';

@Injectable({
  providedIn: 'root'
})
export class TripsService {

  private tripsUrl = 'http://localhost:8181/';

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient,
    private messageService: MessageService) { }

  getTrips(): Observable<Trip[]> {
    return this.http.get<Trip[]>(this.tripsUrl)
  }

  getTrip(trip: Trip): Observable<any> {

    const params = {
      location: trip.dest,
      start: trip.start,
      end: trip.end
    };
    
    return this.http.get<any>('http://localhost:8181/newDestination', { params });
  }


  private log(message: string) {
    this.messageService.add(`HeroService: ${message}`);
  }
}
