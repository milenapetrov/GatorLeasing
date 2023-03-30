import { Component, OnInit } from '@angular/core';
import { Lease } from 'src/app/models/lease';
import { LeaseService } from 'src/app/services/lease.service';
import { ColDef, GridReadyEvent, ICellRendererParams } from 'ag-grid-community';
import { Observable } from 'rxjs';
import { GridCellComponent, MyCellParams } from '../grid-cell/grid-cell.component';

@Component({
  selector: 'app-display-leases',
  templateUrl: './display-leases.component.html',
  styleUrls: ['./display-leases.component.css']
})
export class DisplayLeasesComponent implements OnInit {
  leases: Lease[] = [];

  constructor(private leaseService: LeaseService) {}

  ngOnInit(): void {
    this.leaseService.getLeases().subscribe(leases => this.leases = leases)
  }

  columnDefs = [
    {field: "name",
            cellRenderer: GridCellComponent,
            cellRendererParams: {
              buttonText: 'Update'
            } as MyCellParams},
    {field: "startDate",
            cellRenderer: (params: ICellRendererParams) => {
            return params.value.substring(0,10)
            }},
    {field: "endDate",
            cellRenderer: (params: ICellRendererParams) => {
            return params.value.substring(0,10)
            }},
    {field: "rent",
            cellRenderer: (params: ICellRendererParams) => {
            return `$${params.value} `
            }},
    {field: "utilities"},
    {field: "parkingCost"},
    {field: "squareFootage"},
    {field: "furnished",
            cellRenderer: (params: ICellRendererParams) => {
            if(params.value) {
              return `<b> yes </b>`
            }else{
              return `<b> no </b>`
            }
      }},
    {field: "parking",
            cellRenderer: (params: ICellRendererParams) => {
              if(params.value) {
                return `<b> yes </b>`
              }else{
                return `<b> no </b>`
              }
            }},
    {field: 'beds'},
    {field: "baths"}
  ];

  public defaultColDef: ColDef = {
    sortable: true,
    filter: true,
  };

  public rowData$!: Observable<Lease[]>;

  onGridReady(params: GridReadyEvent) {
    this.rowData$ =  this.leaseService.getLeases();
  }
}



