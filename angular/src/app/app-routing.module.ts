import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {LoginSectionComponent} from "./modules/auth/components/templates/login-section/login-section.component";

const routes: Routes = [
  {path: 'login', component: LoginSectionComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
