import {ComponentFixture, TestBed} from '@angular/core/testing';

import { LoginFormComponent } from './login-form.component';
import {AuthModule} from "../../../auth.module";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {cases} from "jasmine-parameterized";
import {RouterTestingModule} from "@angular/router/testing";

describe('LoginFormComponent', () => {
  let component: LoginFormComponent;
  let fixture: ComponentFixture<LoginFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LoginFormComponent ],
      imports: [
        BrowserAnimationsModule,
        AuthModule,
        RouterTestingModule
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LoginFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  afterEach(() => {
    fixture.destroy();
  });

  it('should create LoginFormComponent', () => {
    expect(component).toBeTruthy();
  });

  it('should render input elements', () => {
    const emailInput = fixture.nativeElement.querySelector('input[formControlName=email]');
    const passwordInput = fixture.nativeElement.querySelector('input[formControlName=password]');
    expect(emailInput).toBeTruthy();
    expect(passwordInput).toBeTruthy();
  });

  cases([
    ["test@test.com", "test", true],
    ["", "test", false],
    ["test@invalid", "test", false],
    ["test@test.com", "", false]
  ]).
  it('should show errors if inputs are invalid', ([email, password, valid]) => {
    const {debugElement} = fixture;
    const emailInput: HTMLInputElement = debugElement.nativeElement.querySelector('input[formControlName=email]');
    const passwordInput: HTMLInputElement = debugElement.nativeElement.querySelector('input[formControlName=password]');

    emailInput.value = email;
    passwordInput.value = password;

    emailInput.dispatchEvent(new Event("input"));
    passwordInput.dispatchEvent(new Event("input"));

    expect(component.validateForm.valid).toBe(valid);
  });
});
