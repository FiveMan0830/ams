window.onload = function() {
    function login() {
        const username = document.getElementById("username-field").value
        const password = document.getElementById("password-field").value
        $.ajax({
            type: "POST",
            url: "http://localhost:9000/api/login",
            data: {
                username: username,
                password: password
            },
            success: function(res) {
                console.log(res)
            }
        })
    }
    $("#login-button")
}
