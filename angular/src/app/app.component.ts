import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <div class="root__container">
      <nz-layout>
        <nz-layout>
          <nz-header class="root__container__header">
            <img class="logo" [src]=assetsPath alt="jigdra logo" [routerLink]="'/'">
            <nz-space [nzSize]="size" class="auth">
              <button *nzSpaceItem nz-button nzType="primary" [routerLink]="'/login'">Login</button>
              <button *nzSpaceItem nz-button nzType="default" >Register</button>
            </nz-space>
          </nz-header>
          <nz-content class="root__container__inner-content">
            <nz-card class="root__container__inner-content__card" nzTitle="Sign in">
              <router-outlet></router-outlet>
            </nz-card>
          </nz-content>
          <nz-footer class="root__container__footer">Jigdra @2021 Implemented By Angular</nz-footer>
        </nz-layout>
        <nz-sider [nzWidth]="600"></nz-sider>
      </nz-layout>
    </div>
  `,
  styleUrls: ["app.component.scss"]
})
export class AppComponent {

  assetsPath: string = 'assets/img/logo.png';

  size: 'small' | 'middle' | 'large' | number = "small";
}
