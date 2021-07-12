import {ComponentFixture, TestBed} from '@angular/core/testing';

import { LoginFormComponent } from './login-form.component';
import {AuthModule} from "../../../auth.module";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
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

  it('should submit form if all inputs are valid', () => {
    const {debugElement} = fixture;
    const emailInput: HTMLInputElement = debugElement.nativeElement.querySelector('input[formControlName=email]');
    const passwordInput: HTMLInputElement = debugElement.nativeElement.querySelector('input[formControlName=password]');

    emailInput.value = "test@test.com";
    passwordInput.value = "test";

    emailInput.dispatchEvent(new Event("input"));
    passwordInput.dispatchEvent(new Event("input"));

    expect(component.validateForm.valid).toBeTrue();
  });

  it('should show errors if inputs are invalid', () => {
    const {debugElement} = fixture;
    const emailInput: HTMLInputElement = debugElement.nativeElement.querySelector('input[formControlName=email]');
    const passwordInput: HTMLInputElement = debugElement.nativeElement.querySelector('input[formControlName=password]');

    // email is empty
    emailInput.value = "";
    passwordInput.value = "test";

    emailInput.dispatchEvent(new Event("input"));
    passwordInput.dispatchEvent(new Event("input"));

    expect(component.validateForm.valid).toBeFalse();

    // email is invalid
    emailInput.value = "test@invalid";
    passwordInput.value = "test";

    emailInput.dispatchEvent(new Event("input"));
    passwordInput.dispatchEvent(new Event("input"));

    expect(component.validateForm.valid).toBeFalse();

    // password is empty
    emailInput.value = "test@test.com";
    passwordInput.value = "";

    emailInput.dispatchEvent(new Event("input"));
    passwordInput.dispatchEvent(new Event("input"));

    expect(component.validateForm.valid).toBeFalse();
  });
});
