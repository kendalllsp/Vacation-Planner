import { Location } from "@angular/common";
import { TestBed, fakeAsync, tick } from "@angular/core/testing";
import { RouterTestingModule } from "@angular/router/testing";
import { Router } from "@angular/router";

import { NavbarComponent } from './navbar.component';
import { routes } from "./navbar.component";
import { Component } from "@angular/core";

describe("NavbarComponent", () => {
  let location: Location;
  let router: Router;
  let fixture;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [RouterTestingModule.withRoutes(routes)],
      declarations: [NavbarComponent]
    });

    router = TestBed.get(Router);
    location = TestBed.get(Location);

    //fixture = TestBed.createComponent(AppComponent);
    router.initialNavigation();
  });

  it("fakeAsync works", fakeAsync(() => {
    let promise = new Promise(resolve => {
      setTimeout(resolve, 10);
    });
    let done = false;
    promise.then(() => (done = true));
    tick(50);
    expect(done).toBeTruthy();
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