import {Component} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'login-component',
  templateUrl: 'login.component.html',
  styleUrls: ['login.component.scss'],
})
export class LoginComponent {
  // Initializing variables to save input given from user within login component
  email!: string;
  password!: string;

  // Injecting HttpClient for http requests and Router for dynamic routing
  constructor(private http: HttpClient, private router: Router) {}

  // Initializing function to be called on "sign in" button click
  loginUser() {
    // Initializing struct of parameters for request body
    const params = { Email: this.email, Password: this.password }

    // Initializing header for JSON file request body abilities
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    };

    // Make post request with appropriate URL, the user/pass given by the user and JSON capabilities
    this.http.post("http://localhost:8181/loginUser", params, httpOptions)
    .subscribe(response => {
        // Print request response to JS console
        console.log(response)
        // Unpacking JSON response so we can refer to it's values
        const jsonData = JSON.parse(JSON.stringify(response));
        // If logged in, go to main page
        if (jsonData.LoggedIn) {
          this.router.navigate(['']);
        }
        // If not logged in, say at login page
        else {
          this.router.navigate(['/login']);
        }
    });
  }
}