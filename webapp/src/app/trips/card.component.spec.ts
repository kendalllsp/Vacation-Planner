import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing'; 
import { ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TripsCardComponent } from './card.component';
import { AppModule } from '../app.module';

describe('TripsCardComponent', () =>  {
    let component: TripsCardComponent;
    let fixture: ComponentFixture<TripsCardComponent>;
    let router: Router;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [TripsCardComponent],
            imports: [AppModule, HttpClientTestingModule, RouterTestingModule, ReactiveFormsModule]
        })
        .compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(TripsCardComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
        router = TestBed.inject(Router);
    });

    it('should create the TripsCardComponent component', () => {
        const fixture = TestBed.createComponent(TripsCardComponent);
        const app = fixture.debugElement.componentInstance;
        expect(app).toBeTruthy();
    });
});