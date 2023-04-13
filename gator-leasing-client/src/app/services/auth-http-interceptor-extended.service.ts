import { HttpEvent, HttpHandler, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthHttpInterceptor } from '@auth0/auth0-angular';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class AuthHttpInterceptorExtendedService extends AuthHttpInterceptor {
  override intercept(
    req: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    if (req.url.endsWith('/leases') && req.method === 'GET') {
      return next.handle(req);
    }
    if (req.url.match('[^ ]+\/leases\/[0-9]+)') && req.method == 'GET') {
      return next.handle(req)
    }
    if (req.url.endsWith('/leases/paged') && req.method == 'POST') {
      return next.handle(req);
    }
      
    return super.intercept(req, next);
  }
}
