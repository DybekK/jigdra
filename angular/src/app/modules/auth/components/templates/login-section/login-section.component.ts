import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login-section',
  template: `
    <div class="login-section-container">
      <app-login-form></app-login-form>
    </div>
  `,
  styles: [
    `
      .login-section-container {
        height: 100%;
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
      }
    `
  ]
})
export class LoginSectionComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
