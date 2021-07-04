import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-auth-header',
  template: `
    <nz-header class="auth-header">
      <img class="logo" [routerLink]="'/'">
      <nz-space [nzSize]="size" class="auth">
        <button *nzSpaceItem nz-button nzType="primary" [routerLink]="'/login'">Login</button>
        <button *nzSpaceItem nz-button nzType="default" [routerLink]="'/register'">Register</button>
      </nz-space>
    </nz-header>
  `,
  styleUrls: ['./auth-header.component.scss']
})
export class AuthHeaderComponent implements OnInit {
  size: 'small' | 'middle' | 'large' | number = "small";
  constructor() { }

  ngOnInit(): void {
  }

}
