import { ComponentFixture, TestBed, fakeAsync, tick } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing'; 
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { SignupComponent } from './signup.component';
import { AppModule } from '../app.module';
import { routes } from "./signup.component";
import { Router } from "@angular/router";
import { Location } from "@angular/common";

describe('SignupComponent', () =>  {
    let component: SignupComponent;
    let fixture: ComponentFixture<SignupComponent>;
    let router: Router;
    let location: Location;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [SignupComponent],
            imports: [AppModule, HttpClientTestingModule, RouterTestingModule, ReactiveFormsModule]
        })
        .compileComponents();

        router = TestBed.get(Router);
        location = TestBed.get(Location);
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(SignupComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
        router = TestBed.inject(Router);
    });

    it('should create the SignupComponent', () => {
        const fixture = TestBed.createComponent(SignupComponent);
        const app = fixture.debugElement.componentInstance;
        expect(app).toBeTruthy();
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