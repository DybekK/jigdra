import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginFormComponent } from './components/organisms/login-form/login-form.component';
import {AuthContentComponent} from './components/templates/auth-content/auth-content.component';
import {ReactiveFormsModule} from "@angular/forms";
import {NzFormModule} from "ng-zorro-antd/form";
import {NzCheckboxModule} from "ng-zorro-antd/checkbox";
import {NzButtonModule} from "ng-zorro-antd/button";
import {NzInputModule} from "ng-zorro-antd/input";
import {NzLayoutModule} from "ng-zorro-antd/layout";
import {RouterModule} from "@angular/router";
import {NzSpaceModule} from "ng-zorro-antd/space";
import {NzCardModule} from "ng-zorro-antd/card";
import {AuthViewComponent} from "./components/views/auth-view/auth-view.component";
import { RegisterFormComponent } from './components/organisms/register-form/register-form.component';
import {AuthRoutingModule} from "./auth-routing.module";
import { AuthHeaderComponent } from './components/atoms/auth-header/auth-header.component';
import {SharedModule} from "../../shared/shared.module";
import {NzDatePickerModule} from "ng-zorro-antd/date-picker";
import {NzSelectModule} from "ng-zorro-antd/select";
import {NzIconModule} from "ng-zorro-antd/icon";
import { UserOutline, LockOutline } from '@ant-design/icons-angular/icons';
import { IconDefinition } from '@ant-design/icons-angular';
import {NzTypographyModule} from "ng-zorro-antd/typography";

const icons: IconDefinition[] = [UserOutline, LockOutline];

@NgModule({
  declarations: [
    LoginFormComponent,
    AuthContentComponent,
    AuthViewComponent,
    RegisterFormComponent,
    AuthHeaderComponent,
  ],
  exports: [
    AuthContentComponent,
    LoginFormComponent,
    AuthViewComponent,
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NzFormModule,
    NzInputModule,
    NzCheckboxModule,
    NzButtonModule,
    NzLayoutModule,
    NzIconModule,
    RouterModule,
    NzSpaceModule,
    NzCardModule,
    AuthRoutingModule,
    SharedModule,
    NzDatePickerModule,
    NzSelectModule,
    NzIconModule.forChild(icons),
    NzTypographyModule
  ]
})
export class AuthModule { }
