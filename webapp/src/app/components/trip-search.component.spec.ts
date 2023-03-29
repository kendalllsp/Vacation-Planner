import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing'; 
import { ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TripSearchComponent } from './trip-search.component';
import { AppModule } from '../app.module';

describe('NavbarComponent', () =>  {
    let component: TripSearchComponent;
    let fixture: ComponentFixture<TripSearchComponent>;
    let router: Router;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [TripSearchComponent],
            imports: [AppModule, HttpClientTestingModule, RouterTestingModule, ReactiveFormsModule]
        })
        .compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(TripSearchComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
        router = TestBed.inject(Router);
    });

    it('should create the navbar component', () => {
        const fixture = TestBed.createComponent(TripSearchComponent);
        const app = fixture.debugElement.componentInstance;
        expect(app).toBeTruthy();
    });
});