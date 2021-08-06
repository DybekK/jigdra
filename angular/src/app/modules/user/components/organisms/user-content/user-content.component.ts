import {Component, OnInit} from '@angular/core';
import {NzMarks, NzSliderValue} from "ng-zorro-antd/slider";
import {FormBuilder, FormGroup, Validators} from "@angular/forms";

@Component({
  selector: 'app-user-content',
  template: `
    <form [formGroup]="validateForm" class="ant-advanced-search-form" (ngSubmit)="submitTask()">
      <div nz-row nzJustify="center" [nzGutter]="[vGutter, hGutter]">
        <div nz-col [nzSpan]="20">
          <nz-form-item>
            <nz-form-control nzErrorTip="You need to specify task name">
              <nz-input-group nzAddOnBefore="Task">
                <input formControlName="taskName" type="text" nz-input placeholder="Task name" />
              </nz-input-group>
            </nz-form-control>
          </nz-form-item>
        </div>
      </div>

      <div nz-row nzJustify="center" [nzGutter]="[vGutter, hGutter]">
        <div nz-col [nzSpan]="20">
          <nz-form-item>
            <nz-form-control nzErrorTip="Task needs to be longer than 0 min">
              <nz-slider formControlName="time" [nzMarks]="timeStamps" [nzStep]="1" nzMax="120"
                         (nzOnAfterChange)="getTimeValue($event)"></nz-slider>
            </nz-form-control>
          </nz-form-item>
        </div>
      </div>

      <div nz-row nzJustify="center" [nzGutter]="[vGutter, hGutter]">
        <div nz-col [nzSpan]="10">
          <nz-form-item>
            <nz-form-control nzErrorTip="Number of breaks needs to be a natural number">
              <nz-input-group nzAddOnBefore="Breaks">
                <input formControlName="breaksNumber" type="number" nz-input min="0"/>
              </nz-input-group>
            </nz-form-control>
          </nz-form-item>
        </div>

        <div nz-col [nzSpan]="10">
          <nz-form-item>
            <nz-form-control nzErrorTip="You need to specify task name">
              <nz-input-group nzAddOnBefore="Lorem">
                <input type="text" nz-input placeholder="Lorem ipsum"/>
              </nz-input-group>
            </nz-form-control>
          </nz-form-item>
        </div>
      </div>

      <div nz-row nzJustify="center" [nzGutter]="[vGutter, hGutter]">
        <div nz-col [nzSpan]="20">
          <nz-form-item>
              <textarea formControlName="comment" nz-input placeholder="Write your comment here!"
                        [nzAutosize]="{ minRows: 2, maxRows: 5 }"></textarea>
          </nz-form-item>
        </div>
      </div>

      <div nz-row nzJustify="center" [nzGutter]="[vGutter, hGutter]">
        <div nz-col [nzSpan]="3">
          <button nz-button [nzType]="'primary'" >Crate task</button>
        </div>
        <div nz-col [nzSpan]="3">
          <button nz-button [nzType]="'default'" (click)="resetTaskForm($event)">Reset</button>
        </div>
      </div>

    </form>
  `,
  styleUrls: ['./user-content.component.scss']
})
export class UserContentComponent implements OnInit {

  validateForm!: FormGroup;

  hGutter = {xs: 8, sm: 16, md: 24};
  vGutter = {xs: 8, sm: 16, md: 24};

  timeStamps: NzMarks = {
    0: '0 min',
    10: '10 min',
    30: '30 min',
    60: '60 min'
  }

  constructor(private fb: FormBuilder) {
  }

  /*
   * Get time value from slider as **number**
   */
  getTimeValue(value: NzSliderValue): void {
    console.log(value);
  }

  ngOnInit(): void {
    this.validateForm = this.fb.group({
      taskName: [null, [Validators.required]],
      time: [null, [Validators.required, Validators.min(1)]],
      breaksNumber: [0, [Validators.required, Validators.min(0)]],
      comment: [null]
    });
  }

  submitTask(): void {
    Object.values(this.validateForm.controls).forEach(control => {
      control.markAsDirty();
      control.updateValueAndValidity();
    })

    if (this.validateForm.valid) {
      console.log('everything is fine')
    }
  }

  resetTaskForm(e: MouseEvent) {
    e.preventDefault();
    this.validateForm.reset();
    Object.values(this.validateForm.controls).forEach( control => {
        control.markAsPristine();
        control.updateValueAndValidity();

    })
    this.validateForm.controls["breaksNumber"].setValue(0);
  }
}
