import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RegisterFormComponent } from './register-form.component';
import {AuthModule} from "../../../auth.module";
import {RouterTestingModule} from "@angular/router/testing";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";

describe('RegisterFormComponent', () => {
  let component: RegisterFormComponent;
  let fixture: ComponentFixture<RegisterFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegisterFormComponent ],
      imports: [
        BrowserAnimationsModule,
        AuthModule,
        RouterTestingModule
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RegisterFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create RegisterFormComponent', () => {
    expect(component).toBeTruthy();
  });

  it('should render input elements', () => {
    const nameInput : HTMLInputElement = fixture.nativeElement.querySelector('input[formControlName=name]');
    const surnameInput : HTMLInputElement = fixture.nativeElement.querySelector('input[formControlName=surname]');
    const usernameInput : HTMLInputElement = fixture.nativeElement.querySelector('input[formControlName=username]');
    const passwordInput : HTMLInputElement = fixture.nativeElement.querySelector('input[formControlName=password]');
    const confirmPasswordInput : HTMLInputElement = fixture.nativeElement.querySelector('input[formControlName=confirmPassword]');
    const emailInput : HTMLInputElement = fixture.nativeElement.querySelector('input[formControlName=email]');

    expect(nameInput).toBeTruthy();
    expect(surnameInput).toBeTruthy();
    expect(usernameInput).toBeTruthy();
    expect(passwordInput).toBeTruthy();
    expect(confirmPasswordInput).toBeTruthy();
    expect(emailInput).toBeTruthy();
  })
});
