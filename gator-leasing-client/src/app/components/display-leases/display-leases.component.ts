import { Component, OnInit } from '@angular/core';
import { Lease } from 'src/app/models/lease';
import { LeaseService } from 'src/app/services/lease.service';
import { ColDef, GridReadyEvent } from 'ag-grid-community';
import { Observable } from 'rxjs';

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
    {field: "name" }, 
    {field: "start_date"}, 
    {field: "end_date"},
    {field: "rent"},
    {field: "utilities"},
    {field: "parking_cost"},
    {field: "square_footage"},
    {field: "furnished"},
    {field: "parking"},
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



