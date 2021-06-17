
import './App.css';
import React, { Component, forwardRef } from "react";
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
      localStorage.setItem('uid', searchParams.get('uid'))
      localStorage.setItem('cn', searchParams.get('cn'))
      localStorage.setItem('sn', searchParams.get('sn'))
      localStorage.setItem('givenName', searchParams.get('givenName'))
      localStorage.setItem('displayName', searchParams.get('displayName'))
      localStorage.setItem('mail', searchParams.get('mail'))
      localStorage.setItem('teamName', searchParams.get('teamName'))

      axios.post(API_HOST + "/team/get/belonging-teams", localStorage.getItem("cn"))
      .then(res => {
          this.setState({teamList: res.data})
          if(localStorage.getItem("teamName") == "null") {
            this.setState({teamName : res.data[0].name});
            localStorage.setItem('teamName', res.data[0].name)
          }
      })
      .catch(err => {
          console.log(err);
      })
    } else {
      if (!!!localStorage.getItem('cn')) {
        var amsURL = process.env.REACT_APP_AMS_LOGIN_URL
        amsURL += '?' + encodeURIComponent('redirect_url=' + process.env.REACT_APP_WEB)
        window.location.replace(amsURL)
        return
      }
    }
  }

  componentDidUpdate(prevProps, prevState) {
    if (prevState.teamName !== this.state.teamName) { // => 比較更新前後的 state 屬性
      this.setState({teamName : this.state.teamName});
      localStorage.setItem('teamName', this.state.teamName)
    }
  }

render(){
  return (
    <div className="App">
       <MyAppBar></MyAppBar>
       {console.log("render state: " + this.state.teamName )}
       {console.log("local: " + localStorage.getItem("teamName") )}
       {console.log("local boolean: " + (localStorage.getItem("teamName") == "null") )}
       {localStorage.getItem("teamName") == "null" ? 
        null
       :  
         <Home teamName = {localStorage.getItem("teamName")} username = {localStorage.getItem("cn")}></Home>
      }
    </div>
  );
}
  
}

export default App;
