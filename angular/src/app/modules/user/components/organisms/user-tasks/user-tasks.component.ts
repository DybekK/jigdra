import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-user-tasks',
  template: `
    <p>
      user-tasks works!
    </p>
  `,
  styleUrls: ['./user-tasks.component.scss']
})
export class UserTasksComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
