import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class MovieSearchService {

  textValue: any;
  constructor() { }

  get data(): any {
    return this.textValue;
  }

  set data(val: any) {
    this.textValue = val;
    console.log(this.textValue);
  }
}
