import { TestBed } from '@angular/core/testing';

import { AuthHttpClient } from './auth-http.client';

describe('AuthHttpService', () => {
  let service: AuthHttpClient;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AuthHttpClient);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
