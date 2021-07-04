import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {AuthViewComponent} from "./modules/auth/components/views/auth-view/auth-view.component";
import {LoginFormComponent} from "./modules/auth/components/organisms/login-form/login-form.component";
import {RegisterFormComponent} from "./modules/auth/components/organisms/register-form/register-form.component";
import {NotFoundComponent} from "./modules/errors/not-found/not-found.component";

const routes: Routes = [
  {
    path: 'home',
    component: AuthViewComponent,
    loadChildren: () => import("./modules/auth/auth.module").then(m => m.AuthModule)
  },
  {
    path: 'error', component: NotFoundComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
