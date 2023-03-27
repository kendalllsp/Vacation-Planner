import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { MaterialModule} from './material/material.module';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';

import { AppComponent } from './app.component';

import { NavbarComponent } from './navbar/navbar.component';
import { NavbarLoggedInComponent } from './navbar-logged-in/navbar-logged-in.component';

import { HomeComponent } from './home/home.component';

import { HomeLoggedInComponent } from './home-logged-in/home-logged-in.component';

import { TripSearchComponent, DestinationResultDialog } from './components/trip-search.component';

import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';

import { TripsComponent } from './trips/trips.component';
import { TripsCardComponent } from './trips/card.component';
import { TripsCardDeckComponent } from './trips/deck.component';

import { MessageService } from './message.service';


import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { FlexLayoutModule } from '@angular/flex-layout';



@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    NavbarLoggedInComponent,
    HomeComponent,
    HomeLoggedInComponent,
    TripSearchComponent,
    DestinationResultDialog,
    LoginComponent,
    SignupComponent,
    TripsComponent,
    TripsCardComponent,
    TripsCardDeckComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    AppRoutingModule,
    MaterialModule,
    FontAwesomeModule,
    FormsModule,
    ReactiveFormsModule,
    FlexLayoutModule
  ],
  providers: [
    MessageService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
