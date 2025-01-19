import React from "react";
import { useTable, useSortBy, useGlobalFilter, usePagination, useRowSelect} from "react-table";
import { GlobalFilterInput } from "./GlobalFilterInput";
import "./table.css"




const IndeterminateCheckbox = React.forwardRef(
    ({ indeterminate, ...rest }, ref) => {
      const defaultRef = React.useRef();
      const resolvedRef = ref || defaultRef;

      React.useEffect(() => {
        resolvedRef.current.indeterminate = indeterminate;
      }, [resolvedRef, indeterminate]);
  
      return (
        <>
          <input type="checkbox" ref={resolvedRef} {...rest} />
        </>
      );
    }
  );

 

const sortDown = '<▼>'
const sortUp = '<▲>'
export const Table = ({ columns, data }) => {
    const {
      getTableProps,
      getTableBodyProps,
      headerGroups,
      page,
      nextPage,
      previousPage,
      prepareRow,
      preGlobalFilteredRows, 
      setGlobalFilter, 
      state,
      
    } = useTable({
      columns,
      data
    }  ,useGlobalFilter,  useSortBy,  usePagination,  useRowSelect,
    (hooks) => {
      hooks.visibleColumns.push((columns) => [
        {
          id: "selection",

          Header: ({ getToggleAllRowsSelectedProps }) => (
            <div>
              <IndeterminateCheckbox {...getToggleAllRowsSelectedProps()} />
            </div>
          ),

          Cell: ({ row }) => (
            <div classname= "checkbox" >
              <IndeterminateCheckbox {...row.getToggleRowSelectedProps()} />
            </div>
          )
        },
        ...columns
      ]);
    }
  );
  
    return (
<div classname= "Table">
        <>
          <GlobalFilterInput
            preGlobalFilteredRows={preGlobalFilteredRows}
            setGlobalFilter={setGlobalFilter} globalFilter={state.globalFilter}
          />
        
          <table {...getTableProps()}>
            <thead>
              {headerGroups.map((headerGroup) => (
                <tr {...headerGroup.getHeaderGroupProps()}>
                  {headerGroup.headers.map((column) => (
                    <th {...column.getHeaderProps(column.getSortByToggleProps())}>{column.render('Header')}
                    <span>
                         {column.isSorted ? (column.isSortedDesc ? sortDown : sortUp ) : ''}
                    </span>
                  </th>
                  ))}
                </tr>
              ))}
            </thead>
            <tbody {...getTableBodyProps()}>
              {
                page.map(row => {
                prepareRow(row)
            
                return (
                  <tr {...row.getRowProps()}>
                    {row.cells.map((cell) => {
                      return (
                        <td {...cell.getCellProps()}>{cell.render('Cell')}</td>
                      );
                    })}
                  </tr>
                );
              })}
            </tbody>
          <div>
            <button onClick={() => previousPage()}>Previous</button>
            <button onClick={() => nextPage()}>Next</button>
          </div>



          
          
            
          
          </table>
          
        </>
 </div>  
  );
}
