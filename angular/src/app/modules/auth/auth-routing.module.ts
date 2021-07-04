import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {LoginFormComponent} from "./components/organisms/login-form/login-form.component";
import {RegisterFormComponent} from "./components/organisms/register-form/register-form.component";
import {AuthViewComponent} from "./components/views/auth-view/auth-view.component";

const routes: Routes = [
  {
    path: '',
    component: AuthViewComponent,
    children: [
      {
        path: "login", component: LoginFormComponent
      },
      {
        path: "register", component: RegisterFormComponent
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AuthRoutingModule { }
