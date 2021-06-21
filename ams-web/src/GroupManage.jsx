
import './Home.css';
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

const API_HOST = process.env.REACT_APP_HOST;

class GroupManage extends Component{
  constructor(props) {
    super(props)
    this.state = {
      teamName : "",
      teamList: [],
      leaderName:"",
      addMemberOpen:false,
    }
    this.initialize = this.initialize.bind(this)
    this.createTeam = this.createTeam.bind(this)
  }

  initialize() {
    console.log("init")
    axios.get(API_HOST+'/team')
        .then(res => {
          const result = []
          res.data.map((team)=>{
            result.push({teamID:team.uid})
          })
          this.setState({teamList: result})
        })
        .catch(err => {
            console.log(err);
        })


  }

  createTeam() {
    const data = {
      teamName: this.state.teamName,
      leaderName: this.state.leaderName
    }
    axios.post("http://localhost:8080/team/create", data)
        .then(res => {
          this.setState({leaderName:res.data})
        })
        .catch(err => {
            console.log(err);
        })
  }

  
  render() {
    return (
      <div>
          <form>
            <FormControl fullWidth={true}>
              <InputLabel htmlFor="description">Description</InputLabel>
              <Input id="description" onChange={(e) => {this.setState({description: e.target.value})}}/>
            </FormControl>

            <FormControl fullWidth={true}>
              <InputLabel htmlFor="description">Description</InputLabel>
              <Input id="description" onChange={(e) => {this.setState({description: e.target.value})}}/>
            </FormControl>
           
            
            <Button className="Create-Team" block size="lg" variant="outline-secondary" onClick={this.createTeam()} value="Submit">
              CREATE TEAM
            </Button>
          </form>


          <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="simple table">
              <TableHead>
                <TableRow>
                  <TableCell align="left">Team Name</TableCell>
                  <TableCell align="left">Leader Name</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {this.state.teamList.map((team,index) => (
                  <TableRow align="left" key={index}>
                    {/* <TableCell align="left">{team.UUID}</TableCell>
                    <TableCell align="left">{team.value}</TableCell> */}
                    
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
