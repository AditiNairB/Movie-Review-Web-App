import { TestBed } from '@angular/core/testing';

import { HeaderComponentService } from './header-component.service';

describe('HeaderComponentService', () => {
  let service: HeaderComponentService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(HeaderComponentService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
