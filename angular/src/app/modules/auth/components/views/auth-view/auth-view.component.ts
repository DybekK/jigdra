import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-auth-view',
  template: `
      <nz-layout class="auth-view">
        <nz-layout class="auth-view__inner-content">
          <app-auth-header></app-auth-header>
          <app-auth-content></app-auth-content>
          <app-footer></app-footer>
        </nz-layout>
        <nz-sider [nzWidth]="600"></nz-sider>
      </nz-layout>
  `,
  styleUrls: ['./auth-view.component.scss']
})
export class AuthViewComponent implements OnInit {


  constructor() { }

  ngOnInit(): void {
  }

}
