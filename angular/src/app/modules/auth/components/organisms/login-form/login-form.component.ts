import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import {emailRegExp} from "../../../../../shared/regexps/regexps";
import StatusValidator, {ValidateStatus} from "../../../../../shared/validators/status-validator";

@Component({
  selector: 'app-login-form',
  template: `
    <form nz-form [formGroup]="validateForm" class="login-form" (ngSubmit)="submitForm()">
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('email')" nzErrorTip="The input is not valid e-mail!">
          <nz-input-group nzPrefixIcon="user">
            <input type="text" nz-input formControlName="email" placeholder="Email" />
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzHasFeedback [nzValidateStatus]="validateStatus('password')" nzErrorTip="Please input your password!">
          <nz-input-group nzPrefixIcon="lock">
            <input type="password" nz-input formControlName="password" placeholder="Password" />
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <div nz-row class="login-form--margin">
        <div nz-col [nzSpan]="12">
          <label nz-checkbox formControlName="remember">
            <span>Remember me</span>
          </label>
        </div>
        <div nz-col>
          <a [routerLink]="'/forgot-password'" class="login-form-forgot">Forgot password</a>
        </div>
      </div>
        <button [nzLoading]="isLoading" nz-button class="login-form__button" [nzType]="'primary'">Log in</button>
        Or <a [routerLink]="'/register'"> register now! </a>
    </form>
  `,
  styleUrls: ["login-form.component.scss"]
})
export class LoginFormComponent implements OnInit {
  isLoading = false;
  validateForm!: FormGroup;
  validateStatus!: ValidateStatus;

  constructor(private fb: FormBuilder) {}

  submitForm(): void {
    Object.values(this.validateForm.controls).forEach(control => {
      control.markAsDirty();
      control.updateValueAndValidity();
    });

    this.fetchRegister();
    if(this.validateForm.valid) {
      console.log("went nice");
    }
  }

  // method only for testing purposes
  fetchRegister() {
    this.isLoading = true;
    setTimeout(() => {
    this.isLoading = false;
    }, 2000);
  }

  ngOnInit(): void {
    this.validateForm = this.fb.group({
      email: [null, [Validators.required, Validators.pattern(emailRegExp)]],
      password: [null, [Validators.required]],
      remember: [true]
    });
    this.validateStatus = StatusValidator.validateStatus(this.validateForm);
  }
}
