import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AuthContentComponent } from './auth-content.component';
import {RouterTestingModule} from "@angular/router/testing";
import {AuthModule} from "../../../auth.module";

describe('LoginSectionComponent', () => {
  let component: AuthContentComponent;
  let fixture: ComponentFixture<AuthContentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AuthContentComponent ],
      imports: [RouterTestingModule, AuthModule]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AuthContentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
