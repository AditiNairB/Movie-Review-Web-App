import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthServiceService } from '../auth-service/auth-service.service';

@Injectable({
  providedIn: 'root'
})
export class ReviewService {

  apiUrl = 'http://localhost:8080/';
  movieID ="";

  setMovieId(id:string){
    this.movieID=id;
  }
  constructor(private http: HttpClient, private authServiceService: AuthServiceService) { }

  getReviewsForMovie(id:String): Observable<any> {
    return this.http.get(this.apiUrl + 'api/review/'+id, {
    });
  }

  addReview(review: Object): Observable<any> {
    console.log(sessionStorage.getItem("id_token"));

    return this.http.post(this.apiUrl + 'api/review/'+this.movieID+'/', review, {
    headers: new HttpHeaders({
      "Content-Type": "application/json",
      'Authorization': 'Bearer '+sessionStorage.getItem("id_token")
    }, )
    });


  }
}
