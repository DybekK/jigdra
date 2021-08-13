import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-calendar',
  template: `
    <p>
      calendar works!
    </p>
  `,
  styleUrls: ['./calendar.component.scss']
})
export class CalendarComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
