import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MovieService {

  apiUrl = 'http://localhost:8080/';

  constructor(private http: HttpClient) { }

  getAllMovies(): Observable<any> {
    return this.http.get(this.apiUrl + 'api/movies/', {
      headers: new HttpHeaders({
        "Content-Type": "application/json"
      })
    });
  }

  getMovieDetail(id: String): Observable<any> {
    return this.http.get(this.apiUrl + 'api/movies/' + id, {
      // headers: new HttpHeaders({
      //   "Content-Type": "application/json"
      // })
    });
  }

  getMovie(param: Object): Observable<any> {
    var audioUrl = "\"" + param + "\"";
    const body = { movieName: param };
    return this.http.post(this.apiUrl + 'api/movies/', body, {
      headers: new HttpHeaders({
        "Content-Type": "application/json"
      })
    });
  }
}
