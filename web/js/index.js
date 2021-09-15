window.onload = () => {
    function login() {
        const username = $("#username-field").val()
        const password = $("#password-field").val() 
        $.ajax({
            type: "POST",
            url: "http://localhost:8080/login",
            data: {
                Username: username,
                Password: password
            },
            success: res => {
                const queryStr = decodeURIComponent(window.location.search.substring(1)).split("&")
                const query = {}
                queryStr.forEach((item, i) => {
                    itemPair = item.split("=")
                    query[itemPair[0]] = itemPair[1]
                })
                console.log(res)
                if (!!query.redirect_url) {
                    var redirectURL = query.redirect_url + "?"
                    // redirectURL += "accessToken=" + res.accessToken + "&"
                    // console.log(redirectURL)
                    // window.location.replace(redirectURL.substring(0, redirectURL.length-1))
                    res.Attributes.forEach(attribute => {
                        if (attribute.Name !== "userPassword")
                            redirectURL += attribute.Name + "=" + attribute.Values[0] + "&"
                    });
                    console.log(redirectURL.substring(0, redirectURL.length - 1))
                    window.location.replace(redirectURL.substring(0, redirectURL.length - 1))
                } else {
                    throw "redirect_url not defined"
                }
            },
            error: (xhr, ajaxOptions, thrownError) => {
                console.log(xhr.status)
                console.log(thrownError)
                var errorMessage = document.getElementById('error')
                var temp = window.getComputedStyle(errorMessage).getPropertyValue("opacity");
                if (temp == 1) {
                    errorMessage.classList.toggle('shake');
                } else {
                    errorMessage.classList.toggle('fade');
                }
            }
        })
    }

    $("#login-button").click(() => {
        console.log("login")
        login()
    })

    $(document).keyup(event => {
        if (event.keyCode === 13) {
            login()
        }
    })
}
