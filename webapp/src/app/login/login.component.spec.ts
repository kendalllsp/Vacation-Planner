import { ComponentFixture, TestBed, fakeAsync, tick } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing'; 
import { ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { LoginComponent } from './login.component';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { AppModule } from '../app.module';
import { Location } from "@angular/common";
import { routes } from './login.component';

describe('LoginComponent', () =>  {
    let location: Location;
    let component: LoginComponent;
    let fixture: ComponentFixture<LoginComponent>;
    let router: Router;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [LoginComponent],
            imports: [AppModule, RouterTestingModule.withRoutes(routes), MatFormFieldModule, MatCardModule, MatToolbarModule, HttpClientTestingModule, RouterTestingModule, ReactiveFormsModule]
        })
        .compileComponents();

        router = TestBed.get(Router);
        location = TestBed.get(Location);
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(LoginComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create the login component', () => {
        const fixture = TestBed.createComponent(LoginComponent);
        const app = fixture.debugElement.componentInstance;
        expect(app).toBeTruthy();
    });

    it('navigate to "trips" takes you to /trips', fakeAsync(() => {
        router.navigate(["/trips"]).then(() => {
          expect(location.path()).toBe("/trips");
        });
    }));

    it('navigate to "login" takes you to /login', fakeAsync(() => {
        router.navigate(["/login"]).then(() => {
          expect(location.path()).toBe("/login");
        });
    }));

    it('navigate to "" takes you to /', fakeAsync(() => {
        router.navigate([""]).then(() => {
          expect(location.path()).toBe("/");
        });
    }));
});