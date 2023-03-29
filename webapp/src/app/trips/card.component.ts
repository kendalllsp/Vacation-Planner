import {Component, Input} from '@angular/core';


@Component({
  selector: 'trips-card-component',
  templateUrl: 'card.component.html',
  styleUrls: ['card.component.scss'],
})
export class TripsCardComponent {

  @Input() location: string = '';
  @Input() date: string = '';

}