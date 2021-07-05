import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-user-content',
  template: `
    <h2>content works</h2>
  `,
  styleUrls: ['./user-content.component.scss']
})
export class UserContentComponent implements OnInit {

  constructor() {
  }

  ngOnInit(): void {
  }

}
