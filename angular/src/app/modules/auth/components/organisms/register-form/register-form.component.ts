import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {emailRegExp} from "../../../../../shared/regexps/regexps";
import customValidator from "../../../../../shared/validators/custom-validator";
import StatusValidator, {ValidateStatus} from "../../../../../shared/validators/status-validator";

@Component({
  selector: 'app-register-form',
  template: `
    <form nz-form [formGroup]="validateForm" class="register-form" (ngSubmit)="submitForm()">
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('username')"
                         nzErrorTip="Please input your username!">
          <nz-input-group>
            <input type="text" nz-input formControlName="username" placeholder="Username"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('name')" nzErrorTip="Please input your name!">
          <nz-input-group>
            <input type="text" nz-input formControlName="name" placeholder="Name"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('surname')"
                         nzErrorTip="Please input your surname!">
          <nz-input-group>
            <input type="text" nz-input formControlName="surname" placeholder="Surname"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('email')"
                         nzErrorTip="The input is not valid e-mail!">
          <nz-input-group>
            <input type="text" nz-input formControlName="email" placeholder="Email"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('password')"
                         nzErrorTip="Please input your password!">
          <nz-input-group nzPrefixIcon="lock">
            <input type="password" nz-input formControlName="password" placeholder="Password"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('confirmPassword')"
                         nzErrorTip="Two passwords that you enter is inconsistent!">
          <nz-input-group nzPrefixIcon="lock">
            <input type="password" nz-input formControlName="confirmPassword" placeholder="Confirm password"/>
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
          <nz-select nzPlaceHolder="Select your gender" formControlName="gender">
            <nz-option data-cy="Male" nzValue="Male" nzLabel="Male"></nz-option>
            <nz-option data-cy="Female" nzValue="Female" nzLabel="Female"></nz-option>
            <nz-option data-cy="Other" nzValue="Other" nzLabel="Other"></nz-option>
            <nz-option data-cy="Unknown" nzValue="Unknown" nzLabel="Prefer not to say"></nz-option>
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
  validateStatus!: ValidateStatus

  constructor(private fb: FormBuilder) {
  }

  submitForm(): void {
    Object.values(this.validateForm.controls).forEach(control => {
      control.markAsDirty();
      control.updateValueAndValidity();
    })

    this.fetchRegister();
    console.log(this.validateForm.valid ? "Works" : "Doesn't work");
  }

  fetchRegister() {
    this.isLoading = true;
    setTimeout(() => {
      this.isLoading = false;
    }, 2000);
  }

  ngOnInit(): void {
    this.validateForm = this.fb.group({
        name: [null, [Validators.required]],
        surname: [null, [Validators.required]],
        username: [null, [Validators.required, Validators.minLength(8)]],
        password: [null, [Validators.required]],
        confirmPassword: [null, [Validators.required]],
        email: [null, [Validators.required, Validators.pattern(emailRegExp)]],
        dateOfBirth: [null],
        gender: [null]
      },
      {
        validators: [customValidator.stringsMatch('password', 'confirmPassword')]
      });
    this.validateStatus = StatusValidator.validateStatus(this.validateForm);
  }
}
