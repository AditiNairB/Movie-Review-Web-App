import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { JwtHelperService } from '@auth0/angular-jwt';

@Injectable({
  providedIn: 'root'
})
export class AuthServiceService {

  apiUrl = 'http://localhost:8080/';
  private userLoggedIn = new Subject<boolean>();
  timeout: number;
  authToken: any;
  user: any;
  refreshTokenValue: any;

  constructor(private http: HttpClient) {
    this.userLoggedIn.next(false);
  }

  setUserLoggedIn(userLoggedIn: boolean) {
    this.userLoggedIn.next(userLoggedIn);
  }

  getUserLoggedIn(): Observable<boolean> {
    return this.userLoggedIn.asObservable();
  }

  authenticateUser(user: Object): Observable<any> {
    return this.http.post(this.apiUrl + 'api/auth/login', user, {
      headers: new HttpHeaders({
        "Content-Type": "application/json"
      })
    });
  }

  registerUser(user: Object): Observable<any> {
    return this.http.post(this.apiUrl + 'api/auth/signup', user, {
      headers: new HttpHeaders({
        "Content-Type": "application/json"
      })
    });
  }

  storeUserData(id, token, user, refreshToken) {
    const jwtHelper = new JwtHelperService();
    this.timeout = jwtHelper.getTokenExpirationDate(token).valueOf() - new Date().valueOf();
    sessionStorage.setItem("id_token", token);
    sessionStorage.setItem("user", JSON.stringify(user));
    sessionStorage.setItem("id", JSON.stringify(id));
    this.authToken = token;
    this.user = user;
    this.refreshTokenValue = refreshToken;
    // this.emit({ username: this.user.username });
    // this.expirationCounter(this.timeout);
  }

  getSessionData(): any {
    return sessionStorage.getItem("user");
  }

  getSessionToken(): any {
    return sessionStorage.getItem("id_token");
  }
}
