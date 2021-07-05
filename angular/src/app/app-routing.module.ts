import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {AuthViewComponent} from "./modules/auth/components/views/auth-view/auth-view.component";
import {NotFoundComponent} from "./modules/errors/not-found/not-found.component";

const routes: Routes = [
  {
    path: 'home',
    component: AuthViewComponent,
    loadChildren: () => import("./modules/auth/auth.module").then(m => m.AuthModule)
  },
  {
    path: 'error', component: NotFoundComponent
  },
  { path: 'user', loadChildren: () => import('./modules/user/user.module').then(m => m.UserModule) },
  {
    path: '**', redirectTo: '/error'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
