import { Component, OnInit } from '@angular/core';
import { Lease } from 'src/app/models/lease';
import { LeaseService } from 'src/app/services/lease.service';

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
}



