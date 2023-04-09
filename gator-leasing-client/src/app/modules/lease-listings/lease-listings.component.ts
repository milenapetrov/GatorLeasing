import { AfterViewInit, Component } from '@angular/core';
import { Grid, GridOptions, IServerSideDatasource } from 'ag-grid-community';
import 'ag-grid-enterprise';
import { Lease } from 'src/app/models/lease';
import { LeaseService } from 'src/app/services/lease.service';
import { format, parseISO } from 'date-fns'

@Component({
  selector: 'app-lease-listings',
  templateUrl: './lease-listings.component.html',
  styleUrls: ['./lease-listings.component.css']
})
export class LeaseListingsComponent implements AfterViewInit {
  gridOptions : GridOptions<Lease> = {
    columnDefs: [
      { field: 'name', sortable: true },
      { field: 'createdAt', sortable: false, cellRenderer: (data) => {
        return format(parseISO(data.value), 'MM/dd/yyyy')
      } },
      { field: 'startDate', sortable: true, cellRenderer: (data) => {
        return format(parseISO(data.value), 'MM/dd/yyyy')
      }},
      { field: 'endDate', sortable: true, cellRenderer: (data) => {
        return format(parseISO(data.value), 'MM/dd/yyyy')
      }},
    ],
    rowModelType: 'serverSide', 
    pagination: true,
    paginationPageSize: 10,
    cacheBlockSize: 10,
  }

  sortToken : string = ""
  sortDirection : number = 1
  paginationToken : string = ""

  constructor(private leaseService : LeaseService) {}

  ngAfterViewInit() {
    var datasource = this.getLeaseDatasource()
    this.gridOptions.api?.setServerSideDatasource(datasource)
  }

  getLeaseDatasource(): IServerSideDatasource {
    return {
      getRows: (params) => {
        console.log(params)
        if (params.request.sortModel.length > 0) {
          const sortModel = params.request.sortModel[params.request.sortModel.length - 1]
          if (sortModel.colId != this.sortToken) {
            this.paginationToken = ""
            this.sortToken = sortModel.colId
            this.gridOptions.api?.setServerSideDatasource(this.getLeaseDatasource())
          }
          const newSortDirection = sortModel.sort == "asc" ? 0 : 1
          if (newSortDirection != this.sortDirection) {
            this.paginationToken = ""
            this.sortDirection = newSortDirection
            this.gridOptions.api?.setServerSideDatasource(this.getLeaseDatasource())
          }
        } else {
          if (this.sortToken != "") {
            this.sortToken = ""
            this.paginationToken = ""
            this.sortDirection = 1
            this.gridOptions.api?.setServerSideDatasource(this.getLeaseDatasource())
          }
        }
        this.leaseService.getPagedLeases(10, this.sortToken,this.paginationToken , this.sortDirection, "").subscribe(paginatedLeasesResult => {
          params.success({
            rowData: paginatedLeasesResult.leases,
            rowCount: paginatedLeasesResult.count
          })
          this.paginationToken = paginatedLeasesResult.paginationToken
        })
      }
    }
  }
}

