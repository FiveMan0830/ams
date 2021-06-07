
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


class App extends Component{
  constructor(props) {
    super(props)
    this.state = {
      teamList: [],
      teamName : 'SSL LAB',
      username : "ssl1321ois",
      addMemberOpen: false,
      memberList : [],
    }
    this.handleAddMemberClose = this.handleAddMemberClose.bind(this);
    this.handleAddMemberOpen = this.handleAddMemberOpen.bind(this);
  }

  componentWillMount() {
    const data = {
      Username: this.state.username
    }
    axios.post("http://localhost:8080/team/get/belonging-teams", this.state.username)
      .then(res => {
          console.log(res)
          this.setState({teamList: res.data})
      })
      .catch(err => {
          console.log(err);
      })

    // axios.post("http://localhost:8080/team/get/name", String Team ID)
    //     .then(res => {
    //         this.setState({team: res.data})
    //     })
    //     .catch(err => {
    //         console.log(err);
    //     })
  }

  handleAddMemberOpen() {
    this.setState({addMemberOpen:true});
  };
  handleAddMemberClose() {
    this.setState({addMemberOpen:false});
  };

render(){
  return (
    <div className="App">
      <div className="selector">
         <FormControl>
          <InputLabel id="demo-simple-select-label">Team</InputLabel>
          <Select
            value={this.state.teamName}
            onChange={(e) => {this.setState({teamName: e.target.value})}}
          >
            {
              this.state.teamList.map((team,index) => {
                return(
                  <MenuItem key={index} value={team.name}>{team.name}</MenuItem>
                )
              })
            }
          </Select>
        </FormControl>
      </div>
      <div className="team-name">
        {this.state.teamName}
        {/* <Button onClick = {this.handleAddMemberOpen}>ADD</Button>
        <AddMember open={this.state.addMemberOpen} handleClose={this.handleAddMemberClose} memberList={this.state.memberList}/> */}
      </div>
      <div className="table">
        <TeamManage teamName = {this.state.teamName} username = {this.state.username}> </TeamManage>
      </div>
    </div>
  );
}
  
}

export default App;
