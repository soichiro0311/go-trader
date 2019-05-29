import React from "react";

import { observer } from "mobx-react";
import RaisedButton from 'material-ui/RaisedButton';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'

@observer
export default class TodoList extends React.Component {
  render() {
    const { amount } = this.props.store;

    return (<div>
    <h1>
      Trader Simulator
    </h1>  
      <MuiThemeProvider>
        <RaisedButton secondary={true}>{amount}</RaisedButton>
      </MuiThemeProvider>
    </div>);
  };
}