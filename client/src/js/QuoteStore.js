import { observable } from "mobx";
import axios from 'axios';

class QuoteStore {
    @observable latestBuyPrice = "";
    @observable latestSellPrice = "";

    pollingQuote(){
        axios.get("http://localhost:8091/quote/latest/BTC/USD").then((results) => {
            var arr = Object.entries(results.data)
            for (let [key, value] of arr) {
                if(value.side == "BUY"){
                    this.latestBuyPrice = value.price
                }else{
                    this.latestSellPrice = value.price
                }
            }
        })
    }
}

var store = new QuoteStore();

setInterval(() => store.pollingQuote(), 1000)

export default store;