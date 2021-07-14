import {ComponentFixture, TestBed} from '@angular/core/testing';
import {RegisterFormComponent} from './register-form.component';
import {AuthModule} from "../../../auth.module";
import {RouterTestingModule} from "@angular/router/testing";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {NzDatePickerComponent} from "ng-zorro-antd/date-picker";
import {NzSelectComponent} from "ng-zorro-antd/select";
import {cases} from "jasmine-parameterized";

describe('RegisterFormComponent', () => {
  let component: RegisterFormComponent;
  let fixture: ComponentFixture<RegisterFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [RegisterFormComponent],
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

  cases([
    'name',
    'surname',
    'username',
    'password',
    'confirmPassword',
    'email'
  ]).it('should render input element', (fieldName) => {
    const inputElement: HTMLInputElement = fixture.nativeElement.querySelector(`input[formControlName=${fieldName}]`);
    expect(inputElement).toBeTruthy();
  });

  it('should render datepicker element', () => {
    const datePickerElement: NzDatePickerComponent = fixture.nativeElement.querySelector('nz-date-picker[formControlName=dateOfBirth]');
    expect(datePickerElement).toBeTruthy();
  });

  it('should render select element', () => {
    const selectElement: NzSelectComponent = fixture.nativeElement.querySelector('nz-select[formControlName=gender]');
    expect(selectElement).toBeTruthy();
  });


  // TODO write select options render test


  // TODO add gender and dateOfBirth validation
  cases([
    ["Adam", "Kowalski", "Kowal", "test@test.com", "Password", "Password", true],
    ["1234567890123456789012345678901234567890", "Kowalski", "Kowal", "test@test.com", "Password", "Password", false],
    ["Adam", "1234567890123456789012345678901234567890", "Kowal", "test@test.com", "Password", "Password", false],
    ["Adam", "Kowalski", "1234567890123456789012345678901234567890", "test@test.com", "Password", "Password", false],    ["Adam", "Kowalski", "Kowal", "test@test.com", "Password", "Password", true],
    ["Adam", "Kowalski", "Kowal", "test@invalid", "Password", "Password", false],
    ["Adam", "Kowalski", "Kowal", "test@test.com", "123", "Password", false],
    ["Adam", "Kowalski", "Kowal", "test@test.com", "Password", "123", false],
    ["Adam", "Kowalski", "Kowal", "test@test.com", "Password", "PasswordNotMatch", false],
  ]).it('should show errors if inputs are invalid', ([name, surname, username, email, password, confirmPassword, valid]) => {
    const {debugElement} = fixture;

    const nameInput = debugElement.nativeElement.querySelector('input[formControlName=name]')
    const surnameInput = debugElement.nativeElement.querySelector('input[formControlName=surname]')
    const usernameInput = debugElement.nativeElement.querySelector('input[formControlName=username]')
    const emailInput = debugElement.nativeElement.querySelector('input[formControlName=email]')
    const passwordInput = debugElement.nativeElement.querySelector('input[formControlName=password]')
    const confirmPasswordInput = debugElement.nativeElement.querySelector('input[formControlName=confirmPassword]')

    nameInput.value = name;
    surnameInput.value = surname;
    usernameInput.value = username;
    emailInput.value = email;
    passwordInput.value = password;
    confirmPasswordInput.value = confirmPassword;

    nameInput.dispatchEvent(new Event("input"));
    surnameInput.dispatchEvent(new Event("input"));
    usernameInput.dispatchEvent(new Event("input"));
    emailInput.dispatchEvent(new Event("input"));
    passwordInput.dispatchEvent(new Event("input"));
    confirmPasswordInput.dispatchEvent(new Event("input"));

    expect(component.validateForm.valid).toBe(valid);
  });
});
