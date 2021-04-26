import React, { Component, forwardRef } from "react";
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import Avatar from '@material-ui/core/Avatar';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemText from '@material-ui/core/ListItemText';
import DialogTitle from '@material-ui/core/DialogTitle';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import PersonIcon from '@material-ui/icons/Person';
import AddIcon from '@material-ui/icons/Add';
import Typography from '@material-ui/core/Typography';
import { blue } from '@material-ui/core/colors';
import axios from 'axios'
import Checkbox from '@material-ui/core/Checkbox';



class AddMember extends Component {
  constructor(props) {
    super(props)
    this.state = {
      memberList : [],
      open : false,
    }
    this.handleSelectMember= this.handleSelectMember.bind(this);
    this.initialize= this.initialize.bind(this);
    this.addMember= this.addMember.bind(this);

  }

  initialize(){
    axios.get("http://localhost:8080/all/username")
    .then(res => {
      const result = []
      res.data.map((user)=>{
        var flag = false
        this.props.memberList.map((member)=>{
          if(member.username == user.username ){
            flag = true;
          }
        })
        if(!flag){
          result.push({username:user.username,name:user.displayname, isSelect:false})
        }
        this.setState({memberList:result})
      })
    })
    .catch(err => {
        console.log(err);
    })
  }
  componentWillMount(){
    this.initialize();
  }

  addMember(){
    this.state.memberList.map((member)=>{
      if(member.isSelect === true){
        member.isSelect = false;
        const addMemberRquest = {
          GroupName: this.props.teamName,
          Username: member.username,
        }
        axios.post("http://localhost:8080/team/add/member",addMemberRquest)
            .then(res => {
              console.log(res);
              this.props.handleClose()
            })
            .catch(err => {
              console.log(err);
              this.props.handleClose()
            })
      }
    })
    // this.props.handleClose()
  }

  componentDidUpdate(prevProps) {
    if (this.props.teamName !== prevProps.teamName) {
      this.initialize();
    }    
    if(this.props.open === true && this.state.open === false){
      this.setState({open: true});
      this.initialize();
    }
    if(this.props.open === false && this.state.open === true){
      this.setState({open: false});
    }
  }

  handleSelectMember(event) {   
    this.state.memberList.map((member)=>{
      if( member.username === event.target.value) {
        member.isSelect = event.target.checked;
      }
    })
    this.setState({ memberList: this.state.memberList});
    console.log(this.state.memberList)
  }

  render() {
    return (
      <div>
        <Dialog 
        onClose={this.props.handleClose} 
        aria-labelledby="simple-dialog-title" 
        open={this.props.open}>
        <DialogTitle id="simple-dialog-title">Add Member</DialogTitle>
        <DialogContent>
            <List>
                {this.state.memberList.map((member) => (
                <ListItem button key={member.name}>
                  <Checkbox  value={member.username} checked={member.isSelect} onChange={this.handleSelectMember}></Checkbox>
                    <ListItemAvatar>
                        <Avatar className="avatar-name" alt={member.name} src="/broken-image.jpg"/>
                    </ListItemAvatar>
                    <ListItemText primary={member.name} secondary={member.username}/>
                </ListItem>
                ))}
            </List>
        </DialogContent>
        <DialogActions>
            <Button onClick={this.addMember} color="primary">
                Submit
            </Button>
        </DialogActions>
        </Dialog>
      </div>
    );
  }
}

export default AddMember