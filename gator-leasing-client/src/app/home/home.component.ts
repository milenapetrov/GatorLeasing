import { Component, OnInit } from '@angular/core';

import { LeaseService } from '../services/lease.service';
import { AuthService } from '@auth0/auth0-angular';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  constructor(private leaseService:LeaseService, public auth: AuthService) {}

  ngOnInit(): void {}
}
