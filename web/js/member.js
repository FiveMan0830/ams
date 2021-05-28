window.onload = () => {

    function addMember() {
        const groupname = $("#groupname-field").val()
        const username = $("#username-field").val()
        
        $.ajax({
            type: "POST",
            url: "http://localhost:8080/add/member",
            data: {
                GroupName: groupname, 
                Username: username
            },
            success: res => {
                const queryStr = decodeURIComponent(window.location.search.substring(1)).split("&")
                const query = {}
                queryStr.forEach((item, i) => {
                    itemPair = item.split("=")
                    query[itemPair[0]] = itemPair[1]
                })
                console.log(res)
                changeTextSuccess("Adding ");
            },
            error: (xhr, ajaxOptions, thrownError) => {
                console.log(xhr.status)
                console.log(thrownError)
                changeTextFail("Adding ");
            }
        })
    }

    function getMembers() {
        const groupname = $("#groupname-field").val()

        $.ajax({
            type: "POST",
            url: "http://localhost:8080/get/members",
            data: {
                GroupName: groupname
            },
            success: res => {
                const queryStr = decodeURIComponent(window.location.search.substring(1)).split("&")
                const query = {}
                queryStr.forEach((item, i) => {
                    itemPair = item.split("=")
                    query[itemPair[0]] = itemPair[1]
                })
                console.log(res);
                const ul = document.getElementById("member");
                ul.innerHTML = '';
                res.forEach(data => {
                    if (data != "") {
                       const li = document.createElement("li");
                        li.textContent = data;
                        ul.appendChild(li); 
                    }
                    
                })
            },
            error: (xhr, ajaxOptions, thrownError) => {
                console.log(xhr.status)
                console.log(thrownError)
            }
        })
    }

    function removeMember() {
        const groupname = $("#groupname-field").val()
        const username = $("#username-field").val()
        
        $.ajax({
            type: "POST",
            url: "http://localhost:8080/remove/member",
            data: {
                GroupName: groupname, 
                Username: username
            },
            success: res => {
                const queryStr = decodeURIComponent(window.location.search.substring(1)).split("&")
                const query = {}
                queryStr.forEach((item, i) => {
                    itemPair = item.split("=")
                    query[itemPair[0]] = itemPair[1]
                })
                console.log(res)
                changeTextSuccess("Removing ");
            },
            error: (xhr, ajaxOptions, thrownError) => {
                console.log(xhr.status)
                console.log(thrownError)
                changeTextFail("Removing ");
            }
        })
    }

    function changeTextSuccess(type) {
        const result = document.getElementById("username-field");
        const li = document.createElement("li");
        li.textContent = type + result.value + ' member success!';
        document.getElementById("member-list").textContent=li.textContent;
    }

    function changeTextFail(type) {
        const result = document.getElementById("username-field");
        const li = document.createElement("li");
        li.textContent = type + result.value + ' member fail!';
        document.getElementById("member-list").textContent=li.textContent;
    }

    $("#add-button").click(() => {
        addMember();
    })

    $("#remove-button").click(() => {
        removeMember();
    })

    $("#get-member-button").click(() => {
        getMembers();
        document.getElementById("member-list").textContent=null;
    })

}
