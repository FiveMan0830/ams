// import React, { Component, forwardRef } from "react";
// import PropTypes from 'prop-types';
// import { makeStyles } from '@material-ui/core/styles';
// import Button from '@material-ui/core/Button';
// import Avatar from '@material-ui/core/Avatar';
// import List from '@material-ui/core/List';
// import ListItem from '@material-ui/core/ListItem';
// import ListItemAvatar from '@material-ui/core/ListItemAvatar';
// import ListItemText from '@material-ui/core/ListItemText';
// import DialogTitle from '@material-ui/core/DialogTitle';
// import Dialog from '@material-ui/core/Dialog';
// import PersonIcon from '@material-ui/icons/Person';
// import AddIcon from '@material-ui/icons/Add';
// import Typography from '@material-ui/core/Typography';
// import { blue } from '@material-ui/core/colors';

// class addMember extends Component {
//   constructor(props) {
//     super(props)
//   }

//   render() {
//     return (
//       <div>
//         <Dialog onClose={this.props.handleClose()} aria-labelledby="simple-dialog-title" open={this.props.open}>
//         <DialogTitle id="simple-dialog-title">Set backup account</DialogTitle>
//         <List>
//             {this.props.memberList.map((member) => (
//             <ListItem button key={member.name}>
//                 <ListItemAvatar>
//                 <Avatar className="avatar-name" alt={member.name} src="/broken-image.jpg"/>
//                 </ListItemAvatar>
//                 <ListItemText primary={member.name} />
//             </ListItem>
//             ))}
//         </List>
//         </Dialog>
//         <Button onClick={this.props.handleClose()} color="primary">
//             Submit
//           </Button>
//       </div>
//     );
//   }
// }

// export default addMember