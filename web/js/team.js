window.onload = () => {
    getGroups();

    function creategroup() {
        const groupname = $("#groupname-field").val()
        $.ajax({
            type: "POST",
            url: "http://localhost:8080/create/team",
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
                console.log(res)
            },
            error: (xhr, ajaxOptions, thrownError) => {
                console.log(xhr.status)
                console.log(thrownError)
            }
        })
    }

    function getGroups() {
        $.ajax({
            type: "POST",
            url: "http://localhost:8080/get/groups",
            success: res => {
                const queryStr = decodeURIComponent(window.location.search.substring(1)).split("&")
                const query = {}
                queryStr.forEach((item, i) => {
                    itemPair = item.split("=")
                    query[itemPair[0]] = itemPair[1]
                })
                console.log(res);
                const ul = document.getElementById("groups");
                ul.innerHTML = '';
                res.forEach(data => {
                    const li = document.createElement("li");
                    li.textContent = data;
                    ul.appendChild(li);
                })
            },
            error: (xhr, ajaxOptions, thrownError) => {
                console.log(xhr.status)
                console.log(thrownError)
            }
        })
    }

    function changeText() {
        const result = document.getElementById("groupname-field");
        const li = document.createElement("li");
        li.textContent = 'Create ' + result.value + ' group success!';
        document.getElementById("group-list").textContent=li.textContent;
    }

    $("#create-button").click(() => {
        creategroup();
        changeText();
        window.location.reload();
    })

    $(document).keyup(event => {
        if (event.keyCode === 13) {
            creategroup();
            getGroups();
            window.location.reload();
        }
    })
}
