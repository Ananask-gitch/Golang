import React, {useMemo} from 'react';
import 'devextreme/dist/css/dx.light.css';
import {
    DataGrid,
    Column,
    Paging,
    Pager,
    Editing,
    Popup
} from 'devextreme-react/data-grid';
import DataSource from 'devextreme/data/data_source';
import { createStore } from 'devextreme-aspnet-data-nojquery';


const url = ""
 
function App() {
    const  dataSource = useMemo (() => {
    return new DataSource({
        store: createStore({
            key: 'ID',
            loadUrl: `${url}/advertisements`,
            insertUrl: `${url}/advertisements`,
            deleteUrl: `${url}/advertisements`,
            onBeforeSend: (operation, Settings) => {
            if (operation ==='load') return;
            if (operation === 'delete') 
            Settings.url = `${Settings.url}/${Settings.data.key}`;
            Settings.data = Settings.data.values;
            Settings.headers = {
            'Content-Type': 'application/json'
            };
        },
    }),
}); }, [url])
console.log(dataSource);

    return (
        <div className="App">
            <DataGrid
                dataSource={dataSource}
                allowColumnReordering={true}
                onRowUpdating = {(options) => options.newData = {...options.oldData, ...options.newData}}>
                <Editing
                    mode="popup"
                    allowUpdating={false}
                    allowAdding={true}
                    allowDeleting={true}
                ></Editing>
                <Paging defaultPageSize={10} />
                <Pager
                visible={true}/>
                <Column dataField="Name" caption={"Имя"}></Column>
                <Column dataField="Comment" caption={"Комментарий"}></Column>
                <Column dataField="PhotoMain" caption={"Фотография"}></Column>
                <Column dataField="Price" caption={"Цена"}></Column>
            </DataGrid>
        </div>
    );
}
 
export default App;
