window.onload = () => {
    
    const API_HOST = "http://localhost:8080";
    getGroups();
    // var Team = function (team, leader) {
    //     var self = this;
    //     self.team = team;
    //     self.leader = leader;
    // }

    // var viewModel = function () {
    //     var self = this;
    //     self.users = ko.observableArray();

    //     //移除使用者事件
    //     self.RemoveUser = deletegroup(self.team);
    //     console.log(this.team);
    // }


   function creategroup(){
        const groupname = $("#groupname-field").val()
        const username = $("#username-field").val()
        const data = {
            GroupName: groupname, 
            Username: username
        }
        console.log(data)
        axios.post(API_HOST + "/create/team", data)
        .then(res => {
            const queryStr = decodeURIComponent(window.location.search.substring(1)).split("&")
                const query = {}
                queryStr.forEach((item, i) => {
                    itemPair = item.split("=")
                    query[itemPair[0]] = itemPair[1]
                })
            console.log(res);
        })
        .catch(err => {
            console.log(err);
        })
    }

    function deletegroup() {
        const groupname = $("#groupname-field").val()
        const data = {
            GroupName: groupname 
        }
        console.log(groupname);
        axios.post(API_HOST + "/delete/team", data)
        .then(res =>{
            const queryStr = decodeURIComponent(window.location.search.substring(1)).split("&")
                const query = {}
                queryStr.forEach((item, i) => {
                    itemPair = item.split("=")
                    query[itemPair[0]] = itemPair[1]
                })
            console.log(res);
        })
        .catch(err => {
            console.log(err);
        })
    }

    function getGroups() {
        axios.post(API_HOST + "/get/groups")
        .then(res => {
            const queryStr = decodeURIComponent(window.location.search.substring(1)).split("&")
                const query = {}
                queryStr.forEach((item, i) => {
                    itemPair = item.split("=")
                    query[itemPair[0]] = itemPair[1]
                })
            console.log(res);
            const table = document.getElementById("groups");
            table.innerHTML = '';
            res.data.forEach(data => {
                const tr = document.createElement("tr");
                const td = document.createElement("td");
                td.textContent = data;
                tr.appendChild(td);
                const td1 = document.createElement("td");
                leaderName = getLeaderName(data)
                leaderName.then((value) => {
                    td1.textContent = value;
                    console.log(value);
                });
                tr.appendChild(td1);
                // const td3 = document.createElement("td");
                // const button = document.createElement("button");
                // button.textContent = "X";
                // button.setAttribute("id", "btn");
                // button.setAttribute("value", data);
                // td3.appendChild(button);
                // tr.appendChild(td3);
                table.appendChild(tr);
            })
        })
        .catch(err => {
            console.log(err);
        })
    }

    function getLeaderName(groupname) {
        return new Promise(function(resolve, reject) {
            const data = {
                GroupName: groupname 
            }
            axios.post(API_HOST + "/get/leader", data)
            .then(res => {
                const queryStr = decodeURIComponent(window.location.search.substring(1)).split("&")
                const query = {}
                queryStr.forEach((item, i) => {
                    itemPair = item.split("=")
                    query[itemPair[0]] = itemPair[1]
                })
    
                // const table = document.getElementById("groups");
                // // table.innerHTML = '';
                // res.data.forEach(data => {
                //     const tr = document.createElement("tr");
                //     // const td = document.createElement("td");
                //     tr.textContent = data;
                //     // tr.appendChild(td);
                //     table.appendChild(tr);
                // })
    
                console.log(res);
                console.log(res.data);
                resolve(res.data);
            })
            .catch(err => {
                console.log(err);
                reject(err);
            })
        })
        

        

        // $.ajax({
        //     type: "POST",
        //     url: "http://localhost:8080/get/leader",
        //     data: {
        //         GroupName: groupname
        //     },
        //     success: res => {
        //         const queryStr = decodeURIComponent(window.location.search.substring(1)).split("&")
        //         const query = {}
        //         queryStr.forEach((item, i) => {
        //             itemPair = item.split("=")
        //             query[itemPair[0]] = itemPair[1]
        //         })
        //         res.forEach(data => {
        //             // console.log(data)
        //             string=data+""
        //             // console.log(string)
        //         })
        //     },
        //     error: (xhr, ajaxOptions, thrownError) => {
        //         console.log(xhr.status)
        //         console.log(thrownError)
        //     }
        // })
        // console.log(string)
        // return string;
    }

    function changeText(type) {
        const result = document.getElementById("groupname-field");
        const li = document.createElement("li");
        li.textContent = type + result.value + ' group success!';
        document.getElementById("group-list").textContent=li.textContent;
    }

    $("#create-button").click(() => {
        creategroup();
        window.location.reload();
    })

    $("#delete-button").click(() => {
        deletegroup();
        window.location.reload();
        // changeText("Delete ");
    })

    // var deleteBtn = document.getElementById("btn");
    // deleteBtn.click(() => {
    //     console.log(deleteBtn.value)
    //     deletegroup(deleteBtn.value);        
    //     // window.location.reload();
    //     changeText("Delete ");
    // })

    $(document).keyup(event => {
        if (event.keyCode === 13) {
            creategroup();
            getGroups();
            window.location.reload();
        }
    })
}
