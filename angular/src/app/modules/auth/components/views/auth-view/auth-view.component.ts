import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-auth-view',
  template: `
      <nz-layout nz-row class="auth-view">
        <nz-layout nz-col [nzSm]="24" [nzMd]="14" class="auth-view__inner-content">
          <app-auth-header></app-auth-header>
          <app-auth-content></app-auth-content>
          <app-footer></app-footer>
        </nz-layout>
        <nz-layout nz-col [nzSm]="24" [nzMd]="10" class="auth-view__sider">
          <h2 nz-typography>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </h2>
        </nz-layout>
      </nz-layout>
  `,
  styleUrls: ['./auth-view.component.scss']
})
export class AuthViewComponent implements OnInit {


  constructor() { }

  ngOnInit(): void {
  }

}
