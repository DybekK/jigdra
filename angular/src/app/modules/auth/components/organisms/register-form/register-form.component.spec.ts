import {ComponentFixture, TestBed} from '@angular/core/testing';
import {RegisterFormComponent} from './register-form.component';
import {AuthModule} from "../../../auth.module";
import {RouterTestingModule} from "@angular/router/testing";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {NzDatePickerComponent} from "ng-zorro-antd/date-picker";
import {cases} from "jasmine-parameterized";
import {NzSelectComponent} from "ng-zorro-antd/select";

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
  ]).it('should render input elements', (fieldName) => {
    const inputElement: HTMLInputElement = fixture.nativeElement.querySelector(`input[formControlName=${fieldName}]`);
    expect(inputElement).toBeTruthy();
  });

  it('should render datepicker element', () => {
    const datePickerElement: NzDatePickerComponent = fixture.nativeElement.querySelector('nz-date-picker[formControlName=dateOfBirth]');
    expect(datePickerElement).toBeTruthy();
  });

  it('should render select element', () => {
    const selectElement: NzSelectComponent = fixture.nativeElement.querySelector('nz-select[formControlName=gender]')
    expect(selectElement).toBeTruthy();
  });

  // TODO test to refactor
  cases([
    "Male",
    "Female",
    "Other",
    "Unknown"
  ]).it('should render gender select options', (optionValue) => {
    const optionElement = Array.from(document.querySelectorAll('.ant-select-item-option-content'))
      .find(el => el.textContent === `${optionValue}`);
    expect(optionElement).toBeTruthy();
  });
});
