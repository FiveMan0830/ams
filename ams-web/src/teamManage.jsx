import React, { Component, forwardRef } from "react";
import MaterialTable from "material-table";
import { AddBox, ArrowDownward, Check, ChevronLeft, ChevronRight,
  Clear, DeleteOutline, Edit, FilterList, FirstPage, LastPage,
  MicNone,
  Remove, SaveAlt, Search, ViewColumn } from '@material-ui/icons';
import SwapHorizontalCircleOutlinedIcon from '@material-ui/icons/SwapHorizontalCircleOutlined';
import RemoveCircleOutlineOutlinedIcon from '@material-ui/icons/RemoveCircleOutlineOutlined';
import addMember from './addMember';
import axios from 'axios'

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
    this.handleAddMemberClose = this.handleAddMemberClose.bind(this);
    this.handleAddMemberOpen = this.handleAddMemberOpen.bind(this);
    this.state = {
      personal: [{unitID: localStorage.getItem("uid"),unitName: "Personal" }],
      selectTeam: [],
      teamName : "",
      startTime: 0,
      endTime: 0,
      addMemberOpen: false,
      columns: [
        { title: 'Display Name', field: 'name'},
        { title: 'Username', field: 'username'},
        
      ],
      data: [
        { name: 'Mehmet', username: 'Baran' },
        { name: 'Zerya Betül', username: 'Baran'},
      ]
    }
  }
  handleAddMemberOpen() {
    this.setState({addMemberOpen:true});
  };

  handleAddMemberClose() {
    this.setState({addMemberOpen:false});
  };

  componentWillMount() {
    const data = {
      GroupName: this.props.teamName
    }
    axios.post("http://localhost:8080/team/get/members", data)
        .then(res => {
          console.log(res);
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


  render() {
    return (
      <div>
        <MaterialTable
          icons={tableIcons}
          title="Member List"
          columns={this.state.columns}
          data={this.state.data}
          actions={[
            {
              icon: SwapHorizontalCircleOutlinedIcon,
              tooltip: 'Hand Over',
              onClick: (event, rowData) => alert("You saved " + rowData.name),
              // disabled: leader,
            },
            rowData => ({
              icon: RemoveCircleOutlineOutlinedIcon,
              tooltip: 'Delete User',
              onClick: (event, rowData) => alert("You want to delete " + rowData.name),
              // disabled: leader,
            })
          ]}
          options={{
            actionsColumnIndex: -1
          }}
          editable={{
            // onRowAdd: this.handleAddMemberOpen()
          }}
        />
        
        {/* <addMember open={this.state.addMemberOpen} handleClose={this.handleAddMemberClose} memberList={this.state.data}></addMember> */}
      </div>
    );
  }
}

export default TeamManage