import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {LoginFormComponent} from "./components/organisms/login-form/login-form.component";
import {RegisterFormComponent} from "./components/organisms/register-form/register-form.component";
import {AuthViewComponent} from "./components/views/auth-view/auth-view.component";
import {MailConfirmationComponent} from "./components/organisms/mail-confirmation/mail-confirmation.component";

export const routes: Routes = [
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full'
  },
  {
    path: '',
    component: AuthViewComponent,
    children: [
      {
        path: "login", component: LoginFormComponent
      },
      {
        path: "register", component: RegisterFormComponent
      },
      {
        path: "emailConfirmation", component: MailConfirmationComponent
      }
    ]
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AuthRoutingModule { }
