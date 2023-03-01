import {Component} from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'signup-component',
  templateUrl: 'signup.component.html',
  styleUrls: ['signup.component.scss'],
})
export class SignupComponent {
  email!: string;
  password!: string;
  password1!: string;

  // Injecting HttpClient for request and Router for next page routing
  constructor(private http: HttpClient, private router: Router) {}

  createUser() {
    // Checking if the user entered in their password correctly both times
    if ((this.password == this.password1) && (this.email != null || this.password != null)){
      // Setting params to validated email and password combo given
      const params = { Email: this.email, Password: this.password }
      
      // Initialized an HttpOptions to pass in a JSON body request
      const httpOptions = {
        headers: new HttpHeaders({
          'Content-Type': 'application/json'
        })
      };

      this.http.post("http://localhost:8181/createUser", params, httpOptions)
      .subscribe(response => {
        // Logging the response
        console.log(response)
        // Also unpacking the response to refer to the boolean values associated
        const jsonData = JSON.parse(JSON.stringify(response));
        // If things went well creating account, send user to login page
        if (jsonData.Success) {
          this.router.navigate(['/login']);
        }
        // If things did not go well creating account, keep them at the sign up page
        else {
          this.router.navigate(['/signup']);
        }
      });
    } else if (this.email == null || this.password == null){ 
      console.log("Email or Password is empty!");
    } else {
      console.log("Passwords don't match!");
    }
    
  }
}
