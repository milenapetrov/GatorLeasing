import { AfterViewInit, Component, EventEmitter, Output } from '@angular/core';
import {
  CellClickedEvent,
  GridOptions,
  IServerSideDatasource,
  IServerSideGetRowsRequest,
  GridApi,
  RowClickedEvent,
  RowSelectedEvent
} from 'ag-grid-community';
import 'ag-grid-enterprise';
import { Lease } from 'src/app/models/lease';
import { LeaseService } from 'src/app/services/lease.service';
import { format, parseISO } from 'date-fns';
import { take } from 'rxjs';
import { SortDirection } from 'src/enums/sort-direction';
import { GridCellComponent,MyCellParams,} from 'src/app/components/grid-cell/grid-cell.component';
import { data } from 'cypress/types/jquery';
import { Router } from '@angular/router';
import { values } from 'cypress/types/lodash';
import { HttpParams } from '@angular/common/http';

@Component({
  selector: 'app-lease-listings',
  templateUrl: './lease-listings.component.html',
  styleUrls: ['./lease-listings.component.css'],
})
export class LeaseListingsComponent implements AfterViewInit {
  gridOptions: GridOptions<Lease> = {
    onRowSelected: (event: RowSelectedEvent) => console.log,
    columnDefs: [
      {
        field: 'name',
        sortable: true,
        filter: 'agTextColumnFilter',
        filterParams: {
          filterOptions: ['contains'],
        },
        /*cellRenderer: GridCellComponent,
        cellRendererParams: {
          buttonText: `view`,
        } as MyCellParams,*/
      },
      {
        field: 'startDate',
        sortable: true,
        cellRenderer: (data) => {
          return format(parseISO(data.value), 'MM/dd/yyyy');
        },
        filter: 'agDateColumnFilter',
        filterParams: {
          filterOptions: ['greaterThanOrEqual'],
        },
      },
      {
        field: 'endDate',
        sortable: true,
        cellRenderer: (data) => {
          return format(parseISO(data.value), 'MM/dd/yyyy');
        },
      },
      
    ],
    rowModelType: 'serverSide',
    pagination: true,
    paginationPageSize: 10,
    cacheBlockSize: 10,
    onRowClicked: (event: RowClickedEvent) => console.log("hey")
  };

  sortToken: string = '';
  sortDirection = SortDirection.descending;
  paginationToken: string = '';
  filter: string = '';

  constructor(private leaseService: LeaseService, private router: Router) {}

  gridApi: any;

  onRowClicked(event: any){
    const val = event.data
    this.router.navigate((['/view']), {queryParams: {fieldParam: val.id}})
  }

  onGridReady(params: any){
    this.gridApi = params.api
  }
  ngAfterViewInit() {
    var datasource = this.getLeaseDatasource();
    this.gridOptions.api?.setServerSideDatasource(datasource);
  }

  getLeaseDatasource(): IServerSideDatasource {
    return {
      getRows: (params) => {
        let needsReset = false;

        const newFilter = this.getFilter(params.request);
        if (this.filter != newFilter) {
          this.paginationToken = '';
          this.filter = newFilter;
          needsReset = true;
        }

        if (params.request.sortModel.length > 0) {
          const sortModel =
            params.request.sortModel[params.request.sortModel.length - 1];
          if (sortModel.colId != this.sortToken) {
            this.paginationToken = '';
            this.sortToken = sortModel.colId;
            needsReset = true;
          }
          const newSortDirection = sortModel.sort == "asc" ? SortDirection.ascending : SortDirection.descending
          if (newSortDirection != this.sortDirection) {
            this.paginationToken = '';
            this.sortDirection = newSortDirection;
            needsReset = true;
          }
        } else {
          if (this.sortToken != '') {
            this.sortToken = '';
            this.paginationToken = '';
            this.sortDirection = SortDirection.descending;
            needsReset = true;
          }
        }

        if (needsReset) {
          this.gridOptions.api?.setServerSideDatasource(
            this.getLeaseDatasource()
          );
        }

        this.leaseService
          .getPagedLeases(
            10,
            this.sortToken,
            this.paginationToken,
            this.sortDirection,
            this.filter
          )
          .pipe(take(1))
          .subscribe((paginatedLeasesResult) => {
            params.success({
              rowData: paginatedLeasesResult.leases,
              rowCount: paginatedLeasesResult.count,
            });
            this.paginationToken = paginatedLeasesResult.paginationToken;
          });
      },
    };
  }

  getFilter(request: IServerSideGetRowsRequest): string {
    let filterArr: string[] = [];
    if (request.filterModel?.name) {
      let nameFilter = '';
      if (request.filterModel.name.condition1) {
        nameFilter += this.getClause(
          'name',
          'text',
          request.filterModel.name.condition1.type,
          request.filterModel.name.condition1.filter
        );
        nameFilter += ' ' + request.filterModel.name.operator + ' ';
        nameFilter += this.getClause(
          'name',
          'text',
          request.filterModel.name.condition2.type,
          request.filterModel.name.condition2.filter
        );
      } else {
        nameFilter += this.getClause(
          'name',
          'text',
          request.filterModel.name.type,
          request.filterModel.name.filter
        );
      }
      filterArr.push(nameFilter);
    }

    if (request.filterModel?.startDate) {
      let startDateFilter = '';
      if (request.filterModel.startDate.condition1) {
        startDateFilter += this.getClause(
          'start_date',
          request.filterModel.startDate.condtion1.type,
          'text',
          request.filterModel.startDate.condition1.dateFrom
        );
        startDateFilter += ' ' + request.filterModel.startDate.operator + ' ';
        startDateFilter += this.getClause(
          'start_date',
          'text',
          request.filterModel.startDate.condition2.type,
          request.filterModel.startDate.condition2.dateFrom
        )
      } else {
        startDateFilter += this.getClause(
          'start_date',
          'text',
          request.filterModel.startDate.type,
          request.filterModel.startDate.dateFrom
        )
      }
      filterArr.push(startDateFilter)
    }

    let filters = filterArr.join(",")
    return filters;
  }

  getClause(column: string, columnType : string, type: string, param: string) {
    switch (type) {
      case 'contains':
        return column + ' LIKE ' + "'%" + param + "%'";
      case 'greaterThanOrEqual':
        let clause = column + " >= "
        if (columnType == "number") {
          clause += param
        } else {
          clause += "'" + param + "'"
        }
        return clause
      default:
        return '';
    }
  }
}
