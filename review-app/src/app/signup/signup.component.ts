import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { first } from 'rxjs/operators';
import { AuthServiceService } from '../services/auth-service/auth-service.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent implements OnInit {

  form: FormGroup;
  loading = false;
  submitted = false;
  token: any;
  refreshToken: any;
  id: any

  constructor(private formBuilder: FormBuilder,
    private route: ActivatedRoute,
    private router: Router,
    private authService: AuthServiceService
  ) { }

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      username: ['', Validators.required],
      password: ['', [Validators.required, Validators.minLength(6)]]
    });
  }
  get f() { return this.form.controls; }

  onSubmit() {
    this.submitted = true;

    // // reset alerts on submit
    // this.alertService.clear();

    // // stop here if form is invalid
    if (this.form.invalid) {
      return;
    }

    this.loading = true;
    this.authService.registerUser({ 'email': this.f.username.value, 'password': this.f.password.value, 'firstName': this.f.firstName.value, 'lastName': this.f.lastName.value })
      .pipe(first())
      .subscribe(
        data => {
          console.log(data);
          if (data != null) {
            this.authService.setUserLoggedIn(true);
            this.router.navigate(['/']);
            this.token = data.token;
            this.id = data.id;
            console.log("ID:" + data.id);
            this.refreshToken = data.refreshToken;
            this.authService.storeUserData(this.id, this.token, this.f.username.value, this.refreshToken)
          }
        },
        error => {
          this.loading = false;
          return;
        });

  }

}
