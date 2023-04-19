import { ComponentFixture, TestBed, fakeAsync, tick } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing'; 
import { Router } from '@angular/router';
import { ReactiveFormsModule, FormBuilder, FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { HomeComponent } from './home.component';
import { TripSearchComponent } from '../components/trip-search.component';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatCardModule } from '@angular/material/card';
import { MatToolbarModule } from '@angular/material/toolbar';
import { AppModule } from '../app.module';
import { MatDialogModule } from '@angular/material/dialog';
import { routes } from "./home.component";
import { Location } from "@angular/common";

describe('HomeComponent', () =>  {
    let component: HomeComponent;
    let fixture: ComponentFixture<HomeComponent>;
    let router: Router;
    let location: Location;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [HomeComponent, TripSearchComponent],
            imports: [AppModule, MatDialogModule, FormsModule, HttpClientModule, MatAutocompleteModule, MatToolbarModule, MatCardModule, ReactiveFormsModule, HttpClientTestingModule, RouterTestingModule],
            providers: [FormBuilder]
        })
        .compileComponents();

        router = TestBed.get(Router);
        location = TestBed.get(Location);
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(HomeComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
        router = TestBed.inject(Router);
    });

    it('should create the HomeComponent ', () => {
        const fixture = TestBed.createComponent(HomeComponent);
        const app = fixture.debugElement.componentInstance;
        expect(component).toBeTruthy();
    });

    it('navigate to "login" redirects you to /login', fakeAsync(() => {
        router.navigate(["/login"]).then(() => {
          expect(location.path()).toBe("/login");
        });
    }));
    
    it('navigate to "trips" takes you to /trips', fakeAsync(() => {
        router.navigate(["/trips"]).then(() => {
          expect(location.path()).toBe("/trips");
        });
    }));
    
    it('navigate to "" takes you to /', fakeAsync(() => {
        router.navigate([""]).then(() => {
          expect(location.path()).toBe("/");
        });
    }));

});