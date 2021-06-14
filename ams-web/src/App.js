
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
import MyAppBar from './Appbar';

const API_HOST = process.env.REACT_APP_HOST;

class App extends Component{
  constructor(props) {
    super(props)
    this.state = {
      teamName : '',
      username : "",
    }
   
  }

  componentWillMount() {
    const searchParams = new URLSearchParams(window.location.search)
    if (searchParams.get('cn') !== null) {
      localStorage.setItem('cn', searchParams.get('cn'))
      localStorage.setItem('displayName', searchParams.get('displayName'))
      localStorage.setItem('teamName', searchParams.get('teamName'))
      
      

      console.log(localStorage.getItem("cn"))
      console.log(localStorage.getItem("displayName"))
      console.log(localStorage.getItem("teamName"))

      axios.post(API_HOST + "/team/get/belonging-teams", localStorage.getItem("cn"))
      .then(res => {
          console.log(res)
          console.log("api:belonging-teams")
          this.setState({teamList: res.data})
          console.log("api:stateTeam" + this.state.teamName)
          if(localStorage.getItem("teamName") == "null") {
            console.log("api:state: "+this.state.teamName)
            console.log("api:res.data.name: "+ res.data[0].name)
            this.setState({teamName : res.data[0].name});
            localStorage.setItem('teamName', res.data[0].name)
          }
      })
      .catch(err => {
          console.log(err);
      })


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

  componentDidUpdate(prevProps, prevState) {
    console.log("componentDidUpdate:state: "+this.state.teamName)
      console.log("componentDidUpdate: " + localStorage.getItem("teamName"))

    if (prevState.teamName !== this.state.teamName) { // => 比較更新前後的 state 屬性
      console.log("api:state: "+this.state.teamName)
      console.log("local: " + localStorage.getItem("teamName"))
      this.setState({teamName : this.state.teamName});
      localStorage.setItem('teamName', this.state.teamName)
    }
  }

render(){
  return (
    <div className="App">
       {/* <Home teamName = {this.state.teamName}  username = {localStorage.getItem("cn")}  ></Home> */}


       <MyAppBar></MyAppBar>
       {console.log("render state: " + this.state.teamName )}
       {console.log("local: " + localStorage.getItem("teamName") )}
       {console.log("local boolean: " + (localStorage.getItem("teamName") == "null") )}
       {localStorage.getItem("teamName") == "null" ? 
        null
       :  
         <Home teamName = {localStorage.getItem("teamName")} username = {localStorage.getItem("cn")}></Home>
        //  <Home teamName = "OIS" username = {localStorage.getItem("cn")}></Home>
      }
    </div>
  );
}
  
}

export default App;
