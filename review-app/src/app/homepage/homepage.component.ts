import { Component, Injectable, OnInit } from '@angular/core';
import { HeaderComponent } from '../header/header.component';
import { AuthServiceService } from '../services/auth-service/auth-service.service';
import { MovieSearchService } from '../services/movie-search-service/movie-search.service';

@Injectable({
  providedIn: 'root'
})

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})
export class HomepageComponent implements OnInit {

  public elements: any[];
  public textValue: any;

  constructor(public headerComponent: HeaderComponent
    , public authServiceService: AuthServiceService, public searchValue: MovieSearchService) {
    this.headerComponent.elements[0].style.visibility = "visible";
    this.headerComponent.elements[1].style.visibility = "hidden";
    if (authServiceService.getSessionData() != null) {
      this.headerComponent.elements[0].style.visibility = "hidden";
      this.headerComponent.elements[1].style.visibility = "visible";
    }

  }

  ngOnInit(): void {
    // this.submitOnClick();
  }

  public submitOnClick() {
    this.textValue = (<HTMLInputElement>document.getElementById('searchText')).value;
    this.searchValue.data = this.textValue;
    console.log("searchValue: " + this.searchValue.data);
  }

  homeSlider = { items: 1, dots: true, nav: true };

}
