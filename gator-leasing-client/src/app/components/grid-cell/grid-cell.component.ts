import { Component } from '@angular/core';
import { ICellRendererAngularComp } from 'ag-grid-angular';
import { ICellRendererParams } from 'ag-grid-community';

export interface MyCellParams {
  buttonText?: string;
}

@Component({
  selector: 'app-grid-cell',
  templateUrl: './grid-cell.component.html',
  /*template: `
    <button routerLink="/update"> {{buttonText}}</button>
  `,*/
  styleUrls: ['./grid-cell.component.css'],
})
export class GridCellComponent {
  value: any;
  buttonText: string = 'Default';

  agInit(params: ICellRendererParams & MyCellParams): void {
    this.value = params.value;
    this.buttonText = params.buttonText ?? 'Default';
  }

  refresh(params: ICellRendererParams & MyCellParams): boolean {
    return false;
  }

  onClick(event: any) {
    //this.params.clicked(this.params.value);
    alert('does not work yet lolll');
  }

  ngOnInit(): void {}
}
