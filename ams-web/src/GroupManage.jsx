
import React, { Component, forwardRef } from "react";
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import axios from 'axios';
import {
  FormControl,
  InputLabel,
  Input,
  Button,
} from '@material-ui/core';
import { ThreeSixty } from '@material-ui/icons';

const API_HOST = process.env.REACT_APP_HOST;

class GroupManage extends Component{
  constructor(props) {
    super(props)
    this.state = {
      deleteTeamName : "",
      teamName : "",
      teamList: [],
      leaderName:"",
      addMemberOpen:false,
    }
    this.initialize = this.initialize.bind(this)
    this.createTeam = this.createTeam.bind(this)
    this.deleteTeam = this.deleteTeam.bind(this)
  }

  initialize() {
    console.log("init")
    axios.get(API_HOST+'/team')
        .then(res => {
          const result = []
          res.data.map((team)=>{
            axios.post(API_HOST+'/team/',team)
              .then(res => {
                axios.post(API_HOST+'/team/get/leader', res.data)
                  .then(res2 => {
                    axios.post(API_HOST+'/team/', res2.data)
                    .then(res3 => {
                      result.push({teamName:res.data,leaderName:res3.data})
                      this.setState({teamList: result})
                    })
                    .catch(err => {
                        console.log(err);
                    })
                  })
                  .catch(err => {
                      console.log(err);
                  })
              })
              .catch(err => {
                  console.log(err);
              })

          })
         
        })
        .catch(err => {
            console.log(err);
        })
  }

  createTeam() {
    console.log("create: " + this.state.teamName)
    console.log("create: " + this.state.leaderName)
    const data = {
      GroupName: this.state.teamName,
      SelfUsername: this.state.leaderName
    }
    axios.post("http://localhost:8080/team/create", data)
        .then(res => {
          this.setState({leaderName:res.data})
        })
        .catch(err => {
            console.log(err);
        })
    this.initialize();
  }

  deleteTeam() {
    console.log("create: " + this.state.deleteTeamName)
    const data = {
      GroupName: this.state.deleteTeamName,
    }
    axios.post("http://localhost:8080/team/delete", data)
        .then(res => {
          // this.setState({leaderName:res.data})
        })
        .catch(err => {
            console.log(err);
        })
    this.initialize();
  }

  componentWillMount() {
    this.initialize();
  }


  
  render() {
    return (
      <div>
          <form>
            <FormControl >
              <InputLabel htmlFor="deleteTeamName">Team Name</InputLabel>
              <Input id="deleteTeamName" onChange={(e) => {this.setState({deleteTeamName: e.target.value})}}/>
            </FormControl>
            
            <Button className="Delete-Team" block size="lg" variant="outline-secondary" onClick={this.deleteTeam} value="Submit">
              DELETE TEAM
            </Button>
          </form>

          <form>
            <FormControl >
              <InputLabel htmlFor="teamName">Team Name</InputLabel>
              <Input id="teamName" onChange={(e) => {this.setState({teamName: e.target.value})}}/>
            </FormControl>
            <FormControl >
              <InputLabel htmlFor="leaderName">Leader Name</InputLabel>
              <Input id="leaderName" onChange={(e) => {this.setState({leaderName: e.target.value})}}/>
            </FormControl>
           
            
            <Button className="Create-Team" block size="lg" variant="outline-secondary" onClick={this.createTeam} value="Submit">
              CREATE TEAM
            </Button>
          </form>


          <TableContainer component={Paper}>
            <Table aria-label="simple table">
              <TableHead>
                <TableRow>
                  <TableCell align="left">Team Name</TableCell>
                  <TableCell align="left">Leader Name</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {this.state.teamList.map((team,index) => (
                  <TableRow align="left" key={index}>
                    <TableCell align="left">{team.teamName}</TableCell>
                    <TableCell align="left">{team.leaderName}</TableCell>
                    
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        
      </div>
    );
  }
  
}

export default GroupManage;
