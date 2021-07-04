import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-auth-content',
  template: `
    <nz-content class="auth-content">
      <nz-card class="auth-content__container__card" nzTitle="Sign in">
        <router-outlet></router-outlet>
      </nz-card>
    </nz-content>
  `,
  styleUrls: ["auth-content.component.scss"]
})
export class AuthContentComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
