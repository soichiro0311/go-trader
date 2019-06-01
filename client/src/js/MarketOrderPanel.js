import React from "react";

import { observer } from "mobx-react";
import RaisedButton from 'material-ui/RaisedButton';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'

@observer
export default class MarketOrderPanel extends React.Component {
  render() {
    const { latestBuyPrice,latestSellPrice } = this.props.store;

    return (<div>
    <h1>
      Trader Simulator
    </h1>  
      <MuiThemeProvider>
      <RaisedButton primary={true}>{latestBuyPrice}</RaisedButton>
        <RaisedButton secondary={true}>{latestSellPrice}</RaisedButton>
      </MuiThemeProvider>
    </div>);
  };
}