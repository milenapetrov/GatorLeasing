import { AfterViewInit, Component } from '@angular/core';
import {
  GridOptions,
  IServerSideDatasource,
  IServerSideGetRowsRequest,
} from 'ag-grid-enterprise';
import { Lease } from 'src/app/models/lease';
import { LeaseService } from 'src/app/services/lease.service';
import { format, parseISO } from 'date-fns';
import { take } from 'rxjs';
import { SortDirection } from 'src/enums/sort-direction';
import { Router } from '@angular/router';
import { CustomColumnDef } from 'src/app/shared/custom-column-def';
import { MatSelectChange } from '@angular/material/select';


@Component({
  selector: 'app-lease-listings',
  templateUrl: './lease-listings.component.html',
  styleUrls: ['./lease-listings.component.css'],
})
export class LeaseListingsComponent implements AfterViewInit {
  columnDefs : CustomColumnDef[] = [
    {
      field: 'name',
      colId: 'name',
      dataType: 'text',
      sortable: true,
      filterable: true,
      filter: 'agTextColumnFilter',
      filterParams: {
        filterOptions: ['contains'],
      },
    },
    {
      field: 'startDate',
      colId: 'startDate',
      dataType: 'date',
      sortable: true,
      filterable: true,
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
      colId: 'endDate',
      dataType: 'date',
      sortable: true,
      filterable: true,
      cellRenderer: (data) => {
        return format(parseISO(data.value), 'MM/dd/yyyy');
      },
    },
    {
      field: 'address.zipCode',
      colId: 'zipCode',
      dataType: 'text',
      headerName: 'Zip Code',
      sortable: true,
      filterable: true,
      filter: 'agTextColumnFilter',
      filterParams: {
        filterOptions: ['contains'],
      },
    }
  ]
  gridOptions: GridOptions<Lease> = {
    columnDefs: this.columnDefs,
    rowModelType: 'serverSide',
    pagination: true,
    paginationPageSize: 10,
    cacheBlockSize: 10,
    suppressPropertyNamesCheck: true
  };

  pageSize = 10
  cacheBlockSize = this.pageSize * 5
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

        const newFilter = this.getFilters(params.request);
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
            this.cacheBlockSize,
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

  getFilters(request: IServerSideGetRowsRequest): string {
    let filterArr: string[] = [];
    if (request.filterModel) {
      this.gridOptions.columnDefs?.forEach(colDef => {
        const customColDef = colDef as CustomColumnDef
        if (customColDef.filterable) {
          const filter = this.getFilter(request.filterModel, customColDef.colId!, customColDef.colId!.replace(/[A-Z]/g, letter => `_${letter.toLowerCase()}`), customColDef.dataType)
          if (filter != '') {
            filterArr.push(filter)
          }
        }
      })
    }
    let filters = filterArr.join(",")
    return filters;
  }

  getFilter(filterModel, fieldName : string, columnName : string, columnType : string) : string {
    let filter = '';
    if (filterModel[fieldName] != undefined) {
      if (filterModel[fieldName].condition1) {
        filter += this.getClause(
          columnName,
          columnType,
          filterModel[fieldName].condition1.type,
          columnType === "date" ? filterModel[fieldName].condition1.dateFrom : filterModel[fieldName].condition1.filter
        );
        filter += ' ' + filterModel[fieldName].operator + ' ';
        filter += this.getClause(
          columnName,
          columnType,
          filterModel[fieldName].condition2.type,
          columnType === "date" ? filterModel[fieldName].condition2.dateFrom : filterModel[fieldName].condition2.filter
        )
      } else {
        filter += this.getClause(
          columnName,
          columnType,
          filterModel[fieldName].type,
          columnType === "date" ? filterModel[fieldName].dateFrom : filterModel[fieldName].filter
        )
      }
    }
    return filter
  }

  getClause(column: string, columnType : string, comparisonType: string, param: string) {
    switch (comparisonType) {
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

  pageSizeChange(evt : MatSelectChange) {
    this.pageSize = evt.value
    this.cacheBlockSize = 5 * this.pageSize
    this.gridOptions.paginationPageSize = this.pageSize
    this.gridOptions.cacheBlockSize = this.cacheBlockSize
    this.gridOptions.api?.setServerSideDatasource(this.getLeaseDatasource())
  }
}
