import React from "react"
import TakeAdvertisements from "./constants/advertisements";
import TakeColumns from "./constants/columns";
import { Table} from "./components/table";
import "./App.css"


    


export default function App() {

  const columns =  TakeColumns(); 
  const data = TakeAdvertisements();

    
  return (
    <div className="App">
      <Table columns={columns} data={data} />
      <div >
       <h3> Add advertisement</h3>
        <form classname = "Createform" action="http://localhost:8080/advertisements" method="post">
          <label>Name</label>
          <input type="text" name="name" placeholder="Введите название объявления" />
          <label>Comment</label>
          <input type="text" name="comment" placeholder="Текст" />
          <label>Photo</label>
          <input type="text" name="photo" placeholder="ссылка на главное фото" />
          <label>PhotoSecond</label>
          <input type="text" name="photosecond" placeholder="ссылка на дополнительное фото" />
          <label>PhotoSecond</label>
          <input type="text" name="photosecond2" placeholder="ссылка на дополнительное фото" />
          <label>Price</label>
          <input type="number" name="price" placeholder="цена" ></input>
          <input type="submit" value="Отправить" />
        </form>
      </div>


    </div>


  );
}