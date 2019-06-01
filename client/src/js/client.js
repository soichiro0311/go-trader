import "../css/index.css";
import React from "react";
import ReactDOM from "react-dom";

import MarketOrderPanel from "./MarketOrderPanel";
import store from "./QuoteStore";

const app = document.getElementById("app");

ReactDOM.render(<MarketOrderPanel store={store} />, app);