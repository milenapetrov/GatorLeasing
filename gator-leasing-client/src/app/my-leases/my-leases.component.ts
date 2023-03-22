import { Component } from '@angular/core';
import { ColDef, GridReadyEvent } from 'ag-grid-community';
import { Observable } from 'rxjs';
import { LeaseService } from '../services/lease.service';
 import { Lease } from '../models/lease';


@Component({
  selector: 'app-my-leases',
  templateUrl: './my-leases.component.html',
  styleUrls: ['./my-leases.component.css']
})
export class MyLeasesComponent {
  leases: Lease[] = [];

  constructor(private leaseService: LeaseService) {}

  /*columnDefs = [{field: "id"}, { field: "name" }];

  public defaultColDef: ColDef = {
    sortable: true,
    filter: true,
  };

  public rowData$!: Observable<Lease[]>;

  onGridReady(params: GridReadyEvent) {
    this.rowData$ =  this.leaseService.getLeases();
  }
  */
}
