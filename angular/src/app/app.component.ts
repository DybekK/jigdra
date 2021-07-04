import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <nz-layout>
      <nz-header>
        <div class="logo"><img src="assets/img/logo.png" alt="jigdra logo"></div>
        <nz-space [nzSize]="size" class="auth">
          <button *nzSpaceItem nz-button nzType="primary">Login</button>
          <button *nzSpaceItem nz-button nzType="default">Register</button>
        </nz-space>
      </nz-header>
      <nz-content>
        <nz-breadcrumb>
          <nz-breadcrumb-item>Home</nz-breadcrumb-item>
        </nz-breadcrumb>
        <div class="inner-content">Content</div>
      </nz-content>
      <nz-footer>Jigdra @2021 Implemented By Angular</nz-footer>
    </nz-layout>
  `,
  styles:[`

    [nz-menu] {
      line-height: 64px;
    }

    nz-breadcrumb {
      margin: 16px 0;
    }

    nz-content {
      padding: 0 50px;
    }

    nz-footer {
      text-align: center;
    }

    .inner-content {
      background: #fff;
      padding: 24px;
      min-height: 100vh;
    }
    .auth{
      float: right;
    }
    .logo {
      width: 32px;
      height: 32px;
      background: rgba(255, 255, 255, 0.2);
      margin: 16px 24px 16px 0;
      float: left;
    }
  `]
})
export class AppComponent {

  assetsPath: string;

  constructor() {
    this.assetsPath = '/assets/img/logo.png'
  }

  size: 'small' | 'middle' | 'large' | number = "small";
}
