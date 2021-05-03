import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { MovieObject } from '../movieobject.model';
import { MovieService } from '../services/movie-service/movie.service';
import { ReviewObject } from '../reviewobject.model';
import { ReviewService } from '../services/review-service/review.service';
import { first } from 'rxjs/operators';
import { AuthServiceService } from '../services/auth-service/auth-service.service';
import { HeaderComponent } from '../header/header.component';

@Component({
  selector: 'app-movie-details',
  templateUrl: './movie-details.component.html',
  styleUrls: ['./movie-details.component.scss']
})
export class MovieDetailsComponent implements OnInit {

  public movieId: string;

  public movieDetails: MovieObject;

  public reviews: ReviewObject[];

  // const fieldset = document.fo
  currentRate = 0;
  msg: String = '';
  private authToken: String;
  showAddReview: boolean;
  averageReview: String;

  constructor(public router: Router, public route: ActivatedRoute, public movieService: MovieService
    , public reviewService: ReviewService, public authServiceService: AuthServiceService, public headerComponent: HeaderComponent) {
    route.url.subscribe(() => {
      this.movieId = route.snapshot.firstChild.url[0].path;
      console.log(this.showAddReview);
    });

    this.headerComponent.elements[0].style.visibility = "visible";
    this.headerComponent.elements[1].style.visibility = "hidden";
    if (authServiceService.getSessionData() != null) {
      this.headerComponent.elements[0].style.visibility = "hidden";
      this.headerComponent.elements[1].style.visibility = "visible";
    }
  }

  ngOnInit(): void {

    this.showAddReview = false;
    this.movieService.getMovieDetail(this.movieId).subscribe((movieDet: MovieObject) => {
      this.movieDetails = movieDet;
      console.log(this.movieDetails);
    });
    this.reviewService.setMovieId(this.movieId)

    this.reviewService.getReviewsForMovie(this.movieId).subscribe((movieReviews: ReviewObject[]) => {
      this.reviews = movieReviews;
      console.log(this.reviews);
      var arrayLength = this.reviews.length;
      var average = 0;
      for (var i = 0; i < arrayLength; i++) {
        average = average + parseInt(this.reviews[i].rating);

      }
      this.averageReview = String(average / arrayLength);
    });
    console.log(this.authServiceService.getSessionToken());
    if (this.authServiceService.getSessionToken() != null) {
      this.showAddReview = true;
      this.authToken = this.authServiceService.getSessionToken();
    }



  }

  onSubmit() {
    // console.log(this.currentRate);
    // console.log(this.msg);



    this.reviewService.addReview({ 'rating': this.currentRate, 'description': this.msg })
      .pipe(first())
      .subscribe(
        data => {
          console.log(data);
          if (data != null) {
            console.log(data);
            location.reload();
          }
        },
        error => {
          return;
        });

  }
}
