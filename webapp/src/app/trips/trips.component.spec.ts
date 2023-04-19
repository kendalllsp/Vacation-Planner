import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing'; 
import { ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TripsComponent } from './trips.component';
import { AppModule } from '../app.module';

describe('TripsComponent', () =>  {
    let component: TripsComponent;
    let fixture: ComponentFixture<TripsComponent>;
    let router: Router;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [TripsComponent],
            imports: [AppModule, HttpClientTestingModule, RouterTestingModule, ReactiveFormsModule]
        })
        .compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(TripsComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
        router = TestBed.inject(Router);
    });

    it('should create the TripsComponent', () => {
        const fixture = TestBed.createComponent(TripsComponent);
        const app = fixture.debugElement.componentInstance;
        expect(app).toBeTruthy();
    });
});