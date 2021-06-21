
import './Home.css';
import TeamManage from './teamManage';
import React, { Component, forwardRef } from "react";
import { makeStyles } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import {Button} from '@material-ui/core'
import AddMember from './addMember';
import axios from 'axios';

const API_HOST = process.env.REACT_APP_HOST;

class GroupManage extends Component{
  constructor(props) {
    super(props)
    this.state = {
      teamList: [],
      teamName : this.props.teamName,
      username : this.props.username,
      addMemberOpen: false,
      memberList : [],
    }
  }


render(){
  return (
    <div>
        
    </div>
  );
}
  
}

export default GroupManage;
