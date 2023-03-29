import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing'; 
import { ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { SignupComponent } from './signup.component';
import { AppModule } from '../app.module';

describe('SignupComponent', () =>  {
    let component: SignupComponent;
    let fixture: ComponentFixture<SignupComponent>;
    let router: Router;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            declarations: [SignupComponent],
            imports: [AppModule, HttpClientTestingModule, RouterTestingModule, ReactiveFormsModule]
        })
        .compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(SignupComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
        router = TestBed.inject(Router);
    });

    it('should create the SignupComponent', () => {
        const fixture = TestBed.createComponent(SignupComponent);
        const app = fixture.debugElement.componentInstance;
        expect(app).toBeTruthy();
    });
});