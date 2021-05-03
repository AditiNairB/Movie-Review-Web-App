import { Component, Input, OnInit } from '@angular/core';
import { MovieObject } from '../movieobject.model';
import { MovieSearchService } from '../services/movie-search-service/movie-search.service';
import { MovieService } from '../services/movie-service/movie.service';
import { HeaderComponent } from '../header/header.component';
import { AuthServiceService } from '../services/auth-service/auth-service.service';

@Component({
  selector: 'app-movie-list',
  templateUrl: './movie-list.component.html',
  styleUrls: ['./movie-list.component.scss']
})

export class MovieListComponent implements OnInit {

  public movies: MovieObject[];
  public textValue: any;

  constructor(public movieService: MovieService, public searchValue: MovieSearchService, public headerComponent: HeaderComponent, public authServiceService: AuthServiceService) {
    this.headerComponent.elements[0].style.visibility = "visible";
    this.headerComponent.elements[1].style.visibility = "hidden";
    if (authServiceService.getSessionData() != null) {
      this.headerComponent.elements[0].style.visibility = "hidden";
      this.headerComponent.elements[1].style.visibility = "visible";
    }
  }

  ngOnInit(): void {
    this.getMovieList();
  }
  getMovieList() {
    this.movieService.getAllMovies().subscribe((movieList: MovieObject[]) => {
      this.movies = movieList;
    })
  }
  public submitOnClick() {
    this.textValue = (<HTMLInputElement>document.getElementById('searchText')).value;
    this.searchValue.data = this.textValue;
    console.log("searchValue: " + this.searchValue.data);
  }

  callMovie() {
    console.log("HERE");
  }
}


