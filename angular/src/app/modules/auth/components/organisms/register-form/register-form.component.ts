import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {emailRegExp} from "../../../../../shared/regexps/regexps";

@Component({
  selector: 'app-register-form',
  template: `
    <form nz-form [formGroup]="validateForm" class="register-form" (ngSubmit)="submitForm()">
      <nz-form-item>
        <nz-form-control nzErrorTip="Please input your Username!">
          <nz-input-group nzPrefixIcon="user">
            <input type="text" nz-input formControlName="username" placeholder="Username"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzErrorTip="Please input your Name!">
          <nz-input-group nzPrefixIcon="user">
            <input type="text" nz-input formControlName="name" placeholder="Name"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzErrorTip="Please input your Surname!">
          <nz-input-group nzPrefixIcon="user">
            <input type="text" nz-input formControlName="surname" placeholder="Surname"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzErrorTip="The input is not valid E-mail!">
          <nz-input-group nzPrefixIcon="user">
            <input type="text" nz-input formControlName="email" placeholder="Email"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzErrorTip="Please input your Password!">
          <nz-input-group nzPrefixIcon="lock">
            <input type="password" nz-input formControlName="password" placeholder="Password"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzErrorTip="Please confirm your Password!">
          <nz-input-group nzPrefixIcon="lock">
            <input type="password" nz-input formControlName="confirmPassword" placeholder="Confirm Password"/>
          </nz-input-group>
        </nz-form-control>
      </nz-form-item>
      <nz-form-item>
        <nz-form-control nzErrorTip="Please enter your Date of Birth!">
          <nz-date-picker class="register-form-datepicker" formControlName="dateOfBirth"></nz-date-picker>
        </nz-form-control>
      </nz-form-item>
      <nz-select class="register-form-select" formControlName="gender" nzPlaceHolder="Select your gender">
        <nz-option nzValue="Male" nzLabel="Male"></nz-option>
        <nz-option nzValue="Female" nzLabel="Female"></nz-option>
        <nz-option nzValue="Other" nzLabel="Other"></nz-option>
        <nz-option nzValue="Unknown" nzLabel="Prefer not to say"></nz-option>
      </nz-select>
      <div nz-row class="login-form-margin">
      </div>
      <button nz-button class="login-form-button login-form-margin" [nzType]="'primary'">Register</button>
      Have an account? <a> Login now! </a>
    </form>
  `,
  styleUrls: ['./register-form.component.scss']
})
export class RegisterFormComponent implements OnInit {
  validateForm!: FormGroup;

  constructor(private fb: FormBuilder) {
  }

  submitForm(): void {
    Object.values(this.validateForm.controls).forEach(control => {
      control.markAsDirty();
      control.updateValueAndValidity();
      console.log(control.value);
    })

    console.log(this.validateForm.valid ? "Works" : "Doesn't work");
  }

  ngOnInit(): void {
    this.validateForm = this.fb.group({
      name: [null, [Validators.required]],
      surname: [null, [Validators.required]],
      password: [null, [Validators.required]],
      username: [null, [Validators.required]],
      confirmPassword: [null, [Validators.required]],
      email: [null, [Validators.required, Validators.pattern(emailRegExp)]],
      dateOfBirth: [null],
      gender: [null]
    });
  }
}
