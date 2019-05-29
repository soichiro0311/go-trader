import { observable } from "mobx";
import axios from 'axios';

class TodoStore {
    @observable amount = "";

    pollingQuote(){
        axios.get("http://localhost:8091/quote/latest/BTC/USD").then((results) => {
            this.amount = results.data.price.value
        })
    }
}

var store = new TodoStore();

setInterval(() => store.pollingQuote(), 1000)

export default store;