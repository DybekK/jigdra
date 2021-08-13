import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-user-profile',
  template: `
    <p>
      user-profile works!
    </p>
  `,
  styleUrls: ['./user-profile.component.scss']
})
export class UserProfileComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
