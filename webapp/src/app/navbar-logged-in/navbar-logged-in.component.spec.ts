import { TestBed, fakeAsync, tick, async } from '@angular/core/testing';
import { NavbarLoggedInComponent } from './navbar-logged-in.component'
import { AppModule } from '../app.module';
import { Router } from "@angular/router";
import { RouterTestingModule } from '@angular/router/testing';
import { routes } from "./navbar-logged-in.component";
import { Location } from "@angular/common";

describe('NavbarLoggedInComponent', () => {
  let location: Location;
  let router: Router;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        NavbarLoggedInComponent
      ],
      imports: [
        AppModule, RouterTestingModule.withRoutes(routes)
      ]
    }).compileComponents();

    router = TestBed.get(Router);
    location = TestBed.get(Location);
  }));

  it('should create the NavbarLoggedInComponent', async(() => {
    const fixture = TestBed.createComponent(NavbarLoggedInComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  }));

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