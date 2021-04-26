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


class AddMember extends Component {
  constructor(props) {
    super(props)
    this.state = {
      memberList : [],
    }
  }

  componentWillMount(){
    axios.get("http://localhost:8080/all/username")
    .then(res => {
      const result = []
      console.log(res.data)
      for(var i = 0;i<res.data.length;i++){
        var flag = false
        for(var j = 0; j<this.props.memberList.length;j++){
          if(this.props.memberList[j].username == res.data[i].username ){
            flag = true;
          }
        }
        if(!flag){
          result.push({username:res.data[i].username,name:res.data[i].displayname, isSelect:false})
        }
      }
      this.setState({memberList:result})
    })
    .catch(err => {
        console.log(err);
    })
  }

  addMember(memberList){
    
    const addMemberRquest = {
      GroupName: this.props.teamName,
      Username: username,
    }
    axios.post("http://localhost:8080/team/add/member",addMemberRquest)
        .then(res => {
          console.log(res);
        })
        .catch(err => {
          console.log(err);
        })
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
                    <ListItemAvatar>
                        <Avatar className="avatar-name" alt={member.name} src="/broken-image.jpg"/>
                    </ListItemAvatar>
                    <ListItemText primary={member.name} secondary={member.username}/>
                </ListItem>
                ))}
            </List>
        </DialogContent>
        <DialogActions>
            <Button onClick={this.props.handleClose} color="primary">
                Submit
            </Button>
        </DialogActions>
        </Dialog>
      </div>
    );
  }
}

export default AddMember