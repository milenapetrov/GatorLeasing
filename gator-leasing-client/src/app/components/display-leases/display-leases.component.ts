import { Component, OnInit, AfterViewInit} from '@angular/core';
import { Lease } from 'src/app/models/lease';
import { LeaseService } from 'src/app/services/lease.service';
import { ColDef, GridReadyEvent, ICellRendererParams } from 'ag-grid-community';
import { format, parseISO } from 'date-fns';
import { take, Observable } from 'rxjs';
import { SortDirection } from 'src/enums/sort-direction';
import {
  Grid,
  GridOptions,
  IServerSideDatasource,
  IServerSideGetRowsRequest,
} from 'ag-grid-community';
import 'ag-grid-enterprise';
import {
  GridCellComponent,
  MyCellParams,
} from '../grid-cell/grid-cell.component';
import { HttpParams } from '@angular/common/http';

@Component({
  selector: 'app-display-leases',
  templateUrl: './display-leases.component.html',
  styleUrls: ['./display-leases.component.css']
})
export class DisplayLeasesComponent implements AfterViewInit{
  ID = 0;
  gridOptions: GridOptions<Lease> = {
  columnDefs : [ 
    {
      field: 'name',
      sortable: true,
      filter: 'agTextColumnFilter',
        filterParams: {
          filterOptions: ['contains'],
        },
      cellRenderer: GridCellComponent,
      cellRendererParams: {
        buttonText: `Edit`,
      } as MyCellParams,
    },
    {
      field: 'startDate',
      cellRenderer: (params) => {
        return format(parseISO(params.value), 'MM/dd/yyyy');
      },
    },
    {
      field: 'endDate',
      cellRenderer: (params) => {
        return format(parseISO(params.value), 'MM/dd/yyyy');
      },
    },
    {
      field: 'rent',
      cellRenderer: (params) => {
        return `$${params.value} `;
      },
    },
    { field: 'utilities' },
    { field: 'parkingCost' },
    { field: 'squareFootage' },
    {
      field: 'furnished',
      cellRenderer: (params) => {
        if (params.value) {
          return `<b> yes </b>`;
        } else {
          return `<b> no </b>`;
        }
      },
    },
    {
      field: 'parking',
      cellRenderer: (params) => {
        if (params.value) {
          return `<b> yes </b>`;
        } else {
          return `<b> no </b>`;
        }
      },
    },
    { field: 'beds' },
    { field: 'baths' },
    /*{ 
      field: 'id',
      cellRenderer: (params) => {
        this.ID = params.value;
      },
    },*/
  ],
    rowModelType: 'serverSide',
    pagination: true,
    paginationPageSize: 10,
    cacheBlockSize: 10,
  }

  getID(): number {
    return this.ID;
  }

  constructor(private leaseService: LeaseService) {}
  
  sortToken: string = '';
  sortDirection = SortDirection.descending;
  paginationToken: string = '';
  filter: string = '';

  public rowData$!: Observable<Lease[]>;

  onGridReady(params: GridReadyEvent) {
    this.rowData$ =  this.leaseService.getLeases();
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




