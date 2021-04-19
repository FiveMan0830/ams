import axios from 'axios'
import moment from 'moment'


export function TeamGetMemberOf(username) { 
    var result = new Array();
    const data = {
        Username: username
    }
    axios.post("http://localhost:8080/team/get/memberOf", data)
        .then(res => {
            console.log(res);
            for(var i = 0;i<res.data.length;i++){
                result.push(res.data[i])
            }
            console.log(result);
            // return result;
        })
        .catch(err => {
            console.log(err);
        })
    return result;
}