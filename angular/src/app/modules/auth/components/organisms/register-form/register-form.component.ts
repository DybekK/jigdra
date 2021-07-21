import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {emailRegExp, onlyLettersRegExp} from "../../../../../shared/regexps/regexps";
import customValidator from "../../../../../shared/validators/custom-validator";
import StatusValidator, {ValidateStatus} from "../../../../../shared/validators/status-validator";
import {AuthService} from "../../../services/auth/auth.service";
import {RegisterDto} from "../../../interfaces/RegisterDto";
import {finalize} from "rxjs/operators";

@Component({
  selector: 'app-register-form',
  template: `
    <form nz-form [formGroup]="validateForm" class="register-form" (ngSubmit)="submitForm()">
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('username')"
                         nzErrorTip="Please input your username!">
          <nz-input-group>
            <input type="text" nz-input formControlName="username" placeholder="Username" [maxLength]="32"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('name')" nzErrorTip="Please input your name!">
          <nz-input-group>
            <input type="text" nz-input formControlName="name" placeholder="Name" [maxLength]="32"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('surname')"
                         nzErrorTip="Please input your surname!">
          <nz-input-group>
            <input type="text" nz-input formControlName="surname" placeholder="Surname" [maxLength]="32"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('email')"
                         nzErrorTip="The input is not valid e-mail!">
          <nz-input-group>
            <input type="text" nz-input formControlName="email" placeholder="Email" [maxLength]="255"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('password')"
                         nzErrorTip="Please input your password!">
          <nz-input-group nzPrefixIcon="lock">
            <input type="password" nz-input formControlName="password" placeholder="Password" [minLength]="6"
                   [maxLength]="20"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('confirmPassword')"
                         nzErrorTip="Two passwords that you enter is inconsistent!">
          <nz-input-group nzPrefixIcon="lock">
            <input type="password" nz-input formControlName="confirmPassword" placeholder="Confirm password"
                   [minLength]="6" [maxLength]="20"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control>
          <nz-date-picker class="register-form__full-width" formControlName="dateOfBirth"></nz-date-picker>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control class="register-form__full-width">
          <nz-select nzPlaceHolder="Select your gender" formControlName="gender" [nzOptions]="Gender">
          </nz-select>
        </nz-form-control>
      </nz-form-item>
      <div nz-row class="register-form--margin">
      </div>
      <button [nzLoading]="isLoading" nz-button class="register-form__button" [nzType]="'primary'">Register</button>
      Have an account? <a [routerLink]="'/login'"> Login now! </a>
    </form>
  `,
  styleUrls: ['./register-form.component.scss']
})
export class RegisterFormComponent implements OnInit {
  isLoading = false;
  validateForm!: FormGroup;
  validateStatus!: ValidateStatus;
  responseError!: string;


  constructor(
    private fb: FormBuilder,
    private authService: AuthService,
  ) {
  }

  Gender: any = [
    {label: 'Male', value: 'Male'},
    {label: 'Female', value: 'Female'},
    {label: 'Other', value: 'Other'},
    {label: 'Prefer not to say', value: 'Unknown'}
  ]

  submitForm(): void {
    Object.values(this.validateForm.controls).forEach(control => {
      control.markAsDirty();
      control.updateValueAndValidity();
    });

    if (this.validateForm.valid) {
      this.isLoading = true;
      const value: RegisterDto = this.validateForm.value;
      this.authService.registerUser(value).pipe(
        finalize(() => this.isLoading = false)
      ).subscribe({
        error: this.handleError
      });
    }
  }

  handleError(error: RegisterError) {
    Object.entries(error).forEach(([key,value]) => {
      const {controls} = this.validateForm;
      if (controls[key]) {
        controls[key].setErrors({error: value});
      }
    })
  }

  ngOnInit(): void {
    this.validateForm = this.fb.group({
        name: [null, [Validators.required, Validators.pattern(onlyLettersRegExp), Validators.maxLength(32)]],
        surname: [null, [Validators.required, Validators.pattern(onlyLettersRegExp), Validators.maxLength(32)]],
        username: [null, [Validators.required, Validators.maxLength(32)]],
        password: [null, [Validators.required, Validators.minLength(6), Validators.maxLength(20)]],
        confirmPassword: [null, [Validators.required, Validators.minLength(6), Validators.maxLength(20)]],
        email: [null, [Validators.required, Validators.maxLength(255), Validators.pattern(emailRegExp)]],
        dateOfBirth: [null],
        gender: [null]
      },
      {
        validators: [customValidator.stringsMatch('password', 'confirmPassword')]
      });
    this.validateStatus = StatusValidator.validateStatus(this.validateForm);
  }
}


// TODO: Wait for correct backend error response
type RegisterErrorKey = 'username' | 'email'
type RegisterError = {
  [key in RegisterErrorKey]?: string
}

