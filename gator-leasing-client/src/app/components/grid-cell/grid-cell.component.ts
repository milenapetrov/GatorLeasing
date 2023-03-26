import { Component } from '@angular/core';
import { ICellRendererAngularComp } from 'ag-grid-angular';
import { ICellRendererParams } from 'ag-grid-community';

export interface MyCellParams {
  buttonText?: string;
}

@Component({
  selector: 'app-grid-cell',
  templateUrl: './grid-cell.component.html',
  styleUrls: ['./grid-cell.component.css']
})
export class GridCellComponent {
  value: any;
  buttonText: string = 'Default';

  agInit(params: ICellRendererParams & MyCellParams): void {
    this.value = params.value
    this.buttonText = params.buttonText ?? 'Default';
   }

   refresh(params: ICellRendererParams & MyCellParams): boolean {
    return false;
  }

  onClick(event: any){
    alert('Cell value is ' + this.value);
  }

  ngOnInit(): void {
  }
}
