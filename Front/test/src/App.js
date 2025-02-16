import React, {useMemo} from 'react';
import 'devextreme/dist/css/dx.light.css';
import {
    DataGrid,
    Column,
    Paging,
    Pager,
    Editing,
    Popup,
    MasterDetail
} from 'devextreme-react/data-grid';
import DataSource from 'devextreme/data/data_source';
import { createStore } from 'devextreme-aspnet-data-nojquery';


const url = ""


  const renderDetail = (props) => {
    const { PhotoSecond, PhotoSecond2 } = props.data.Photos;
    return (
      <div className="employee-info">
        <p>Доп фото: {PhotoSecond}</p>
        <p>Доп фото: {PhotoSecond2}</p>
      </div>
    );
  };
 
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
                <Column dataField="Photos.PhotoMain" caption={"Фотография"}></Column>
                <Column dataField="Price" caption={"Цена"}></Column>
                <Column dataField="Photos.PhotoSecond" caption={"допфото"} visible={false}></Column>
                <Column dataField="Photos.PhotoSecond2" caption={"допфото"} visible={false}></Column>
                
                <MasterDetail
                    enabled={true}
                    render={renderDetail}
                />
            </DataGrid>
        </div>
    );
}
 
export default App;
