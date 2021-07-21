import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MailConfirmationComponent } from './mail-confirmation.component';

describe('MailConfirmationComponent', () => {
  let component: MailConfirmationComponent;
  let fixture: ComponentFixture<MailConfirmationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MailConfirmationComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MailConfirmationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
