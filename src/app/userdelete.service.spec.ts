import { TestBed } from '@angular/core/testing';

import { UserdeleteService } from './userdelete.service';

describe('UserdeleteService', () => {
  let service: UserdeleteService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UserdeleteService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
