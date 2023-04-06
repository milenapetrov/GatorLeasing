import { Component } from '@angular/core';
import { ICellRendererAngularComp } from 'ag-grid-angular';
import { ICellRendererParams } from 'ag-grid-community';
import { Router } from '@angular/router';

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

  constructor(private router: Router) { }
  agInit(params: ICellRendererParams & MyCellParams): void {
    this.value = params.value
    this.buttonText = params.buttonText ?? 'Default';
   }

   refresh(params: ICellRendererParams & MyCellParams): boolean {
    return true;
  }

  onClick(event: any){
    this.router.navigateByUrl('/update');
  }

  ngOnInit(): void {
  }
}
