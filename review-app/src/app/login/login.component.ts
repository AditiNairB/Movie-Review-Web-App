import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuthServiceService } from '../services/auth-service/auth-service.service';
import { first } from 'rxjs/operators';
import { HeaderComponent } from '../header/header.component';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  loginForm: FormGroup;
  submitted = false;
  loading = false;
  token: any;
  refreshToken: any;
  id: any;

  constructor(
    private formBuilder: FormBuilder,
    private router: Router,
    private authService: AuthServiceService,
    private headerComponent: HeaderComponent
  ) { }

  ngOnInit(): void {
    this.loginForm = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
    if (this.authService.getSessionData() != null) {
      this.headerComponent.elements[0].style.visibility = "visible";
      this.headerComponent.elements[1].style.visibility = "hidden";
      sessionStorage.clear();
    }
  }

  clearSession() {
    sessionStorage.clear();
  }

  get f() { return this.loginForm.controls; }

  onSubmit() {
    this.submitted = true;
    // console.log(this.loginForm);
    if (this.loginForm.invalid) {
      return;
    }
    // console.log(this.f.username.value);
    this.loading = true;

    this.authService.authenticateUser({ 'email': this.f.username.value, 'password': this.f.password.value })
      .pipe(first())
      .subscribe(
        data => {
          console.log(data);
          this.authService.setUserLoggedIn(true);
          this.router.navigate(['/']);
          this.token = data.token;
          this.id = data.id;
          this.refreshToken = data.refreshToken;
          console.log(this.token);
          this.authService.storeUserData(this.id, this.token, this.f.username.value, this.refreshToken);
        },
        error => {
          this.loading = false;
          return;
        });

  }

}
