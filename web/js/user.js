window.onload = () => {
    function getGroups() {
        const username = $("#username-field").val()

        $.ajax({
            type: "POST",
            url: "http://localhost:8080/get/groups/byuser",
            data: {
                Username: username
            },
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

    $("#search-button").click(() => {
        getGroups();
    })

    $(document).keyup(event => {
        if (event.keyCode === 13) {
            getGroups();
        }
    })
}
