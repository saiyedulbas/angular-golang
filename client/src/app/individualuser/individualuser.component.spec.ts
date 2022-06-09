import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IndividualuserComponent } from './individualuser.component';

describe('IndividualuserComponent', () => {
  let component: IndividualuserComponent;
  let fixture: ComponentFixture<IndividualuserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IndividualuserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(IndividualuserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
