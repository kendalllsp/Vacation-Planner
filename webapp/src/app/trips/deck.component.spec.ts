import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing'; 
import { ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TripsCardDeckComponent } from './deck.component';
import { AppModule } from '../app.module';

describe('TripsCardDeckComponent', () =>  {
    let component: TripsCardDeckComponent;
    let fixture: ComponentFixture<TripsCardDeckComponent>;
    let router: Router;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [TripsCardDeckComponent],
            imports: [AppModule, HttpClientTestingModule, RouterTestingModule, ReactiveFormsModule]
        })
        .compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(TripsCardDeckComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
        router = TestBed.inject(Router);
    });

    it('should create the TripsCardDeckComponent component', () => {
        const fixture = TestBed.createComponent(TripsCardDeckComponent);
        const app = fixture.debugElement.componentInstance;
        expect(app).toBeTruthy();
    });
});