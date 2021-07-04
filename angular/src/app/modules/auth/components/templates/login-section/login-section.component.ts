import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login-section',
  template: `
    <div class="login-section-container">
      <app-login-form></app-login-form>
    </div>
  `,
  styleUrls: ["login-section.component.scss"]
})
export class LoginSectionComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
