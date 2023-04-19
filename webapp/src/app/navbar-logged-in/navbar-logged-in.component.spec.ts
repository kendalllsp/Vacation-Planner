import { TestBed, async } from '@angular/core/testing';
import { NavbarLoggedInComponent } from './navbar-logged-in.component'
import { AppModule } from '../app.module';

describe('NavbarLoggedInComponent', () => {

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        NavbarLoggedInComponent
      ],
      imports: [
        AppModule
      ]
    }).compileComponents();
  }));

  it('should create the NavbarLoggedInComponent', async(() => {
    const fixture = TestBed.createComponent(NavbarLoggedInComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  }));
});