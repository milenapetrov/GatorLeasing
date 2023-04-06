import { Component, OnInit } from '@angular/core';

import { LeaseService } from '../services/lease.service';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { Observable } from 'rxjs';
import { map, shareReplay, take } from 'rxjs/operators';
import { AuthService } from '@auth0/auth0-angular';
import {  Router } from '@angular/router';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  search: String = "";

  isHandset$: Observable<boolean> = this.breakpointObserver.observe(Breakpoints.Handset)
    .pipe(
      map(result => result.matches),
      shareReplay()
    );

  constructor(private breakpointObserver: BreakpointObserver, public auth: AuthService, private router: Router) {}

  loginOrRoute(requested: string) {
    this.auth.isAuthenticated$.pipe(take(1)).subscribe(isLoggedIn => {
      if (isLoggedIn) {
        this.router.navigate(['/create'])
      }
      else {
        this.auth.loginWithRedirect({
          appState: {
            target: requested
          }
        })
      }
    })
  }

  ngOnInit(): void {}
}
