import React, { Component, forwardRef } from "react";
import MaterialTable from "material-table";
import { AddBox, ArrowDownward, Check, ChevronLeft, ChevronRight,
  Clear, DeleteOutline, Edit, FilterList, FirstPage, LastPage,

  Remove, SaveAlt, Search, ViewColumn } from '@material-ui/icons';
import SwapHorizontalCircleOutlinedIcon from '@material-ui/icons/SwapHorizontalCircleOutlined';
import RemoveCircleOutlineOutlinedIcon from '@material-ui/icons/RemoveCircleOutlineOutlined';
import axios from 'axios'
import StarIcon from '@material-ui/icons/Star';
import {Button} from '@material-ui/core';
import AddMember from './addMember';

  const tableIcons = {
    Add: forwardRef((props, ref) => <AddBox {...props} ref={ref} />),
    Check: forwardRef((props, ref) => <Check {...props} ref={ref} />),
    Clear: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
    Delete: forwardRef((props, ref) => <DeleteOutline {...props} ref={ref} />),
    DetailPanel: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
    Edit: forwardRef((props, ref) => <Edit {...props} ref={ref} />),
    Export: forwardRef((props, ref) => <SaveAlt {...props} ref={ref} />),
    Filter: forwardRef((props, ref) => <FilterList {...props} ref={ref} />),
    FirstPage: forwardRef((props, ref) => <FirstPage {...props} ref={ref} />),
    LastPage: forwardRef((props, ref) => <LastPage {...props} ref={ref} />),
    NextPage: forwardRef((props, ref) => <ChevronRight {...props} ref={ref} />),
    PreviousPage: forwardRef((props, ref) => <ChevronLeft {...props} ref={ref} />),
    ResetSearch: forwardRef((props, ref) => <Clear {...props} ref={ref} />),
    Search: forwardRef((props, ref) => <Search {...props} ref={ref} />),
    SortArrow: forwardRef((props, ref) => <ArrowDownward {...props} ref={ref} />),
    ThirdStateCheck: forwardRef((props, ref) => <Remove {...props} ref={ref} />),
    ViewColumn: forwardRef((props, ref) => <ViewColumn {...props} ref={ref} />)
  };

class TeamManage extends Component {
  constructor(props) {
    super(props)
    this.state = {
      teamName : "",
      columns: [
        { title: '', field: 'isLeader',render: rowData => <StarIcon /> },
        { title: 'Display Name', field: 'name'},
        { title: 'Username', field: 'username'},
      ],
      memberList: [],
      leaderName:"",
    }
    this.initialize = this.initialize.bind(this)
    this.handoverRole = this.handoverRole.bind(this)
    this.handleAddMemberClose = this.handleAddMemberClose.bind(this);
    this.handleAddMemberOpen = this.handleAddMemberOpen.bind(this);

  }
  
  handleAddMemberOpen() {
    this.setState({addMemberOpen:true});
  };
  handleAddMemberClose() {
    this.setState({addMemberOpen:false});
  };
  handoverRole(username){
    const groupRquest = {
      GroupName: this.props.teamName,
      SelfUsername: this.props.username,
      InputUsername: username,
    }
    axios.post("http://localhost:8080/team/leader/handover",groupRquest)
        .then(res => {
          console.log(res);
        })
        .catch(err => {
          console.log(err);
        })
  }

  deleteMember(username){
    const removeMemberRquest = {
      GroupName: this.props.teamName,
      Username: username,
    }
    axios.post("http://localhost:8080/remove/member",removeMemberRquest)
        .then(res => {
          console.log(res);
        })
        .catch(err => {
          console.log(err);
        })
  }

  initialize(){
    console.log(this.props.teamName)
    const data = {
      GroupName: this.props.teamName
    }
    axios.post("http://localhost:8080/team/get/member/name", data)
        .then(res => {
          const result = []
          for(var i = 0;i<res.data.length;i++){
            result.push({username:res.data[i].username,name:res.data[i].displayname})
          }
          this.setState({memberList:result})
        })
        .catch(err => {
            console.log(err);
        })

    axios.post("http://localhost:8080/team/get/leader", data)
        .then(res => {
          this.setState({leaderName:res.data})
        })
        .catch(err => {
            console.log(err);
        })
  }
  componentWillMount() {
    this.initialize();
  }

  componentDidUpdate(prevProps) {
    if (this.props.teamName !== prevProps.teamName) {
      this.initialize();
    }
  }
  
  render() {
    return (
      <div>
        <div className="team-name">
        {this.state.teamName}
          <Button onClick = {this.handleAddMemberOpen}>ADD</Button>
          <AddMember open={this.state.addMemberOpen} handleClose={this.handleAddMemberClose} memberList={this.state.memberList}/>
        </div>
        <MaterialTable
          icons={tableIcons}
          title="Member List"
          columns={[
            { title: '', field: 'isLeader',render: rowData =>rowData.username==this.state.leaderName?<StarIcon />:"" },
            { title: 'Display Name', field: 'name'},
            { title: 'Username', field: 'username'},
          ]}
          data={this.state.memberList}
          actions={[
            rowData => ({
              icon: SwapHorizontalCircleOutlinedIcon,
              tooltip: 'Hand Over',
              onClick: (event, rowData) => this.handoverRole(rowData.username),
              disabled: rowData.username==this.state.leaderName? true:false,
              hidden: this.props.username==this.state.leaderName?false:true,
            }),
            rowData => ({
              icon: RemoveCircleOutlineOutlinedIcon,
              tooltip: 'Delete User',
              onClick: (event, rowData) => this.deleteMember(rowData.username),
              disabled: rowData.username==this.state.leaderName? true:false,
              hidden: this.props.username==this.state.leaderName?false:true,
            })
          ]}
          options={{
            actionsColumnIndex: -1,
          }}
          localization={{
            header: {
              actions: ''
            },  
          }}
        />
      </div>
    );
  }
}

export default TeamManage