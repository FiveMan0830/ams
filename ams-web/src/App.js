
import './App.css';
import TeamManage from './teamManage';
import React, { Component, forwardRef } from "react";
import { makeStyles } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import {Button} from '@material-ui/core'
import AddMember from './addMember';
import axios from 'axios'
import Home from './home';

class App extends Component{
  constructor(props) {
    super(props)
    this.state = {
      teamName : '',
      username : "",
    }
   
  }

  componentDidMount() {
    const searchParams = new URLSearchParams(window.location.search)
    if (searchParams.get('cn') !== null) {
      localStorage.setItem('cn', searchParams.get('cn'))
      localStorage.setItem('displayName', searchParams.get('displayName'))
      localStorage.setItem('teamName', searchParams.get('teamName'))
      // this.setState({teamName : searchParams.get('teamName')});
    } else {
      if (!!!localStorage.getItem('cn')) {
        var amsURL = process.env.REACT_APP_AMS_LOGIN_URL
        amsURL += '?' + encodeURIComponent('redirect_url=' + window.location.href)
        window.location.replace(amsURL)
        return
      }
    }
    
  }

render(){
  return (
    <div className="App">
       {/* <Home teamName = {this.state.teamName}  username = {localStorage.getItem("cn")}  ></Home> */}
       <Home teamName = {localStorage.getItem("teamName")}   username = {localStorage.getItem("cn")}  ></Home>
    </div>
  );
}
  
}

export default App;
