import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginFormComponent } from './components/organisms/login-form/login-form.component';
import { LoginSectionComponent } from './components/templates/login-section/login-section.component';
import {ReactiveFormsModule} from "@angular/forms";
import {NzFormModule} from "ng-zorro-antd/form";
import {NzCheckboxModule} from "ng-zorro-antd/checkbox";
import {NzButtonModule} from "ng-zorro-antd/button";
import {NzInputModule} from "ng-zorro-antd/input";



@NgModule({
  declarations: [
    LoginFormComponent,
    LoginSectionComponent
  ],
  exports: [
    LoginSectionComponent
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NzFormModule,
    NzInputModule,
    NzCheckboxModule,
    NzButtonModule
  ]
})
export class AuthModule { }
