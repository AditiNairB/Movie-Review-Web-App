import { Component, Injectable, OnInit } from '@angular/core';
import { HeaderComponentService } from '../services/header-component-service/header-component.service';

@Injectable({
  providedIn: 'root'
})

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  public elements: any[];

  constructor(public headerComponentService: HeaderComponentService) {
    this.elements = [document.getElementById("login"), document.getElementById("logout")]
    this.headerComponentService.data = this.elements;
    console.log(headerComponentService.data)
  }

  ngOnInit(): void {
  }

}
