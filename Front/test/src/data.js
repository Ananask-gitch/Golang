import {useState, useEffect} from "react";
import axios from "axios"

export default function TakeAdvertisements () {
    const [data, setData] = useState([]);
    
    useEffect(() => {
    axios
      .get("/advertisements")
      .then((response) => {
        setData(response.data);
      })
      .catch((error) => {
        console.log(error);
      });
    }, []);
    
    return data
    }

export function AddAdvertisement (data) {
  console.log(data)      
}